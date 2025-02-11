package auth

import (
	"fmt"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	govalidatormappervalidations "github.com/ralvarezdev/go-validator/struct/mapper/validator"
	"net/http"
)

var (
	ErrSignUpUsernameAlreadyRegistered = gonethttpresponse.NewFailResponseError(
		"username",
		"username is already registered",
		nil,
		http.StatusBadRequest,
	)
	ErrSignUpEmailAlreadyRegistered = gonethttpresponse.NewFailResponseError(
		"email",
		"email is already registered",
		nil,
		http.StatusBadRequest,
	)
	ErrLogInInvalidUsername = gonethttpresponse.NewFailResponseError(
		"username",
		"user not found by username",
		nil,
		http.StatusUnauthorized,
	)
	ErrLogInTooManyFailedAttempts = gonethttpresponse.NewFailResponseError(
		"password",
		"too many failed login attempts, try again later",
		nil,
		http.StatusUnauthorized,
	)
	ErrLogInInvalidPassword = gonethttpresponse.NewFailResponseError(
		"password",
		"invalid password",
		nil,
		http.StatusUnauthorized,
	)
	ErrLogInInvalidTOTPCode = gonethttpresponse.NewFailResponseError(
		"totp_code",
		"invalid TOTP code",
		nil,
		http.StatusUnauthorized,
	)
	ErrLogInInvalidTOTPRecoveryCode = gonethttpresponse.NewFailResponseError(
		"totp_code",
		"invalid TOTP recovery code",
		nil,
		http.StatusUnauthorized,
	)
	ErrLogInRequiredTOTPCode = gonethttpresponse.NewFailResponseError(
		"totp_code",
		fmt.Sprintf(govalidatormappervalidations.ErrRequiredField, "totp_code"),
		nil,
		http.StatusBadRequest,
	)
	ErrLogInRequiredIsTOTPRecoveryCode = gonethttpresponse.NewFailResponseError(
		"is_totp_recovery_code",
		fmt.Sprintf(
			govalidatormappervalidations.ErrRequiredField,
			"is_totp_recovery_code",
		),
		nil,
		http.StatusBadRequest,
	)
	ErrGenerateTOTPUrlAlreadyVerified = gonethttpresponse.NewFailResponseError(
		"totp",
		"TOTP is already verified",
		nil,
		http.StatusBadRequest,
	)
	ErrVerifyTOTPNotGenerated = gonethttpresponse.NewFailResponseError(
		"totp",
		"user has not generated TOTP",
		nil,
		http.StatusBadRequest,
	)
	ErrVerifyTOTPInvalidTOTPCode = gonethttpresponse.NewFailResponseError(
		"totp_code",
		"invalid TOTP code",
		nil,
		http.StatusBadRequest,
	)
	ErrVerifyTOTPAlreadyVerified = gonethttpresponse.NewFailResponseError(
		"totp",
		"TOTP is already verified",
		nil,
		http.StatusBadRequest,
	)
	ErrGetRefreshTokenNotFound = gonethttpresponse.NewFailResponseError(
		"id",
		"refresh token not found",
		nil,
		http.StatusNotFound,
	)
	ErrVerifyEmailTokenNotFound = gonethttpresponse.NewFailResponseError(
		"token_id",
		"email verification token not found",
		nil,
		http.StatusNotFound,
	)
	ErrSendEmailVerificationTokenAlreadyVerified = gonethttpresponse.NewFailResponseError(
		"email",
		"email is already verified",
		nil,
		http.StatusBadRequest,
	)
	ErrChangeEmailAlreadyRegistered = gonethttpresponse.NewFailResponseError(
		"email",
		"email is already registered",
		nil,
		http.StatusBadRequest,
	)
	ErrVerifyEmailInvalidToken = gonethttpresponse.NewFailResponseError(
		"token",
		"email has already been verified, token has expired, or token is invalid",
		nil,
		http.StatusBadRequest,
	)
	ErrResetPasswordInvalidToken = gonethttpresponse.NewFailResponseError(
		"token",
		"token has expired or is invalid",
		nil,
		http.StatusBadRequest,
	)
	ErrChangePasswordInvalidOldPassword = gonethttpresponse.NewFailResponseError(
		"old_password",
		"invalid old password",
		nil,
		http.StatusBadRequest,
	)
	ErrChangePasswordSamePassword = gonethttpresponse.NewFailResponseError(
		"new_password",
		"new password is the same as the old password",
		nil,
		http.StatusBadRequest,
	)
)
