package http_test

import (
	"github.com/ZergsLaw/back-template/cmd/user/internal/app"
	"net/http"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	httpapi "github.com/ZergsLaw/back-template/cmd/user/internal/api/http"

	"github.com/ZergsLaw/back-template/internal/testhelper"
)

var (
	token   = app.Token{Value: "token"}
	session = app.Session{
		ID:     uuid.Must(uuid.NewV4()),
		UserID: uuid.Must(uuid.NewV4()),
	}
	fileID = uuid.Must(uuid.NewV4())
)

func start(t *testing.T) (http.Handler, *Mockapplication, *require.Assertions) {
	t.Helper()

	ctrl := gomock.NewController(t)
	app := NewMockapplication(ctrl)

	ctx := testhelper.Context(t)
	handler := httpapi.New(ctx, app)

	return handler, app, require.New(t)
}
