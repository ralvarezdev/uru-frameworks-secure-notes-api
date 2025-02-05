package note

import (
	gonethttpctx "github.com/ralvarezdev/go-net/http/context"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	gonethttpstatusresponse "github.com/ralvarezdev/go-net/http/status/response"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"net/http"
)

type (
	// controller is the structure for the API V1 note controller
	controller struct{}
)

// CreateUserNote creates a user note
// @Summary Create a user note
// @Description Creates a user note
// @Tags api v1 note
// @Accept json
// @Produce json
// @Param request body CreateUserNoteRequest true "Create User UserNote Request"
// @Success 201 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note [post]
func (c *controller) CreateUserNote(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}

// UpdateUserNote updates a user note
// @Summary Update a user note
// @Description Updates a user note
// @Tags api v1 note
// @Accept json
// @Produce json
// @Param request body UpdateUserNoteRequest true "Update User UserNote Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note [put]
func (c *controller) UpdateUserNote(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}

// DeleteUserNote deletes a user note
// @Summary Delete a user note
// @Description Deletes a user note
// @Tags api v1 note
// @Accept json
// @Produce json
// @Param request body DeleteUserNoteRequest true "Delete User UserNote Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note [delete]
func (c *controller) DeleteUserNote(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}

// GetUserNote gets a user note
// @Summary Get a user note
// @Description Gets a user note
// @Tags api v1 note
// @Accept json
// @Produce json
// @Param request body GetUserNoteRequest true "Get User UserNote Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note [get]
func (c *controller) GetUserNote(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}

// UpdateUserNotePin updates a user note as pinned or unpinned
// @Summary Update a user note as pinned or unpinned
// @Description Updates a user note as pinned or unpinned
// @Tags api v1 note
// @Accept json
// @Produce json
// @Param request body UpdateUserNotePinRequest true "Update User UserNote Pin Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note/pin [put]
func (c *controller) UpdateUserNotePin(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*UpdateUserNotePinRequest)

	// Update the user note pin
	userID := Service.UpdateUserNotePin(r, body)

	// Log the user note pin update
	internallogger.Api.UpdateUserNotePin(userID, body.NoteID, body.Pin)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(nil, http.StatusOK),
	)
}

// UpdateUserNoteArchive updates a user note as archived or unarchived
// @Summary Update a user note as archived or unarchived
// @Description Updates a user note as archived or unarchived
// @Tags api v1 note
// @Accept json
// @Produce json
// @Param request body UpdateUserNoteArchiveRequest true "Update User UserNote Archive Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note/archive [put]
func (c *controller) UpdateUserNoteArchive(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*UpdateUserNoteArchiveRequest)

	// Update the user note archive
	userID := Service.UpdateUserNoteArchive(r, body)

	// Log the user note archive update
	internallogger.Api.UpdateUserNotePin(userID, body.NoteID, body.Archive)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(nil, http.StatusOK),
	)
}

// UpdateUserNoteTrash updates a user note as trashed or untrashed
// @Summary Update a user note as trashed or untrashed
// @Description Updates a user note as trashed or untrashed
// @Tags api v1 note
// @Accept json
// @Produce json
// @Param request body UpdateUserNoteTrashRequest true "Update User UserNote Trash Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note/trash [put]
func (c *controller) UpdateUserNoteTrash(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*UpdateUserNoteTrashRequest)

	// Update the user note trash
	userID := Service.UpdateUserNoteTrash(r, body)

	// Log the user note trash update
	internallogger.Api.UpdateUserNoteTrash(userID, body.NoteID, body.Trash)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(nil, http.StatusOK),
	)
}

// UpdateUserNoteStar updates a user note as starred or unstarred
// @Summary Update a user note as starred or unstarred
// @Description Updates a user note as starred or unstarred
// @Tags api v1 note
// @Accept json
// @Produce json
// @Param request body UpdateUserNoteStarRequest true "Update User UserNote Star Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note/star [put]
func (c *controller) UpdateUserNoteStar(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*UpdateUserNoteStarRequest)

	// Update the user note star
	userID := Service.UpdateUserNoteStar(r, body)

	// Log the user note star update
	internallogger.Api.UpdateUserNoteStar(userID, body.NoteID, body.Star)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(nil, http.StatusOK),
	)
}
