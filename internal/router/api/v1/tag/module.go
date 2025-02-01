package tag

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
	"net/http"
)

var (
	Service    = &service{}
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Pattern:    "/tag",
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
				Controller.CreateTag, internalmiddleware.Validate(
					&CreateTagRequest{},
				),
			)
			m.RegisterExactRoute(
				"PUT /",
				Controller.UpdateTag,
				internalmiddleware.Validate(
					&UpdateTagRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /",
				Controller.DeleteTag,
				internalmiddleware.Validate(
					&DeleteTagRequest{},
				),
			)
			m.RegisterExactRoute(
				"GET /",
				Controller.GetTag,
				internalmiddleware.Validate(
					&GetTagRequest{},
				),
			)
		},
	}
)
