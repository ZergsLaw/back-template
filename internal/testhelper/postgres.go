package testhelper

import (
	"context"
	"fmt"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/sipki-tech/database"
	"github.com/stretchr/testify/require"
)

// Default values for making postgres test container.
const (
	PostgresDBImage           = `postgres`
	PostgresDBVersion         = `16.1`
	PostgresDBDefaultEndpoint = ``
)

type postgresConnector struct {
	dsn string
}

// DSN implements database.Connector.
func (p postgresConnector) DSN() (string, error) { return p.dsn, nil }

// Postgres build and run test container with database.
//
// Notice:
//   - Starts postgres container.
func Postgres(
	ctx context.Context,
	t *testing.T,
	assert *require.Assertions,
) string {
	t.Helper()

	opt := &dockertest.RunOptions{
		Repository: PostgresDBImage,
		Tag:        PostgresDBVersion,
		Env: []string{
			"POSTGRES_PASSWORD=postgres",
			"POSTGRES_USER=postgres",
			"POSTGRES_DB=postgres",
		},
		//
		//Platform: "darwin/arm64",
	}

	servicePort := 0
	err := runContainer(ctx, t, assert, opt, "5432/tcp", PostgresDBDefaultEndpoint, func(port int) (err error) {
		defer func() {
			if err != nil {
				t.Logf("connection problem: %s", err)
			}
		}()

		cfg := postgresConnector{
			dsn: fmt.Sprintf("postgresql://postgres:postgres@localhost:%d/postgres?sslmode=disable", port),
		}

		db, err := database.NewSQL(ctx, "postgres", database.SQLConfig{}, cfg)
		if err != nil {
			return fmt.Errorf("database.NewSQL: %w", err)
		}

		t.Cleanup(func() { assert.NoError(db.Close()) })

		servicePort = port

		return nil
	})
	assert.NoError(err)

	return fmt.Sprintf("postgresql://postgres:postgres@localhost:%d/postgres?sslmode=disable", servicePort)
}
