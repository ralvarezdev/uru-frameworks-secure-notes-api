package auth

import (
	gonethttpctx "github.com/ralvarezdev/go-net/http/context"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	gonethttpstatusresponse "github.com/ralvarezdev/go-net/http/status/response"
	gostringsconvert "github.com/ralvarezdev/go-strings/convert"
	internalcookie "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/cookie"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"net/http"
)

type (
	// controller is the structure for the API V1 auth controller
	controller struct{}
)

// GetTokenWildcard gets the token wildcard
func (c *controller) GetTokenWildcard(
	w http.ResponseWriter,
	r *http.Request,
	dest *string,
) bool {
	return internalhandler.Handler.ParseWildcard(
		w, r, "token", dest,
		gostringsconvert.ToString,
	)
}

// SignUp signs up a new user
// @Summary Sign up a new user
// @Description Creates a new user account with the provided details
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param request body SignUpRequest true "Sign Up Request"
// @Success 201 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/signup [post]
func (c *controller) SignUp(w http.ResponseWriter, r *http.Request) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*SignUpRequest)

	// Sign up the user
	userID := Service.SignUp(body)

	// Log the user sign up
	internallogger.Api.SignUp(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil, http.StatusCreated,
		),
	)
}

// LogIn logs in a user
// @Summary Log in a user
// @Description Authenticates a user and returns a seed token
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param request body LogInRequest true "Log In Request"
// @Success 201 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailResponse
// @Failure 401 {object} gonethttpresponse.JSendFailResponse
// @Failure 500 {object} gonethttpresponse.JSendErrorResponse
// @Router /api/v1/auth/login [post]
func (c *controller) LogIn(w http.ResponseWriter, r *http.Request) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*LogInRequest)

	// Log in the user
	userID := Service.LogIn(w, r, body)

	// Log the successful login
	internallogger.Api.LogIn(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusCreated,
		),
	)
}

// ListRefreshTokens gets a user's refresh tokens
// @Summary Get a user's refresh tokens
// @Description Gets a user's refresh tokens
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/refresh-tokens [get]
func (c *controller) ListRefreshTokens(w http.ResponseWriter, r *http.Request) {
	// Get the user's refresh tokens
	userID, data := Service.ListRefreshTokens(r)

	// Log the successful fetch of the user's refresh tokens
	internallogger.Api.ListRefreshTokens(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			data,
			http.StatusOK,
		),
	)
}

// GetRefreshToken gets a user's refresh token
// @Summary Get a user's refresh token
// @Description Gets a user's refresh token
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param request body GetRefreshTokenRequest true "Get Refresh Token Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/refresh-token [get]
func (c *controller) GetRefreshToken(w http.ResponseWriter, r *http.Request) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*GetRefreshTokenRequest)

	// Get the user's refresh token by ID
	userID, data := Service.GetRefreshToken(
		r,
		body.RefreshTokenID,
	)

	// Log the successful fetch of the user's refresh token
	internallogger.Api.GetRefreshToken(userID, body.RefreshTokenID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			data,
			http.StatusOK,
		),
	)

}

// RevokeRefreshToken revokes a user's refresh token
// @Summary Revoke a user's refresh token
// @Description Revokes a user's refresh token
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param request body RevokeRefreshTokenRequest true "Revoke Refresh Token Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 404 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/refresh-token [delete]
func (c *controller) RevokeRefreshToken(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*RevokeRefreshTokenRequest)

	// Revoke the user's refresh token
	Service.RevokeRefreshToken(w, r, body.RefreshTokenID)

	// Log the successful token revocation
	internallogger.Api.RevokeRefreshToken(body.RefreshTokenID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusOK,
		),
	)
}

// LogOut logs out a user
// @Summary Log out a user
// @Description Logs out a user
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/logout [post]
func (c *controller) LogOut(w http.ResponseWriter, r *http.Request) {
	// Log out the user
	userID := Service.LogOut(w, r)

	// Log the successful logout
	internallogger.Api.LogOut(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusOK,
		),
	)
}

// RevokeRefreshTokens revokes a user's refresh tokens
// @Summary Revoke a user's refresh tokens
// @Description Revokes a user's refresh tokens
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/refresh-tokens [delete]
func (c *controller) RevokeRefreshTokens(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Revoke the user's refresh tokens
	userID := Service.RevokeRefreshTokens(w, r)

	// Log the successful token revocation
	internallogger.Api.RevokeRefreshTokens(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusOK,
		),
	)
}

// RefreshToken refreshes a user token
// @Summary Refresh a user token
// @Description Refreshes a user token
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Success 201 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/refresh-token [post]
func (c *controller) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Refresh the token
	userID := internalcookie.RefreshTokenFn(w, r)

	// Log the successful token refresh
	internallogger.Api.RefreshToken(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusCreated,
		),
	)
}

