package user

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	"net/http"
)

var (
	ErrDeleteUserInvalidPassword = gonethttpresponse.NewFieldError(
		"password",
		"invalid password",
		nil,
		http.StatusUnauthorized,
	)
	ErrChangeUsernameAlreadyRegistered = gonethttpresponse.NewFieldError(
		"username",
		"username is already registered",
		nil,
		http.StatusBadRequest,
	)
)
