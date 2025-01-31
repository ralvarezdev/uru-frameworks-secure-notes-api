package auth

import (
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
	gonethttp "github.com/ralvarezdev/go-net/http"
	govalidatormappervalidation "github.com/ralvarezdev/go-validator/struct/mapper/validation"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
)

var (
	Service    = &service{}
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Path:       "/auth",
		Service:    Service,
		Controller: Controller,
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterRoute(
				"POST /signup",
				Controller.SignUp,
				internalmiddleware.Validate(
					&SignUpRequest{},
					func(
						body *SignUpRequest,
						validations *govalidatormappervalidation.StructValidations,
					) {
						internalvalidator.Service.Email(
							"email",
							body.Email,
							validations,
						)
					},
				),
			)
			m.RegisterRoute(
				"POST /login",
				Controller.LogIn,
				internalmiddleware.Validate(
					&LogInRequest{},
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
				"GET /refresh-token/{token_id}",
				Controller.GetRefreshToken,
				internalmiddleware.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"GET /refresh-tokens",
				Controller.ListRefreshTokens,
				internalmiddleware.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"DELETE /refresh-token/{token_id}",
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
				),
			)
			m.RegisterRoute(
				"DELETE /totp",
				Controller.RevokeTOTP,
				internalmiddleware.Authenticate(gojwtinterception.AccessToken),
			)
			m.RegisterRoute(
				"PUT /password",
				Controller.ChangePassword,
				internalmiddleware.Authenticate(gojwtinterception.AccessToken),
				internalmiddleware.Validate(
					&ChangePasswordRequest{},
				),
			)
			m.RegisterRoute(
				"POST /password/forgot",
				Controller.ForgotPassword,
				internalmiddleware.Validate(
					&ForgotPasswordRequest{},
				),
			)
			m.RegisterRoute(
				"POST /password/reset",
				Controller.ResetPassword,
				internalmiddleware.Validate(
					&ResetPasswordRequest{},
				),
			)
		},
	}
)
