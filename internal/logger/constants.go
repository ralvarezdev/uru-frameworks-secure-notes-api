package logger

import (
	godatabases "github.com/ralvarezdev/go-databases"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	goloaderenv "github.com/ralvarezdev/go-loader/env"
	goloaderlistener "github.com/ralvarezdev/go-loader/http/listener"
	gologger "github.com/ralvarezdev/go-logger"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	internalapiv1 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/api/v1"
)

var (
	// Flag is the logger for the flag
	Flag, _ = goflagsmode.NewLogger(gologger.NewDefaultLogger("FLAG", nil))

	// Listener is the logger for the listener
	Listener, _ = goloaderlistener.NewLogger(
		gologger.NewDefaultLogger(
			"NET LISTENER",
			nil,
		),
	)

	// Environment is the logger for the environment
	Environment, _ = goloaderenv.NewLogger(
		gologger.NewDefaultLogger(
			"ENV",
			nil,
		),
	)

	// Postgres is the logger for the Postgres client
	Postgres, _ = godatabases.NewLogger(
		gologger.NewDefaultLogger(
			"POSTGRES",
			nil,
		),
	)

	// Route is the logger for the router
	Route, _ = gonethttproute.NewLogger(
		gologger.NewDefaultLogger(
			"ROUTER",
			nil,
		),
	)

	// ApiV1 is the logger for the API V1 endpoints
	ApiV1, _ = internalapiv1.NewLogger(gologger.NewDefaultLogger("API V1", nil))

	// JwtValidator is the logger for the JWT validator
	JwtValidator, _ = gojwtvalidator.NewLogger(
		gologger.NewDefaultLogger(
			"JWT VALIDATOR",
			nil,
		),
	)
)
