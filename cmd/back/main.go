package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hellofresh/health-go/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sipki-tech/database/connectors"
	"google.golang.org/grpc/grpclog"
	"gopkg.in/yaml.v3"

	user_pb "github.com/ZergsLaw/back-template/api/user/v1"
	"github.com/ZergsLaw/back-template/cmd/back/internal/adapters/files"
	"github.com/ZergsLaw/back-template/cmd/back/internal/adapters/repo"
	"github.com/ZergsLaw/back-template/cmd/back/internal/api/grpc"
	"github.com/ZergsLaw/back-template/cmd/back/internal/api/http"
	"github.com/ZergsLaw/back-template/cmd/back/internal/app"
	"github.com/ZergsLaw/back-template/cmd/back/internal/auth"
	"github.com/ZergsLaw/back-template/internal/flags"
	"github.com/ZergsLaw/back-template/internal/grpchelper"
	"github.com/ZergsLaw/back-template/internal/logger"
	"github.com/ZergsLaw/back-template/internal/metrics"
	"github.com/ZergsLaw/back-template/internal/password"
	"github.com/ZergsLaw/back-template/internal/serve"
	"github.com/ZergsLaw/back-template/internal/version"
)

const (
	exitCode       = 2
	configFileSize = 1024 * 1024
)

type (
	config struct {
		Server    server          `yaml:"server" env:"SERVER"`
		DB        dbConfig        `yaml:"db" env:"DB"`
		FileStore fileStoreConfig `yaml:"file_store" env:"FILE_STORE"`
		AuthKey   string          `yaml:"auth_key" env:"AUTH_KEY"`
		DevMode   bool            `yaml:"dev_mode" env:"DEV_MODE"`
	}
	server struct {
		Host string `yaml:"host" env:"HOST"`
		Port ports  `yaml:"port" env:"PORT"`
	}
	ports struct {
		GRPC   uint16 `yaml:"grpc" env:"GRPC"`
		Metric uint16 `yaml:"metric" env:"METRIC"`
		GW     uint16 `yaml:"gw" env:"GW"`
		Files  uint16 `yaml:"files" env:"FILES"`
	}
	dbConfig struct {
		MigrateDir string `yaml:"migrate_dir" env:"MIGRATE_DIR"`
		Driver     string `yaml:"driver" env:"DRIVER"`
		Postgres   string `yaml:"postgres" env:"POSTGRES"`
	}
	fileStoreConfig struct {
		Secure       bool   `yaml:"secure" env:"SECURE"`
		Endpoint     string `yaml:"endpoint" env:"ENDPOINT"`
		AccessKey    string `yaml:"access_key" env:"ACCESS_KEY"`
		SecretKey    string `yaml:"secret_key" env:"SECRET_KEY"`
		SessionToken string `yaml:"session_token" env:"SESSION_TOKEN"`
		Region       string `yaml:"region" env:"REGION"`
	}
)

