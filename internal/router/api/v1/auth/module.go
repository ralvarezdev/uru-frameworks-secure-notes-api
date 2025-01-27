package auth

import (
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
)

var (
	Service    = &service{}
	Validator  = &validator{}
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Path:       "/auth",
		Service:    Service,
		Validator:  Validator,
		Controller: Controller,
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterRoute(
				"POST /signup",
				Controller.SignUp,
				internalmiddleware.Validate(
					&SignUpRequest{},
					Validator.SignUp,
				),
			)
			m.RegisterRoute(
				"POST /login",
				Controller.LogIn,
				internalmiddleware.Validate(
					&LogInRequest{},
					Validator.LogIn,
				),
			)
			m.RegisterRoute(
				"POST /refresh-token",
				Controller.RefreshToken,
				internalmiddleware.Authenticate(gojwtinterception.RefreshToken),
			)
			m.RegisterRoute(
				"POST /logout",
				Controller.LogOut,
				internalmiddleware.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"GET /refresh-token/{id}",
				Controller.GetRefreshToken,
				internalmiddleware.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"GET /refresh-tokens",
				Controller.ListRefreshTokens,
				internalmiddleware.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"DELETE /refresh-token/{id}",
				Controller.RevokeRefreshToken,
				internalmiddleware.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"DELETE /refresh-tokens",
				Controller.RevokeRefreshTokens,
				internalmiddleware.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"POST /totp/generate",
				Controller.GenerateTOTPUrl,
				internalmiddleware.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"POST /totp/verify",
				Controller.VerifyTOTP,
				internalmiddleware.Authenticate(gojwtinterception.AccessToken),
				internalmiddleware.Validate(
					&VerifyTOTPRequest{},
					Validator.VerifyTOTP,
				),
			)
			m.RegisterRoute(
				"DELETE /totp",
				Controller.RevokeTOTP,
				internalmiddleware.Authenticate(gojwtinterception.AccessToken),
			)
		},
	}
)
