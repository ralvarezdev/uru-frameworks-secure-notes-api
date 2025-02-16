package tag

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
)

var (
	Service    = &service{}
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Pattern:    "/tag",
		Service:    Service,
		Controller: Controller,
		BeforeLoadFn: func(m *gonethttp.Module) {
			m.Middlewares = gonethttp.NewMiddlewares(
				internalmiddleware.AuthenticateAccessToken,
			)
		},
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterExactRoute(
				"POST /",
				Controller.CreateUserTag, internalmiddleware.Validate(
					&CreateUserTagRequest{},
				),
			)
			m.RegisterExactRoute(
				"PUT /",
				Controller.UpdateUserTag,
				internalmiddleware.Validate(
					&UpdateUserTagRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /",
				Controller.DeleteUserTag,
				internalmiddleware.Validate(
					&DeleteUserTagRequest{},
				),
			)
			m.RegisterExactRoute(
				"GET /",
				Controller.GetUserTagByID,
				internalmiddleware.Validate(
					&GetUserTagByIDRequest{},
				),
			)
		},
	}
)
