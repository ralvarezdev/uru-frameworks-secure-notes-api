package user

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	"net/http"
)

var (
	ErrSignUpUsernameAlreadyRegistered = gonethttpresponse.NewFieldError(
		"username",
		"username is already registered",
		http.StatusBadRequest,
		nil,
	)
	ErrSignUpEmailAlreadyRegistered = gonethttpresponse.NewFieldError(
		"email",
		"email is already registered",
		http.StatusBadRequest,
		nil,
	)
)