// GenerateTOTPUrl generates a TOTP URL
// @Summary Generate a TOTP URL
// @Description Generates a TOTP URL
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Success 201 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/totp/generate [post]
func (c *controller) GenerateTOTPUrl(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Generate the TOTP URL
	userID, totpUrl := Service.GenerateTOTPUrl(r)

	// Log the successful TOTP URL generation
	internallogger.Api.GenerateTOTPUrl(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			GenerateTOTPUrlResponse{
				TOTPUrl: *totpUrl,
			},
			http.StatusCreated,
		),
	)
}

// VerifyTOTP verifies a TOTP code
// @Summary Verify a TOTP code
// @Description Verifies a TOTP code
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param request body VerifyTOTPRequest
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/totp/verify [post]
func (c *controller) VerifyTOTP(w http.ResponseWriter, r *http.Request) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*VerifyTOTPRequest)

	// Verify the TOTP code
	userID, recoveryCodes := Service.VerifyTOTP(r, body)

	// Log the successful TOTP verification
	internallogger.Api.VerifyTOTP(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			VerifyTOTPResponse{
				RecoveryCodes: *recoveryCodes,
			},
			http.StatusOK,
		),
	)
}

// RevokeTOTP revokes a user's TOTP
// @Summary Revoke a user's TOTP
// @Description Revokes a user's TOTP
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/totp [delete]
func (c *controller) RevokeTOTP(w http.ResponseWriter, r *http.Request) {
	// Revoke the user's TOTP
	userID := Service.RevokeTOTP(r)

	// Log the successful TOTP revocation
	internallogger.Api.RevokeTOTP(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusOK,
		),
	)
}

// ChangeEmail changes the email of the authenticated user
// @Summary Changes the email of the authenticated user
// @Description Changes the email of the authenticated user
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param request body ChangeEmailRequest true "Change Email Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/email [put]
func (c *controller) ChangeEmail(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*ChangeEmailRequest)

	// Change the email
	userID := Service.ChangeEmail(r, body)

	// Log the successful email change
	internallogger.Api.ChangeEmail(userID, body.Email)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusOK,
		),
	)
}

// SendEmailVerificationToken sends an email verification token to the authenticated user
// @Summary Sends an email verification token to the authenticated user
// @Description Sends an email verification token to the authenticated user
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/email/send-verification [post]
func (c *controller) SendEmailVerificationToken(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Send the email verification token
	userID := Service.SendEmailVerificationToken(r)

	// Log the successful email verification token request
	internallogger.Api.SendEmailVerificationToken(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusOK,
		),
	)
}

// VerifyEmail verifies the email of the authenticated user
// @Summary Verifies the email of the authenticated user
// @Description Verifies the email of the authenticated user
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param token path string true "Token"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/email/verify/{token} [post]
func (c *controller) VerifyEmail(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the token from the URL
	var token string
	if !c.GetTokenWildcard(w, r, &token) {
		return
	}

	// Verify the email
	userID := Service.VerifyEmail(token)

	// Log the successful email verification
	internallogger.Api.VerifyEmail(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusOK,
		),
	)
}

// ChangePassword changes a user's password
// @Summary Change a user's password
// @Description Changes a user's password
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param request body ChangePasswordRequest
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/password [put]
func (c *controller) ChangePassword(w http.ResponseWriter, r *http.Request) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*ChangePasswordRequest)

	// Change the password
	userID := Service.ChangePassword(r, body)

	// Log the successful password change
	internallogger.Api.ChangePassword(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusOK,
		),
	)
}

// ForgotPassword sends a password reset email
// @Summary Send a password reset email
// @Description Sends a password reset email
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param request body ForgotPasswordRequest
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/password/forgot [post]
func (c *controller) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*ForgotPasswordRequest)

	// Send the reset password email
	userID := Service.ForgotPassword(body)

	// Log the successful reset password email request
	internallogger.Api.ForgotPassword(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusOK,
		),
	)
}

// ResetPassword resets a user's password
// @Summary Reset a user's password
// @Description Resets a user's password
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param token path string true "Token"
// @Param request body ResetPasswordRequest true "Reset Password Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/password/reset/{token} [post]
func (c *controller) ResetPassword(w http.ResponseWriter, r *http.Request) {
	// Get the token from the URL
	var token string
	if !c.GetTokenWildcard(w, r, &token) {
		return
	}

	// Get the body from the context
	body, _ := gonethttpctx.GetCtxBody(r).(*ResetPasswordRequest)

	// Reset the password
	userID := Service.ResetPassword(token, body)

	// Log the successful password reset
	internallogger.Api.ResetPassword(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusOK,
		),
	)
}

// ChangePhoneNumber changes the phone number of the authenticated user
// @Summary Changes the phone number of the authenticated user
// @Description Changes the phone number of the authenticated user
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param request body ChangePhoneNumberRequest true "Change Phone Number Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/phone-number [put]
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
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/phone-number/send-verification [post]
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
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param request body VerifyPhoneNumberRequest true "Verify Phone Number Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/phone-number/verify [post]
func (c *controller) VerifyPhoneNumber(
	w http.ResponseWriter,
	r *http.Request,
) {
	internalhandler.Handler.HandleResponse(
		w, gonethttpstatusresponse.NewJSendNotImplemented(nil),
	)
}
