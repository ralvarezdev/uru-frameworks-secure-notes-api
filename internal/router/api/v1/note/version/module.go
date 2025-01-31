package version

import (
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
	"net/http"
)

var (
	Service    = &service{}
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Path:       "/version",
		Service:    Service,
		Controller: Controller,
		Middlewares: &[]func(http.Handler) http.Handler{
			internalmiddleware.Authenticate(gojwtinterception.AccessToken),
		},
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterRoute(
				"POST /",
				Controller.CreateNoteVersion,
				internalmiddleware.Validate(
					&CreateNoteVersionRequest{},
				),
			)
			m.RegisterRoute(
				"PUT /",
				Controller.UpdateNoteVersion,
				internalmiddleware.Validate(
					&UpdateNoteVersionRequest{},
				),
			)
			m.RegisterRoute(
				"DELETE /",
				Controller.DeleteNoteVersion,
				internalmiddleware.Validate(
					&DeleteNoteVersionRequest{},
				),
			)
			m.RegisterRoute(
				"GET /",
				Controller.GetNoteVersion,
				internalmiddleware.Validate(
					&GetNoteVersionRequest{},
				),
			)
		},
	}
)
