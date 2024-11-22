package app

import (
	"time"

	"github.com/gofrs/uuid"
)

type (
	Post struct {
		ID          uuid.UUID
		POST        string
		CreatedAt   time.Time
		UpdatedAt   time.Time
		IsPublished bool
		UserId      uuid.UUID
	}
)
