// Package repo contains wrapper for database abstraction.
package repo

import (
	"context"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sipki-tech/database"
	"github.com/sipki-tech/database/connectors"
	"github.com/sipki-tech/database/migrations"

	"github.com/ZergsLaw/back-template/cmd/user/internal/app"
)

var _ app.Repo = &Repo{}

type (
	// Config provide connection info for database.
	Config struct {
		Postgres   connectors.Raw
		MigrateDir string
		Driver     string
	}
	// Repo provided data from and to database.
	Repo struct {
		sql *database.SQL
	}
)

// New build and returns user db.
func New(ctx context.Context, reg *prometheus.Registry, namespace string, cfg Config) (*Repo, error) {
	const subsystem = "repo"
	m := database.NewMetrics(reg, namespace, subsystem, new(app.Repo))

	returnErrs := []error{ // List of app.Err… returned by Repo methods.
		app.ErrNotFound,
		app.ErrDuplicate,
		app.ErrUsernameExist,
		app.ErrEmailExist,
		app.ErrUserIDAndFileIDExist,
	}

	migrates, err := migrations.Parse(cfg.MigrateDir)
	if err != nil {
		return nil, fmt.Errorf("migrations.Parse: %w", err)
	}

	err = migrations.Run(ctx, cfg.Driver, &cfg.Postgres, migrations.Up, migrates)
	if err != nil {
		return nil, fmt.Errorf("migrations.Run: %w", err)
	}

	conn, err := database.NewSQL(ctx, cfg.Driver, database.SQLConfig{
		Metrics:    m,
		ReturnErrs: returnErrs,
	}, &cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("librepo.NewCockroach: %w", err)
	}

	return &Repo{
		sql: conn,
	}, nil
}

// Close implements io.Closer.
func (r *Repo) Close() error {
	return r.sql.Close()
}
