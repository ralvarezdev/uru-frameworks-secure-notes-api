package v1

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
)

type (
	// SyncByLastSyncedAtResponseData is the response DTO to sync user notes by last synced at timestamp
	SyncByLastSyncedAtResponseData struct {
		SyncTags  []*internalpostgresmodel.UserTagWithID      `json:"sync_tags"`
		SyncNotes []*internalpostgresmodel.SyncUserNoteWithID `json:"sync_notes"`
	}

	// SyncByLastSyncedAtResponseBody is the response body to sync user notes by last synced at timestamp
	SyncByLastSyncedAtResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data SyncByLastSyncedAtResponseData `json:"data"`
	}
)
