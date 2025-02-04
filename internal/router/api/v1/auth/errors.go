package auth

import (
	"fmt"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	govalidatormappervalidations "github.com/ralvarezdev/go-validator/struct/mapper/validator"
	"net/http"
)

var (
	ErrSignUpUsernameAlreadyRegistered = gonethttpresponse.NewFieldError(
		"username",
		"username is already registered",
		nil,
		http.StatusBadRequest,
	)
	ErrSignUpEmailAlreadyRegistered = gonethttpresponse.NewFieldError(
		"email",
		"email is already registered",
		nil,
		http.StatusBadRequest,
	)
	ErrLogInInvalidUsername = gonethttpresponse.NewFieldError(
		"username",
		"user not found by username",
		nil,
		http.StatusUnauthorized,
	)
	ErrLogInInvalidPassword = gonethttpresponse.NewFieldError(
		"password",
		"invalid password",
		nil,
		http.StatusUnauthorized,
	)
	ErrLogInInvalidTOTPCode = gonethttpresponse.NewFieldError(
		"totp_code",
		"invalid TOTP code",
		nil,
		http.StatusUnauthorized,
	)
	ErrLogInInvalidTOTPRecoveryCode = gonethttpresponse.NewFieldError(
		"totp_code",
		"invalid TOTP recovery code",
		nil,
		http.StatusUnauthorized,
	)
	ErrLogInRequiredTOTPCode = gonethttpresponse.NewFieldError(
		"totp_code",
		fmt.Sprintf(govalidatormappervalidations.ErrRequiredField, "totp_code"),
		nil,
		http.StatusBadRequest,
	)
	ErrLogInRequiredIsTOTPRecoveryCode = gonethttpresponse.NewFieldError(
		"is_totp_recovery_code",
		fmt.Sprintf(
			govalidatormappervalidations.ErrRequiredField,
			"is_totp_recovery_code",
		),
		nil,
		http.StatusBadRequest,
	)
	ErrGenerateTOTPUrlAlreadyVerified = gonethttpresponse.NewFieldError(
		"totp",
		"TOTP is already verified",
		nil,
		http.StatusBadRequest,
	)
	ErrVerifyTOTPNotGenerated = gonethttpresponse.NewFieldError(
		"totp",
		"user has not generated TOTP",
		nil,
		http.StatusBadRequest,
	)
	ErrVerifyTOTPInvalidTOTPCode = gonethttpresponse.NewFieldError(
		"totp_code",
		"invalid TOTP code",
		nil,
		http.StatusBadRequest,
	)
	ErrVerifyTOTPAlreadyVerified = gonethttpresponse.NewFieldError(
		"totp",
		"TOTP is already verified",
		nil,
		http.StatusBadRequest,
	)
	ErrGetRefreshTokenNotFound = gonethttpresponse.NewFieldError(
		"id",
		"refresh token not found",
		nil,
		http.StatusNotFound,
	)
	ErrVerifyEmailTokenNotFound = gonethttpresponse.NewFieldError(
		"token_id",
		"email verification token not found",
		nil,
		http.StatusNotFound,
	)
	ErrSendEmailVerificationTokenAlreadyVerified = gonethttpresponse.NewFieldError(
		"email",
		"email is already verified",
		nil,
		http.StatusBadRequest,
	)
	ErrChangeEmailAlreadyRegistered = gonethttpresponse.NewFieldError(
		"email",
		"email is already registered",
		nil,
		http.StatusBadRequest,
	)
	ErrVerifyEmailInvalidToken = gonethttpresponse.NewFieldError(
		"token",
		"email has already been verified, token has expired, or token is invalid",
		nil,
		http.StatusBadRequest,
	)
	ErrResetPasswordInvalidToken = gonethttpresponse.NewFieldError(
		"token",
		"token has expired or is invalid",
		nil,
		http.StatusBadRequest,
	)
	ErrChangePasswordInvalidOldPassword = gonethttpresponse.NewFieldError(
		"old_password",
		"invalid old password",
		nil,
		http.StatusBadRequest,
	)
	ErrChangePasswordSamePassword = gonethttpresponse.NewFieldError(
		"new_password",
		"new password is the same as the old password",
		nil,
		http.StatusBadRequest,
	)
)
