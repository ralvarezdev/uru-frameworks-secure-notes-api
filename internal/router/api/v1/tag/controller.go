package tag

import (
	gonethttpctx "github.com/ralvarezdev/go-net/http/context"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"net/http"
)

type (
	// controller is the structure for the API V1 tag controller
	controller struct{}
)

// CreateUserTag creates a user tag
// @Summary Create a user tag
// @Description Creates a user tag
// @Tags api v1 tag
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body CreateUserTagRequest true "Create User Tag Request"
// @Success 201 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/tag [post]
func (c *controller) CreateUserTag(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*CreateUserTagRequest)

	// Create the user tag
	userID, data := Service.CreateUserTag(r, requestBody)

	// Log the user tag creation
	internallogger.Api.CreateUserTag(userID, data)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(data, http.StatusCreated),
	)
}

// UpdateUserTag updates a user tag
// @Summary Update a user tag
// @Description Updates a user tag
// @Tags api v1 tag
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body UpdateUserTagRequest true "Update User Tag Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/tag [put]
func (c *controller) UpdateUserTag(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*UpdateUserTagRequest)

	// Update the user tag
	userID := Service.UpdateUserTag(r, requestBody)

	// Log the user tag update
	internallogger.Api.UpdateUserTag(userID, requestBody.TagID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(nil, http.StatusOK),
	)
}

// DeleteUserTag deletes a user tag
// @Summary Delete a user tag
// @Description Deletes a user tag
// @Tags api v1 tag
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body DeleteUserTagRequest true "Delete User Tag Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/tag [delete]
func (c *controller) DeleteUserTag(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*DeleteUserTagRequest)

	// Delete the user tag
	userID := Service.DeleteUserTag(r, requestBody)

	// Log the user tag deletion
	internallogger.Api.DeleteUserTag(userID, requestBody.TagID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(nil, http.StatusOK),
	)
}

// GetUserTagByID gets a user tag by tag ID
// @Summary Get a user tag by tag ID
// @Description Gets a user tag by tag ID
// @Tags api v1 tag
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body GetUserTagByIDRequest true "Get User Tag By ID Request"
// @Success 200 {object} GetUserTagByIDResponseBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/tag [get]
func (c *controller) GetUserTagByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*GetUserTagByIDRequest)

	// Get the user tag by tag ID
	userID, responseBody := Service.GetUserTagByID(r, requestBody)

	// Log the user tag retrieval
	internallogger.Api.GetUserTagByID(userID, requestBody.TagID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewResponse(responseBody, http.StatusOK),
	)
}
