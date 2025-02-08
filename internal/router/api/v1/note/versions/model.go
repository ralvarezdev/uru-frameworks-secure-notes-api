package versions

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
)

type (
	// ListUserNoteVersionsRequest is the request DTO to list user note versions
	ListUserNoteVersionsRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// ListUserNoteVersionsResponseData is the response data DTO to list user note versions
	ListUserNoteVersionsResponseData struct {
		NoteVersionsID []int64 `json:"note_versions_id"`
	}

	// ListUserNoteVersionsResponseBody is the response body DTO to list user note versions
	ListUserNoteVersionsResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data ListUserNoteVersionsResponseData `json:"data"`
	}
)
