package auth

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	govalidatormappervalidation "github.com/ralvarezdev/go-validator/struct/mapper/validation"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
)

var (
	Service    = &service{}
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Pattern:    "/auth",
		Service:    Service,
		Controller: Controller,
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterExactRoute(
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
			m.RegisterExactRoute(
				"POST /login",
				Controller.LogIn,
				internalmiddleware.Validate(
					&LogInRequest{},
				),
			)
			m.RegisterExactRoute(
				"POST /refresh-token",
				Controller.RefreshToken,
				internalmiddleware.AuthenticateRefreshToken,
			)
			m.RegisterExactRoute(
				"POST /logout",
				Controller.LogOut,
				internalmiddleware.AuthenticateAccessToken,
			)
			m.RegisterExactRoute(
				"GET /refresh-token",
				Controller.GetRefreshToken,
				internalmiddleware.AuthenticateAccessToken,
				internalmiddleware.Validate(
					&GetRefreshTokenRequest{},
				),
			)
			m.RegisterExactRoute(
				"GET /refresh-tokens",
				Controller.ListRefreshTokens,
				internalmiddleware.AuthenticateAccessToken,
			)
			m.RegisterExactRoute(
				"DELETE /refresh-token",
				Controller.RevokeRefreshToken,
				internalmiddleware.AuthenticateAccessToken,
				internalmiddleware.Validate(
					&RevokeRefreshTokenRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /refresh-tokens",
				Controller.RevokeRefreshTokens,
				internalmiddleware.AuthenticateAccessToken,
			)
			m.RegisterExactRoute(
				"POST /totp/generate",
				Controller.GenerateTOTPUrl,
				internalmiddleware.AuthenticateAccessToken,
			)
			m.RegisterExactRoute(
				"POST /totp/verify",
				Controller.VerifyTOTP,
				internalmiddleware.AuthenticateAccessToken,
				internalmiddleware.Validate(
					&VerifyTOTPRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /totp",
				Controller.RevokeTOTP,
				internalmiddleware.AuthenticateAccessToken,
			)
			m.RegisterExactRoute(
				"PUT /password",
				Controller.ChangePassword,
				internalmiddleware.AuthenticateAccessToken,
				internalmiddleware.Validate(
					&ChangePasswordRequest{},
				),
			)
			m.RegisterExactRoute(
				"POST /password/forgot",
				Controller.ForgotPassword,
				internalmiddleware.Validate(
					&ForgotPasswordRequest{},
				),
			)
			m.RegisterExactRoute(
				"POST /password/reset",
				Controller.ResetPassword,
				internalmiddleware.Validate(
					&ResetPasswordRequest{},
				),
			)
			m.RegisterExactRoute(
				"PUT /email",
				Controller.ChangeEmail,
				internalmiddleware.Validate(&ChangeEmailRequest{}),
			)
			m.RegisterExactRoute(
				"POST /email/send-verification",
				Controller.SendEmailVerificationToken,
			)
			m.RegisterExactRoute(
				"POST /email/verify/{token_id}",
				Controller.VerifyEmail,
			)
			m.RegisterExactRoute(
				"PUT /phone-number",
				Controller.ChangePhoneNumber,
				internalmiddleware.Validate(&ChangePhoneNumberRequest{}),
			)
			m.RegisterExactRoute(
				"POST /phone-number/send-verification",
				Controller.SendPhoneNumberVerificationCode,
			)
			m.RegisterExactRoute(
				"POST /phone-number/verify",
				Controller.VerifyPhoneNumber,
				internalmiddleware.Validate(&VerifyPhoneNumberRequest{}),
			)
		},
	}
)
