package tag

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	"net/http"
)

var (
	ErrCreateUserTagAlreadyExists = gonethttpresponse.NewFailResponseError(
		"name",
		"tag already exists",
		nil,
		http.StatusBadRequest,
	)
	ErrUpdateUserTagNotFound = gonethttpresponse.NewFailResponseError(
		"tag_id",
		"tag not found",
		nil,
		http.StatusNotFound,
	)
	ErrDeleteUserTagNotFound = gonethttpresponse.NewFailResponseError(
		"tag_id",
		"tag not found",
		nil,
		http.StatusNotFound,
	)
	ErrGetUserTagByTagIDNotFound = gonethttpresponse.NewFailResponseError(
		"tag_id",
		"tag not found",
		nil,
		http.StatusNotFound,
	)
)
