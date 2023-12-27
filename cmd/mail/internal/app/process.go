package app

import (
	"context"
	"github.com/ZergsLaw/back-template/internal/dom"
	"github.com/ZergsLaw/back-template/internal/logger"
	"log/slog"
)

func (a *App) Process(ctx context.Context) error {
	log := logger.FromContext(ctx)

	for {
		var err error
		select {
		case <-ctx.Done():
			return nil
		case msg := <-a.queue.UpUserStatus():
			err = a.handleUserStatus(ctx, msg)
		}
		if err != nil {
			log.Error("couldn't handle event", slog.String(logger.Error.String(), err.Error()))

			continue
		}
	}
}

func (a *App) handleUserStatus(ctx context.Context, event dom.Event[UserStatus]) error {

	subject := ""
	content := ""
	to := event.Body().Email

	//TODO: добавить логику получения файлов из репозитория с файлами

	err := a.mailer.SendEmail(ctx, subject, content, to, nil)
	if err != nil {
		return err
	}

	return nil
}
