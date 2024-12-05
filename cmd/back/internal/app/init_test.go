package app_test

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ZergsLaw/back-template/cmd/user/internal/app"
	"github.com/ZergsLaw/back-template/internal/testhelper"
)

const pngFilePath = `./testdata/test.png`

var (
	errAny = errors.New("any error")
	origin = app.Origin{
		IP:        net.ParseIP("192.100.10.4"),
		UserAgent: "UserAgent",
	}
	ownerID = uuid.Must(uuid.NewV4())
	fileID  = uuid.Must(uuid.NewV4())
)

type mocks struct {
	hasher   *MockPasswordHash
	repo     *MockRepo
	sessions *MockSessions
	file     *MockFileStore
	auth     *MockAuth
	id       *MockID
}

func start(t *testing.T) (context.Context, *app.App, *mocks, *require.Assertions) {
	t.Helper()
	ctrl := gomock.NewController(t)

	mockRepo := NewMockRepo(ctrl)
	mockHasher := NewMockPasswordHash(ctrl)
	mockSession := NewMockSessions(ctrl)
	mockFileStore := NewMockFileStore(ctrl)
	mockAuth := NewMockAuth(ctrl)
	mockID := NewMockID(ctrl)

	module := app.New(mockRepo, mockHasher, mockAuth, mockID, mockSession, mockFileStore)

	mocks := &mocks{
		hasher:   mockHasher,
		repo:     mockRepo,
		sessions: mockSession,
		file:     mockFileStore,
		auth:     mockAuth,
		id:       mockID,
	}

	return testhelper.Context(t), module, mocks, require.New(t)
}
