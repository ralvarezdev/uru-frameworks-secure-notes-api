package notes

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"net/http"
)

type (
	// controller is the structure for the API V1 notes controller
	controller struct{}
)

// ListUserNotes lists user notes
// @Summary List user notes
// @Description List user notes
// @Tags api v1 notes
// @Accept json
// @Produce json
// @Security CookieAuth
// @Success 200 {object} ListUserNotesResponseBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/notes [get]
func (c *controller) ListUserNotes(
	w http.ResponseWriter,
	r *http.Request,
) {
	// List the user notes
	userID, responseBody := Service.ListUserNotes(r)

	// Log the user notes list
	internallogger.Api.ListUserNotes(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewResponse(responseBody, http.StatusOK),
	)
}

// SyncUserNotesByLastSyncedAt syncs user notes by last synced at timestamp
// @Summary Sync user notes by last synced at timestamp
// @Description Sync user notes by last synced at timestamp
// @Tags api v1 notes
// @Accept json
// @Produce json
// @Security CookieAuth
// @Success 200 {object} SyncUserNotesByLastSyncedAtResponseBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/notes/sync [post]
func (c *controller) SyncUserNotesByLastSyncedAt(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Synchronize the list of notes by last synced at timestamp
	userID, userRefreshTokenID, lastSyncedAt, responseBody := Service.SyncUserNotesByLastSyncedAt(
		w,
		r,
	)

	// Log the list of notes synchronization by last synced at timestamp
	internallogger.Api.SyncUserNotesByLastSyncedAt(
		userID,
		lastSyncedAt,
		userRefreshTokenID,
	)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewResponse(responseBody, http.StatusOK),
	)
}
