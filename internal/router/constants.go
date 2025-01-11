package router

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	gosecurityheadersnethttp "github.com/ralvarezdev/go-security-headers/net/http"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
)

var (
	// Router is the base router for server
	Router = gonethttproute.NewRouter(
		"",
		goflagsmode.ModeFlag,
		internallogger.Router,
		gosecurityheadersnethttp.Handler,
	)
)
