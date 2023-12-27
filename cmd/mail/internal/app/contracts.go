package app

import (
	"context"
	"github.com/ZergsLaw/back-template/internal/dom"
	"io"
)

type (
	MailSender interface {
		SendEmail(ctx context.Context, subject, content, to string, files io.Reader) error
	}
	Queue interface {
		UpUserStatus() <-chan dom.Event[UserStatus]
	}
)
