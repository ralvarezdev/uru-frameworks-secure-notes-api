package logger

import (
	godatabases "github.com/ralvarezdev/go-databases"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	goloaderenv "github.com/ralvarezdev/go-loader/env"
	goloaderlistener "github.com/ralvarezdev/go-loader/http/listener"
	gologger "github.com/ralvarezdev/go-logger"
	gologgermode "github.com/ralvarezdev/go-logger/mode"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	govalidatormapper "github.com/ralvarezdev/go-validator/structs/mapper"
	govalidatormappervalidator "github.com/ralvarezdev/go-validator/structs/mapper/validator"
)

var (
	// BaseLogger is the base logger for the server
	BaseLogger = gologger.NewDefaultLogger()

	// ModeLogger is the extended logger for the server with mode support
	ModeLogger, _ = gologgermode.NewDefaultLogger(BaseLogger)

	// Listener is the logger for the listener
	Listener, _ = goloaderlistener.NewLogger(
		"NET LISTENER",
		ModeLogger,
	)

	// Environment is the logger for the environment
	Environment, _ = goloaderenv.NewLogger(
		"ENV",
		ModeLogger,
	)

	// Postgres is the logger for the Postgres client
	Postgres, _ = godatabases.NewLogger(
		"POSTGRES",
		ModeLogger,
	)

	// Router is the logger for the router
	Router, _ = gonethttproute.NewLogger(
		"ROUTER",
		ModeLogger,
	)

	// Api is the logger for the API endpoints
	Api, _ = NewLogger("API", ModeLogger)

	// JwtValidator is the logger for the JWT validator
	JwtValidator, _ = gojwtvalidator.NewLogger(
		"JWT VALIDATOR",
		ModeLogger,
	)

	// MapperGenerator is the logger for the mapper generator
	MapperGenerator, _ = govalidatormapper.NewLogger(
		"MAPPER GENERATOR",
		ModeLogger,
	)

	// MapperValidator is the logger for the mapper validator
	MapperValidator, _ = govalidatormappervalidator.NewLogger(
		"MAPPER VALIDATOR",
		ModeLogger,
	)
)
