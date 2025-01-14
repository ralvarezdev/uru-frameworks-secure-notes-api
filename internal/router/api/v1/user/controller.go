package user

import (
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttphandler "github.com/ralvarezdev/go-net/http/handler"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
	"net/http"
)

type (
	// Controller is the structure for the API V1 user controller
	Controller struct {
		handler            gonethttphandler.Handler
		authenticator      gonethttpmiddlewareauth.Authenticator
		postgresService    *internalpostgres.Service
		service            *Service
		validator          *Validator
		logger             *internallogger.Logger
		jwtValidatorLogger *gojwtvalidator.Logger
		gonethttproute.Controller
	}
)

// NewController creates a new API V1 user controller
func NewController(
	baseRouter gonethttproute.RouterWrapper,
	authenticator gonethttpmiddlewareauth.Authenticator,
	postgresService *internalpostgres.Service,
) *Controller {
	// Load the validator mappers
	LoadMappers()

	return &Controller{
		Controller: gonethttproute.Controller{
			RouterWrapper: baseRouter.NewGroup(BasePath),
		},
		handler:            internalhandler.Handler,
		authenticator:      authenticator,
		postgresService:    postgresService,
		service:            &Service{PostgresService: postgresService},
		validator:          &Validator{Service: internalvalidator.ValidationsService},
		logger:             internallogger.Api,
		jwtValidatorLogger: internallogger.JwtValidator,
	}
}

// RegisterRoutes registers the routes for the API V1 user controller
func (c *Controller) RegisterRoutes() {
	c.RegisterRoute(
		"POST /signup",
		c.SignUp,
	)
}

// RegisterGroups registers the router groups for the API V1 user controller
func (c *Controller) RegisterGroups() {}

// SignUp signs up a new user
// @Summary Sign up a new user
// @Description Creates a new user account with the provided details
// @Tags api v1 user
// @Accept json
// @Produce json
// @Param request body SignUpRequest true "Sign Up Request"
// @Success 201 {object} internalapiv1common.BasicResponse
// @Failure 400 {object} gonethttpresponse.JSendResponse
// @Failure 500 {object} gonethttpresponse.JSendResponse
// @Router /api/v1/user/signup [post]
func (c *Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	// Decode the request body and validate the request
	var body SignUpRequest
	if !c.handler.DecodeAndValidate(
		w,
		r,
		&body,
		c.validator.ValidateSignUpRequest,
	) {
		return
	}

	// Sign up the user
	userID, err := c.service.SignUp(r, &body)
	if err != nil {
		c.handler.HandleError(w, err)
	}

	// Log the user sign up
	c.logger.SignUp(*userID)

	// Handle the response
	c.handler.HandleResponse(
		w, gonethttpresponse.NewSuccessResponse(
			nil, http.StatusCreated,
		),
	)
}
