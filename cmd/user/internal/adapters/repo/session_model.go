package repo

import (
	"github.com/ZergsLaw/back-template/cmd/user/internal/app"

	"github.com/gofrs/uuid"
	"net"
	"time"
)

const (
	requestUpdateStatus = `StatusUpdate`
)

type (
	session struct {
		ID        uuid.UUID `db:"id"`
		Token     string    `db:"token"`
		IP        string    `db:"ip"`
		UserAgent string    `db:"user_agent"`
		UserID    uuid.UUID `db:"user_id"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)

func convertSession(s app.Session) *session {
	return &session{
		ID:        s.ID,
		Token:     s.Token.Value,
		IP:        s.Origin.IP.String(),
		UserAgent: s.Origin.UserAgent,
		UserID:    s.UserID,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func (s session) convert() *app.Session {
	return &app.Session{
		ID: s.ID,
		Origin: app.Origin{
			IP:        net.ParseIP(s.IP),
			UserAgent: s.UserAgent,
		},
		Token: app.Token{
			Value: s.Token,
		},
		UserID:    s.UserID,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}
