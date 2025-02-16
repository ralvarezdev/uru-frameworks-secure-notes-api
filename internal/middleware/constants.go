package middleware

import (
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttpmiddlewareerrorhandler "github.com/ralvarezdev/go-net/http/middleware/error_handler"
	gonethttpmiddlewaresizelimiter "github.com/ralvarezdev/go-net/http/middleware/size_limiter"
	gonethttpmiddlewarevalidator "github.com/ralvarezdev/go-net/http/middleware/validator"
	internalcookie "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/cookie"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internaljwt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
	"net/http"
)

var (
	// EnvBodyLimit is the environment variable key for the body limit
	EnvBodyLimit = "URU_FRAMEWORKS_SECURE_NOTES_BODY_LIMIT"
)

var (
	// BodyLimit is the API body limit
	BodyLimit int

	// HandleError is the API error handler middleware function
	HandleError func(next http.Handler) http.Handler

	// Limit is the API body limit middleware function
	Limit func(next http.Handler) http.Handler

	// Authenticate is the API authenticator middleware function
	Authenticate func(token gojwttoken.Token) func(next http.Handler) http.Handler

	// AuthenticateAccessToken is the API authenticator middleware function for access tokens
	AuthenticateAccessToken func(next http.Handler) http.Handler

	// AuthenticateRefreshToken is the API authenticator middleware function for refresh tokens
	AuthenticateRefreshToken func(next http.Handler) http.Handler

	// Validate is the API request validator middleware function
	Validate func(
		body interface{},
		auxiliaryValidatorFns ...interface{},
	) func(next http.Handler) http.Handler
)

// Load loads the API middlewares
func Load() {
	// Load the body limit
	err := internalloader.Loader.LoadIntVariable(EnvBodyLimit, &BodyLimit)
	if err != nil {
		panic(err)
	}

	// Create API error handler middleware
	errorHandler, _ := gonethttpmiddlewareerrorhandler.NewMiddleware(internalhandler.Handler)
	HandleError = errorHandler.HandleError

	// Create API body limit middleware
	sizeLimiter := gonethttpmiddlewaresizelimiter.NewMiddleware()
	Limit = sizeLimiter.Limit(int64(BodyLimit))

	// Create API authenticator middleware
	authenticator, _ := gonethttpmiddlewareauth.NewMiddleware(
		internaljwt.Validator,
		internalhandler.Handler,
	)
	Authenticate = func(token gojwttoken.Token) func(next http.Handler) http.Handler {
		return authenticator.AuthenticateFromCookie(
			token,
			internalcookie.RefreshToken.Name,
			internalcookie.AccessToken.Name,
			func(
				w http.ResponseWriter,
				r *http.Request,
			) (*map[gojwttoken.Token]string, error) {
				_, rawTokens := internalcookie.RefreshTokenFn(token)(
					w,
					r,
				)
				return rawTokens, nil
			},
		)
	}
	AuthenticateAccessToken = Authenticate(gojwttoken.AccessToken)
	AuthenticateRefreshToken = Authenticate(gojwttoken.RefreshToken)

	// Create API request validator middleware
	validator, _ := gonethttpmiddlewarevalidator.NewMiddleware(
		internalhandler.Handler,
		internalvalidator.Service,
		internalvalidator.JSONGenerator,
	)
	Validate = validator.Validate
}
