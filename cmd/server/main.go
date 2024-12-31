package main

import (
	"flag"
	"github.com/joho/godotenv"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	goloaderenv "github.com/ralvarezdev/go-loader/env"
	goloaderlistener "github.com/ralvarezdev/go-loader/http/listener"
	gonethttpjson "github.com/ralvarezdev/go-net/http/json"
	gonethttpjwtvalidator "github.com/ralvarezdev/go-net/http/jwt/validator"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
	govalidatorvalidations "github.com/ralvarezdev/go-validator/structs/mapper/validations"
	internalapiv1 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/api/v1"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/database/postgres"
	internaljwt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt"
	internalapiv1jwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	internallistener "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/listener"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// Load environment variables
func init() {
	// Declare flags and parse them
	goflagsmode.SetFlag()
	flag.Parse()
	internallogger.Flag.ModeFlagSet(goflagsmode.Mode)

	// Check if the environment is production
	if goflagsmode.Mode != nil && goflagsmode.Mode.IsProd() {
		return
	}

	if err := godotenv.Load(); err != nil {
		panic(goloaderenv.ErrFailedToLoadEnvironmentVariables)
	}
}

func main() {
	// Get the port listener
	port, err := goloaderlistener.LoadPort(
		"0.0.0.0",
		internallistener.PortKey,
	)
	if err != nil {
		panic(err)
	}
	internallogger.Environment.EnvironmentVariableLoaded(internallistener.PortKey)

	// Get the Postgres URI
	postgresqlDbUri, err := goloaderenv.LoadVariable(internalpostgres.UriKey)
	if err != nil {
		panic(err)
	}
	internallogger.Environment.EnvironmentVariableLoaded(internalpostgres.UriKey)

	// Get the required Postgres database name
	postgresqlDbName, err := goloaderenv.LoadVariable(internalpostgres.DbNameKey)
	if err != nil {

		panic(err)
	}
	internallogger.Environment.EnvironmentVariableLoaded(internalpostgres.DbNameKey)

	// Get the JWT keys
	var jwtKeys = make(map[string]string)
	for _, key := range []string{
		internaljwt.PublicKey,
		internaljwt.PrivateKey,
	} {
		jwtKey, err := goloaderenv.LoadVariable(key)
		if err != nil {
			panic(err)
		}
		internallogger.Environment.EnvironmentVariableLoaded(key)
		jwtKeys[key] = jwtKey
	}

	// Get the JWT tokens duration
	var jwtTokensDuration = make(map[string]time.Duration)
	for key, value := range map[string]string{
		internaljwt.AccessToken:  internaljwt.AccessTokenDuration,
		internaljwt.RefreshToken: internaljwt.RefreshTokenDuration,
	} {
		jwtTokenDuration, err := goloaderenv.LoadVariable(value)
		if err != nil {
			panic(err)
		}
		internallogger.Environment.EnvironmentVariableLoaded(value)

		// Parse the duration
		parsedJwtTokenDuration, err := time.ParseDuration(jwtTokenDuration)
		if err != nil {
			panic(err)
		}
		jwtTokensDuration[key] = parsedJwtTokenDuration
	}

	// Create the Postgres DSN
	postgresqlDsn := postgresqlDbUri + "/" + postgresqlDbName + "?sslmode=require"

	// Connect to Postgres with GORM
	postgresqlDb, err := gorm.Open(postgres.Open(postgresqlDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Retrieve the underlying SQL database connection
	postgresqlConn, err := postgresqlDb.DB()
	if err != nil {
		panic(err)
	}
	internallogger.Postgres.ConnectedToDatabase()

	// Create the Postgres database service
	postgresService, err := internalpostgres.NewService(
		postgresqlDb,
		postgresqlConn,
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = postgresService.Close(); err != nil {
			panic(err)
		}
	}()

	// Create the JWT claims validator
	jwtClaimsValidator, _ := internalapiv1jwtclaims.NewDefaultValidator(
		postgresService, nil,
	)

	// Create the JWT validator with ED25519 public key
	jwtValidator, err := gojwtvalidator.NewEd25519Validator(
		[]byte(jwtKeys[internaljwt.PublicKey]),
		jwtClaimsValidator,
		goflagsmode.Mode,
	)
	if err != nil {
		panic(err)
	}

	// Create the JWT issuer with ED25519 private key
	jwtIssuer, err := gojwtissuer.NewEd25519Issuer([]byte(jwtKeys[internaljwt.PrivateKey]))
	if err != nil {
		panic(err)
	}

	// Create the API V1 service
	apiV1Service, _ := internalapiv1.NewService(
		jwtIssuer,
		postgresService,
	)

	// Create the JSON encoder and decoder
	jsonEncoder := gonethttpjson.NewDefaultEncoder(goflagsmode.Mode)
	jsonDecoder := gonethttpjson.NewDefaultDecoder(goflagsmode.Mode)

	// Create the response handler
	responseHandler, _ := gonethttpresponse.NewDefaultHandler(
		goflagsmode.Mode,
		jsonEncoder,
	)

	// Create the JWT validator handler
	jwtValidatorErrorHandler, _ := gonethttpjwtvalidator.NewDefaultErrorHandler(jsonEncoder)

	// Create API authenticator middleware
	authenticator, _ := gonethttpmiddlewareauth.NewMiddleware(
		jwtValidator,
		responseHandler,
		jwtValidatorErrorHandler,
	)
	if err != nil {
		panic(err)
	}

	// Create the mapper validations validator
	validationsValidator := govalidatorvalidations.NewDefaultValidator(
		goflagsmode.Mode,
	)

	// Create the mapper validations generator
	validationsGenerator := govalidatorvalidations.NewDefaultGenerator(nil)

	// Create the validator service
	validatorService, err := govalidatorservice.NewDefaultService(
		validationsGenerator,
		validationsValidator,
		goflagsmode.Mode,
	)
	if err != nil {
		panic(err)
	}

	// Create the API V1 validator
	apiV1Validator := internalapiv1.NewValidator(
		apiV1Service,
		validatorService,
	)

	// Create the base router
	baseRouter := gonethttproute.NewRouter()

	// Create the API router
	apiRouter, _ := gonethttproute.NewRouterGroup(
		baseRouter,
		"/api",
	)

	// Create the API V1 router group
	apiV1Router, _ := gonethttproute.NewRouterGroup(
		apiRouter,
		"/v1",
	)

	// Create the API V1 controller
	apiV1Controller, err := internalapiv1.NewController(
		apiV1Router,
		apiV1Service,
		apiV1Validator,
		responseHandler,
		authenticator,
		jsonEncoder,
		jsonDecoder,
		internallogger.ApiV1,
		internallogger.JwtValidator,
	)
	if err != nil {
		panic(err)
	}

	// Register the API V1 controller routes
	apiV1Controller.RegisterRoutes()
	apiV1Controller.RegisterRouteGroups()

	// Serve the API server
	internallogger.Listener.ServerStarted(port.Port)
	if err = http.ListenAndServe(
		":"+port.Port,
		baseRouter.Handler(),
	); err != nil {
		panic(goloaderlistener.ErrFailedToServe)
	}
}
