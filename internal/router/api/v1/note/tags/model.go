package tags

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
)

type (
	// AddUserNoteTagsRequest is the request DTO to add a tags to a user note
	AddUserNoteTagsRequest struct {
		NoteID int64  `json:"note_id"`
		TagsID string `json:"tags_id"`
	}

	// RemoveUserNoteTagsRequest is the request DTO to remove a tags from a user note
	RemoveUserNoteTagsRequest struct {
		NoteID int64  `json:"note_id"`
		TagsID string `json:"tags_id"`
	}

	// ListUserNoteTagsRequest is the request DTO to list user note tags
	ListUserNoteTagsRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// ListUserNoteTagsResponseData is the response data DTO to list user note tags
	ListUserNoteTagsResponseData struct {
		NoteTags []*internalpostgresmodel.UserNoteTagWithID `json:"note_tags"`
	}

	// ListUserNoteTagsResponseBody is the response body DTO to list user note tags
	ListUserNoteTagsResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data ListUserNoteTagsResponseData `json:"data"`
	}
)
