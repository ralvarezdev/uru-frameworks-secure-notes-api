package note

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	"net/http"
)

var (
	ErrUpdateUserNoteStarNotFound = gonethttpresponse.NewFieldError(
		"note_id",
		"Note ID not found",
		nil,
		http.StatusNotFound,
	)
	ErrUpdateUserNoteTrashNotFound = gonethttpresponse.NewFieldError(
		"note_id",
		"Note ID not found",
		nil,
		http.StatusNotFound,
	)
	ErrUpdateUserNoteArchiveNotFound = gonethttpresponse.NewFieldError(
		"note_id",
		"Note ID not found",
		nil,
		http.StatusNotFound,
	)
	ErrUpdateUserNotePinNotFound = gonethttpresponse.NewFieldError(
		"note_id",
		"Note ID not found",
		nil,
		http.StatusNotFound,
	)
)
