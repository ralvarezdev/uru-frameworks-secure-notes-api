package tags

import (
	gonethttpctx "github.com/ralvarezdev/go-net/http/context"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"net/http"
)

type (
	// controller is the structure for the API V1 note tags controller
	controller struct{}
)

// ListUserNoteTags lists user note tags
// @Summary List user note tags
// @Description List user note tags
// @Tags api v1 note
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body ListUserNoteTagsRequest true "List User UserNote Tags Request"
// @Success 200 {object} ListUserNoteTagsResponseBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note/tags [get]
func (c *controller) ListUserNoteTags(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*ListUserNoteTagsRequest)

	// List the user note tags
	userID, responseBody := Service.ListUserNoteTags(r, requestBody)

	// Log the user note tags listing
	internallogger.Api.ListUserNoteTags(
		userID,
		requestBody.NoteID,
	)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w,
		gonethttpresponse.NewJSendSuccessResponse(responseBody, http.StatusOK),
	)
}

// AddUserNoteTags adds user note tags
// @Summary Add user note tags
// @Description Adds user note tags
// @Tags api v1 note
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body AddUserNoteTagsRequest true "Add User UserNote Tags Request"
// @Success 201 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note/tags [patch]
func (c *controller) AddUserNoteTags(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*AddUserNoteTagsRequest)

	// Add the user note tags
	userID := Service.AddUserNoteTags(r, requestBody)

	// Log the user note tags addition
	internallogger.Api.AddUserNoteTags(
		userID,
		requestBody.NoteID,
	)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(nil, http.StatusCreated),
	)
}

// RemoveUserNoteTags removes user note tags
// @Summary Remove user note tags
// @Description Removes user note tags
// @Tags api v1 note
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body RemoveUserNoteTagsRequest true "Remove User UserNote Tags Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note/tags [delete]
func (c *controller) RemoveUserNoteTags(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*RemoveUserNoteTagsRequest)

	// Remove the user note tags
	userID := Service.RemoveUserNoteTags(r, requestBody)

	// Log the user note tags removal
	internallogger.Api.RemoveUserNoteTags(
		userID,
		requestBody.NoteID,
	)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(nil, http.StatusOK),
	)
}
