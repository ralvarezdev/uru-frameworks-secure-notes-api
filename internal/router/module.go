package router

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	gosecurityheadersnethttp "github.com/ralvarezdev/go-security-headers/net/http"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
	internalrouterapi "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api"
	"net/http"
)

var (
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Pattern:    "/",
		Controller: Controller,
		BeforeLoadFn: func(m *gonethttp.Module) {
			m.Middlewares = &[]func(http.Handler) http.Handler{
				internalmiddleware.HandleError,
				gosecurityheadersnethttp.Handler,
			}
		},
		Submodules: gonethttp.NewSubmodules(
			internalrouterapi.Module,
		),
	}
)
