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
