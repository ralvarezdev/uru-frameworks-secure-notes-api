package tag

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
		Path:       "/tag",
		Service:    Service,
		Controller: Controller,
		Middlewares: &[]func(http.Handler) http.Handler{
			internalmiddleware.Authenticate(gojwtinterception.AccessToken),
		},
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterRoute(
				"POST /",
				Controller.CreateTag, internalmiddleware.Validate(
					&CreateTagRequest{},
				),
			)
			m.RegisterRoute(
				"PUT /",
				Controller.UpdateTag,
				internalmiddleware.Validate(
					&UpdateTagRequest{},
				),
			)
			m.RegisterRoute(
				"DELETE /",
				Controller.DeleteTag,
				internalmiddleware.Validate(
					&DeleteTagRequest{},
				),
			)
			m.RegisterRoute(
				"GET /",
				Controller.GetTag,
				internalmiddleware.Validate(
					&GetTagRequest{},
				),
			)
		},
	}
)
