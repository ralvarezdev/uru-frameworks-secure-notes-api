package note

import (
	gonethttpstatusresponse "github.com/ralvarezdev/go-net/http/status/response"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
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
