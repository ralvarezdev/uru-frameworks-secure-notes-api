package notes

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	gonethttpstatusresponse "github.com/ralvarezdev/go-net/http/status/response"
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
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/notes [get]
func (c *controller) ListUserNotes(
	w http.ResponseWriter,
	r *http.Request,
) {
	// List the user notes
	userID, data := Service.ListUserNotes(r)

	// Log the user notes list
	internallogger.Api.ListUserNotes(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(data, http.StatusOK),
	)
}

// SyncUserNotesByLastSyncedAt syncs user notes by last synced at timestamp
// @Summary Sync user notes by last synced at timestamp
// @Description Sync user notes by last synced at timestamp
// @Tags api v1 notes
// @Accept json
// @Produce json
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/notes/sync [post]
func (c *controller) SyncUserNotesByLastSyncedAt(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}
