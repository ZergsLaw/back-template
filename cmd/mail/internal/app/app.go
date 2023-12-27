package app

type App struct {
	mailer MailSender
	queue  Queue
}

func New(m MailSender, q Queue) *App {
	return &App{
		mailer: m,
		queue:  q,
	}
}
