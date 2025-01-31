package versions

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
		Path:       "/versions",
		Service:    Service,
		Controller: Controller,
		Middlewares: &[]func(http.Handler) http.Handler{
			internalmiddleware.Authenticate(gojwtinterception.AccessToken),
		},
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterRoute(
				"GET /",
				Controller.ListNoteVersions,
			)
			m.RegisterRoute(
				"POST /sync",
				Controller.SyncNoteVersions,
				internalmiddleware.Validate(&SyncNoteVersionsRequest{}),
			)
		},
	}
)
