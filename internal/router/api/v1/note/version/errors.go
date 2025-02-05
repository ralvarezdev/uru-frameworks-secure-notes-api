package version

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	"net/http"
)

var (
	ErrCreateUserNoteVersionNoteIDIsNotValid = gonethttpresponse.NewFieldError(
		"note_id",
		"note not found",
		nil,
		http.StatusNotFound,
	)
	ErrDeleteUserNoteVersionNotFound = gonethttpresponse.NewFieldError(
		"note_version_id",
		"note version not found",
		nil,
		http.StatusNotFound,
	)
	ErrGetUserNoteVersionByNoteVersionIDNotFound = gonethttpresponse.NewFieldError(
		"note_version_id",
		"note version not found",
		nil,
		http.StatusNotFound,
	)
)
