package user

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	govalidatormappervalidation "github.com/ralvarezdev/go-validator/struct/mapper/validation"
	"github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
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
				internalmiddleware.AuthenticateAccessToken,
			)
		},
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterExactRoute(
				"PUT /profile",
				Controller.UpdateProfile,
				internalmiddleware.Validate(
					&UpdateProfileRequest{},
					func(
						body *UpdateProfileRequest,
						validations *govalidatormappervalidation.StructValidations,
					) {
						if body.Birthdate != nil {
							internalvalidator.Service.Birthdate(
								"birthdate",
								*body.Birthdate,
								internal.BirthdateOptions,
								validations,
							)
						}
					},
				),
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
