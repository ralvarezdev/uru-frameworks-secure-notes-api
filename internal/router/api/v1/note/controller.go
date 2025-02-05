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

// CreateNote creates a note
// @Summary Create a note
// @Description Creates a note
// @Tags api v1 note
// @Accept json
// @Produce json
// @Param request body CreateNoteRequest true "Create Note Request"
// @Success 201 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note [post]
func (c *controller) CreateNote(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}

// UpdateNote updates a note
// @Summary Update a note
// @Description Updates a note
// @Tags api v1 note
// @Accept json
// @Produce json
// @Param request body UpdateNoteRequest true "Update Note Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note [put]
func (c *controller) UpdateNote(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}

// DeleteNote deletes a note
// @Summary Delete a note
// @Description Deletes a note
// @Tags api v1 note
// @Accept json
// @Produce json
// @Param request body DeleteNoteRequest true "Delete Note Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note [delete]
func (c *controller) DeleteNote(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}

// GetNote gets a note
// @Summary Get a note
// @Description Gets a note
// @Tags api v1 note
// @Accept json
// @Produce json
// @Param request body GetNoteRequest true "Get Note Request"
// @Success 200 {object} GetNoteResponse
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note [get]
func (c *controller) GetNote(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}

// ListNoteTags lists note tags
// @Summary List note tags
// @Description List note tags
// @Tags api v1 note
// @Accept json
// @Produce json
// @Param request body ListNoteTagsRequest true "List Note Tags Request"
// @Success 200 {object} ListNoteTagsResponse
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/note/tags [get]
func (c *controller) ListNoteTags(
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
// @Param request body UpdateUserNotePinRequest true "Update User Note Pin Request"
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
// @Param request body UpdateUserNoteArchiveRequest true "Update User Note Archive Request"
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
// @Param request body UpdateUserNoteTrashRequest true "Update User Note Trash Request"
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
// @Param request body UpdateUserNoteStarRequest true "Update User Note Star Request"
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
