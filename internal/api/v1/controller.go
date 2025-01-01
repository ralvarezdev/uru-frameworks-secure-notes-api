package v1

import (
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttpjson "github.com/ralvarezdev/go-net/http/json"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
	"net/http"
)

type (
	// Controller is the structure for the API V1 controller
	Controller struct {
		service            *Service
		validatorService   govalidatorservice.Service
		responseHandler    gonethttpresponse.Handler
		authenticator      gonethttpmiddlewareauth.Authenticator
		jsonEncoder        gonethttpjson.Encoder
		jsonDecoder        gonethttpjson.Decoder
		logger             *Logger
		jwtValidatorLogger *gojwtvalidator.Logger
		gonethttproute.Controller
	}
)

// NewController creates a new API V1 controller
func NewController(
	routeGroup gonethttproute.RouterWrapper,
	service *Service,
	validatorService govalidatorservice.Service,
	responseHandler gonethttpresponse.Handler,
	authenticator gonethttpmiddlewareauth.Authenticator,
	jsonEncoder gonethttpjson.Encoder,
	jsonDecoder gonethttpjson.Decoder,
	logger *Logger,
	jwtValidatorLogger *gojwtvalidator.Logger,
) (*Controller, error) {
	return &Controller{
		Controller: gonethttproute.Controller{
			RouterWrapper: routeGroup,
		},
		service:            service,
		responseHandler:    responseHandler,
		authenticator:      authenticator,
		validatorService:   validatorService,
		jsonEncoder:        jsonEncoder,
		jsonDecoder:        jsonDecoder,
		logger:             logger,
		jwtValidatorLogger: jwtValidatorLogger,
	}, nil
}

// RegisterRoutes registers the routes for the API V1 controller
func (c *Controller) RegisterRoutes() {
	c.RegisterRoute(
		"GET /ping",
		c.Ping,
	)
}

// RegisterRouteGroups registers the route groups for the API V1 controller
func (c *Controller) RegisterRouteGroups() {}

// Ping pings the service
func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) {
	// Get the ping response
	response := c.service.Ping()

	// Handle the success response
	c.responseHandler.HandleSuccess(w, response)
}
