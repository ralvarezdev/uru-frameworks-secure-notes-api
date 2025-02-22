package auth

import (
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gonethttpctx "github.com/ralvarezdev/go-net/http/context"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	gonethttpstatusresponse "github.com/ralvarezdev/go-net/http/status/response"
	internalcookie "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/cookie"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"net/http"
)

type (
	// controller is the structure for the API V1 auth controller
	controller struct{}
)

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
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*SignUpRequest)

	// Sign up the user
	userID := Service.SignUp(requestBody)

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
// @Success 201 {object} LogInResponseBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/login [post]
func (c *controller) LogIn(w http.ResponseWriter, r *http.Request) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*LogInRequest)

	// Log in the user
	userID, response := Service.LogIn(w, r, requestBody)

	// Log the successful login
	internallogger.Api.LogIn(userID)

	// Handle the response
	if response != nil {
		internalhandler.Handler.HandleResponse(
			w, response,
		)
	}
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
// @Security CookieAuth
// @Success 200 {object} ListRefreshTokensResponseBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/refresh-tokens [get]
func (c *controller) ListRefreshTokens(w http.ResponseWriter, r *http.Request) {
	// Get the user's refresh tokens
	userID, responseBody := Service.ListRefreshTokens(r)

	// Log the successful fetch of the user's refresh tokens
	internallogger.Api.ListRefreshTokens(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewResponse(
			responseBody,
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
// @Security CookieAuth
// @Param request body GetRefreshTokenRequest true "Get Refresh Token Request"
// @Success 200 {object} GetRefreshTokenResponseBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/refresh-token [get]
func (c *controller) GetRefreshToken(w http.ResponseWriter, r *http.Request) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*GetRefreshTokenRequest)

	// Get the user's refresh token by ID
	userID, responseBody := Service.GetRefreshToken(
		r,
		requestBody.RefreshTokenID,
	)

	// Log the successful fetch of the user's refresh token
	internallogger.Api.GetRefreshToken(userID, requestBody.RefreshTokenID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewResponse(
			responseBody,
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
// @Security CookieAuth
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
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*RevokeRefreshTokenRequest)

	// Revoke the user's refresh token
	Service.RevokeRefreshToken(w, r, requestBody.RefreshTokenID)

	// Log the successful token revocation
	internallogger.Api.RevokeRefreshToken(requestBody.RefreshTokenID)

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
// @Security CookieAuth
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
// @Security CookieAuth
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
// @Security CookieAuth
// @Success 201 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/refresh-token [post]
func (c *controller) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Refresh the token
	userID, _ := internalcookie.RefreshTokenFn(gojwttoken.RefreshToken)(w, r)

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

// Generate2FATOTPUrl generates a 2FA TOTP URL
// @Summary Generate a 2FA TOTP URL
// @Description Generates a 2FA TOTP URL
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Security CookieAuth
// @Success 201 {object} Generate2FATOTPUrlResponseBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/2fa/totp/generate [post]
func (c *controller) Generate2FATOTPUrl(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Generate the 2FA TOTP URL
	userID, responseBody := Service.Generate2FATOTPUrl(r)

	// Log the successful 2FA TOTP URL generation
	internallogger.Api.Generate2FATOTPUrl(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewResponse(
			responseBody,
			http.StatusCreated,
		),
	)
}

// Verify2FATOTP verifies a 2FA TOTP code
// @Summary Verify a 2FA TOTP code
// @Description Verifies a 2FA TOTP code
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body Verify2FATOTPRequest true "Verify 2FA TOTP Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/2fa/totp/verify [post]
func (c *controller) Verify2FATOTP(w http.ResponseWriter, r *http.Request) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*Verify2FATOTPRequest)

	// Verify the 2FA TOTP code
	userID := Service.Verify2FATOTP(r, requestBody)

	// Log the successful 2FA TOTP verification
	internallogger.Api.Verify2FATOTP(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusOK,
		),
	)
}

