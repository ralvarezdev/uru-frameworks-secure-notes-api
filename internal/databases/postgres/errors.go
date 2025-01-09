package postgres

import (
	"errors"
)

var (
	ErrNilService   = errors.New("database service cannot be nil")
	ErrNilRow       = errors.New("row cannot be nil")
	ErrUserNotFound = errors.New("user not found")
)
