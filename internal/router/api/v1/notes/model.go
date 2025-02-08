package notes

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
)

type (
	// ListUserNotesResponseData is the response data DTO to list user notes
	ListUserNotesResponseData struct {
		NotesID []int64 `json:"notes"`
	}

	// ListUserNotesResponseBody is the response body DTO to list user notes
	ListUserNotesResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data ListUserNotesResponseData `json:"data"`
	}

	// SyncUserNotesByLastSyncedAtResponseData is the response data DTO to sync user notes by last synced at timestamp
	SyncUserNotesByLastSyncedAtResponseData struct {
		SyncNotes []*internalpostgresmodel.SyncUserNoteWithID `json:"sync_notes"`
	}

	// SyncUserNotesByLastSyncedAtResponseBody is the response body DTO to sync user notes by last synced at timestamp
	SyncUserNotesByLastSyncedAtResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data SyncUserNotesByLastSyncedAtResponseData `json:"data"`
	}
)
