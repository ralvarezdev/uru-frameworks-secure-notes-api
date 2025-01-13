package auth

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
)

var (
	ErrInvalidPassword = gonethttpresponse.NewFieldError(
		"password",
		"invalid password",
	)
	ErrInvalidTOTPCode = gonethttpresponse.NewFieldError(
		"totp_code",
		"invalid TOTP code",
	)
	ErrInvalidTOTPRecoveryCode = gonethttpresponse.NewFieldError(
		"totp_code",
		"invalid TOTP recovery code",
	)
	ErrMissingTOTPCode = gonethttpresponse.NewFieldError(
		"totp_code",
		"missing TOTP code",
	)
	ErrMissingIsTOTPRecoveryCode = gonethttpresponse.NewFieldError(
		"is_totp_recovery_code",
		"missing is TOTP recovery code",
	)
)
