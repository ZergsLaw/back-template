package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sipki-tech/database/connectors"
	"google.golang.org/grpc/grpclog"
	"gopkg.in/yaml.v3"

	user_pb "github.com/Bar-Nik/back-template/api/user/v1"
	session_client "github.com/Bar-Nik/back-template/cmd/session/client"
	"github.com/Bar-Nik/back-template/cmd/user/internal/adapters/files"
	"github.com/Bar-Nik/back-template/cmd/user/internal/adapters/queue"
	"github.com/Bar-Nik/back-template/cmd/user/internal/adapters/repo"
	"github.com/Bar-Nik/back-template/cmd/user/internal/api/grpc"
	"github.com/Bar-Nik/back-template/cmd/user/internal/api/http"
	"github.com/Bar-Nik/back-template/cmd/user/internal/app"
	session_adapter "github.com/Bar-Nik/back-template/internal/adapters/session"
	"github.com/Bar-Nik/back-template/internal/flags"
	"github.com/Bar-Nik/back-template/internal/grpchelper"
	"github.com/Bar-Nik/back-template/internal/logger"
	"github.com/Bar-Nik/back-template/internal/metrics"
	"github.com/Bar-Nik/back-template/internal/password"
	"github.com/Bar-Nik/back-template/internal/serve"
)

type (
	config struct {
		Server    server          `yaml:"server"`
		Clients   clients         `yaml:"clients"`
		DB        dbConfig        `yaml:"db"`
		FileStore fileStoreConfig `yaml:"file_store"`
		Queue     queueConfig     `yaml:"queue"`
		DevMode   bool            `yaml:"dev_mode"`
	}
	server struct {
		Host string `yaml:"host"`
		Port ports  `yaml:"port"`
	}
	ports struct {
		GRPC   uint16 `yaml:"grpc"`
		Metric uint16 `yaml:"metric"`
		GW     uint16 `yaml:"gw"`
		Files  uint16 `yaml:"files"`
	}
	dbConfig struct {
		MigrateDir string                 `yaml:"migrate_dir"`
		Driver     string                 `yaml:"driver"`
		Cockroach  connectors.CockroachDB `yaml:"cockroach"`
	}
	fileStoreConfig struct {
		Secure       bool   `yaml:"secure"`
		Endpoint     string `yaml:"endpoint"`
		AccessKey    string `yaml:"access_key"`
		SecretKey    string `yaml:"secret_key"`
		SessionToken string `yaml:"session_token"`
		Region       string `yaml:"region"`
	}
	clients struct {
		Session string `yaml:"session"`
	}
	queueConfig struct {
		URLs     []string `yaml:"urls"`
		Username string   `yaml:"username"`
		Password string   `yaml:"password"`
	}
)

var (
	cfgFile  = &flags.File{DefaultPath: "config.yml", MaxSize: 1024 * 1024}
	logLevel = &flags.Level{Level: slog.LevelDebug}
)

const version = "v0.1.0"

func main() {
	flag.Var(cfgFile, "cfg", "path to config file") // как работает?
	flag.Var(logLevel, "log_level", "log level")    // как работает?
	flag.Parse()                                    // как работает?

	log := buildLogger(logLevel.Level) // создаем logger передаем в него минимальный уровень логирования
	grpclog.SetLoggerV2(grpchelper.NewLogger(log))

	appName := filepath.Base(os.Args[0])                                                                                              // Из пути берем последнее имя для названия приложения
	ctxParent := logger.NewContext(context.Background(), log.With(slog.String(logger.Version.String(), version)))                     // Создаем контекст с добавлением логгера
	ctx, cancel := signal.NotifyContext(ctxParent, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM) // Зачем?
	defer cancel()                                                                                                                // ?
	go forceShutdown(ctx)                                                                                                             // ?

	err := start(ctx, cfgFile, appName) 
	if err != nil {
		log.Error("shutdown",
			slog.String(logger.Error.String(), err.Error()),
		)
		os.Exit(2)
	}
}

func start(ctx context.Context, cfgFile io.Reader, appName string) error {
	cfg := config{}
	err := yaml.NewDecoder(cfgFile).Decode(&cfg) // десериализация конфига
	if err != nil {
		return fmt.Errorf("yaml.NewDecoder.Decode: %w", err)
	}

	reg := prometheus.NewPedanticRegistry() // Зачем?

	return run(ctx, cfg, reg, appName) // сервиc user передаем контекст, конфиг и название и reg?
}

