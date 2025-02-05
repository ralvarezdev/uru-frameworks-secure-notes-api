package versions

import (
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
)

type (
	// ListUserNoteVersionsRequest is the request DTO to list user note versions
	ListUserNoteVersionsRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// ListUserNoteVersionsResponse is the response DTO to list user note versions
	ListUserNoteVersionsResponse struct {
		NoteVersionsID []int64 `json:"note_versions_id"`
	}

	// SyncUserNoteVersionsRequest is the request DTO to sync note versions
	SyncUserNoteVersionsRequest struct {
		NoteID              int64 `json:"note_id"`
		LatestNoteVersionID int64 `json:"latest_note_version_id"`
	}

	// SyncUserNoteVersionsResponse is the response DTO to sync user note versions
	SyncUserNoteVersionsResponse struct {
		NoteVersions []*internalpostgresmodel.NoteVersionWithID `json:"note_versions"`
	}
)
