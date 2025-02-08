package tags

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
)

type (
	// ListUserTagsResponseData is the response DTO to list tags
	ListUserTagsResponseData struct {
		Tags []*internalpostgresmodel.UserTagWithID `json:"tags"`
	}

	// ListUserTagsResponseBody is the response body DTO to list tags
	ListUserTagsResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data ListUserTagsResponseData `json:"data"`
	}

	// SyncUserTagsByLastSyncedAtResponseData is the response DTO to sync tags by last synced at timestamp
	SyncUserTagsByLastSyncedAtResponseData struct {
		SyncTags []*internalpostgresmodel.UserTagWithID `json:"sync_tags"`
	}

	// SyncUserTagsByLastSyncedAtResponseBody is the response body DTO to sync tags by last synced at timestamp
	SyncUserTagsByLastSyncedAtResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data SyncUserTagsByLastSyncedAtResponseData `json:"data"`
	}
)
