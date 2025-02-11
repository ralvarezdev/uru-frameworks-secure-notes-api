package version

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	"net/http"
)

var (
	ErrCreateUserNoteVersionNoteIDIsNotValid = gonethttpresponse.NewFailResponseError(
		"note_id",
		"note not found",
		nil,
		http.StatusNotFound,
	)
	ErrDeleteUserNoteVersionNotFound = gonethttpresponse.NewFailResponseError(
		"note_version_id",
		"note version not found",
		nil,
		http.StatusNotFound,
	)
	ErrGetUserNoteVersionByIDNotFound = gonethttpresponse.NewFailResponseError(
		"note_version_id",
		"note version not found",
		nil,
		http.StatusNotFound,
	)
)
