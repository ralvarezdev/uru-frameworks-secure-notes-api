package model

import (
	"time"
)

type (
	// UserRefreshToken is the response DTO for the user refresh token
	UserRefreshToken struct {
		IssuedAt  time.Time  `json:"issued_at"`
		ExpiresAt time.Time  `json:"expires_at"`
		RevokedAt *time.Time `json:"revoked_at,omitempty"`
		IPAddress string     `json:"ip_address"`
	}

	// UserRefreshTokenWithID is the response DTO for the user refresh token with ID
	UserRefreshTokenWithID struct {
		ID int64 `json:"id"`
		UserRefreshToken
	}

	// UserTag is the response DTO for the user tag
	UserTag struct {
		Name      string     `json:"name"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at,omitempty"`
	}

	// UserTagWithID is the response DTO for the user tag with ID
	UserTagWithID struct {
		ID int64 `json:"id"`
		UserTag
	}

	// UserNote is the response DTO for the user note
	UserNote struct {
		Title               string     `json:"title"`
		NoteTagsID          []string   `json:"note_tags_id"`
		Color               *string    `json:"color,omitempty"`
		CreatedAt           time.Time  `json:"created_at"`
		UpdatedAt           *time.Time `json:"updated_at,omitempty"`
		PinnedAt            *time.Time `json:"pinned_at,omitempty"`
		StarredAt           *time.Time `json:"starred_at,omitempty"`
		ArchivedAt          *time.Time `json:"archived_at,omitempty"`
		TrashedAt           *time.Time `json:"trashed_at,omitempty"`
		LatestNoteVersionID *int64     `json:"latest_note_version_id,omitempty"`
	}

	// UserNoteWithID is the response DTO for the user note with ID
	UserNoteWithID struct {
		ID int64 `json:"id"`
		UserNote
	}

	// UserNoteVersion is the response DTO for the user note version
	UserNoteVersion struct {
		NoteID           *int64     `json:"note_id,omitempty"`
		EncryptedContent string     `json:"encrypted_content"`
		CreatedAt        time.Time  `json:"created_at"`
		DeletedAt        *time.Time `json:"deleted_at,omitempty"`
	}

	// UserNoteVersionWithID is the response DTO for the user note version with ID
	UserNoteVersionWithID struct {
		ID int64 `json:"id"`
		UserNoteVersion
	}

	// UserNoteTag is the response DTO for the user note tag
	UserNoteTag struct {
		TagID      int64      `json:"tag_id"`
		AssignedAt time.Time  `json:"assigned_at"`
		DeletedAt  *time.Time `json:"deleted_at,omitempty"`
	}

	// UserNoteTagWithID is the response DTO for the user note tag with ID
	UserNoteTagWithID struct {
		ID int64 `json:"id"`
		UserNoteTag
	}

	// SyncUserNote is the response DTO for the sync user note
	SyncUserNote struct {
		Title            *string                  `json:"title,omitempty"`
		Color            *string                  `json:"color,omitempty"`
		CreatedAt        *time.Time               `json:"created_at,omitempty"`
		UpdatedAt        *time.Time               `json:"updated_at,omitempty"`
		PinnedAt         *time.Time               `json:"pinned_at,omitempty"`
		StarredAt        *time.Time               `json:"starred_at,omitempty"`
		ArchivedAt       *time.Time               `json:"archived_at,omitempty"`
		TrashedAt        *time.Time               `json:"trashed_at,omitempty"`
		SyncNoteVersions []*UserNoteVersionWithID `json:"sync_note_versions"`
		SyncNoteTags     []*UserNoteTag           `json:"sync_note_tags"`
	}

	// SyncUserNoteWithID is the response DTO for the sync user note with ID
	SyncUserNoteWithID struct {
		ID int64 `json:"id"`
		SyncUserNote
	}
)
