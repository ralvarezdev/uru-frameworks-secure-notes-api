package tags

import (
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
)

type (
	// ListUserTagsResponse is the response DTO to list tags
	ListUserTagsResponse struct {
		Tags []*internalpostgresmodel.UserTagWithID `json:"tags"`
	}

	// SyncUserTagsResponse is the response DTO to sync tags
	SyncUserTagsResponse struct {
		SyncTags []*internalpostgresmodel.UserTagWithID `json:"sync_tags"`
	}
)
