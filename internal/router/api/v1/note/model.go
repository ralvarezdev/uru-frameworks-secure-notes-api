package note

import (
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
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

	// DeleteUserNoteRequest is the request DTO to delete a user note
	DeleteUserNoteRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// GetUserNoteByIDRequest is the request DTO to get a user note by ID
	GetUserNoteByIDRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// GetUserNoteByIDResponse is the response DTO to get a user note by ID
	GetUserNoteByIDResponse struct {
		Note internalpostgresmodel.UserNote `json:"note"`
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
