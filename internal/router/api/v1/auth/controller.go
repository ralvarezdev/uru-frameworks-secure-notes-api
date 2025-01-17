package auth

import (
	gojwtcache "github.com/ralvarezdev/go-jwt/cache"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttphandler "github.com/ralvarezdev/go-net/http/handler"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	gostringsconvert "github.com/ralvarezdev/go-strings/convert"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
	"net/http"
)

type (
	// Controller is the structure for the API V1 auth controller
	Controller struct {
		handler            gonethttphandler.Handler
		authenticator      gonethttpmiddlewareauth.Authenticator
		postgresService    *internalpostgres.Service
		jwtIssuer          gojwtissuer.Issuer
		jwtTokenValidator  gojwtcache.TokenValidator
		service            *Service
		validator          *Validator
		logger             *internallogger.Logger
		jwtValidatorLogger *gojwtvalidator.Logger
		gonethttproute.Controller
	}
)

// NewController creates a new API V1 auth controller
func NewController(
	baseRouter gonethttproute.RouterWrapper,
	authenticator gonethttpmiddlewareauth.Authenticator,
	postgresService *internalpostgres.Service,
	jwtIssuer gojwtissuer.Issuer,
	jwtTokenValidator gojwtcache.TokenValidator,
) *Controller {
	// Load the validator mappers
	LoadMappers()

	return &Controller{
		Controller: gonethttproute.Controller{
			RouterWrapper: baseRouter.NewGroup(BasePath),
		},
		handler:           internalhandler.Handler,
		authenticator:     authenticator,
		postgresService:   postgresService,
		jwtIssuer:         jwtIssuer,
		jwtTokenValidator: jwtTokenValidator,
		service: &Service{
			jwtIssuer:         jwtIssuer,
			postgresService:   postgresService,
			jwtTokenValidator: jwtTokenValidator,
		},
		validator:          &Validator{Service: internalvalidator.ValidationsService},
		logger:             internallogger.Api,
		jwtValidatorLogger: internallogger.JwtValidator,
	}
}

// RegisterRoutes registers the routes for the API V1 auth controller
func (c *Controller) RegisterRoutes() {
	c.RegisterRoute(
		"POST /login",
		c.LogIn,
	)
	c.RegisterRoute(
		"POST /refresh-token",
		c.RefreshToken,
		c.authenticator.Authenticate(gojwtinterception.RefreshToken),
	)
	c.RegisterRoute(
		"POST /logout",
		c.LogOut,
		c.authenticator.Authenticate(gojwtinterception.AccessToken),
	)
	c.RegisterRoute(
		"GET /refresh-token/{id}",
		c.GetRefreshToken,
		c.authenticator.Authenticate(gojwtinterception.AccessToken),
	)
	c.RegisterRoute(
		"GET /refresh-tokens",
		c.ListRefreshTokens,
		c.authenticator.Authenticate(gojwtinterception.AccessToken),
	)
	c.RegisterRoute(
		"DELETE /refresh-token/{id}",
		c.RevokeRefreshToken,
		c.authenticator.Authenticate(gojwtinterception.AccessToken),
	)
	c.RegisterRoute(
		"DELETE /refresh-tokens",
		c.RevokeRefreshTokens,
		c.authenticator.Authenticate(gojwtinterception.AccessToken),
	)
	c.RegisterRoute(
		"POST /totp/generate",
		c.GenerateTOTPUrl,
		c.authenticator.Authenticate(gojwtinterception.AccessToken),
	)
	c.RegisterRoute(
		"POST /totp/verify",
		c.VerifyTOTP,
		c.authenticator.Authenticate(gojwtinterception.AccessToken),
	)
	c.RegisterRoute(
		"DELETE /totp",
		c.RevokeTOTP,
		c.authenticator.Authenticate(gojwtinterception.AccessToken),
	)
}

// RegisterGroups registers the router groups for the API V1 auth controller
func (c *Controller) RegisterGroups() {}

