package version

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
	"net/http"
)

var (
	Service    = &service{}
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Pattern:    "/version",
		Service:    Service,
		Controller: Controller,
		BeforeLoadFn: func(m *gonethttp.Module) {
			m.Middlewares = &[]func(http.Handler) http.Handler{
				internalmiddleware.AuthenticateAccessToken,
			}
		},
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterExactRoute(
				"POST /",
				Controller.CreateNoteVersion,
				internalmiddleware.Validate(
					&CreateNoteVersionRequest{},
				),
			)
			m.RegisterExactRoute(
				"PUT /",
				Controller.UpdateNoteVersion,
				internalmiddleware.Validate(
					&UpdateNoteVersionRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /",
				Controller.DeleteNoteVersion,
				internalmiddleware.Validate(
					&DeleteNoteVersionRequest{},
				),
			)
			m.RegisterExactRoute(
				"GET /",
				Controller.GetNoteVersion,
				internalmiddleware.Validate(
					&GetNoteVersionRequest{},
				),
			)
		},
	}
)
