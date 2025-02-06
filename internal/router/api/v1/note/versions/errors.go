package versions

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	"net/http"
)

var (
	ErrListUserNoteVersionsNotFound = gonethttpresponse.NewFieldError(
		"note_id",
		"note not found",
		nil,
		http.StatusNotFound,
	)
)
