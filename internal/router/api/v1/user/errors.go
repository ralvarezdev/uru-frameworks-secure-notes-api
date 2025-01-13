package user

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
)

var (
	ErrUsernameAlreadyRegistered = gonethttpresponse.NewFieldError(
		"username",
		"username is already registered",
	)
	ErrEmailAlreadyRegistered = gonethttpresponse.NewFieldError(
		"email",
		"email is already registered",
	)
)
