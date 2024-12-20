package password_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ZergsLaw/back-template/internal/password"
)

var pass = "pass"

func TestManager_Smoke(t *testing.T) {
	t.Parallel()

	passwords := password.New()
	assert := require.New(t)
	hashPass, err := passwords.Hashing(pass)
	assert.NoError(err)
	compare := passwords.Compare(hashPass, []byte(pass))
	assert.True(compare)
}