func run(ctx context.Context, cfg config, reg *prometheus.Registry, namespace string) error {
	log := logger.FromContext(ctx)   // извлекаем логер
	m := metrics.New(reg, namespace) // Зачем?

	r, err := repo.New(ctx, reg, namespace, repo.Config{ // создаем bd user?
		Cockroach:  cfg.DB.Cockroach,
		MigrateDir: cfg.DB.MigrateDir,
		Driver:     cfg.DB.Driver,
	})
	if err != nil {
		return fmt.Errorf("repo.New: %w", err)
	}
	defer func() { // что это и для чего?
		err := r.Close()
		if err != nil {
			log.Error("close database connection", slog.String(logger.Error.String(), err.Error()))
		}
	}()

	fileStore, err := files.New(ctx, reg, namespace, files.Config{ // что это и для чего?
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

	client, err := session_client.New(ctx, log, reg, namespace, cfg.Clients.Session) // создаем сессию клиента
	if err != nil {
		return fmt.Errorf("session_client.New: %w", err)
	}
	sessionSvc := session_adapter.New(client, convertErr) // Записываем данные клиента? проверка на ошибки

	q, err := queue.New(ctx, reg, namespace, queue.Config{ // это что- то с очередями
		URLs:     cfg.Queue.URLs,
		Username: cfg.Queue.Username,
		Password: cfg.Queue.Password,
	})
	if err != nil {
		return fmt.Errorf("queue.New: %w", err)
	}
	defer func() {
		err = q.Close()
		if err != nil {
			log.Error("close queue connection", slog.String(logger.Error.String(), err.Error()))
		}
	}()

	ph := password.New() // проверка пароля

	module := app.New(r, ph, sessionSvc, fileStore, q)  // ?
	grpcAPI := grpc.New(ctx, m, module, reg, namespace) // Создаем grps server

	httpAPI := http.New(ctx, module) // ?

	gwCfg := serve.GateWayConfig{ // Конфигурация прокси сервер?
		FS:             user_pb.OpenAPI,
		Spec:           "user.swagger.json",
		GRPCServerPort: cfg.Server.Port.GRPC,
		Reg:            reg,
		Namespace:      namespace,
		GRPCGWPattern:  "/",
		DocsUIPattern:  "/user/api/v1/docs/",
		Register:       user_pb.RegisterUserExternalAPIHandler,
		DevMode:        cfg.DevMode,
	}

	return serve.Start( // Запускаем прокси сервер
		ctx,
		serve.Metrics(log.With(slog.String(logger.Module.String(), "metric")), cfg.Server.Host, cfg.Server.Port.Metric, reg), // ?
		serve.GRPC(log.With(slog.String(logger.Module.String(), "gRPC")), cfg.Server.Host, cfg.Server.Port.GRPC, grpcAPI),
		serve.GRPCGateWay(log.With(slog.String(logger.Module.String(), "gRPC-Gateway")), cfg.Server.Host, cfg.Server.Port.GW, gwCfg),
		serve.HTTP(log.With(slog.String(logger.Module.String(), "files")), cfg.Server.Host, cfg.Server.Port.Files, httpAPI),
		q.Monitor,
		module.Process,
	)
}

func buildLogger(level slog.Level) *slog.Logger {
	return slog.New(
		slog.NewJSONHandler( // создается logger который будет писать в JSON?
			os.Stdout,
			&slog.HandlerOptions{ //nolint:exhaustruct
				AddSource: true, // что значит вычисляет позицию кода?
				Level:     level,
			},
		),
	)
}

func forceShutdown(ctx context.Context) { // ?
	log := logger.FromContext(ctx)
	const shutdownDelay = 15 * time.Second

	<-ctx.Done()
	time.Sleep(shutdownDelay)

	log.Error("failed to graceful shutdown")
	os.Exit(2)
}

func convertErr(err error) error {
	switch {
	case errors.Is(err, session_client.ErrNotFound):
		return app.ErrNotFound
	case errors.Is(err, session_client.ErrInvalidArgument):
		return app.ErrInvalidArgument
	default:
		return err
	}
}
