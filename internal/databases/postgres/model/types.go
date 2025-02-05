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

	// Tag is the response DTO for the tag
	Tag struct {
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	// TagWithID is the response DTO for the tag with ID
	TagWithID struct {
		ID int64 `json:"id"`
		Tag
	}

	// Note is the response DTO for the note
	Note struct {
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

	// NoteWithID is the response DTO for the note with ID
	NoteWithID struct {
		ID int64 `json:"id"`
		Note
	}

	// NoteVersion is the response DTO for the note version
	NoteVersion struct {
		NoteID           *int64    `json:"note_id,omitempty"`
		EncryptedContent string    `json:"encrypted_content"`
		CreatedAt        time.Time `json:"created_at"`
	}

	// NoteVersionWithID is the response DTO for the note version with ID
	NoteVersionWithID struct {
		ID int64 `json:"id"`
		NoteVersion
	}

	// SyncNote is the response DTO for the sync note
	SyncNote struct {
		Title            *string           `json:"title,omitempty"`
		NoteTags         []string          `json:"note_tags"`
		IsPinned         *bool             `json:"is_pinned,omitempty"`
		Color            *string           `json:"color,omitempty"`
		CreatedAt        *time.Time        `json:"created_at,omitempty"`
		UpdatedAt        *time.Time        `json:"updated_at,omitempty"`
		SyncNoteVersions []SyncNoteVersion `json:"sync_note_versions"`
	}

	// NoteTag is the response DTO for the note tag
	NoteTag struct {
		TagID      int64     `json:"tag_id"`
		AssignedAt time.Time `json:"assigned_at"`
	}
)
