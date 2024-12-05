//go:build integration

package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"

	pb "github.com/ZergsLaw/back-template/api/user/v1"
	"github.com/ZergsLaw/back-template/internal/grpchelper"
	"github.com/ZergsLaw/back-template/internal/testhelper"
)

const (
	queueCfgPath  = `testdata/nats.conf`
	queueUsername = `test_svc`
	queuePassword = `test_pass`
	migrateDir    = `./migrate`
	caCrtPath     = `../../certs/cockroach/ca.crt`
	caKeyPath     = `../../certs/cockroach/ca.key`
	nodeCrtPath   = `../../certs/cockroach/nodes/node1/node.crt`
	nodeKeyPath   = `../../certs/cockroach/nodes/node1/node.key`
	clientCrtPath = `../../certs/cockroach/client.root.crt`
	clientKeyPath = `../../certs/cockroach/client.root.key`
)

const (
	username1  = `username`
	fullName   = `full name`
	email      = `email@email.com`
	pass1      = `11111111`
	username2  = `username2`
	email2     = `email2@email.com`
	pass2      = `22222222`
	avatarPath = `testdata/test.jpg`
)

func initService(t *testing.T, ctx context.Context) (*require.Assertions, pb.UserExternalAPIClient, *config) {
	t.Helper()

	assert := require.New(t)

	reg := prometheus.NewPedanticRegistry()
	namespace := testhelper.Namespace(t)
	subsystem := testhelper.Namespace(t) + "_subsystem"

	devLogger := slog.New(slog.NewJSONHandler(
		os.Stderr, &slog.HandlerOptions{ //nolint:exhaustruct
			AddSource: true,
			Level:     slog.LevelDebug,
		}),
	)
	grpclog.SetLoggerV2(grpchelper.NewLogger(devLogger))

	postgresCfg := testhelper.Postgres(
		ctx,
		t,
		assert,
	)

	username := "test_svc"
	pass := "test_pass"

	endpoint := testhelper.Minio(
		ctx,
		t,
		assert,
		username, pass, "", false,
		"local-1",
	)

	cfg := config{
		Server: server{
			Host: testhelper.Host,
			Port: ports{
				GRPC:   testhelper.UnusedTCPPort(t, assert, testhelper.Host),
				Metric: testhelper.UnusedTCPPort(t, assert, testhelper.Host),
				GW:     testhelper.UnusedTCPPort(t, assert, testhelper.Host),
				Files:  testhelper.UnusedTCPPort(t, assert, testhelper.Host),
			},
		},
		DB: dbConfig{
			MigrateDir: migrateDir,
			Driver:     "postgres",
			Postgres:   postgresCfg,
		},
		FileStore: fileStoreConfig{
			Secure:    false,
			Endpoint:  endpoint,
			AccessKey: username,
			SecretKey: pass,
			Region:    "local-1",
		},
		AuthKey: "super-duper-secret-key-qwertyuio",
	}
	addr := net.JoinHostPort(cfg.Server.Host, fmt.Sprintf("%d", cfg.Server.Port.GRPC))

	errc := make(chan error)
	ctxShutdown, shutdown := context.WithCancel(ctx)
	go func() { errc <- run(ctxShutdown, cfg, reg, namespace) }()
	t.Cleanup(func() {
		shutdown()
		assert.NoError(<-errc)
	})
	assert.NoError(testhelper.WaitTCPPort(ctx, addr))

	clientMetric := grpchelper.NewClientMetrics(reg, namespace, subsystem)
	conn, err := grpchelper.Dial(ctx, addr, devLogger, clientMetric,
		[]grpc.UnaryClientInterceptor{},
		[]grpc.StreamClientInterceptor{},
		[]grpc.DialOption{},
	)
	assert.NoError(err)

	return assert, pb.NewUserExternalAPIClient(conn), &cfg
}

func authCtx(ctx context.Context, token string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.MD{
		"authorization": {fmt.Sprintf("Bearer %s", token)},
	})
}
