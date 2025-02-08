package user

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	"net/http"
)

var (
	ErrDeleteUserInvalidPassword = gonethttpresponse.NewFailResponseError(
		"password",
		"invalid password",
		nil,
		http.StatusUnauthorized,
	)
	ErrChangeUsernameAlreadyRegistered = gonethttpresponse.NewFailResponseError(
		"username",
		"username is already registered",
		nil,
		http.StatusBadRequest,
	)
)
