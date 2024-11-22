package repo_test

import (
	"testing"
	"time"

	"github.com/Bar-Nik/back-template/cmd/twitter/internal/app"
	"github.com/gofrs/uuid"
)

func TestRepo(t *testing.T) {
	t.Parallel()

	ctx, r, assert := start(t)

	post := app.Post{
		ID:          uuid.Must(uuid.NewV4()),
		POST:        string("Text"),
		UserId:      uuid.Must(uuid.NewV4()),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsPublished: bool(true),
	}

	err := r.Save(ctx, post)
	assert.NoError(err)

	err = r.Update(ctx, post)
	assert.NoError(err)

	err = r.Delete(ctx, post.ID)
	assert.NoError(err)

	err = r.PostById(ctx, post.ID)
	assert.NoError(err)

	err = r.PostsByUser(ctx, post.UserId, offset, limit)
	assert.NoError(err)
}
