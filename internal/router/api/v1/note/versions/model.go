package versions

import (
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/common"
)

// ListNoteVersionsResponse is the response DTO to list note versions
type ListNoteVersionsResponse struct {
	Message        string   `json:"message"`
	NoteVersionsID []string `json:"note_versions_id"`
}

// ListLastNoteVersionsWithContentResponse is the response DTO to list last note versions with their content
type ListLastNoteVersionsWithContentResponse struct {
	Message      string                                  `json:"message"`
	NoteVersions []internalapiv1common.NoteVersionWithID `json:"note_versions"`
}

// SyncNoteVersionsRequest is the request DTO to sync note versions
type SyncNoteVersionsRequest struct {
	NoteID               uint   `json:"note_id"`
	LoadedNoteVersionsID []uint `json:"loaded_note_versions_id"`
}

// SyncNoteVersionsResponse is the response DTO to sync note versions
type SyncNoteVersionsResponse struct {
	Message          string                                `json:"message"`
	SyncNoteVersions []internalapiv1common.SyncNoteVersion `json:"sync_note_versions"`
}
