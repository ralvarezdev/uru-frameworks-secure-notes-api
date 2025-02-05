package versions

import (
	gonethttpctx "github.com/ralvarezdev/go-net/http/context"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"net/http"
)

type (
	// controller is the structure for the API V1 versions controller
	controller struct{}
)

// ListUserNoteVersions lists user note versions
// @Summary List user note versions
// @Description List user note versions
// @Tags api v1 note versions
// @Accept json
// @Produce json
// @Param request body ListUserNoteVersionsRequest true "List User UserNote Versions Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note/versions [get]
func (c *controller) ListUserNoteVersions(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*ListUserNoteVersionsRequest)

	// List the user note versions
	userID, data := Service.ListUserNoteVersions(r, body)

	// Log the user note versions listing
	internallogger.Api.ListUserNoteVersions(userID, body.NoteID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(data, http.StatusOK),
	)
}
