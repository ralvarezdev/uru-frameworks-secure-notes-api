package logger

import (
	godatabases "github.com/ralvarezdev/go-databases"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	goloaderenv "github.com/ralvarezdev/go-loader/env"
	goloaderlistener "github.com/ralvarezdev/go-loader/http/listener"
	gologger "github.com/ralvarezdev/go-logger"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
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

	// Router is the logger for the router
	Router, _ = gonethttproute.NewLogger(
		gologger.NewDefaultLogger(
			"ROUTER",
			nil,
		),
	)

	// Api is the logger for the API endpoints
	Api, _ = NewLogger(gologger.NewDefaultLogger("API", nil))

	// JwtValidator is the logger for the JWT validator
	JwtValidator, _ = gojwtvalidator.NewLogger(
		gologger.NewDefaultLogger(
			"JWT VALIDATOR",
			nil,
		),
	)
)
