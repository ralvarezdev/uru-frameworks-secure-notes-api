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
		Title     string     `json:"title"`
		NoteTags  []string   `json:"note_tags"`
		IsPinned  *bool      `json:"is_pinned,omitempty"`
		Color     *string    `json:"color,omitempty"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}

	// NoteWithID is the response DTO for the note with ID
	NoteWithID struct {
		ID int64 `json:"id"`
		Note
	}

	// NoteVersion is the response DTO for the note version
	NoteVersion struct {
		NoteID           uint       `json:"note_id"`
		EncryptedContent string     `json:"encrypted_content"`
		CreatedAt        time.Time  `json:"created_at"`
		UpdatedAt        *time.Time `json:"updated_at,omitempty"`
	}

	// NoteVersionWithID is the response DTO for the note version with ID
	NoteVersionWithID struct {
		ID int64 `json:"id"`
		NoteVersion
	}

	// LoadedNote is the request DTO for the sync note
	LoadedNote struct {
		NoteID               uint   `json:"note_id"`
		LoadedNoteVersionsID []uint `json:"loaded_note_versions_id"`
	}

	// SyncNoteVersion is the response DTO for the sync note version
	SyncNoteVersion struct {
		NoteVersionID    uint       `json:"note_version_id"`
		EncryptedContent string     `json:"encrypted_content"`
		CreatedAt        time.Time  `json:"created_at"`
		UpdatedAt        *time.Time `json:"updated_at,omitempty"`
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
)
