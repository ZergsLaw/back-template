package session

import (
	"context"
	"github.com/ZergsLaw/back-template/cmd/user/internal/app"
)

type ctxMarker struct{}

// NewContext returns context with slog.Logger.
func NewContext(ctx context.Context, session *app.Session) context.Context {
	return context.WithValue(ctx, ctxMarker{}, session)
}

// FromContext returns slog.Logger from context.
func FromContext(ctx context.Context) *app.Session {
	l, ok := ctx.Value(ctxMarker{}).(*app.Session)
	if !ok {
		return nil
	}

	return l
}
