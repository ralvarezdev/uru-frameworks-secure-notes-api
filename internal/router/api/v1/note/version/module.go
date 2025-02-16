package version

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
)

var (
	Service    = &service{}
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Pattern:    "/version",
		Service:    Service,
		Controller: Controller,
		BeforeLoadFn: func(m *gonethttp.Module) {
			m.Middlewares = gonethttp.NewMiddlewares(internalmiddleware.AuthenticateAccessToken)
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
				Controller.GetUserNoteVersionByID,
				internalmiddleware.Validate(
					&GetUserNoteVersionByIDRequest{},
				),
			)
		},
	}
)
