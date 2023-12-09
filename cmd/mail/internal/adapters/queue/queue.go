package queue

import (
	"context"
	"fmt"
	user_pb "github.com/ZergsLaw/back-template/api/user/v1"
	"github.com/ZergsLaw/back-template/internal/queue"
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

type (
	Config struct {
		URLs     []string
		Username string
		Password string
	}

	Client struct {
		m     Metrics
		queue *queue.Queue
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
		queue: client,
		m:     m,
	}, nil
}
func (c *Client) Close() error {
	return c.queue.Drain()
}
