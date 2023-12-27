package mailer

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
	"io"
	"net/smtp"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type GmailConfig struct {
	GmailSenderName          string
	GmailSenderEmailAddress  string
	GmailSenderEmailPassword string
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewSenderGmail(cfg GmailConfig) *GmailSender {
	return &GmailSender{
		name:              cfg.GmailSenderName,
		fromEmailAddress:  cfg.GmailSenderEmailAddress,
		fromEmailPassword: cfg.GmailSenderEmailPassword,
	}
}
func (g *GmailSender) SendEmail(ctx context.Context, subject, content string, to string, files io.Reader) error {

	e := email.NewEmail()
	e.Subject = subject
	e.From = fmt.Sprintf("%s <%s>", g.name, g.fromEmailAddress)
	e.HTML = []byte(content)
	e.To = []string{to}

	//TODO:одбавить логику присваивания файлов полю e.Attach

	smtpAuth := smtp.PlainAuth("", g.fromEmailAddress, g.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}
