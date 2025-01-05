package main

import (
	"flag"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	goloaderlistener "github.com/ralvarezdev/go-loader/http/listener"
	gonethttpjwtvalidator "github.com/ralvarezdev/go-net/http/jwt/validator"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	internalbcrypt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/bcrypt"
	internalpbkdf2 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/pbkdf2"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internaljson "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/json"
	internaljwt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt"
	internalapiv1jwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	internallistener "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/listener"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalrouter "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router"
	"log"
	"net/http"
)

// init initializes the flags and calls the load functions
func init() {
	// Parse the flags
	flag.Parse()

	// Call the load functions
	internalbcrypt.Load()
	internalpbkdf2.Load()
	internalpostgres.Load()
	internaljwt.Load()
	internallistener.Load()
}

func main() {
	// Log the mode flag
	log.Printf("Running in %s mode...\n", goflagsmode.ModeFlag.Value())

	// Create the Postgres database service
	postgresService, err := internalpostgres.NewService(
		internalpostgres.Database,
		internalpostgres.Connection,
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		internallogger.Postgres.DisconnectedFromDatabase()
		if err = postgresService.Close(); err != nil {
			panic(err)
		}
	}()
	internallogger.Postgres.ConnectedToDatabase()

	// Create the JWT claims validator
	jwtClaimsValidator, _ := internalapiv1jwtclaims.NewDefaultValidator(
		postgresService, nil,
	)

	// Create the JWT validator with ED25519 public key
	jwtValidator, err := gojwtvalidator.NewEd25519Validator(
		[]byte(internaljwt.Keys[internaljwt.EnvPublicKey]),
		jwtClaimsValidator,
		goflagsmode.ModeFlag,
	)
	if err != nil {
		panic(err)
	}

	// Create the JWT issuer with ED25519 private key
	jwtIssuer, err := gojwtissuer.NewEd25519Issuer(
		[]byte(internaljwt.Keys[internaljwt.EnvPrivateKey]),
	)
	if err != nil {
		panic(err)
	}

	// Create the JWT validator handler
	jwtValidatorErrorHandler, _ := gonethttpjwtvalidator.NewDefaultErrorHandler(internaljson.Encoder)

	// Create API authenticator middleware
	authenticator, _ := gonethttpmiddlewareauth.NewMiddleware(
		jwtValidator,
		internalhandler.Handler,
		jwtValidatorErrorHandler,
	)
	if err != nil {
		panic(err)
	}

	// Create the router controller
	routerController := internalrouter.NewController(
		authenticator,
		postgresService,
		jwtIssuer,
	)
	if err != nil {
		panic(err)
	}

	// Register the router controller routes
	routerController.RegisterRoutes()
	routerController.RegisterGroups()

	// Serve the API server
	internallogger.Listener.ServerStarted(internallistener.Port.Port)
	if err = http.ListenAndServe(
		":"+internallistener.Port.Port,
		routerController.Handler(),
	); err != nil {
		panic(goloaderlistener.ErrFailedToServe)
	}
}
