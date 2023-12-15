package queue

import (
	"context"
	"errors"
	"fmt"
	user_pb "github.com/ZergsLaw/back-template/api/user/v1"
	"github.com/ZergsLaw/back-template/internal/dom"
	"github.com/ZergsLaw/back-template/internal/logger"
	"github.com/ZergsLaw/back-template/internal/queue"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"strings"
)

type (
	Config struct {
		URLs     []string
		Username string
		Password string
	}

	Client struct {
		consumerName string
		m            Metrics
		queue        *queue.Queue
	}
)

func New(ctx context.Context, reg *prometheus.Registry, namespace string, cfg Config) (*Client, error) {
	const subsystem = "queue"
	m := NewMetrics(reg, namespace, subsystem, []string{})

	client, err := queue.Connect(ctx, strings.Join(cfg.URLs, ", "), namespace, cfg.Username, cfg.Password)
	if err != nil {
		return nil, fmt.Errorf("queue.ConnectNATS: %w", err)
	}

	err = client.Migrate(user_pb.Migrate)
	if err != nil {
		return nil, fmt.Errorf("client.Migrate: %w", err)
	}

	return &Client{
		consumerName: namespace,
		queue:        client,
		m:            m,
	}, nil
}

// Close Закрываем подключение к очереди
func (c *Client) Close() error {
	return c.queue.Drain()
}

// Monitor мониторим подключение к очереди
func (c *Client) Monitor(ctx context.Context) error {
	return c.queue.Monitor(ctx)
}

// Process запускаем процесс накопления событий из очереди
func (c *Client) Process(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	subjects := []string{
		user_pb.TopicAdd,
	}

	//Перебираем темы для поиска в очереди и получаем событие
	for i := range subjects {
		i := i
		group.Go(func() error {
			return c.queue.Subscribe(ctx, subjects[i], c.consumerName, c.handleEvent)
		})
	}

	return group.Wait()
}

// Обрабатываем полученное из очереди событие
func (c *Client) handleEvent(ctx context.Context, msg queue.Message) error {
	ack := make(chan dom.AcknowledgeKind)

	log := logger.FromContext(ctx)

	var err error
	//Записываем первую ошибку (отмена контекста, ошибка при обработке события) в переменную
	switch {
	case ctx.Err() != nil:
		return nil
	case msg.Subject() == user_pb.TopicAdd:
		//TODO: Добавить логику обработки события
		log.Info("Find event by topic:", slog.With(slog.String(user_pb.TopicAdd, msg.Subject())))
		return nil
	default:
		err = fmt.Errorf("%w: unknown topic %s", errors.New("invalid argument"), msg.Subject())
	}
	if err != nil {
		return err
	}

	//Отправляем очереди сообщение о получении события, либо повторый запрос на него
	//Сейчас из канала читать нечего, но после добавления логики обработки события все будет ок
	select {
	case <-ctx.Done():
		return nil
	case ackKind := <-ack:
		switch ackKind {
		case dom.AcknowledgeKindAck:
			err = msg.Ack(ctx)
		case dom.AcknowledgeKindNack:
			err = msg.Nack(ctx)
		}
		if err != nil {
			return fmt.Errorf("msg.Ack|Nack: %w", err)
		}
	}

	return nil
}
