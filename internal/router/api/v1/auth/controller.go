package auth

import (
	"errors"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttperrors "github.com/ralvarezdev/go-net/http/errors"
	gonethttphandler "github.com/ralvarezdev/go-net/http/handler"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/_common"
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
) *Controller {
	// Load the validator mappers
	LoadMappers()

	return &Controller{
		Controller: gonethttproute.Controller{
			RouterWrapper: baseRouter.NewGroup(BasePath),
		},
		handler:         internalhandler.Handler,
		authenticator:   authenticator,
		postgresService: postgresService,
		jwtIssuer:       jwtIssuer,
		service: &Service{
			JwtIssuer:       jwtIssuer,
			PostgresService: postgresService,
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
		"POST /logout",
		c.LogOut,
		c.authenticator.Authenticate(gojwtinterception.AccessToken),
	)
	c.RegisterRoute(
		"POST /refresh-token",
		c.RefreshToken,
		c.authenticator.Authenticate(gojwtinterception.RefreshToken),
	)
	c.RegisterRoute(
		"DELETE /refresh-token",
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
}

// RegisterGroups registers the router groups for the API V1 auth controller
func (c *Controller) RegisterGroups() {}

// LogIn logs in a user
// @Summary Log in a user
// @Description Authenticates a user and returns a seed token
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param request body LogInRequest true "Log In Request"
// @Success 200 {object} LogInResponse
// @Failure 400 {object} gonethttphandler.JSendResponse
// @Failure 401 {object} gonethttphandler.JSendResponse
// @Failure 500 {object} gonethttphandler.JSendResponse
// @Router /api/v1/auth/login [post]
func (c *Controller) LogIn(w http.ResponseWriter, r *http.Request) {
	// Decode the request body and validate the request
	var body LogInRequest
	if !c.handler.HandleRequestAndValidations(
		w,
		r,
		&body,
		c.validator.ValidateLogInRequest,
	) {
		return
	}

	// Log in the user
	userID, userTokens, err := c.service.LogIn(r, &body)
	if err == nil {
		// Log the successful login
		c.logger.LogIn(*userID)

		// Handle the response
		c.handler.HandleResponse(
			w, gonethttpresponse.NewSuccessResponse(
				RefreshTokenResponse{
					RefreshToken: (*userTokens)[gojwttoken.RefreshToken],
					AccessToken:  (*userTokens)[gojwttoken.AccessToken],
				},
				http.StatusCreated,
			),
		)
		return
	}

	// Handle the error
	anErrorOccurred := false
	data := make(map[string]*[]string)
	if errors.Is(err, internalapiv1common.UserNotFoundByUsername) {
		data["username"] = &[]string{err.Error()}
	} else if errors.Is(err, ErrInvalidPassword) {
		data["password"] = &[]string{err.Error()}
	} else if errors.Is(err, ErrInvalidTOTPCode) || errors.Is(
		err,
		ErrMissingTOTPCode,
	) {
		data["totp_code"] = &[]string{err.Error()}
	} else if errors.Is(err, ErrInvalidTOTPRecoveryCode) || errors.Is(
		err,
		ErrMissingIsTOTPRecoveryCode,
	) {
		data["is_totp_recovery_code"] = &[]string{err.Error()}
	} else {
		anErrorOccurred = true
	}

	// Check if an error occurred
	if !anErrorOccurred {
		c.handler.HandleResponse(
			w, gonethttpresponse.NewFailResponse(
				&data,
				nil,
				http.StatusUnauthorized,
			),
		)
		return
	}

	c.handler.HandleResponse(
		w, gonethttpresponse.NewDebugErrorResponse(
			gonethttperrors.InternalServerError,
			err,
			nil,
			nil,
			http.StatusInternalServerError,
		),
	)
}

// RevokeRefreshToken revokes a user's refresh token
// @Summary Revoke a user's refresh token
// @Description Revokes a user's refresh token
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param request body RevokeRefreshTokenRequest
// @Success 200 {object} internalapiv1common.BasicResponse
// @Failure 401 {object} gonethttphandler.JSendResponse
// @Failure 500 {object} gonethttphandler.JSendResponse
// @Router /api/v1/auth/refresh-token [delete]
func (c *Controller) RevokeRefreshToken(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Decode the request body and validate the request
	var body RevokeRefreshTokenRequest
	if !c.handler.HandleRequestAndValidations(
		w,
		r,
		&body,
		c.validator.ValidateRevokeRefreshTokenRequest,
	) {
		return
	}

	// Revoke the user's refresh token
	err := c.service.RevokeRefreshToken(r, body.UserRefreshTokenID)
	if err == nil {
		// Log the successful token revocation
		c.logger.RevokeRefreshToken(body.UserRefreshTokenID)

		// Handle the response
		c.handler.HandleResponse(
			w, gonethttpresponse.NewSuccessResponse(
				nil,
				http.StatusOK,
			),
		)
		return
	}

	// Handle the error
	c.handler.HandleResponse(
		w, gonethttpresponse.NewDebugErrorResponse(
			gonethttperrors.InternalServerError,
			err,
			nil, nil, http.StatusInternalServerError,
		),
	)
}

// LogOut logs out a user
// @Summary Log out a user
// @Description Logs out a user
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Success 200 {object} internalapiv1common.BasicResponse
// @Failure 401 {object} gonethttphandler.JSendResponse
// @Failure 500 {object} gonethttphandler.JSendResponse
// @Router /api/v1/auth/logout [post]
func (c *Controller) LogOut(w http.ResponseWriter, r *http.Request) {
	// Log out the user
	userID, err := c.service.LogOut(r)
	if err == nil {
		// Log the successful logout
		c.logger.LogOut(*userID)

		// Handle the response
		c.handler.HandleResponse(
			w, gonethttpresponse.NewSuccessResponse(
				nil,
				http.StatusOK,
			),
		)
		return
	}

	// Handle the error
	c.handler.HandleResponse(
		w, gonethttpresponse.NewDebugErrorResponse(
			gonethttperrors.InternalServerError,
			err,
			nil, nil, http.StatusInternalServerError,
		),
	)
}

// RevokeRefreshTokens revokes a user's refresh tokens
// @Summary Revoke a user's refresh tokens
// @Description Revokes a user's refresh tokens
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param request body RevokeRefreshTokensRequest
// @Success 200 {object} internalapiv1common.BasicResponse
// @Failure 401 {object} gonethttphandler.JSendResponse
// @Failure 500 {object} gonethttphandler.JSendResponse
// @Router /api/v1/auth/refresh-tokens [delete]
func (c *Controller) RevokeRefreshTokens(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Revoke the user's refresh tokens
	userID, err := c.service.RevokeRefreshTokens(r)
	if err == nil {
		// Log the successful token revocation
		c.logger.RevokeRefreshTokens(*userID)

		// Handle the response
		c.handler.HandleResponse(
			w, gonethttpresponse.NewSuccessResponse(
				nil,
				http.StatusOK,
			),
		)
		return
	}

	// Handle the error
	c.handler.HandleResponse(
		w, gonethttpresponse.NewDebugErrorResponse(
			gonethttperrors.InternalServerError,
			err,
			nil, nil, http.StatusInternalServerError,
		),
	)
}

// RefreshToken refreshes a user token
// @Summary Refresh a user token
// @Description Refreshes a user token
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Param request body RefreshTokenRequest
// @Success 200 {object} RefreshTokenResponse
// @Failure 401 {object} gonethttphandler.JSendResponse
// @Failure 500 {object} gonethttphandler.JSendResponse
// @Router /api/v1/auth/refresh-token [post]
func (c *Controller) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Refresh the token
	userID, userTokens, err := c.service.RefreshToken(r)
	if err == nil {
		// Log the successful token refresh
		c.logger.RefreshToken(*userID)

		// Handle the response
		c.handler.HandleResponse(
			w, gonethttpresponse.NewSuccessResponse(
				RefreshTokenResponse{
					RefreshToken: (*userTokens)[gojwttoken.RefreshToken],
					AccessToken:  (*userTokens)[gojwttoken.AccessToken],
				},
				http.StatusCreated,
			),
		)
		return
	}

	// Handle the error
	c.handler.HandleResponse(
		w, gonethttpresponse.NewDebugErrorResponse(
			gonethttperrors.InternalServerError,
			err,
			nil, nil, http.StatusInternalServerError,
		),
	)
}

// GenerateTOTPUrl generates a TOTP URL
// @Summary Generate a TOTP URL
// @Description Generates a TOTP URL
// @Tags api v1 auth
// @Accept json
// @Produce json
// @Success 200 {object} GenerateTOTPUrlResponse
// @Failure 401 {object} gonethttphandler.JSendResponse
// @Failure 500 {object} gonethttphandler.JSendResponse
// @Router /api/v1/auth/totp/generate [post]
func (c *Controller) GenerateTOTPUrl(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Generate the TOTP URL
	userID, totpUrl, err := c.service.GenerateTOTPUrl(r)
	if err == nil {
		// Log the successful TOTP URL generation
		c.logger.GenerateTOTPUrl(*userID)

		// Handle the response
		c.handler.HandleResponse(
			w, gonethttpresponse.NewSuccessResponse(
				GenerateTOTPUrlResponse{
					TOTPUrl: *totpUrl,
				},
				http.StatusCreated,
			),
		)
		return
	}

	// Handle the error
	c.handler.HandleResponse(
		w, gonethttpresponse.NewDebugErrorResponse(
			gonethttperrors.InternalServerError,
			err,
			nil, nil, http.StatusInternalServerError,
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
// @Success 200 {object} VerifyTOTPResponse
// @Failure 401 {object} gonethttphandler.JSendResponse
// @Failure 500 {object} gonethttphandler.JSendResponse
// @Router /api/v1/auth/totp/verify [post]
func (c *Controller) VerifyTOTP(w http.ResponseWriter, r *http.Request) {
	// Decode the request body and validate the request
	var body VerifyTOTPRequest
	if !c.handler.HandleRequestAndValidations(
		w,
		r,
		&body,
		c.validator.ValidateVerifyTOTPRequest,
	) {
		return
	}

	// Verify the TOTP code
	userID, recoveryCodes, err := c.service.VerifyTOTP(r, &body)
	if err == nil {
		// Log the successful TOTP verification
		c.logger.VerifyTOTP(*userID)

		// Handle the response
		c.handler.HandleResponse(
			w, gonethttpresponse.NewSuccessResponse(
				VerifyTOTPResponse{
					IsVerified:    recoveryCodes != nil,
					RecoveryCodes: *recoveryCodes,
				},
				http.StatusOK,
			),
		)
		return
	}

	// Handle the error
	anErrorOccurred := false
	data := make(map[string]*[]string)
	if errors.Is(err, ErrInvalidTOTPCode) {
		data["totp_code"] = &[]string{err.Error()}
	} else if errors.Is(
		err,
		internalapiv1common.UserTOTPSecretNotFoundByUserID,
	) {
		data["totp_secret"] = &[]string{err.Error()}
	} else {
		anErrorOccurred = true
	}

	// Check if an error occurred
	if !anErrorOccurred {
		c.handler.HandleResponse(
			w, gonethttpresponse.NewFailResponse(
				&data,
				nil,
				http.StatusUnauthorized,
			),
		)
		return
	}

	c.handler.HandleResponse(
		w, gonethttpresponse.NewDebugErrorResponse(
			gonethttperrors.InternalServerError,
			err,
			nil, nil, http.StatusInternalServerError,
		),
	)
}
