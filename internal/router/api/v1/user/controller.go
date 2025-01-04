package user

import (
	"errors"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttp "github.com/ralvarezdev/go-net/http"
	gonethttphandler "github.com/ralvarezdev/go-net/http/handler"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalcommon "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/common"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
	"net/http"
	"strconv"
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
// @Tags User
// @Accept json
// @Produce json
// @Param request body SignUpRequest true "Sign Up Request"
// @Success 201 {object} internalcommon.BasicResponse
// @Failure 400 {object} gonethttpresponse.JSONErrorResponse
// @Failure 500 {object} gonethttpresponse.JSONErrorResponse
// @Router /api/v1/user/signup [post]
func (c *Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	// Decode the request body
	var body SignUpRequest
	if err := c.handler.HandleRequest(w, r, &body); err != nil {
		return
	}

	// Validate the request body
	if err := c.validator.ValidateSignUpRequest(&body); err != nil {
		c.handler.HandleResponse(
			w,
			gonethttpresponse.NewErrorResponseWithCode(
				err,
				http.StatusBadRequest,
			),
		)
		return
	}

	// Sign up the user
	user, err := c.service.SignUp(&body)
	if err != nil && !errors.Is(
		err,
		ErrEmailAlreadyRegistered,
	) && !errors.Is(err, ErrUsernameAlreadyRegistered) {
		c.handler.HandleResponse(
			w,
			gonethttpresponse.NewDebugErrorResponseWithCode(
				errors.New(gonethttp.InternalServerError),
				err,
				http.StatusInternalServerError,
			),
		)
		return
	}

	// Check if the email or username is already registered
	if err != nil {
		c.handler.HandleResponse(
			w,
			gonethttpresponse.NewErrorResponseWithCode(
				err,
				http.StatusBadRequest,
			),
		)
		return
	}

	// Log the user sign up
	c.logger.SignUp(strconv.Itoa(int(user.ID)))

	// Handle the response
	c.handler.HandleResponse(
		w, gonethttpresponse.NewResponseWithCode(
			&internalcommon.BasicResponse{
				Message: SignUpSuccess,
			}, http.StatusCreated,
		),
	)
}
