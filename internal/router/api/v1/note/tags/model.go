package tags

import (
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
)

type (
	// AddUserNoteTagsRequest is the request DTO to add a tags to a user note
	AddUserNoteTagsRequest struct {
		NoteID int64  `json:"note_id"`
		TagID  string `json:"tag_id"`
	}

	// RemoveUserNoteTagsRequest is the request DTO to remove a tags from a user note
	RemoveUserNoteTagsRequest struct {
		NoteID int64  `json:"note_id"`
		TagID  string `json:"tag_id"`
	}

	// ListUserNoteTagsRequest is the request DTO to list user note tags
	ListUserNoteTagsRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// ListNoteTagsResponse is the response DTO to list user note tags
	ListNoteTagsResponse struct {
		NoteTags []internalpostgresmodel.UserNoteTag `json:"note_tags"`
	}

	// SyncUserNoteTagsRequest is the request DTO to sync note tags
	SyncUserNoteTagsRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// SyncUserNoteTagsResponse is the response DTO to sync user note tags
	SyncUserNoteTagsResponse struct {
		NoteTags []*internalpostgresmodel.UserNoteTagWithID `json:"note_tags"`
	}
)
