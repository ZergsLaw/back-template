package repo

import (
	"time"

	"github.com/Bar-Nik/back-template/cmd/twitter/internal/app"
	"github.com/gofrs/uuid"
)

type post struct {
	ID          uuid.UUID `json:"id"`
	POST        string    `json:"post"`
	CreatedAt   time.Time `json:"created at"`
	UpdatedAt   time.Time `json:"updated at "`
	IsPublished bool      `json:"is published"`
	UserId      uuid.UUID `json:"user id"`
}

func convert(p app.Post) *post {
	return &post{
		ID:          p.ID,
		POST:        p.POST,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		IsPublished: p.IsPublished,
		UserId:      p.UserId,
	}
}
func (p post) convert() *app.Post {
	return &app.Post{
		ID:          p.ID,
		POST:        p.POST,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		IsPublished: p.IsPublished,
		UserId:      p.UserId,
	}
}
