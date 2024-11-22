//go:build integration

package queue_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	user_pb "github.com/ZergsLaw/back-template/api/user/v1"
	"github.com/ZergsLaw/back-template/cmd/user/internal/app"
	"github.com/ZergsLaw/back-template/internal/dom"
	que "github.com/ZergsLaw/back-template/internal/queue"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
)

func TestQueue_Smoke(t *testing.T) {

	ctx, client, assert, cliQ := start(t)

	consureName := t.Name()

	msgId, err := uuid.NewV4()
	require.NoError(t, err)

	usrId, err := uuid.NewV4()
	require.NoError(t, err)

	user := app.User{
		ID:        usrId,
		Email:     "email@gmail.com",
		FullName:  "username",
		Name:      "Elon Musk",
		PassHash:  []byte("pass"),
		AvatarID:  uuid.Nil,
		Status:    dom.UserStatusDefault,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user2 := user

	err = client.AddUser(ctx, msgId, user)
	assert.NoError(err)

	user2.FullName = "username2"
	user2.Email = "email2@gmail.com"

	err = client.UpdateUser(ctx, msgId, user2)
	assert.NoError(err)

	err = client.DeleteUser(ctx, msgId, user2)
	assert.NoError(err)

	subscribeCtx, subscribeCtxCancel := context.WithTimeout(ctx, time.Second*2)
	t.Cleanup(subscribeCtxCancel)

	err = cliQ.Subscribe(subscribeCtx, user_pb.SubscribeToAllEvents, consureName, func(ctx context.Context, message que.Message) error {
		err := message.Ack(ctx)
		if err != nil {
			return fmt.Errorf("message.Ack: %w", err)
		}

		assert.Equal(user_pb.TopicAdd, message.Subject())
		eventAdd := user_pb.Event{}
		err = message.Unmarshal(&eventAdd)
		assert.NoError(err)
		assert.Equal(user.ID, eventAdd.GetAdd().User.Id)
		assert.Equal(user.Email, eventAdd.GetAdd().User.Email)
		assert.Equal(user2.ID, eventAdd.GetUpdate().User.Id)
		assert.Equal(user2.Email, eventAdd.GetUpdate().User.Email)
		assert.Equal(user2.ID, eventAdd.GetDelete().UserId)

		return nil
	})

	require.NoError(t, err)

}