func main() {
	var (
		cfgFile  = &flags.File{DefaultPath: "/config.yml", MaxSize: configFileSize}
		logLevel = &flags.Level{Level: slog.LevelDebug}
	)

	flag.Var(cfgFile, "cfg", "path to config file")
	flag.Var(logLevel, "log_level", "log level")
	flag.Parse()

	log := buildLogger(logLevel.Level)
	grpclog.SetLoggerV2(grpchelper.NewLogger(log))

	appName := filepath.Base(os.Args[0])
	ctxParent := logger.NewContext(context.Background(), log.With(slog.String(logger.Version.String(), version.System())))
	ctx, cancel := signal.NotifyContext(ctxParent, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	defer cancel()
	go forceShutdown(ctx)

	err := start(ctx, cfgFile, appName)
	if err != nil {
		log.Error("shutdown",
			slog.String(logger.Error.String(), err.Error()),
		)
		os.Exit(exitCode)
	}
}

func start(ctx context.Context, cfgFile io.Reader, appName string) error {
	cfg := config{}
	err := yaml.NewDecoder(cfgFile).Decode(&cfg)
	if err != nil {
		return fmt.Errorf("yaml.NewDecoder.Decode: %w", err)
	}

	reg := prometheus.NewPedanticRegistry()

	return run(ctx, cfg, reg, appName)
}

func run(ctx context.Context, cfg config, reg *prometheus.Registry, namespace string) error {
	log := logger.FromContext(ctx)
	m := metrics.New(reg, namespace)

	r, err := repo.New(ctx, reg, namespace, repo.Config{
		Postgres: connectors.Raw{
			Query: cfg.DB.Postgres,
		},
		MigrateDir: cfg.DB.MigrateDir,
		Driver:     cfg.DB.Driver,
	})
	if err != nil {
		return fmt.Errorf("repo.New: %w", err)
	}
	defer func() {
		err := r.Close()
		if err != nil {
			log.Error("close database connection", slog.String(logger.Error.String(), err.Error()))
		}
	}()

	fileStore, err := files.New(ctx, reg, namespace, files.Config{
		Secure:       cfg.FileStore.Secure,
		Endpoint:     cfg.FileStore.Endpoint,
		AccessKey:    cfg.FileStore.AccessKey,
		SecretKey:    cfg.FileStore.SecretKey,
		SessionToken: cfg.FileStore.SessionToken,
		Region:       cfg.FileStore.Region,
	})
	if err != nil {
		return fmt.Errorf("files.New: %w", err)
	}

	ph := password.New()

	authModule := auth.New(cfg.AuthKey)
	module := app.New(r, ph, authModule, idGenerator{}, r, fileStore)
	grpcAPI := grpc.New(ctx, m, module, reg, namespace)

	httpAPI := http.New(ctx, module)

	const healthTimeout = 5 * time.Second

	// add some checks on instance creation
	h, err := health.New(
		health.WithComponent(
			health.Component{
				Name:    namespace,
				Version: version.System(),
			},
		),
		health.WithChecks(
			health.Config{
				Name:    "postgres",
				Timeout: healthTimeout,
				Check:   r.Health,
			},
			health.Config{
				Name:    "minio",
				Timeout: healthTimeout,
				Check:   fileStore.HealthCheck,
			},
		),
	)
	if err != nil {
		return fmt.Errorf("health.New: %w", err)
	}

	gwCfg := serve.GateWayConfig{
		FS:             user_pb.OpenAPI,
		Spec:           "user.swagger.json",
		GRPCServerPort: cfg.Server.Port.GRPC,
		Reg:            reg,
		Namespace:      namespace,
		GRPCGWPattern:  "/",
		DocsUIPattern:  "/user/api/v1/docs/",
		Register:       user_pb.RegisterUserExternalAPIHandler,
		Healthcheck:    h.Handler(),
		DevMode:        cfg.DevMode,
	}

	return serve.Start(
		ctx,
		serve.Metrics(log.With(slog.String(logger.Module.String(), "metric")), cfg.Server.Host, cfg.Server.Port.Metric, reg),
		serve.GRPC(log.With(slog.String(logger.Module.String(), "gRPC")), cfg.Server.Host, cfg.Server.Port.GRPC, grpcAPI),
		serve.GRPCGateWay(log.With(slog.String(logger.Module.String(), "gRPC-Gateway")), cfg.Server.Host, cfg.Server.Port.GW, gwCfg),
		serve.HTTP(log.With(slog.String(logger.Module.String(), "files")), cfg.Server.Host, cfg.Server.Port.Files, httpAPI),
		module.Process,
	)
}

func buildLogger(level slog.Level) *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{ //nolint:exhaustruct
				AddSource: true,
				Level:     level,
			},
		),
	)
}

func forceShutdown(ctx context.Context) {
	log := logger.FromContext(ctx)
	const shutdownDelay = 15 * time.Second

	<-ctx.Done()
	time.Sleep(shutdownDelay)

	log.Error("failed to graceful shutdown")
	os.Exit(exitCode)
}

var _ app.ID = &idGenerator{}

type idGenerator struct{}

// New implements app.ID.
func (idGenerator) New() uuid.UUID {
	return uuid.Must(uuid.NewV4())
}
