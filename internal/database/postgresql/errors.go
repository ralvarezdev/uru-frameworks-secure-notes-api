package postgresql

import (
	"errors"
)

var (
	ErrNilService    = errors.New("database service cannot be nil")
	ErrNilDatabase   = errors.New("database cannot be nil")
	ErrNilConnection = errors.New("database connection cannot be nil")
)
