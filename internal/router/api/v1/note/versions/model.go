package versions

import (
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/_common"
)

// ListNoteVersionsResponse is the response DTO to list note versions
type ListNoteVersionsResponse struct {
	NoteVersionsID []string `json:"note_versions_id"`
}

// SyncNoteVersionsRequest is the request DTO to sync note versions
type SyncNoteVersionsRequest struct {
	NoteID               uint   `json:"note_id"`
	LoadedNoteVersionsID []uint `json:"loaded_note_versions_id"`
}

// SyncNoteVersionsResponse is the response DTO to sync note versions
type SyncNoteVersionsResponse struct {
	SyncNoteVersions []internalapiv1common.SyncNoteVersion `json:"sync_note_versions"`
}
