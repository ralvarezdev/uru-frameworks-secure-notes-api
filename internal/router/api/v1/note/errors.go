package note

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	"net/http"
)

var (
	ErrUpdateUserNoteStarNotFound = gonethttpresponse.NewFailResponseError(
		"note_id",
		"note not found",
		nil,
		http.StatusNotFound,
	)
	ErrUpdateUserNoteTrashNotFound = gonethttpresponse.NewFailResponseError(
		"note_id",
		"note not found",
		nil,
		http.StatusNotFound,
	)
	ErrUpdateUserNoteArchiveNotFound = gonethttpresponse.NewFailResponseError(
		"note_id",
		"note not found",
		nil,
		http.StatusNotFound,
	)
	ErrUpdateUserNotePinNotFound = gonethttpresponse.NewFailResponseError(
		"note_id",
		"note not found",
		nil,
		http.StatusNotFound,
	)
	ErrGetUserNoteNotFound = gonethttpresponse.NewFailResponseError(
		"note_id",
		"note not found",
		nil,
		http.StatusNotFound,
	)
	ErrUpdateUserNoteNotFound = gonethttpresponse.NewFailResponseError(
		"note_id",
		"note not found",
		nil,
		http.StatusNotFound,
	)
	ErrDeleteUserNoteNotFound = gonethttpresponse.NewFailResponseError(
		"note_id",
		"note not found",
		nil,
		http.StatusNotFound,
	)
)
