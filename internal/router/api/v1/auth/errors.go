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
		http.StatusBadRequest,
		nil,
	)
	ErrSignUpEmailAlreadyRegistered = gonethttpresponse.NewFieldError(
		"email",
		"email is already registered",
		http.StatusBadRequest,
		nil,
	)
	ErrLogInInvalidUsername = gonethttpresponse.NewFieldError(
		"username",
		"user not found by username",
		http.StatusUnauthorized,
		nil,
	)
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
	ErrLogInRequiredTOTPCode = gonethttpresponse.NewFieldError(
		"totp_code",
		fmt.Sprintf(govalidatormappervalidations.ErrRequiredField, "totp_code"),
		http.StatusBadRequest,
		nil,
	)
	ErrLogInRequiredIsTOTPRecoveryCode = gonethttpresponse.NewFieldError(
		"is_totp_recovery_code",
		fmt.Sprintf(
			govalidatormappervalidations.ErrRequiredField,
			"is_totp_recovery_code",
		),
		http.StatusBadRequest,
		nil,
	)
	ErrGenerateTOTPUrlAlreadyVerified = gonethttpresponse.NewFieldError(
		"totp",
		"TOTP is already verified",
		http.StatusBadRequest,
		nil,
	)
	ErrVerifyTOTPNotGenerated = gonethttpresponse.NewFieldError(
		"totp",
		"user has not generated TOTP",
		http.StatusBadRequest,
		nil,
	)
	ErrVerifyTOTPInvalidTOTPCode = gonethttpresponse.NewFieldError(
		"totp_code",
		"invalid TOTP code",
		http.StatusBadRequest,
		nil,
	)
	ErrVerifyTOTPAlreadyVerified = gonethttpresponse.NewFieldError(
		"totp",
		"TOTP is already verified",
		http.StatusBadRequest,
		nil,
	)
	ErrGetRefreshTokenNotFound = gonethttpresponse.NewFieldError(
		"id",
		"refresh token not found",
		http.StatusNotFound,
		nil,
	)
	ErrVerifyEmailTokenNotFound = gonethttpresponse.NewFieldError(
		"token_id",
		"email verification token not found",
		http.StatusNotFound,
		nil,
	)
	ErrSendEmailVerificationTokenAlreadyVerified = gonethttpresponse.NewFieldError(
		"email",
		"email is already verified",
		http.StatusBadRequest,
		nil,
	)
	ErrChangeEmailAlreadyRegistered = gonethttpresponse.NewFieldError(
		"email",
		"email is already registered",
		http.StatusBadRequest,
		nil,
	)
	ErrVerifyEmailInvalidToken = gonethttpresponse.NewFieldError(
		"token",
		"email has already been verified, token has expired, or token is invalid",
		http.StatusBadRequest,
		nil,
	)
	ErrResetPasswordInvalidToken = gonethttpresponse.NewFieldError(
		"token",
		"token has expired or is invalid",
		http.StatusBadRequest,
		nil,
	)
	ErrChangePasswordInvalidOldPassword = gonethttpresponse.NewFieldError(
		"old_password",
		"invalid old password",
		http.StatusBadRequest,
		nil,
	)
	ErrChangePasswordSamePassword = gonethttpresponse.NewFieldError(
		"new_password",
		"new password is the same as the old password",
		http.StatusBadRequest,
		nil,
	)
)
