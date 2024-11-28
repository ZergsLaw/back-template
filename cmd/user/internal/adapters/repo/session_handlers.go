package repo

import (
	"context"
	"fmt"
	"github.com/ZergsLaw/back-template/cmd/user/internal/app"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

// SessionSave for implements app.Repo.
func (r *Repo) SessionSave(ctx context.Context, session app.Session) error {
	return r.sql.NoTx(func(db *sqlx.DB) error {
		newSession := convertSession(session)

		const query = `
		insert into 
		sessions 
		    (id, token, ip, user_agent, user_id) 
		values 
			($1, $2, $3, $4, $5)`

		_, err := db.ExecContext(ctx, query, newSession.ID, newSession.Token, newSession.IP, newSession.UserAgent, newSession.UserID)
		if err != nil {
			return fmt.Errorf("db.ExecContext: %w", err)
		}

		return nil
	})
}

// SessionByID for implements app.Repo.
func (r *Repo) SessionByID(ctx context.Context, sessionID uuid.UUID) (s *app.Session, err error) {
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `select * from sessions where id = $1`

		res := session{}
		err = db.GetContext(ctx, &res, query, sessionID)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}

		s = res.convert()

		return nil
	})
	if err != nil {
		return nil, err
	}

	return s, nil
}

// SessionDelete for implements app.Repo.
func (r *Repo) SessionDelete(ctx context.Context, sessionID uuid.UUID) error {
	return r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `
		delete
		from sessions
		where id = $1 returning *`

		err := db.GetContext(ctx, &session{}, query, sessionID)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}

		return nil
	})
}
