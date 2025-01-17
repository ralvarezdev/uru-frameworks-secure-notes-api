package main

import (
	"database/sql"
	"flag"
	_ "github.com/jackc/pgx/v5/stdlib"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttpjwtvalidator "github.com/ralvarezdev/go-net/http/jwt/validator"
	gonethttpmiddlewareauth "github.com/ralvarezdev/go-net/http/middleware/auth"
	internalbcrypt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/bcrypt"
	internaltotp "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/otp/totp"
	internalpbkdf2 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/pbkdf2"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	internaljson "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/json"
	internaljwt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt"
	internaljwtcache "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/cache"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	internallistener "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/listener"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalrouter "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
	"log"
	"net/http"
)

// init initializes the flags and calls the load functions
func init() {
	// Parse the flags
	flag.Parse()

	// Log the mode flag
	mode := goflagsmode.ModeFlag
	log.Printf("Running in %s mode...\n", mode.Value())

	// Call the load functions
	internalloader.Load()
	internalbcrypt.Load()
	internaltotp.Load()
	internalpbkdf2.Load()
	internalpostgres.Load()
	internaljwt.Load()
	internaljwtcache.Load(mode)
	internallistener.Load()
	internalvalidator.Load(mode)
}

func main() {
	// Connect to the database
	db, err := internalpostgres.Config.Connect(
		"pgx",
		internalpostgres.DataSourceName,
	)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			panic(err)
		}
		internallogger.Postgres.DisconnectedFromDatabase()
	}(db)
	internallogger.Postgres.ConnectedToDatabase()

	// Create the Postgres database service
	postgresService, err := internalpostgres.NewService(
		db,
	)
	if err != nil {
		panic(err)
	}

	// Create the JWT claims validator
	jwtClaimsValidator, _ := internaljwtclaims.NewDefaultValidator(
		postgresService, internaljwtcache.TokenValidator,
	)

	// Create the JWT validator with ED25519 public key
	mode := goflagsmode.ModeFlag
	jwtValidator, err := gojwtvalidator.NewEd25519Validator(
		[]byte(internaljwt.Keys[internaljwt.EnvPublicKey]),
		jwtClaimsValidator,
		mode,
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
	jwtValidatorFailHandler, _ := gonethttpjwtvalidator.NewDefaultFailHandler(internaljson.Encoder)

	// Create API authenticator middleware
	authenticator, _ := gonethttpmiddlewareauth.NewMiddleware(
		jwtValidator,
		internalhandler.Handler,
		jwtValidatorFailHandler,
	)
	if err != nil {
		panic(err)
	}

	// Create the router controller
	routerController := internalrouter.NewController(
		authenticator,
		postgresService,
		jwtIssuer,
		internaljwtcache.TokenValidator,
	)
	if err != nil {
		panic(err)
	}

	// Register the router controller routes
	routerController.RegisterRoutes()
	routerController.RegisterGroups()

	// Serve the API server
	internallogger.Api.ServerStarted(internallistener.Port)
	if err = http.ListenAndServe(
		":"+internallistener.Port,
		routerController.Handler(),
	); err != nil {
		panic(err)
	}
}
