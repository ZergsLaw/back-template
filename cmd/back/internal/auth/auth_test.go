package auth_test

import (
	"github.com/ZergsLaw/back-template/cmd/user/internal/auth"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
)

func TestAuth_TokenAndSubject(t *testing.T) {
	t.Parallel()

	assert := require.New(t)
	a := auth.New("super-duper-secret-key-qwertyuio")

	subject := uuid.Must(uuid.NewV4())
	appToken, err := a.Token(subject)
	assert.NoError(err)
	assert.NotNil(appToken)

	res, err := a.Subject(appToken.Value)
	assert.NoError(err)
	assert.Equal(subject, res)
}
