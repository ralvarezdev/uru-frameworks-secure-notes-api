package notes

import (
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
)

type (
	// ListUserNotesResponse is the response DTO to list user notes
	ListUserNotesResponse struct {
		NotesID []int64 `json:"notes"`
	}

	// SyncUserNotesResponse is the response DTO to sync user notes
	SyncUserNotesResponse struct {
		SyncNotes []*internalpostgresmodel.SyncUserNoteWithID `json:"sync_notes"`
	}
)