// Revoke2FATOTP revokes a user's 2FA TOTP
// @Summary Revoke a user's 2FA TOTP
// @Description Revokes a user's 2FA TOTP
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Security CookieAuth
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/2fa/totp [delete]
func (c *controller) Revoke2FATOTP(w http.ResponseWriter, r *http.Request) {
	// Revoke the user's 2FA TOTP
	userID := Service.Revoke2FATOTP(r)

	// Log the successful 2FA TOTP revocation
	internallogger.Api.Revoke2FATOTP(userID)

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
// @Security CookieAuth
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
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*ChangeEmailRequest)

	// Change the email
	userID := Service.ChangeEmail(r, requestBody)

	// Log the successful email change
	internallogger.Api.ChangeEmail(userID, requestBody.Email)

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
// @Security CookieAuth
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
// @Security CookieAuth
// @Param request body VerifyEmailRequest true "Verify Email Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/email/verify [post]
func (c *controller) VerifyEmail(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*VerifyEmailRequest)

	// Verify the email
	userID := Service.VerifyEmail(requestBody)

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
// @Security CookieAuth
// @Param request body ChangePasswordRequest true "Change Password Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/password [put]
func (c *controller) ChangePassword(w http.ResponseWriter, r *http.Request) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*ChangePasswordRequest)

	// Change the password
	userID := Service.ChangePassword(r, requestBody)

	// Log the successful password change
	internallogger.Api.ChangePassword(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewResponse(
			gonethttpresponse.NewJSendSuccessBody(
				nil,
			),
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
// @Security CookieAuth
// @Param request body ForgotPasswordRequest true "Forgot Password Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/password/forgot [post]
func (c *controller) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*ForgotPasswordRequest)

	// Send the reset password email
	userID := Service.ForgotPassword(requestBody)

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
// @Security CookieAuth
// @Param request body ResetPasswordRequest true "Reset Password Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/password/reset [post]
func (c *controller) ResetPassword(w http.ResponseWriter, r *http.Request) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*ResetPasswordRequest)

	// Reset the password
	userID := Service.ResetPassword(requestBody)

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
// @Security CookieAuth
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
// @Security CookieAuth
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
// @Security CookieAuth
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

// EnableUser2FA enables 2FA for the authenticated user
// @Summary Enable 2FA for the authenticated user
// @Description Enables 2FA for the authenticated user
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body EnableUser2FARequest true "Enable User 2FA Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/2fa/enable [post]
func (c *controller) EnableUser2FA(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*EnableUser2FARequest)

	// Enable 2FA for the user
	userID, responseBody := Service.EnableUser2FA(r, requestBody)

	// Log the successful 2FA enablement
	internallogger.Api.EnableUser2FA(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewResponse(
			responseBody,
			http.StatusOK,
		),
	)
}

// DisableUser2FA disables 2FA for the authenticated user
// @Summary Disable 2FA for the authenticated user
// @Description Disables 2FA for the authenticated user
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body DisableUser2FARequest true "Disable User 2FA Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/2fa/disable [post]
func (c *controller) DisableUser2FA(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*DisableUser2FARequest)

	// Disable 2FA for the user
	userID := Service.DisableUser2FA(r, requestBody)

	// Log the successful 2FA disablement
	internallogger.Api.DisableUser2FA(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusOK,
		),
	)
}

// RegenerateUser2FARecoveryCodes regenerates the 2FA recovery codes for the authenticated user
// @Summary Regenerate 2FA recovery codes for the authenticated user
// @Description Regenerates the 2FA recovery codes for the authenticated user
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body RegenerateUser2FARecoveryCodesRequest true "Regenerate User 2FA Recovery Codes Request"
// @Success 200 {object} RegenerateUser2FARecoveryCodesResponseBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/2fa/recovery-codes/regenerate [post]
func (c *controller) RegenerateUser2FARecoveryCodes(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*RegenerateUser2FARecoveryCodesRequest)

	// Regenerate the 2FA recovery codes for the user
	userID, responseBody := Service.RegenerateUser2FARecoveryCodes(
		r,
		requestBody,
	)

	// Log the successful 2FA recovery codes regeneration
	internallogger.Api.RegenerateUser2FARecoveryCodes(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewResponse(
			responseBody,
			http.StatusOK,
		),
	)
}

// SendUser2FAEmailCode sends a 2FA email code to the authenticated user
// @Summary Send 2FA email code to the authenticated user
// @Description Sends a 2FA email code to the authenticated user
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body SendUser2FAEmailCodeRequest true "Send User 2FA Email Code Request"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/2fa/email/send-code [post]
func (c *controller) SendUser2FAEmailCode(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the body from the context
	requestBody, _ := gonethttpctx.GetCtxBody(r).(*SendUser2FAEmailCodeRequest)

	// Send the 2FA email code
	userID := Service.SendUser2FAEmailCode(r, requestBody)

	// Log the successful 2FA email code send
	internallogger.Api.SendUser2FAEmailCode(userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewResponse(
			nil,
			http.StatusOK,
		),
	)
}
