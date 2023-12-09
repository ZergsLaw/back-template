package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ZergsLaw/back-template/cmd/mail/internal/adapters/queue"
	"github.com/ZergsLaw/back-template/cmd/mail/internal/adapters/repo"
	"github.com/ZergsLaw/back-template/internal/flags"
	"github.com/ZergsLaw/back-template/internal/grpchelper"
	"github.com/ZergsLaw/back-template/internal/logger"
	"github.com/ZergsLaw/back-template/internal/serve"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sipki-tech/database/connectors"
	"google.golang.org/grpc/grpclog"
	"gopkg.in/yaml.v3"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

type (
	config struct {
		Server server      `yaml:"server"`
		DB     dbConfig    `yaml:"db"`
		Queue  queueConfig `yaml:"queue"`
	}
	server struct {
		Host string `yaml:"host"`
		Port ports  `yaml:"port"`
	}
	ports struct {
		GRPC   uint16 `yaml:"grpc"`
		Metric uint16 `yaml:"metric"`
	}
	dbConfig struct {
		MigrateDir string                 `yaml:"migrate_dir"`
		Driver     string                 `yaml:"driver"`
		Cockroach  connectors.CockroachDB `yaml:"cockroach"`
	}
	queueConfig struct {
		URLs     []string `yaml:"urls"`
		Username string   `yaml:"username"`
		Password string   `yaml:"password"`
	}
)

// Обьявляем и инициализируем переменные флагов
var (
	cfgFile  = &flags.File{DefaultPath: "config.yml", MaxSize: 1024 * 1024}
	logLevel = &flags.Level{Level: slog.LevelDebug}
)

const version = "v0.1.0"

func main() {

	//Флаги с указанным именем и значением переменных
	flag.Var(cfgFile, "cfg", "path to configuration file")
	flag.Var(logLevel, "log_level", "logger level")
	flag.Parse()

	//Собираем логгер
	log := buildLogger(logLevel.Level)
	grpclog.SetLoggerV2(grpchelper.NewLogger(log))

	//Имя сервиса
	appName := filepath.Base(os.Args[0])
	//Родительский контекст содержащий наш логгер
	ctxParent := logger.NewContext(context.Background(), log.With(slog.String(logger.Version.String(), version)))
	//Контекст слушающий данный набор системных сигналов
	ctx, cancel := signal.NotifyContext(ctxParent, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	defer cancel()

	//Принудительное завершение
	go forceShutdown(ctx)

	err := start(ctx, cfgFile, appName)
	if err != nil {
		log.Error("shutdown",
			slog.String(logger.Error.String(), err.Error()),
		)
		os.Exit(2)
	}

}
func start(ctx context.Context, cfgFile io.Reader, appName string) error {
	//Инициализируем переменную конфига пустой структурой
	cfg := config{}

	//Декодируем из переменной флага в переменную конфига
	err := yaml.NewDecoder(cfgFile).Decode(cfg)
	if err != nil {
		return fmt.Errorf("yaml.NewDecoder.Decode: %w", err)
	}
	//Реестр для метрик
	reg := prometheus.NewPedanticRegistry()

	return run(ctx, cfg, reg, appName)
}
func run(ctx context.Context, cfg config, reg *prometheus.Registry, namespace string) error {
	//Получаем логгер из контекста
	log := logger.FromContext(ctx)

	//Получаем обьект Repo с нашей БД
	r, err := repo.New(ctx, reg, namespace, repo.Config{
		Cockroach:  cfg.DB.Cockroach,
		MigrateDir: cfg.DB.MigrateDir,
		Driver:     cfg.DB.Driver,
	})
	if err != nil {
		return fmt.Errorf("repo.New: %w", err)
	}

	//Отложенное закрытие подключения к БД
	defer func() {
		err := r.Close()
		if err != nil {
			log.Error("close database connection", slog.String(logger.Error.String(), err.Error()))
		}
	}()

	//оплучение обьекта очереди
	q, err := queue.New(ctx, reg, namespace, queue.Config{
		URLs:     cfg.Queue.URLs,
		Username: cfg.Queue.Username,
		Password: cfg.Queue.Password,
	})
	//Отложенное закрытие подключения к очереди
	defer func() {
		err := q.Close()
		if err != nil {
			log.Error("close queue connection", slog.String(logger.Error.String(), err.Error()))
		}
	}()

	//Запуск служб (сервер для метрик)
	err = serve.Start(
		ctx,
		serve.Metrics(log.With(slog.String(logger.Module.String(), "metric")), cfg.Server.Host, cfg.Server.Port.Metric, reg),
	)
	if err != nil {
		return fmt.Errorf("serve.Start: %w", err)
	}

	return nil
}

// Собираем логгер с JSON обработчиком
func buildLogger(level slog.Level) *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource: false,
				Level:     level,
			},
		),
	)
}
func forceShutdown(ctx context.Context) {
	//Достаем логгер из контекста
	log := logger.FromContext(ctx)
	const shutdownDelay = 15 * time.Second
	//Ожидаем завершения контекста
	<-ctx.Done()
	time.Sleep(shutdownDelay)

	log.Error("failed to graceful shutdown")
	os.Exit(2)
}
