package note

import (
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/common"
)

// CreateNoteRequest is the request DTO to create a note
type CreateNoteRequest struct {
	Title    string   `json:"title"`
	NoteTags []string `json:"note_tags"`
	Color    *string  `json:"color,omitempty"`
}

// UpdateNoteRequest is the request DTO to update a note
type UpdateNoteRequest struct {
	NoteID     uint     `json:"note_id"`
	IsPinned   *bool    `json:"is_pinned,omitempty"`
	Title      *string  `json:"title,omitempty"`
	NoteTagsID []string `json:"note_tags_id,omitempty"`
	Color      *string  `json:"color,omitempty"`
}

// DeleteNoteRequest is the request DTO to delete a note
type DeleteNoteRequest struct {
	NoteID uint `json:"note_id"`
}

// GetNoteRequest is the request DTO to get a note
type GetNoteRequest struct {
	NoteID uint `json:"note_id"`
}

// GetNoteResponse is the response DTO to get a note
type GetNoteResponse struct {
	Message string                   `json:"message"`
	Note    internalapiv1common.Note `json:"note"`
}

// ListNoteTagsResponse is the response DTO to list note tags
type ListNoteTagsResponse struct {
	Message  string                          `json:"message"`
	NoteTags []internalapiv1common.TagWithID `json:"note_tags"`
}
