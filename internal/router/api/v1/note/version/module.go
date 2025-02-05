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
				Controller.CreateUserNoteVersion,
				internalmiddleware.Validate(
					&CreateUserNoteVersionRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /",
				Controller.DeleteUserNoteVersion,
				internalmiddleware.Validate(
					&DeleteUserNoteVersionRequest{},
				),
			)
			m.RegisterExactRoute(
				"GET /",
				Controller.GetUserNoteVersionByNoteVersionID,
				internalmiddleware.Validate(
					&GetUserNoteVersionByNoteVersionIDRequest{},
				),
			)
		},
	}
)
