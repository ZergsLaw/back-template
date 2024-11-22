package repo

import (
	"database/sql"
	"errors"

	"github.com/Bar-Nik/back-template/cmd/twitter/internal/app"
)

func convertErr(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return app.ErrNotFound
	default:
		return err
	}
}
