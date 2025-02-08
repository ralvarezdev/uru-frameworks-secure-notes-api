package versions

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	"net/http"
)

var (
	ErrListUserNoteVersionsNotFound = gonethttpresponse.NewFailResponseError(
		"note_id",
		"note not found",
		nil,
		http.StatusNotFound,
	)
)
