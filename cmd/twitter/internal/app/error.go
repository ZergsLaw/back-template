package app

import (
	"errors"
)

// Errors.
var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrAccessDenied    = errors.New("access denied")
)
