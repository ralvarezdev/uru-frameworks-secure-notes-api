package user

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
)

var (
	Service    = &service{}
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Pattern:    "/user",
		Service:    Service,
		Controller: Controller,
		BeforeLoadFn: func(m *gonethttp.Module) {
			m.Middlewares = gonethttp.NewMiddlewares(
				internalmiddleware.Authenticate,
			)
		},
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterExactRoute(
				"PUT /profile",
				Controller.UpdateProfile,
				internalmiddleware.Validate(&UpdateProfileRequest{}),
			)
			m.RegisterExactRoute(
				"GET /profile",
				Controller.GetMyProfile,
			)
			m.RegisterExactRoute(
				"PUT /username",
				Controller.ChangeUsername,
				internalmiddleware.Validate(&ChangeUsernameRequest{}),
			)
			m.RegisterExactRoute(
				"DELETE /",
				Controller.DeleteUser,
				internalmiddleware.Validate(&DeleteUserRequest{}),
			)
		},
	}
)
