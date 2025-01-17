package logger

import (
	godatabases "github.com/ralvarezdev/go-databases"
	gojwtcache "github.com/ralvarezdev/go-jwt/cache"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	goloaderenv "github.com/ralvarezdev/go-loader/env"
	gologger "github.com/ralvarezdev/go-logger"
	gologgermode "github.com/ralvarezdev/go-logger/mode"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	govalidatormapper "github.com/ralvarezdev/go-validator/struct/mapper"
	govalidatormapperparserjson "github.com/ralvarezdev/go-validator/struct/mapper/parser/json"
	govalidatormappervalidator "github.com/ralvarezdev/go-validator/struct/mapper/validator"
)

var (
	// BaseLogger is the base logger for the server
	BaseLogger = gologger.NewDefaultLogger()

	// ModeLogger is the extended logger for the server with mode support
	ModeLogger, _ = gologgermode.NewDefaultLogger(BaseLogger)

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

	// CacheTokenValidator is the logger for the cache token validator
	CacheTokenValidator, _ = gojwtcache.NewLogger(
		"CACHE TOKEN VALIDATOR",
		ModeLogger,
	)

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

	// MapperParser is the logger for the mapper parser
	MapperParser, _ = govalidatormapperparserjson.NewLogger(
		"MAPPER PARSER",
		ModeLogger,
	)
)
