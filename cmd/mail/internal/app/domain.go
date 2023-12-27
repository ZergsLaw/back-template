package app

import (
	"github.com/ZergsLaw/back-template/internal/dom"
	"github.com/gofrs/uuid"
)

type (
	UserStatus struct {
		UserID uuid.UUID
		Status dom.UserStatus
		Email  string
	}
)
