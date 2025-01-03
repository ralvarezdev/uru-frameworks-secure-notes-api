package notes

import (
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/common"
)

// ListNotesResponse is the response DTO to list notes
type ListNotesResponse struct {
	Message string                           `json:"message"`
	Notes   []internalapiv1common.NoteWithID `json:"notes"`
}

// SyncNotesRequest is the request DTO to sync notes
type SyncNotesRequest struct {
	LoadedNotes []internalapiv1common.LoadedNote `json:"loaded_notes"`
}

// SyncNotesResponse is the response DTO to sync notes
type SyncNotesResponse struct {
	Message   string                         `json:"message"`
	SyncNotes []internalapiv1common.SyncNote `json:"sync_notes"`
}
