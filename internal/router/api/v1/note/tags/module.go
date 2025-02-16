package tags

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
)

var (
	Service    = &service{}
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Pattern:    "/tags",
		Service:    Service,
		Controller: Controller,
		BeforeLoadFn: func(m *gonethttp.Module) {
			m.Middlewares = gonethttp.NewMiddlewares(
				internalmiddleware.AuthenticateAccessToken,
			)
		},
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterExactRoute(
				"PATCH /tags",
				Controller.AddUserNoteTags,
				internalmiddleware.Validate(
					&AddUserNoteTagsRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /tags",
				Controller.RemoveUserNoteTags,
				internalmiddleware.Validate(
					&RemoveUserNoteTagsRequest{},
				),
			)
			m.RegisterExactRoute(
				"GET /tags",
				Controller.ListUserNoteTags,
				internalmiddleware.Validate(
					&ListUserNoteTagsRequest{},
				),
			)
		},
	}
)
