package tags

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"net/http"
)

type (
	// controller is the structure for the API V1 tags controller
	controller struct{}
)

// ListUserTags lists tags of the authenticated user
// @Summary List tags of the authenticated user
// @Description Lists tags of the authenticated user
// @Tags api v1 tags
// @Accept json
// @Produce json
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/tags [get]
func (c *controller) ListUserTags(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the list of tags
	userID, data := Service.ListUserTags(r)

	// Log the list of tags retrieval
	internallogger.Api.ListUserTags(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(data, http.StatusOK),
	)
}

// SyncUserTagsByLastSyncedAt synchronizes tags of the authenticated user by last synced at timestamp
// @Summary Synchronize tags of the authenticated user by last synced at timestamp
// @Description Synchronizes tags of the authenticated user by last synced at timestamp
// @Tags api v1 tags
// @Accept json
// @Produce json
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/tags/sync [post]
func (c *controller) SyncUserTagsByLastSyncedAt(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Synchronize the list of tags by last synced at timestamp
	userID, userRefreshTokenID, lastSyncedAt, data := Service.SyncUserTagsByLastSyncedAt(
		w,
		r,
	)

	// Log the list of tags synchronization by last synced at timestamp
	internallogger.Api.SyncUserTagsByLastSyncedAt(
		userID,
		lastSyncedAt,
		userRefreshTokenID,
	)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(data, http.StatusOK),
	)
}
