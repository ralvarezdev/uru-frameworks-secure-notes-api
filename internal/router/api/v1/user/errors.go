package user

import (
	"errors"
)

var (
	ErrUsernameAlreadyRegistered = errors.New("username is already registered")
	ErrEmailAlreadyRegistered    = errors.New("email is already registered")
)
