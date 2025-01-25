package auth

import (
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	gostringsconvert "github.com/ralvarezdev/go-strings/convert"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internaljwt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"net/http"
)

type (
	// controller is the structure for the API V1 auth controller
	controller struct {
		Service   *service
		Validator *validator
		gonethttpfactory.Controller
	}
)

// RegisterRoutes registers the routes for the API V1 auth controller
func (c *controller) RegisterRoutes() {
	c.RegisterRoute(
		"POST /signup",
		c.SignUp,
	)
	c.RegisterRoute(
		"POST /login",
		c.LogIn,
	)
	c.RegisterRoute(
		"POST /refresh-token",
		c.RefreshToken,
		internaljwt.Authenticate(gojwtinterception.RefreshToken),
	)
	c.RegisterRoute(
		"POST /logout",
		c.LogOut,
		internaljwt.Authenticate(gojwtinterception.AccessToken),
	)
	c.RegisterRoute(
		"GET /refresh-token/{id}",
		c.GetRefreshToken,
		internaljwt.Authenticate(gojwtinterception.AccessToken),
	)
	c.RegisterRoute(
		"GET /refresh-tokens",
		c.ListRefreshTokens,
		internaljwt.Authenticate(gojwtinterception.AccessToken),
	)
	c.RegisterRoute(
		"DELETE /refresh-token/{id}",
		c.RevokeRefreshToken,
		internaljwt.Authenticate(gojwtinterception.AccessToken),
	)
	c.RegisterRoute(
		"DELETE /refresh-tokens",
		c.RevokeRefreshTokens,
		internaljwt.Authenticate(gojwtinterception.AccessToken),
	)
	c.RegisterRoute(
		"POST /totp/generate",
		c.GenerateTOTPUrl,
		internaljwt.Authenticate(gojwtinterception.AccessToken),
	)
	c.RegisterRoute(
		"POST /totp/verify",
		c.VerifyTOTP,
		internaljwt.Authenticate(gojwtinterception.AccessToken),
	)
	c.RegisterRoute(
		"DELETE /totp",
		c.RevokeTOTP,
		internaljwt.Authenticate(gojwtinterception.AccessToken),
	)
}

// RegisterGroups registers the router groups for the API V1 auth controller
func (c *controller) RegisterGroups() {}

// getRefreshTokenID gets the refresh token ID from the path
func (c *controller) getRefreshTokenID(
	w http.ResponseWriter,
	r *http.Request,
	refreshTokenID *int64,
) bool {
	// Get the refresh token ID from the path
	return internalhandler.Handler.ParseWildcard(
		w, r, "id", refreshTokenID,
		gostringsconvert.ToInt64,
	)
}

