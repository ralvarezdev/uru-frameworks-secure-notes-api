package version

import (
	gonethttpctx "github.com/ralvarezdev/go-net/http/context"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"net/http"
)

type (
	// controller is the structure for the API V1 version controller
	controller struct{}
)

// CreateUserNoteVersion creates a user note version
// @Summary Create a user note version
// @Description Creates a user note version
// @Tags api v1 note version
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body CreateUserNoteVersionRequest true "Create User Note Version Request"
// @Success 201 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note/version [post]
func (c *controller) CreateUserNoteVersion(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*CreateUserNoteVersionRequest)

	// Create the user note version
	userID, userNoteVersionID := Service.CreateUserNoteVersion(r, requestBody)

	// Log the user note version creation
	internallogger.Api.CreateUserNoteVersion(
		userID,
		requestBody.NoteID,
		userNoteVersionID,
	)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(nil, http.StatusCreated),
	)
}

// DeleteUserNoteVersion deletes a user note version
// @Summary Delete a user note version
// @Description Deletes a user note version
// @Tags api v1 note version
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body DeleteUserNoteVersionRequest true "Delete User Note Version Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note/version [delete]
func (c *controller) DeleteUserNoteVersion(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*DeleteUserNoteVersionRequest)

	// Delete the user note version
	userID := Service.DeleteUserNoteVersion(r, requestBody)

	// Log the user note version deletion
	internallogger.Api.DeleteUserNoteVersion(userID, requestBody.NoteVersionID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(nil, http.StatusOK),
	)
}

// GetUserNoteVersionByID gets a user note version by note version ID
// @Summary Get a user note version by note version ID
// @Description Gets a user note version by note version ID
// @Tags api v1 note version
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body GetUserNoteVersionByIDRequest true "Get User Note Version By ID Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note/version [get]
func (c *controller) GetUserNoteVersionByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*GetUserNoteVersionByIDRequest)

	// Get the user note version by note version ID
	userID, responseBody := Service.GetUserNoteVersionByNoteVersionID(
		r,
		requestBody,
	)

	// Log the user note version by note version ID
	internallogger.Api.GetUserNoteVersionByID(
		userID,
		requestBody.NoteVersionID,
	)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewResponse(responseBody, http.StatusOK),
	)
}
