package jwt

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	"github.com/ralvarezdev/go-jwt/token/interception"
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttpjwtvalidator "github.com/ralvarezdev/go-net/http/jwt/validator"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internaljson "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/json"
	internaljwtcache "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/cache"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	"net/http"
	"time"
)

const (
	// EnvPublicKey is the key of the JWT public key
	EnvPublicKey = "URU_FRAMEWORKS_SECURE_NOTES_JWT_PUBLIC_KEY"

	// EnvPrivateKey is the key of the JWT private key
	EnvPrivateKey = "URU_FRAMEWORKS_SECURE_NOTES_JWT_PRIVATE_KEY"

	// EnvAccessTokenDuration is the key of the access token duration
	EnvAccessTokenDuration = "URU_FRAMEWORKS_SECURE_NOTES_ACCESS_TOKEN_DURATION"

	// EnvRefreshTokenDuration is the key of the refresh token duration
	EnvRefreshTokenDuration = "URU_FRAMEWORKS_SECURE_NOTES_REFRESH_TOKEN_DURATION"
)

var (
	// Keys are the JWT keys
	Keys = make(map[string]string)

	// Durations are the JWT tokens duration
	Durations = make(map[gojwttoken.Token]time.Duration)

	// Issuer is the JWT issuer
	Issuer gojwtissuer.Issuer

	// Authenticate is the API authenticator middleware function
	Authenticate func(interception interception.Interception) func(next http.Handler) http.Handler
)

// Load loads the JWT constants
func Load() {
	// Get the JWT keys
	for _, env := range []string{
		EnvPublicKey,
		EnvPrivateKey,
	} {
		var key string
		if err := internalloader.Loader.LoadVariable(env, &key); err != nil {
			panic(err)
		}
		Keys[env] = key
	}

	// Get the JWT tokens duration
	for key, env := range map[gojwttoken.Token]string{
		gojwttoken.AccessToken:  EnvAccessTokenDuration,
		gojwttoken.RefreshToken: EnvRefreshTokenDuration,
	} {
		var tokenDuration time.Duration
		if err := internalloader.Loader.LoadDurationVariable(
			env,
			&tokenDuration,
		); err != nil {
			panic(err)
		}
		Durations[key] = tokenDuration
	}

	// Create the JWT claims validator
	claimsValidator, _ := internaljwtclaims.NewDefaultValidator(
		internalpostgres.DBService, internaljwtcache.TokenValidator,
	)

	// Create the JWT validator with ED25519 public key
	validator, err := gojwtvalidator.NewEd25519Validator(
		[]byte(Keys[EnvPublicKey]),
		claimsValidator,
		goflagsmode.ModeFlag,
	)
	if err != nil {
		panic(err)
	}

	// Create the JWT issuer with ED25519 private key
	issuer, err := gojwtissuer.NewEd25519Issuer(
		[]byte(Keys[EnvPrivateKey]),
	)
	if err != nil {
		panic(err)
	}
	Issuer = issuer

	// Create the JWT validator handler
	validatorFailHandler, _ := gonethttpjwtvalidator.NewDefaultFailHandler(internaljson.Encoder)

	// Create API authenticator middleware
	authenticator, _ := gonethttpmiddlewareauth.NewMiddleware(
		validator,
		internalhandler.Handler,
		validatorFailHandler,
	)
	Authenticate = authenticator.Authenticate
}
