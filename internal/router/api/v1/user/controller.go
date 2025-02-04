package user

import (
	gonethttpctx "github.com/ralvarezdev/go-net/http/context"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"net/http"
)

type (
	// controller is the structure for the API V1 user controller
	controller struct{}
)

// UpdateProfile updates the profile of the authenticated user
// @Summary Updates the profile of the authenticated user
// @Description Updates the profile of the authenticated user
// @Tags api v1 user
// @Accept json
// @Produce json
// @Param request body UpdateProfileRequest true "Update Profile Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/user/profile [put]
func (c *controller) UpdateProfile(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*UpdateProfileRequest)

	// Update the profile
	userID := Service.UpdateProfile(r, body)

	// Log the profile update
	internallogger.Api.UpdateProfile(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(nil, http.StatusOK),
	)
}

// GetMyProfile gets the profile of the authenticated user
// @Summary Gets the profile of the authenticated user
// @Description Gets the profile of the authenticated user
// @Tags api v1 user
// @Accept json
// @Produce json
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/user/profile [get]
func (c *controller) GetMyProfile(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the user profile
	userID, data := Service.GetMyProfile(r)

	// Log the profile retrieval
	internallogger.Api.GetMyProfile(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(data, http.StatusOK),
	)
}

// ChangeUsername changes the username of the authenticated user
// @Summary Changes the username of the authenticated user
// @Description Changes the username of the authenticated user
// @Tags api v1 user
// @Accept json
// @Produce json
// @Param request body ChangeUsernameRequest true "Change Username Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/user/username [put]
func (c *controller) ChangeUsername(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*ChangeUsernameRequest)

	// Change the username
	userID := Service.ChangeUsername(r, body)

	// Log the username change
	internallogger.Api.ChangeUsername(userID, body.Username)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(nil, http.StatusOK),
	)
}

// DeleteUser deletes the authenticated user
// @Summary Deletes the authenticated user
// @Description Deletes the authenticated user
// @Tags api v1 user
// @Accept json
// @Produce json
// @Param request body DeleteUserRequest true "Delete User Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/user [delete]
func (c *controller) DeleteUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*DeleteUserRequest)

	// Delete the user
	userID := Service.DeleteUser(r, body)

	// Log the user deletion
	internallogger.Api.DeleteUser(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(nil, http.StatusOK),
	)
}
