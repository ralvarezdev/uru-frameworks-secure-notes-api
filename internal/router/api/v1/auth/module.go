package auth

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
						internalvalidator.Service.Password(
							"password",
							body.Password,
							internal.PasswordOptions,
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
				internalmiddleware.Authenticate,
			)
			m.RegisterExactRoute(
				"POST /logout",
				Controller.LogOut,
				internalmiddleware.Authenticate,
			)
			m.RegisterExactRoute(
				"GET /refresh-token",
				Controller.GetRefreshToken,
				internalmiddleware.Authenticate,
				internalmiddleware.Validate(
					&GetRefreshTokenRequest{},
				),
			)
			m.RegisterExactRoute(
				"GET /refresh-tokens",
				Controller.ListRefreshTokens,
				internalmiddleware.Authenticate,
			)
			m.RegisterExactRoute(
				"DELETE /refresh-token",
				Controller.RevokeRefreshToken,
				internalmiddleware.Authenticate,
				internalmiddleware.Validate(
					&RevokeRefreshTokenRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /refresh-tokens",
				Controller.RevokeRefreshTokens,
				internalmiddleware.Authenticate,
			)
			m.RegisterExactRoute(
				"POST /totp/generate",
				Controller.GenerateTOTPUrl,
				internalmiddleware.Authenticate,
			)
			m.RegisterExactRoute(
				"POST /totp/verify",
				Controller.VerifyTOTP,
				internalmiddleware.Authenticate,
				internalmiddleware.Validate(
					&VerifyTOTPRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /totp",
				Controller.RevokeTOTP,
				internalmiddleware.Authenticate,
			)
			m.RegisterExactRoute(
				"PUT /password",
				Controller.ChangePassword,
				internalmiddleware.Authenticate,
				internalmiddleware.Validate(
					&ChangePasswordRequest{},
					func(
						body *ChangePasswordRequest,
						validations *govalidatormappervalidation.StructValidations,
					) {
						internalvalidator.Service.Password(
							"new_password",
							body.NewPassword,
							internal.PasswordOptions,
							validations,
						)
					},
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
				"POST /password/reset/{token}",
				Controller.ResetPassword,
				internalmiddleware.Validate(
					&ResetPasswordRequest{},
					func(
						body *ResetPasswordRequest,
						validations *govalidatormappervalidation.StructValidations,
					) {
						internalvalidator.Service.Password(
							"new_password",
							body.NewPassword,
							internal.PasswordOptions,
							validations,
						)
					},
				),
			)
			m.RegisterExactRoute(
				"PUT /email",
				Controller.ChangeEmail,
				internalmiddleware.Authenticate,
				internalmiddleware.Validate(&ChangeEmailRequest{}),
			)
			m.RegisterExactRoute(
				"POST /email/send-verification",
				Controller.SendEmailVerificationToken,
				internalmiddleware.Authenticate,
			)
			m.RegisterExactRoute(
				"POST /email/verify/{token}",
				Controller.VerifyEmail,
			)
			m.RegisterExactRoute(
				"PUT /phone-number",
				Controller.ChangePhoneNumber,
				internalmiddleware.Authenticate,
				internalmiddleware.Validate(&ChangePhoneNumberRequest{}),
			)
			m.RegisterExactRoute(
				"POST /phone-number/send-verification",
				Controller.SendPhoneNumberVerificationCode,
				internalmiddleware.Authenticate,
			)
			m.RegisterExactRoute(
				"POST /phone-number/verify",
				Controller.VerifyPhoneNumber,
				internalmiddleware.Authenticate,
				internalmiddleware.Validate(&VerifyPhoneNumberRequest{}),
			)
		},
	}
)
