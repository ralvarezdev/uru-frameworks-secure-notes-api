package internal

import (
	"errors"
)

var (
	// InDevelopment is the response when a request is made to an endpoint that is not implemented yet
	InDevelopment = errors.New("this endpoint is not implemented yet")
)
