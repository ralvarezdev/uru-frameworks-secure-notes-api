package notes

import (
	gonethttpstatusresponse "github.com/ralvarezdev/go-net/http/status/response"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
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
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}

// SyncUserNotes syncs user notes
// @Summary Sync user notes
// @Description Sync user notes
// @Tags api v1 notes
// @Accept json
// @Produce json
// @Param request body SyncUserNotesRequest true "Sync User Notes Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/notes/sync [post]
func (c *controller) SyncUserNotes(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}
