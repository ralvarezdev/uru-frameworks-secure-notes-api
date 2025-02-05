package note

import (
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/_common"
)

type (
	// CreateNoteRequest is the request DTO to create a note
	CreateNoteRequest struct {
		Title    string   `json:"title"`
		NoteTags []string `json:"note_tags"`
		Color    *string  `json:"color,omitempty"`
	}

	// UpdateNoteRequest is the request DTO to update a note
	UpdateNoteRequest struct {
		NoteID     int64    `json:"note_id"`
		Title      *string  `json:"title,omitempty"`
		NoteTagsID []string `json:"note_tags_id,omitempty"`
		Color      *string  `json:"color,omitempty"`
	}

	// DeleteNoteRequest is the request DTO to delete a note
	DeleteNoteRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// GetNoteRequest is the request DTO to get a note
	GetNoteRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// GetNoteResponse is the response DTO to get a note
	GetNoteResponse struct {
		Note internalapiv1common.Note `json:"note"`
	}

	// ListNoteTagsRequest is the request DTO to list note tags
	ListNoteTagsRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// ListNoteTagsResponse is the response DTO to list note tags
	ListNoteTagsResponse struct {
		NoteTags []internalapiv1common.TagWithID `json:"note_tags"`
	}

	// PinNoteRequest is the request DTO to pin a note
	PinNoteRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// UnpinNoteRequest is the request DTO to unpin a note
	UnpinNoteRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// ArchiveNoteRequest is the request DTO to archive a note
	ArchiveNoteRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// UnarchiveNoteRequest is the request DTO to unarchive a note
	UnarchiveNoteRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// StarNoteRequest is the request DTO to star a note
	StarNoteRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// UnstarNoteRequest is the request DTO to unstar a note
	UnstarNoteRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// TrashNoteRequest is the request DTO to trash a note
	TrashNoteRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// UntrashNoteRequest is the request DTO to untrash a note
	UntrashNoteRequest struct {
		NoteID int64 `json:"note_id"`
	}
)
