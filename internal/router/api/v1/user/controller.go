package user

import (
	gonethttpstatusresponse "github.com/ralvarezdev/go-net/http/status/response"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
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
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
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
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
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
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}

// ChangeEmail changes the email of the authenticated user
// @Summary Changes the email of the authenticated user
// @Description Changes the email of the authenticated user
// @Tags api v1 user
// @Accept json
// @Produce json
// @Param request body ChangeEmailRequest true "Change Email Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/user/email [put]
func (c *controller) ChangeEmail(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}

// SendEmailVerificationToken sends an email verification token to the authenticated user
// @Summary Sends an email verification token to the authenticated user
// @Description Sends an email verification token to the authenticated user
// @Tags api v1 user
// @Accept json
// @Produce json
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/user/email/send-verification [post]
func (c *controller) SendEmailVerificationToken(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}

// VerifyEmail verifies the email of the authenticated user
// @Summary Verifies the email of the authenticated user
// @Description Verifies the email of the authenticated user
// @Tags api v1 user
// @Accept json
// @Produce json
// @Param request body VerifyEmailRequest true "Verify Email Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/user/email/verify [post]
func (c *controller) VerifyEmail(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}

// ChangePhoneNumber changes the phone number of the authenticated user
// @Summary Changes the phone number of the authenticated user
// @Description Changes the phone number of the authenticated user
// @Tags api v1 user
// @Accept json
// @Produce json
// @Param request body ChangePhoneNumberRequest true "Change Phone Number Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/user/phone-number [put]
func (c *controller) ChangePhoneNumber(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}

// SendPhoneNumberVerificationCode sends a phone number verification code to the authenticated user
// @Summary Sends a phone number verification code to the authenticated user
// @Description Sends a phone number verification code to the authenticated user
// @Tags api v1 user
// @Accept json
// @Produce json
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/user/phone-number/send-verification [post]
func (c *controller) SendPhoneNumberVerificationCode(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}

// VerifyPhoneNumber verifies the phone number of the authenticated user
// @Summary Verifies the phone number of the authenticated user
// @Description Verifies the phone number of the authenticated user
// @Tags api v1 user
// @Accept json
// @Produce json
// @Param request body VerifyPhoneNumberRequest true "Verify Phone Number Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/user/phone-number/verify [post]
func (c *controller) VerifyPhoneNumber(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
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
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}
