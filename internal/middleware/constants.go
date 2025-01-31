package middleware

import (
	"github.com/ralvarezdev/go-jwt/token/interception"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttpmiddlewareerrorhandler "github.com/ralvarezdev/go-net/http/middleware/error_handler"
	gonethttpmiddlewarevalidator "github.com/ralvarezdev/go-net/http/middleware/validator"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internaljwt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
	"net/http"
)

var (
	// HandleError is the API error handler middleware function
	HandleError func(next http.Handler) http.Handler

	// Authenticate is the API authenticator middleware function
	Authenticate func(interception interception.Interception) func(next http.Handler) http.Handler

	// Validate is the API request validator middleware function
	Validate func(
		body interface{},
		auxiliaryValidatorFns ...interface{},
	) func(next http.Handler) http.Handler
)

// Load loads the API middlewares
func Load() {
	// Create API error handler middleware
	errorHandler, _ := gonethttpmiddlewareerrorhandler.NewMiddleware(internalhandler.Handler)
	HandleError = errorHandler.HandleError

	// Create API authenticator middleware
	authenticator, _ := gonethttpmiddlewareauth.NewMiddleware(
		internaljwt.Validator,
		internalhandler.Handler,
		internaljwt.ValidatorFailHandler,
	)
	Authenticate = authenticator.Authenticate

	// Create API request validator middleware
	validator, _ := gonethttpmiddlewarevalidator.NewMiddleware(
		internalhandler.Handler,
		internalvalidator.Service,
		internalvalidator.JSONGenerator,
	)
	Validate = validator.Validate
}
