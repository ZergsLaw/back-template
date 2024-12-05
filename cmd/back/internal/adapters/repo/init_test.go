//go:build integration

package repo_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sipki-tech/database/connectors"
	"github.com/stretchr/testify/require"

	"github.com/ZergsLaw/back-template/cmd/user/internal/adapters/repo"
	"github.com/ZergsLaw/back-template/internal/logger"
	"github.com/ZergsLaw/back-template/internal/testhelper"
)

const (
	migrateDir = `../../../migrate`
)

func start(t *testing.T) (context.Context, *repo.Repo, *require.Assertions) {
	t.Helper()
	ctx := testhelper.Context(t)
	assert := require.New(t)

	namespace := testhelper.Namespace(t)

	postgresDSN := testhelper.Postgres(
		ctx,
		t,
		assert,
	)

	reg := prometheus.NewPedanticRegistry()
	r, err := repo.New(ctx, reg, namespace, repo.Config{
		Postgres: connectors.Raw{
			Query: postgresDSN,
		},
		MigrateDir: migrateDir,
		Driver:     "postgres",
	})
	assert.NoError(err)
	t.Cleanup(func() {
		assert.NoError(r.Close())
	})

	log := slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{ //nolint:exhaustruct
				AddSource: true,
				Level:     slog.LevelDebug,
			},
		),
	)

	return logger.NewContext(ctx, log), r, assert
}
