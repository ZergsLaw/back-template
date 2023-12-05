package repo

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sipki-tech/database"
	"github.com/sipki-tech/database/connectors"
	"github.com/sipki-tech/database/migrations"
)

type (
	// Config Информация для подключения к БД
	Config struct {
		Cockroach  connectors.CockroachDB
		MigrateDir string
		Driver     string
	}
	// Repo Информация о БД, ошибках, метриках
	Repo struct {
		sql *database.SQL
	}
)

func New(ctx context.Context, reg *prometheus.Registry, namespace string, cfg Config) (*Repo, error) {
	const subsystem = "repo"
	//Инициализация метрик для БД
	m := database.NewMetrics(reg, namespace, subsystem, nil)

	//Получаем миграции по пути из конфига
	migrates, err := migrations.Parse(cfg.MigrateDir)
	if err != nil {
		return nil, fmt.Errorf("migrations.Parse: %w", err)
	}

	//Тут должны быть какие-то возвращаемые ошибки...
	returnErrs := []error{}

	//Запускаем миграции
	err = migrations.Run(ctx, cfg.Driver, &cfg.Cockroach, migrations.Up, migrates)
	if err != nil {
		return nil, fmt.Errorf("migrations.Run: %w", err)
	}

	//Создаем подключение БД используя метрики и возвращаемые ошибки, присвоив полям SQL конфига
	conn, err := database.NewSQL(ctx, cfg.Driver, database.SQLConfig{
		Metrics:    m,
		ReturnErrs: returnErrs,
	}, &cfg.Cockroach)
	if err != nil {
		return nil, fmt.Errorf("librepo.NewCockroach: %w", err)
	}

	//Возвращаем структуру Repo с БД и ошибку
	return &Repo{
		sql: conn,
	}, nil
}
func (r *Repo) Close() error {
	return r.sql.Close()
}
