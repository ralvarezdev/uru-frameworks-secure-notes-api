package v1

import (
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttphandler "github.com/ralvarezdev/go-net/http/handler"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/database/postgres"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalrouterauth "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/auth"
	internalrouternote "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note"
	internalrouternotes "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/notes"
	internalroutertag "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/tag"
	internalrouteruser "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/user"
	"net/http"
)

type (
	// Controller is the structure for the API V1 controller
	Controller struct {
		handler            gonethttphandler.Handler
		authenticator      gonethttpmiddlewareauth.Authenticator
		validatorService   govalidatorservice.Service
		postgresService    *internalpostgres.Service
		jwtIssuer          gojwtissuer.Issuer
		service            *Service
		validator          *Validator
		logger             *internallogger.Logger
		jwtValidatorLogger *gojwtvalidator.Logger
		gonethttproute.Controller
	}
)

// NewController creates a new API V1 controller
func NewController(
	baseRouter gonethttproute.RouterWrapper,
	handler gonethttphandler.Handler,
	authenticator gonethttpmiddlewareauth.Authenticator,
	validatorService govalidatorservice.Service,
	postgresService *internalpostgres.Service,
	jwtIssuer gojwtissuer.Issuer,
) (*Controller, error) {
	return &Controller{
		Controller: gonethttproute.Controller{
			RouterWrapper: baseRouter.NewGroup(BasePath),
		},

		handler:            handler,
		authenticator:      authenticator,
		validatorService:   validatorService,
		postgresService:    postgresService,
		jwtIssuer:          jwtIssuer,
		service:            &Service{},
		validator:          &Validator{Service: validatorService},
		logger:             internallogger.Api,
		jwtValidatorLogger: internallogger.JwtValidator,
	}, nil
}

// RegisterRoutes registers the routes for the API V1 controller
func (c *Controller) RegisterRoutes() {
	c.RegisterRoute(
		"GET /ping",
		c.Ping,
	)
}

// RegisterGroups registers the router groups for the API V1 controller
func (c *Controller) RegisterGroups() {
	// Create the controllers
	authController := internalrouterauth.NewController(
		c.NewGroup(internalrouterauth.BasePath),
		c.handler,
		c.authenticator,
		c.validatorService,
		c.postgresService,
		c.jwtIssuer,
	)
	noteController := internalrouternote.NewController(
		c.NewGroup(internalrouternote.BasePath),
		c.handler,
		c.authenticator,
		c.validatorService,
		c.postgresService,
	)
	notesController := internalrouternotes.NewController(
		c.NewGroup(internalrouternotes.BasePath),
		c.handler,
		c.authenticator,
		c.validatorService,
		c.postgresService,
	)
	tagController := internalroutertag.NewController(
		c.NewGroup(internalroutertag.BasePath),
		c.handler,
		c.authenticator,
		c.validatorService,
		c.postgresService,
	)
	userController := internalrouteruser.NewController(
		c.NewGroup(internalrouteruser.BasePath),
		c.handler,
		c.authenticator,
		c.validatorService,
		c.postgresService,
	)

	// Register the controllers routes
	for _, controller := range []gonethttproute.ControllerWrapper{
		authController,
		noteController,
		notesController,
		tagController,
		userController,
	} {
		controller.RegisterRoutes()
		controller.RegisterGroups()
	}
}

// Ping pings the service
// @Summary Ping the service
// @Description Returns a pong response to check if the service is running
// @Tags v1
// @Accept json
// @Produce json
// @Success 200 {object} BasicResponse
// @Router /ping [get]
func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) {
	// Get the ping response
	response := c.service.Ping()

	// Handle the success response
	c.handler.HandleSuccessResponse(w, response)
}
