package v1

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalrouterapiv1notes "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/notes"
	internalrouterapiv1tags "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/tags"
	"net/http"
	"time"
)

type (
	// service is the structure for the API V1 service
	service struct{}
)

// SyncByLastSyncedAt synchronizes the notes and tags of the authenticated user by last synced at timestamp
func (s *service) SyncByLastSyncedAt(
	w http.ResponseWriter,
	r *http.Request,
) (
	int64,
	int64,
	*time.Time,
	*time.Time,
	*SyncByLastSyncedAtResponseBody,
) {
	// Synchronize the list of tags by last synced at timestamp
	userID, userRefreshTokenID, userTagsLastSyncedAt, syncUserTags := internalrouterapiv1tags.Service.SyncUserTagsByLastSyncedAt(
		w,
		r,
	)

	// Synchronize the list of notes by last synced at timestamp
	_, _, userNotesLastSyncedAt, syncUserNotes := internalrouterapiv1notes.Service.SyncUserNotesByLastSyncedAt(
		w,
		r,
	)

	return userID, userRefreshTokenID, userTagsLastSyncedAt, userNotesLastSyncedAt, &SyncByLastSyncedAtResponseBody{
		BaseJSendSuccessBody: *gonethttpresponse.NewBaseJSendSuccessBody(),
		Data: SyncByLastSyncedAtResponseData{
			SyncTags:  syncUserTags.Data.SyncTags,
			SyncNotes: syncUserNotes.Data.SyncNotes,
		},
	}
}
