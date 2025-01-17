package router

import (
	gojwtcache "github.com/ralvarezdev/go-jwt/cache"
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttphandler "github.com/ralvarezdev/go-net/http/handler"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalrouterapi "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api"
)

type (
	// Controller is the structure for the API V1 controller
	Controller struct {
		handler            gonethttphandler.Handler
		authenticator      gonethttpmiddlewareauth.Authenticator
		postgresService    *internalpostgres.Service
		jwtIssuer          gojwtissuer.Issuer
		jwtTokenValidator  gojwtcache.TokenValidator
		logger             *internallogger.Logger
		jwtValidatorLogger *gojwtvalidator.Logger
		gonethttproute.Controller
	}
)

// NewController creates a new API V1 controller
func NewController(
	authenticator gonethttpmiddlewareauth.Authenticator,
	postgresService *internalpostgres.Service,
	jwtIssuer gojwtissuer.Issuer,
	jwtTokenValidator gojwtcache.TokenValidator,
) *Controller {
	return &Controller{
		Controller: gonethttproute.Controller{
			RouterWrapper: Router,
		},
		handler:            internalhandler.Handler,
		authenticator:      authenticator,
		postgresService:    postgresService,
		jwtIssuer:          jwtIssuer,
		jwtTokenValidator:  jwtTokenValidator,
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
		c.authenticator,
		c.postgresService,
		c.jwtIssuer,
		c.jwtTokenValidator,
	)

	// Register the controllers routes
	for _, controller := range []gonethttproute.ControllerWrapper{
		apiController,
	} {
		controller.RegisterRoutes()
		controller.RegisterGroups()
	}
}
