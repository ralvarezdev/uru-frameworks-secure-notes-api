package v1

import (
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
)

type (
	// SyncResponse is the response DTO to sync user notes
	SyncResponse struct {
		SyncTags  []*internalpostgresmodel.UserTagWithID `json:"sync_tags"`
		SyncNotes []*internalpostgresmodel.SyncUserNote  `json:"sync_notes"`
	}
)
