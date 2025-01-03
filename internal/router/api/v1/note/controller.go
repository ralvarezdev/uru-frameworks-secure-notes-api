package note

import (
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttphandler "github.com/ralvarezdev/go-net/http/handler"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/database/postgres"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalrouterversion "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note/version"
	internalrouterversions "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note/versions"
)

type (
	// Controller is the structure for the API V1 note controller
	Controller struct {
		service            *Service
		validator          *Validator
		handler            gonethttphandler.Handler
		authenticator      gonethttpmiddlewareauth.Authenticator
		validatorService   govalidatorservice.Service
		postgresService    *internalpostgres.Service
		logger             *internallogger.Logger
		jwtValidatorLogger *gojwtvalidator.Logger
		gonethttproute.Controller
	}
)

// NewController creates a new API V1 note controller
func NewController(
	baseRouter gonethttproute.RouterWrapper,
	handler gonethttphandler.Handler,
	authenticator gonethttpmiddlewareauth.Authenticator,
	validatorService govalidatorservice.Service,
	postgresService *internalpostgres.Service,
) *Controller {
	return &Controller{
		Controller: gonethttproute.Controller{
			RouterWrapper: baseRouter.NewGroup(BasePath),
		},
		handler:            handler,
		authenticator:      authenticator,
		validatorService:   validatorService,
		postgresService:    postgresService,
		service:            &Service{PostgresService: postgresService},
		validator:          &Validator{Service: validatorService},
		logger:             internallogger.Api,
		jwtValidatorLogger: internallogger.JwtValidator,
	}
}

// RegisterRoutes registers the routes for the API V1 note controller
func (c *Controller) RegisterRoutes() {}

// RegisterGroups registers the router groups for the API V1 note controller
func (c *Controller) RegisterGroups() {
	// Create the controllers
	versionController := internalrouterversion.NewController(
		c.RouterWrapper,
		c.handler,
		c.authenticator,
		c.validatorService,
		c.postgresService,
	)
	versionsController := internalrouterversions.NewController(
		c.RouterWrapper,
		c.handler,
		c.authenticator,
		c.validatorService,
		c.postgresService,
	)

	// Register the controllers routes
	for _, controller := range []gonethttproute.ControllerWrapper{
		versionController,
		versionsController,
	} {
		controller.RegisterRoutes()
		controller.RegisterGroups()
	}
}
