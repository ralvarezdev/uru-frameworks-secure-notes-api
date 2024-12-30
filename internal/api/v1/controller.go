package v1

import (
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttpjson "github.com/ralvarezdev/go-net/http/json"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
)

type (
	// Controller is the structure for the API V1 controller
	Controller struct {
		authMiddleware   gonethttpmiddlewareauth.Middleware
		logger           Logger
		authLogger       gojwtvalidator.Logger
		validatorService govalidatorservice.Service
		jsonEncoder      gonethttpjson.Encoder
		jsonDecoder      gonethttpjson.Decoder
		gonethttproute.Controller
	}
)

// NewController creates a new API V1 controller
func NewController(
	routeGroup gonethttproute.RouterWrapper,
	service Service,
	authMiddleware gonethttpmiddlewareauth.Middleware,
	validatorService govalidatorservice.Service,
	jsonEncoder gonethttpjson.Encoder,
	jsonDecoder gonethttpjson.Decoder,
) (*Controller, error) {
	return &Controller{
		Controller: gonethttproute.Controller{
			Service:       service,
			RouterWrapper: routeGroup,
		},
		authMiddleware:   authMiddleware,
		validatorService: validatorService,
		jsonEncoder:      jsonEncoder,
		jsonDecoder:      jsonDecoder,
		internallogger.ApiV1,
		internallogger.JwtValidator,
	}, nil
}

// RegisterRoutes registers the routes for the API V1 controller
func (c *Controller) RegisterRoutes() {}

// RegisterRouteGroups registers the route groups for the API V1 controller
func (c *Controller) RegisterRouteGroups() {}