// getRefreshTokenID gets the refresh token ID from the path
func (c *Controller) getRefreshTokenID(
	w http.ResponseWriter,
	r *http.Request,
	refreshTokenID *int64,
) bool {
	// Get the refresh token ID from the path
	return c.handler.ParseWildcard(
		w, r, "id", refreshTokenID,
		gostringsconvert.ToInt64,
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
func (c *Controller) LogIn(w http.ResponseWriter, r *http.Request) {
	// Decode the request body and validate the request
	var requestBody LogInRequest
	if !c.handler.DecodeAndValidate(
		w,
		r,
		&requestBody,
		c.validator.ValidateLogInRequest,
	) {
		return
	}

	// Log in the user
	userID, userTokens, err := c.service.LogIn(r, &requestBody)
	if err != nil {
		c.handler.HandleError(w, err)
		return
	}

	// Log the successful login
	c.logger.LogIn(*userID)

	// Handle the response
	c.handler.HandleResponse(
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
func (c *Controller) ListRefreshTokens(w http.ResponseWriter, r *http.Request) {
	// Get the user's refresh tokens
	userID, userRefreshTokens, err := c.service.ListRefreshTokens(r)
	if err != nil {
		c.handler.HandleError(w, err)
		return
	}

	// Log the successful fetch of the user's refresh tokens
	c.logger.ListRefreshTokens(*userID)

	// Handle the response
	c.handler.HandleResponse(
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
func (c *Controller) GetRefreshToken(w http.ResponseWriter, r *http.Request) {
	// Get the refresh token ID from the path
	var refreshTokenID int64
	if !c.getRefreshTokenID(w, r, &refreshTokenID) {
		return
	}

	// Get the user's refresh token by ID
	userID, userRefreshToken, err := c.service.GetRefreshToken(
		r,
		refreshTokenID,
	)
	if err != nil {
		c.handler.HandleError(w, err)
		return
	}

	// Log the successful fetch of the user's refresh token
	c.logger.GetRefreshToken(*userID, refreshTokenID)

	// Handle the response
	c.handler.HandleResponse(
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
func (c *Controller) RevokeRefreshToken(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get the refresh token ID from the path
	var refreshTokenID int64
	if !c.getRefreshTokenID(w, r, &refreshTokenID) {
		return
	}

	// Revoke the user's refresh token
	err := c.service.RevokeRefreshToken(r, refreshTokenID)
	if err != nil {
		c.handler.HandleError(w, err)
		return
	}

	// Log the successful token revocation
	c.logger.RevokeRefreshToken(refreshTokenID)

	// Handle the response
	c.handler.HandleResponse(
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
func (c *Controller) LogOut(w http.ResponseWriter, r *http.Request) {
	// Log out the user
	userID, err := c.service.LogOut(r)
	if err != nil {
		c.handler.HandleError(w, err)
		return
	}

	// Log the successful logout
	c.logger.LogOut(*userID)

	// Handle the response
	c.handler.HandleResponse(
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
func (c *Controller) RevokeRefreshTokens(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Revoke the user's refresh tokens
	userID, err := c.service.RevokeRefreshTokens(r)
	if err != nil {
		c.handler.HandleError(w, err)
		return
	}

	// Log the successful token revocation
	c.logger.RevokeRefreshTokens(*userID)

	// Handle the response
	c.handler.HandleResponse(
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
func (c *Controller) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Refresh the token
	userID, userTokens, err := c.service.RefreshToken(r)
	if err != nil {
		c.handler.HandleError(w, err)
		return
	}

	// Log the successful token refresh
	c.logger.RefreshToken(*userID)

	// Handle the response
	c.handler.HandleResponse(
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
func (c *Controller) GenerateTOTPUrl(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Generate the TOTP URL
	userID, totpUrl, err := c.service.GenerateTOTPUrl(r)
	if err != nil {
		c.handler.HandleError(w, err)
		return
	}

	// Log the successful TOTP URL generation
	c.logger.GenerateTOTPUrl(*userID)

	// Handle the response
	c.handler.HandleResponse(
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
func (c *Controller) VerifyTOTP(w http.ResponseWriter, r *http.Request) {
	// Decode the request body and validate the request
	var requestBody VerifyTOTPRequest
	if !c.handler.DecodeAndValidate(
		w,
		r,
		&requestBody,
		c.validator.ValidateVerifyTOTPRequest,
	) {
		return
	}

	// Verify the TOTP code
	userID, recoveryCodes, err := c.service.VerifyTOTP(r, &requestBody)
	if err != nil {
		c.handler.HandleError(w, err)
		return
	}

	// Log the successful TOTP verification
	c.logger.VerifyTOTP(*userID)

	// Handle the response
	c.handler.HandleResponse(
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
func (c *Controller) RevokeTOTP(w http.ResponseWriter, r *http.Request) {
	// Revoke the user's TOTP
	userID, err := c.service.RevokeTOTP(r)
	if err != nil {
		c.handler.HandleError(w, err)
		return
	}

	// Log the successful TOTP revocation
	c.logger.RevokeTOTP(*userID)

	// Handle the response
	c.handler.HandleResponse(
		w, gonethttpresponse.NewJSendSuccessResponse(
			nil,
			http.StatusOK,
		),
	)
}
