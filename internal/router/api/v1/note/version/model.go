package version

import (
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
)

type (
	// CreateUserNoteVersionRequest is the request DTO to create a user note version
	CreateUserNoteVersionRequest struct {
		NoteID           int64  `json:"note_id"`
		EncryptedContent string `json:"encrypted_content"`
	}

	// DeleteUserNoteVersionRequest is the request DTO to delete a user note version
	DeleteUserNoteVersionRequest struct {
		NoteVersionID int64 `json:"note_version_id"`
	}

	// GetUserNoteVersionByIDRequest is the request DTO to get a user note version
	GetUserNoteVersionByIDRequest struct {
		NoteVersionID int64 `json:"note_version_id"`
	}

	// GetUserNoteVersionByIDResponse is the response DTO to get a user note version by note version ID
	GetUserNoteVersionByIDResponse struct {
		NoteVersion *internalpostgresmodel.UserNoteVersion `json:"note_version"`
	}
)
