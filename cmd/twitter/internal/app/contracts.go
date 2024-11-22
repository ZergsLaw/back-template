package app

import (
	"context"

	"github.com/gofrs/uuid"
)

type Repo interface {
	Save(context.Context, Post) (uuid.UUID, error)
	Update(context.Context, Post) (*Post, error)
	Delete(context.Context, uuid.UUID) error
	PostById(context.Context, uuid.UUID) (*Post, error)
	PostsByUser(context.Context, uuid.UUID, int, int) ([]Post, error)
}
