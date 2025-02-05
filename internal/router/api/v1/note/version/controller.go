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
	body, _ := gonethttpctx.GetCtxBody(r).(*CreateUserNoteVersionRequest)

	// Create the user note version
	userID, userNoteVersionID := Service.CreateUserNoteVersion(r, body)

	// Log the user note version creation
	internallogger.Api.CreateUserNoteVersion(
		userID,
		body.NoteID,
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
	body, _ := gonethttpctx.GetCtxBody(r).(*DeleteUserNoteVersionRequest)

	// Delete the user note version
	userID := Service.DeleteUserNoteVersion(r, body)

	// Log the user note version deletion
	internallogger.Api.DeleteUserNoteVersion(userID, body.NoteVersionID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(nil, http.StatusOK),
	)
}

// GetUserNoteVersionByNoteVersionID gets a user note version by note version ID
// @Summary Get a user note version by note version ID
// @Description Gets a user note version by note version ID
// @Tags api v1 note version
// @Accept json
// @Produce json
// @Param request body GetUserNoteVersionByNoteVersionIDRequest true "Get User Note Version By Note Version ID Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note/version [get]
func (c *controller) GetUserNoteVersionByNoteVersionID(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*GetUserNoteVersionByNoteVersionIDRequest)

	// Get the user note version by note version ID
	userID, data := Service.GetUserNoteVersionByNoteVersionID(r, body)

	// Log the user note version by note version ID
	internallogger.Api.GetUserNoteVersionByNoteVersionID(
		userID,
		body.NoteVersionID,
	)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(data, http.StatusOK),
	)
}
