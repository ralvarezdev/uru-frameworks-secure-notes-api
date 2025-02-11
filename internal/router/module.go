package router

import (
	"fmt"
	gonethttp "github.com/ralvarezdev/go-net/http"
	gosecurityheadersnethttp "github.com/ralvarezdev/go-security-headers/net/http"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
	internalrouterapi "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api"
)

var (
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Pattern:    "/",
		Controller: Controller,
		BeforeLoadFn: func(m *gonethttp.Module) {
			m.Middlewares = gonethttp.NewMiddlewares(
				internalmiddleware.Limit,
				internalmiddleware.HandleError,
				gosecurityheadersnethttp.Handler,
			)
			fmt.Println("1a", m.Middlewares)
		},
		AfterLoadFn: func(m *gonethttp.Module) {
			fmt.Println("1a", m.GetRouter().GetMiddlewares())
		},
		Submodules: gonethttp.NewSubmodules(
			internalrouterapi.Module,
		),
	}
)
