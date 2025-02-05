package tag

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	"net/http"
)

var (
	ErrCreateUserTagAlreadyExists = gonethttpresponse.NewFieldError(
		"name",
		"tag already exists",
		nil,
		http.StatusBadRequest,
	)
	ErrUpdateUserTagNotFound = gonethttpresponse.NewFieldError(
		"tag_id",
		"tag not found",
		nil,
		http.StatusNotFound,
	)
	ErrDeleteUserTagNotFound = gonethttpresponse.NewFieldError(
		"tag_id",
		"tag not found",
		nil,
		http.StatusNotFound,
	)
	ErrGetUserTagByTagIDNotFound = gonethttpresponse.NewFieldError(
		"tag_id",
		"tag not found",
		nil,
		http.StatusNotFound,
	)
)
