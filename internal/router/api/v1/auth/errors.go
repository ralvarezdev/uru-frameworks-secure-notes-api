package auth

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	"net/http"
)

var (
	ErrLogInInvalidPassword = gonethttpresponse.NewFieldError(
		"password",
		"invalid password",
		http.StatusUnauthorized,
		nil,
	)
	ErrLogInInvalidTOTPCode = gonethttpresponse.NewFieldError(
		"totp_code",
		"invalid TOTP code",
		http.StatusUnauthorized,
		nil,
	)
	ErrLogInInvalidTOTPRecoveryCode = gonethttpresponse.NewFieldError(
		"totp_code",
		"invalid TOTP recovery code",
		http.StatusUnauthorized,
		nil,
	)
	ErrLogInMissingTOTPCode = gonethttpresponse.NewFieldError(
		"totp_code",
		"missing TOTP code",
		http.StatusUnauthorized,
		nil,
	)
	ErrLogInMissingIsTOTPRecoveryCode = gonethttpresponse.NewFieldError(
		"is_totp_recovery_code",
		"missing is TOTP recovery code",
		http.StatusUnauthorized,
		nil,
	)
	ErrVerifyTOTPInvalidTOTPCode = gonethttpresponse.NewFieldError(
		"totp_code",
		"TOTP code is required",
		http.StatusUnauthorized,
		nil,
	)
)
