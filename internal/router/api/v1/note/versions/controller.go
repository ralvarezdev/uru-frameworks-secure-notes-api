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
// @Security CookieAuth
// @Param request body ListUserNoteVersionsRequest true "List User Note Versions Request"
// @Success 200 {object} ListUserNoteVersionsResponseBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note/versions [get]
func (c *controller) ListUserNoteVersions(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*ListUserNoteVersionsRequest)

	// List the user note versions
	userID, responseBody := Service.ListUserNoteVersions(r, requestBody)

	// Log the user note versions listing
	internallogger.Api.ListUserNoteVersions(userID, requestBody.NoteID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewResponse(responseBody, http.StatusOK),
	)
}
