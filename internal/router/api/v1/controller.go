package v1

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"net/http"
)

type (
	// controller is the structure for the API V1 controller
	controller struct{}
)

// Ping pings the service
// @Summary Ping the service
// @Description Returns a pong response to check if the service is running
// @Tags api v1
// @Accept json
// @Produce json
// @Success 200 {object} BasicResponse
// @Router /api/v1/ping [get]
func (c *controller) Ping(w http.ResponseWriter, r *http.Request) {
	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewSuccessResponse(
			nil, http.StatusOK,
		),
	)
}

// SyncByLastSyncedAt synchronizes user notes and tags by the last synced at timestamp
// @Summary Synchronize user notes and tags by the last synced at timestamp
// @Description Synchronizes user notes and tags by the last synced at timestamp
// @Tags api v1
// @Accept json
// @Produce json
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/sync [post]
func (c *controller) SyncByLastSyncedAt(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Synchronize the list of notes and tags by last synced at timestamp
	userID, userRefreshTokenID, userTagsLastSyncedAt, userNotesLastSyncedAt, data := Service.SyncByLastSyncedAt(
		w,
		r,
	)

	// Log the list of notes and tags synchronization by last synced at timestamp
	internallogger.Api.SyncByLastSyncedAt(
		userID,
		userTagsLastSyncedAt,
		userNotesLastSyncedAt,
		userRefreshTokenID,
	)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(data, http.StatusOK),
	)
}
