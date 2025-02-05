package note

import (
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/_common"
)

type (
	// CreateUserNoteRequest is the request DTO to create a user note
	CreateUserNoteRequest struct {
		Title            string   `json:"title"`
		NoteTagsID       []string `json:"note_tags_id,omitempty"`
		Color            *string  `json:"color,omitempty"`
		Pinned           bool     `json:"pinned"`
		Archived         bool     `json:"archived"`
		Starred          bool     `json:"starred"`
		Trashed          bool     `json:"trashed"`
		EncryptedContent string   `json:"encrypted_content"`
	}

	// UpdateUserNoteRequest is the request DTO to update a user note
	UpdateUserNoteRequest struct {
		NoteID int64   `json:"note_id"`
		Title  *string `json:"title,omitempty"`
		Color  *string `json:"color,omitempty"`
	}

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

	// DeleteUserNoteRequest is the request DTO to delete a user note
	DeleteUserNoteRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// GetUserNoteRequest is the request DTO to get a user note
	GetUserNoteRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// GetUserNoteResponse is the response DTO to get a user note
	GetUserNoteResponse struct {
		Note internalpostgresmodel.Note `json:"note"`
	}

	// ListUserNoteTagsRequest is the request DTO to list user note tags
	ListUserNoteTagsRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// ListNoteTagsResponse is the response DTO to list user note tags
	ListNoteTagsResponse struct {
		NoteTags []internalpostgresmodel.NoteTag `json:"note_tags"`
	}

	// UpdateUserNotePinRequest is the request DTO to pin/unpin a note
	UpdateUserNotePinRequest struct {
		NoteID int64 `json:"note_id"`
		Pin    bool  `json:"pin"`
	}

	// UpdateUserNoteArchiveRequest is the request DTO to archive/unarchive a note
	UpdateUserNoteArchiveRequest struct {
		NoteID  int64 `json:"note_id"`
		Archive bool  `json:"archive"`
	}

	// UpdateUserNoteStarRequest is the request DTO to star/unstar a note
	UpdateUserNoteStarRequest struct {
		NoteID int64 `json:"note_id"`
		Star   bool  `json:"star"`
	}

	// UpdateUserNoteTrashRequest is the request DTO to trash/untrash a note
	UpdateUserNoteTrashRequest struct {
		NoteID int64 `json:"note_id"`
		Trash  bool  `json:"trash"`
	}
)
