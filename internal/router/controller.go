package router

import (
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttphandler "github.com/ralvarezdev/go-net/http/handler"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/database/postgres"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalrouterapi "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api"
)

type (
	// Controller is the structure for the API V1 controller
	Controller struct {
		handler            gonethttphandler.Handler
		authenticator      gonethttpmiddlewareauth.Authenticator
		validatorService   govalidatorservice.Service
		postgresService    *internalpostgres.Service
		jwtIssuer          gojwtissuer.Issuer
		logger             *internallogger.Logger
		jwtValidatorLogger *gojwtvalidator.Logger
		gonethttproute.Controller
	}
)

// NewController creates a new API V1 controller
func NewController(
	handler gonethttphandler.Handler,
	authenticator gonethttpmiddlewareauth.Authenticator,
	validatorService govalidatorservice.Service,
	postgresService *internalpostgres.Service,
	jwtIssuer gojwtissuer.Issuer,
) *Controller {
	return &Controller{
		Controller: gonethttproute.Controller{
			RouterWrapper: Router,
		},
		handler:            handler,
		authenticator:      authenticator,
		validatorService:   validatorService,
		postgresService:    postgresService,
		jwtIssuer:          jwtIssuer,
		logger:             internallogger.Api,
		jwtValidatorLogger: internallogger.JwtValidator,
	}
}

// RegisterRoutes registers the routes for the API V1 controller
func (c *Controller) RegisterRoutes() {}

// RegisterGroups registers the router groups for the API V1 controller
func (c *Controller) RegisterGroups() {
	// Create the controllers
	apiController := internalrouterapi.NewController(
		c.RouterWrapper,
		c.handler,
		c.authenticator,
		c.validatorService,
		c.postgresService,
		c.jwtIssuer,
	)

	// Register the controllers routes
	for _, controller := range []gonethttproute.ControllerWrapper{
		apiController,
	} {
		controller.RegisterRoutes()
		controller.RegisterGroups()
	}
}
