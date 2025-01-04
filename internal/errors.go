package internal

import (
	"errors"
)

var (
	ErrNilRequestBody = errors.New("request body cannot be nil")
)
