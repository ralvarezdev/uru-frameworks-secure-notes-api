package user

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
		Path:       "/user",
		Service:    Service,
		Controller: Controller,
		Middlewares: &[]func(http.Handler) http.Handler{
			internalmiddleware.Authenticate(gojwtinterception.AccessToken),
		},
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterRoute(
				"PUT /profile",
				Controller.UpdateProfile,
				internalmiddleware.Validate(&UpdateProfileRequest{}),
			)
			m.RegisterRoute(
				"GET /profile",
				Controller.GetMyProfile,
			)
			m.RegisterRoute(
				"PUT /username",
				Controller.ChangeUsername,
				internalmiddleware.Validate(&ChangeUsernameRequest{}),
			)
			m.RegisterRoute(
				"PUT /email",
				Controller.ChangeEmail,
				internalmiddleware.Validate(&ChangeEmailRequest{}),
			)
			m.RegisterRoute(
				"POST /email/send-verification",
				Controller.SendEmailVerificationToken,
			)
			m.RegisterRoute(
				"POST /email/verify",
				Controller.VerifyEmail,
				internalmiddleware.Validate(&VerifyEmailRequest{}),
			)
			m.RegisterRoute(
				"PUT /phone-number",
				Controller.ChangePhoneNumber,
				internalmiddleware.Validate(&ChangePhoneNumberRequest{}),
			)
			m.RegisterRoute(
				"POST /phone-number/send-verification",
				Controller.SendPhoneNumberVerificationCode,
			)
			m.RegisterRoute(
				"POST /phone-number/verify",
				Controller.VerifyPhoneNumber,
				internalmiddleware.Validate(&VerifyPhoneNumberRequest{}),
			)
			m.RegisterRoute(
				"DELETE /",
				Controller.DeleteUser,
				internalmiddleware.Validate(&DeleteUserRequest{}),
			)
		},
	}
)
