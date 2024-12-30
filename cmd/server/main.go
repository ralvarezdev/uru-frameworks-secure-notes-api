package main

import (
	"flag"
	"github.com/joho/godotenv"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	goloaderenv "github.com/ralvarezdev/go-loader/env"
	goloaderlistener "github.com/ralvarezdev/go-loader/http/listener"
	goloadertls "github.com/ralvarezdev/go-loader/http/tls"
	gonethttpjson "github.com/ralvarezdev/go-net/http/json"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
	govalidatorvalidations "github.com/ralvarezdev/go-validator/structs/mapper/validations"
	internalapiv1 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/api/v1"
	internalpostgresql "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/database/postgresql"
	internaljwt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt"
	internaljwtvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/validator"
	internallistener "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/listener"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// Load environment variables
func init() {
	// Declare flags and parse them
	goflagsmode.SetFlag()
	flag.Parse()
	internallogger.Flag.ModeFlagSet(goflagsmode.Flag)

	// Check if the environment is production
	if goflagsmode.Flag != nil && goflagsmode.Flag.IsProd() {
		return
	}

	if err := godotenv.Load(); err != nil {
		panic(goloaderenv.FailedToLoadEnvironmentVariablesError)
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

	// Get the PostgreSQL URI
	postgresqlDbUri, err := goloaderenv.LoadVariable(internalpostgresql.UriKey)
	if err != nil {
		panic(err)
	}
	internallogger.Environment.EnvironmentVariableLoaded(internalpostgresql.UriKey)

	// Get the required PostgreSQL database name
	postgresqlDbName, err := goloaderenv.LoadVariable(internalpostgresql.DbNameKey)
	if err != nil {

		panic(err)
	}
	internallogger.Environment.EnvironmentVariableLoaded(internalpostgresql.DbNameKey)

	// Get the JWT keys
	var jwtKeys = make(map[string]string)
	for _, key := range []string{
		internaljwt.PublicKey,
		internaljwt.PrivateKey,
	} {
		jwtKey, err := commonenv.LoadVariable(key)
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
		jwtTokenDuration, err := commonenv.LoadVariable(value)
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

	// Create the PostgreSQL DSN
	postgresqlDsn := postgresqlDbUri + "/" + postgresqlDbName + "?sslmode=required"

	// Connect to PostgreSQL with GORM
	postgresqlDb, err := gorm.Open(postgres.Open(postgresqlDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Retrieve the underlying SQL database connection
	postgresqlConn, err := postgresqlDb.DB()
	if err != nil {
		panic(err)
	}
	internallogger.PostgreSQL.ConnectedToDatabase()

	// Load transport credentials
	var transportCredentials credentials.TransportCredentials

	if goflagsmode.Flag != nil && goflagsmode.Flag.IsDev() {
		transportCredentials = insecure.NewCredentials()
	} else {
		// Load system certificates
		transportCredentials, err = goloadertls.LoadSystemCredentials()
		if err != nil {
			panic(err)
		}
	}

	// Create the PostgreSQL database service
	postgresqlService, err := internalpostgresql.NewService(
		postgresqlDb,
		postgresqlConn,
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = postgresqlService.Close(); err != nil {
			panic(err)
		}
	}()

	// Create token validator
	tokenValidator, err := internaljwtvalidator.NewDefaultTokenValidator(
		postgresqlService,
		nil,
	)
	if err != nil {
		panic(err)
	}

	// Create JWT validator with ED25519 public key
	jwtValidator, err := gojwtvalidator.NewEd25519Validator(
		[]byte(jwtKeys[internaljwt.PublicKey]),
		tokenValidator,
		goflagsmode.Flag,
	)
	if err != nil {
		panic(err)
	}

	// Create the JWT issuer with ED25519 private key
	jwtIssuer, err := gojwtissuer.NewEd25519Issuer([]byte(jwtKeys[internaljwt.PrivateKey]))
	if err != nil {
		panic(err)
	}

	// Create the JSON encoder and decoder
	jsonEncoder := gonethttpjson.NewDefaultEcoder()
	jsonDecoder := gonethttpjson.NewDefaultDecoder()

	// Create API authenticator middleware
	authMiddleware, err := gonethttpmiddlewareauth.NewMiddleware(
		jwtValidator,
	)
	if err != nil {
		panic(err)
	}

	// Create the API V1 service
	apiV1Service := internalapiv1.NewService(
		jwtIssuer,
		postgresqlService,
	)

	// Create the mapper validations validator
	validationsValidator := govalidatorvalidations.NewDefaultValidator(
		goflagsmode.Flag,
	)

	// Create the mapper validations generator
	validationsGenerator := govalidatorvalidations.NewDefaultGenerator(nil)

	// Create the validator service
	validatorService, err := govalidatorservice.NewDefaultService(
		validationsGenerator,
		validationsValidator,
		goflagsmode.Flag,
	)
	if err != nil {
		panic(err)
	}

	// Create the API V1 validator
	apiV1Validator, err := internalapiv1.NewValidator(
		apiV1Service,
		validatorService,
	)
	if err != nil {
		panic(err)
	}

	// Create the base router
	baseRouter := gonethttproute.NewRouter(
		http.NewServeMux(),
	)

	// Create the API router
	apiRouter := gonethttproute.NewRouterGroup(
		"/api",
		baseRouter,
	)

	// Create the API V1 router group
	apiV1Router := gonethttproute.NewRouterGroup(
		"/v1",
		apiRouter,
	)

	// Create and initialize the API V1 controller
	apiV1Controller := internalapiv1.NewController(
		apiV1Router,
		apiV1Service,
		apiV1Validator,
		authMiddleware,
		jsonEncoder,
		jsonDecoder,
	)
	apiV1Controller.RegisterRoutes()
	apiV1Controller.RegisterRouteGroups()

	// Serve the API server
	internallogger.Listener.ServerStarted(port.Port)
	if err = http.ListenAndServe(":"+port.Port, baseRouter); err != nil {
		panic(goloaderlistener.FailedToServeError)
	}
}
