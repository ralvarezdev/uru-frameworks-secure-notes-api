package user

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
	"net/http"
)

var (
	Service    = &service{}
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Pattern:    "/user",
		Service:    Service,
		Controller: Controller,
		BeforeLoadFn: func(m *gonethttp.Module) {
			m.Middlewares = &[]func(http.Handler) http.Handler{
				internalmiddleware.AuthenticateAccessToken,
			}
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