// SignUp signs up a new user
// @Summary Sign up a new user
// @Description Creates a new user account with the provided details
// @Tags api v1 user
// @Accept json
// @Produce json
// @Param request body SignUpRequest true "Sign Up Request"
// @Success 201 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/user/signup [post]
func (c *controller) SignUp(w http.ResponseWriter, r *http.Request) {
	// Decode the request body and validate the request
	var body SignUpRequest
	if !internalhandler.Handler.Parse(
		w,
		r,
		&body,
		c.Validator.SignUp(&body),
	) {
		return
	}

	// Sign up the user
	userID, err := c.Service.SignUp(r, &body)
	if err != nil {
		internalhandler.Handler.HandleError(w, err)
		return
	}

	// Log the user sign up
	internallogger.Api.SignUp(*userID)

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
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Failure 400 {object} gonethttpresponse.JSendFailResponse
// @Failure 401 {object} gonethttpresponse.JSendFailResponse
// @Failure 500 {object} gonethttpresponse.JSendErrorResponse
// @Router /api/v1/auth/login [post]
func (c *controller) LogIn(w http.ResponseWriter, r *http.Request) {
	// Decode the request body and validate the request
	var body LogInRequest
	if !internalhandler.Handler.Parse(
		w,
		r,
		&body,
		c.Validator.LogIn(&body),
	) {
		return
	}

	// Log in the user
	userID, userTokens, err := c.Service.LogIn(r, &body)
	if err != nil {
		internalhandler.Handler.HandleError(w, err)
		return
	}

	// Log the successful login
	internallogger.Api.LogIn(*userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			RefreshTokenResponse{
				RefreshToken: (*userTokens)[gojwttoken.RefreshToken],
				AccessToken:  (*userTokens)[gojwttoken.AccessToken],
			},
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
// @Success 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/refresh-tokens [get]
func (c *controller) ListRefreshTokens(w http.ResponseWriter, r *http.Request) {
	// Get the user's refresh tokens
	userID, userRefreshTokens, err := c.Service.ListRefreshTokens(r)
	if err != nil {
		internalhandler.Handler.HandleError(w, err)
		return
	}

	// Log the successful fetch of the user's refresh tokens
	internallogger.Api.ListRefreshTokens(*userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			ListRefreshTokensResponse{
				RefreshTokens: *userRefreshTokens,
			},
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
// @Param id path string true "Refresh Token ID"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Success 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/refresh-token/{id} [get]
func (c *controller) GetRefreshToken(w http.ResponseWriter, r *http.Request) {
	// Get the refresh token ID from the path
	var refreshTokenID int64
	if !c.getRefreshTokenID(w, r, &refreshTokenID) {
		return
	}

	// Get the user's refresh token by ID
	userID, userRefreshToken, err := c.Service.GetRefreshToken(
		r,
		refreshTokenID,
	)
	if err != nil {
		internalhandler.Handler.HandleError(w, err)
		return
	}

	// Log the successful fetch of the user's refresh token
	internallogger.Api.GetRefreshToken(*userID, refreshTokenID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			GetRefreshTokenResponse{
				RefreshToken: userRefreshToken,
			},
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
// @Param id path string true "Refresh Token ID"
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Success 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/refresh-token/{id} [delete]
func (c *controller) RevokeRefreshToken(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the refresh token ID from the path
	var refreshTokenID int64
	if !c.getRefreshTokenID(w, r, &refreshTokenID) {
		return
	}

	// Revoke the user's refresh token
	err := c.Service.RevokeRefreshToken(r, refreshTokenID)
	if err != nil {
		internalhandler.Handler.HandleError(w, err)
		return
	}

	// Log the successful token revocation
	internallogger.Api.RevokeRefreshToken(refreshTokenID)

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
// @Success 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/logout [post]
func (c *controller) LogOut(w http.ResponseWriter, r *http.Request) {
	// Log out the user
	userID, err := c.Service.LogOut(r)
	if err != nil {
		internalhandler.Handler.HandleError(w, err)
		return
	}

	// Log the successful logout
	internallogger.Api.LogOut(*userID)

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
// @Success 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/refresh-tokens [delete]
func (c *controller) RevokeRefreshTokens(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Revoke the user's refresh tokens
	userID, err := c.Service.RevokeRefreshTokens(r)
	if err != nil {
		internalhandler.Handler.HandleError(w, err)
		return
	}

	// Log the successful token revocation
	internallogger.Api.RevokeRefreshTokens(*userID)

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
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Success 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/refresh-token [post]
func (c *controller) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Refresh the token
	userID, userTokens, err := c.Service.RefreshToken(r)
	if err != nil {
		internalhandler.Handler.HandleError(w, err)
		return
	}

	// Log the successful token refresh
	internallogger.Api.RefreshToken(*userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			RefreshTokenResponse{
				RefreshToken: (*userTokens)[gojwttoken.RefreshToken],
				AccessToken:  (*userTokens)[gojwttoken.AccessToken],
			},
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
// @Success 200 {object} gonethttpresponse.JSendSuccessBody
// @Success 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/totp/generate [post]
func (c *controller) GenerateTOTPUrl(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Generate the TOTP URL
	userID, totpUrl, err := c.Service.GenerateTOTPUrl(r)
	if err != nil {
		internalhandler.Handler.HandleError(w, err)
		return
	}

	// Log the successful TOTP URL generation
	internallogger.Api.GenerateTOTPUrl(*userID)

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
// @Success 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/totp/verify [post]
func (c *controller) VerifyTOTP(w http.ResponseWriter, r *http.Request) {
	// Decode the request body and validate the request
	var body VerifyTOTPRequest
	if !internalhandler.Handler.Parse(
		w,
		r,
		&body,
		c.Validator.VerifyTOTP(&body),
	) {
		return
	}

	// Verify the TOTP code
	userID, recoveryCodes, err := c.Service.VerifyTOTP(r, &body)
	if err != nil {
		internalhandler.Handler.HandleError(w, err)
		return
	}

	// Log the successful TOTP verification
	internallogger.Api.VerifyTOTP(*userID)

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
// @Success 401 {object} gonethttpresponse.JSendFailBody
// @Failure 500 {object} gonethttpresponse.JSendErrorBody
// @Router /api/v1/auth/totp [delete]
func (c *controller) RevokeTOTP(w http.ResponseWriter, r *http.Request) {
	// Revoke the user's TOTP
	userID, err := c.Service.RevokeTOTP(r)
	if err != nil {
		internalhandler.Handler.HandleError(w, err)
		return
	}

	// Log the successful TOTP revocation
	internallogger.Api.RevokeTOTP(*userID)

	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusOK,
		),
	)
}
