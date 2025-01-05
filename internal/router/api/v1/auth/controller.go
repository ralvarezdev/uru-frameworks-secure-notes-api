package auth

import (
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttphandler "github.com/ralvarezdev/go-net/http/handler"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
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
	/*
		c.RegisterRoute(
			"POST /login",
			c.LogIn,
		)
	*/
}

// RegisterGroups registers the router groups for the API V1 auth controller
func (c *Controller) RegisterGroups() {}

// LogIn logs in a user
// @Summary Log in a user
// @Description Authenticates a user and returns a seed token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LogInRequest true "Log In Request"
// @Success 200 {object} LogInResponse
// @Failure 400 {object} gonethttphandler.ErrorResponse
// @Failure 401 {object} gonethttphandler.ErrorResponse
// @Failure 500 {object} gonethttphandler.ErrorResponse
// @Router /api/v1/auth/login [post]
/*
func (c *Controller) LogIn(w http.ResponseWriter, r *http.Request) {
	// Decode the request body and va
	var body LogInRequest
	if err := c.handler.HandleRequest(w, r, &body); err != nil {
		return
	}
}
*/
