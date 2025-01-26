package auth

import (
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
	gonethttp "github.com/ralvarezdev/go-net/http"
	internaljwt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
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
		LoadFn: func(m *gonethttp.Module) {
			LogInRequestMapper, _ = internalvalidator.JSONGenerator.NewMapper(&LogInRequest{})
			VerifyTOTPRequestMapper, _ = internalvalidator.JSONGenerator.NewMapper(&VerifyTOTPRequest{})
			SignUpRequestMapper, _ = internalvalidator.JSONGenerator.NewMapper(&SignUpRequest{})
		},
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterRoute(
				"POST /signup",
				Controller.SignUp,
			)
			m.RegisterRoute(
				"POST /login",
				Controller.LogIn,
			)
			m.RegisterRoute(
				"POST /refresh-token",
				Controller.RefreshToken,
				internaljwt.Authenticate(gojwtinterception.RefreshToken),
			)
			m.RegisterRoute(
				"POST /logout",
				Controller.LogOut,
				internaljwt.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"GET /refresh-token/{id}",
				Controller.GetRefreshToken,
				internaljwt.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"GET /refresh-tokens",
				Controller.ListRefreshTokens,
				internaljwt.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"DELETE /refresh-token/{id}",
				Controller.RevokeRefreshToken,
				internaljwt.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"DELETE /refresh-tokens",
				Controller.RevokeRefreshTokens,
				internaljwt.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"POST /totp/generate",
				Controller.GenerateTOTPUrl,
				internaljwt.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"POST /totp/verify",
				Controller.VerifyTOTP,
				internaljwt.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"DELETE /totp",
				Controller.RevokeTOTP,
				internaljwt.Authenticate(gojwtinterception.AccessToken),
			)
		},
	}
)
