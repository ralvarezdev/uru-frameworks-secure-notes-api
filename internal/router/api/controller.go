package api

import (
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttphandler "github.com/ralvarezdev/go-net/http/handler"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalrouterv1 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1"
)

type (
	// Controller is the structure for the API controller
	Controller struct {
		handler            gonethttphandler.Handler
		authenticator      gonethttpmiddlewareauth.Authenticator
		postgresService    *internalpostgres.Service
		jwtIssuer          gojwtissuer.Issuer
		logger             *internallogger.Logger
		jwtValidatorLogger *gojwtvalidator.Logger
		gonethttproute.Controller
	}
)

// NewController creates a new API controller
func NewController(
	baseRouter gonethttproute.RouterWrapper,
	authenticator gonethttpmiddlewareauth.Authenticator,
	postgresService *internalpostgres.Service,
	jwtIssuer gojwtissuer.Issuer,
) *Controller {
	return &Controller{
		Controller: gonethttproute.Controller{
			RouterWrapper: baseRouter.NewGroup(BasePath),
		},
		handler:            internalhandler.Handler,
		authenticator:      authenticator,
		postgresService:    postgresService,
		jwtIssuer:          jwtIssuer,
		logger:             internallogger.Api,
		jwtValidatorLogger: internallogger.JwtValidator,
	}
}

// RegisterRoutes registers the routes for the API controller
func (c *Controller) RegisterRoutes() {}

// RegisterGroups registers the router groups for the API controller
func (c *Controller) RegisterGroups() {
	// Create the controllers
	v1Controller := internalrouterv1.NewController(
		c.RouterWrapper,
		c.authenticator,
		c.postgresService,
		c.jwtIssuer,
	)

	// Register the controllers routes
	for _, controller := range []gonethttproute.ControllerWrapper{
		v1Controller,
	} {
		controller.RegisterRoutes()
		controller.RegisterGroups()
	}
}
