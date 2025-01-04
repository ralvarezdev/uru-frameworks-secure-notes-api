package note

import (
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttphandler "github.com/ralvarezdev/go-net/http/handler"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalrouterversion "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note/version"
	internalrouterversions "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note/versions"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
)

type (
	// Controller is the structure for the API V1 note controller
	Controller struct {
		service            *Service
		validator          *Validator
		handler            gonethttphandler.Handler
		authenticator      gonethttpmiddlewareauth.Authenticator
		postgresService    *internalpostgres.Service
		logger             *internallogger.Logger
		jwtValidatorLogger *gojwtvalidator.Logger
		gonethttproute.Controller
	}
)

// NewController creates a new API V1 note controller
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

// RegisterRoutes registers the routes for the API V1 note controller
func (c *Controller) RegisterRoutes() {}

// RegisterGroups registers the router groups for the API V1 note controller
func (c *Controller) RegisterGroups() {
	// Create the controllers
	versionController := internalrouterversion.NewController(
		c.RouterWrapper,
		c.authenticator,
		c.postgresService,
	)
	versionsController := internalrouterversions.NewController(
		c.RouterWrapper,
		c.authenticator,
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
