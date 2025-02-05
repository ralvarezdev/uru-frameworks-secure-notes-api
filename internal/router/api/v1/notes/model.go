package notes

import (
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/_common"
)

type (
	// ListUserNotesResponse is the response DTO to list user notes
	ListUserNotesResponse struct {
		NotesID []int64 `json:"notes"`
	}

	// SyncNotesResponse is the response DTO to sync user notes
	SyncNotesResponse struct {
		SyncNotes []internalpostgresmodel.SyncNote `json:"sync_notes"`
	}
)
