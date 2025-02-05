package version

import (
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/_common"
)

type (
	// CreateNoteVersionRequest is the request DTO to create a note version
	CreateNoteVersionRequest struct {
		NoteID           int64  `json:"note_id"`
		EncryptedContent string `json:"encrypted_content"`
	}

	// UpdateNoteVersionRequest is the request DTO to update a note version
	UpdateNoteVersionRequest struct {
		NoteVersionID    int64   `json:"note_version_id"`
		EncryptedContent *string `json:"encrypted_content,omitempty"`
	}

	// DeleteNoteVersionRequest is the request DTO to delete a note version
	DeleteNoteVersionRequest struct {
		NoteVersionID int64 `json:"note_version_id"`
	}

	// GetNoteVersionRequest is the request DTO to get a note version
	GetNoteVersionRequest struct {
		NoteVersionID int64 `json:"note_version_id"`
	}

	// GetNoteVersionResponse is the response DTO to get a note version
	GetNoteVersionResponse struct {
		NoteVersion internalapiv1common.NoteVersion `json:"note_version"`
	}
)
