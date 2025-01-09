package v1

import (
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttphandler "github.com/ralvarezdev/go-net/http/handler"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/_common"
	internalrouterauth "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/auth"
	internalrouternote "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note"
	internalrouternotes "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/notes"
	internalroutertag "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/tag"
	internalrouteruser "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/user"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
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
		service:            &Service{},
		validator:          &Validator{Service: internalvalidator.ValidationsService},
		logger:             internallogger.Api,
		jwtValidatorLogger: internallogger.JwtValidator,
	}
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
		c.RouterWrapper,
		c.authenticator,
		c.postgresService,
		c.jwtIssuer,
	)
	noteController := internalrouternote.NewController(
		c.RouterWrapper,
		c.authenticator,
		c.postgresService,
	)
	notesController := internalrouternotes.NewController(
		c.RouterWrapper,
		c.authenticator,
		c.postgresService,
	)
	tagController := internalroutertag.NewController(
		c.RouterWrapper,
		c.authenticator,
		c.postgresService,
	)
	userController := internalrouteruser.NewController(
		c.RouterWrapper,
		c.authenticator,
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
// @Router /api/v1/ping [get]
func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) {
	// Handle the response
	c.handler.HandleResponse(
		w, gonethttpresponse.NewSuccessResponse(
			&internalapiv1common.BasicResponse{
				Message: "pong",
			}, http.StatusOK,
		),
	)
}
