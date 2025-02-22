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
	ErrLogInRequired2FACode = gonethttpresponse.NewFailResponseError(
		"2fa_code",
		fmt.Sprintf(govalidatormappervalidations.ErrRequiredField, "2fa_code"),
		nil,
		http.StatusBadRequest,
	)
	ErrLogInInvalid2FAMethod = gonethttpresponse.NewFailResponseError(
		"2fa_method",
		"invalid 2FA method",
		nil,
		http.StatusBadRequest,
	)
	ErrLogInInvalid2FARecoveryCode = gonethttpresponse.NewFailResponseError(
		"2fa_code",
		"invalid 2FA recovery code",
		nil,
		http.StatusUnauthorized,
	)
	ErrLogInInvalid2FATOTPCode = gonethttpresponse.NewFailResponseError(
		"2fa_code",
		"invalid 2FA TOTP code",
		nil,
		http.StatusUnauthorized,
	)
	ErrLogInInvalid2FAEmailCode = gonethttpresponse.NewFailResponseError(
		"2fa_code",
		"invalid 2FA email code",
		nil,
		http.StatusUnauthorized,
	)
	ErrGenerate2FATOTP2FAIsNotEnabled = gonethttpresponse.NewFailResponseError(
		"2fa",
		"2FA is not enabled",
		nil,
		http.StatusBadRequest,
	)
	ErrGenerate2FATOTPUrlAlreadyVerified = gonethttpresponse.NewFailResponseError(
		"totp",
		"2FA TOTP is already verified",
		nil,
		http.StatusBadRequest,
	)
	ErrVerify2FATOTPNotGenerated = gonethttpresponse.NewFailResponseError(
		"totp",
		"user has not generated 2FA TOTP",
		nil,
		http.StatusBadRequest,
	)
	ErrVerify2FATOTPInvalidTOTPCode = gonethttpresponse.NewFailResponseError(
		"totp_code",
		"invalid 2FA TOTP code",
		nil,
		http.StatusBadRequest,
	)
	ErrVerify2FATOTPAlreadyVerified = gonethttpresponse.NewFailResponseError(
		"totp",
		"2FA TOTP is already verified",
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
	ErrEnableUser2FAInvalidPassword = gonethttpresponse.NewFailResponseError(
		"password",
		"invalid password",
		nil,
		http.StatusBadRequest,
	)
	ErrEnableUser2FAEmailNotVerified = gonethttpresponse.NewFailResponseError(
		"email",
		"email is not verified",
		nil,
		http.StatusBadRequest,
	)
	ErrorEnableUser2FA2FAIsAlreadyEnabled = gonethttpresponse.NewFailResponseError(
		"2fa",
		"2FA is already enabled",
		nil,
		http.StatusBadRequest,
	)
	ErrDisableUser2FAInvalidPassword = gonethttpresponse.NewFailResponseError(
		"password",
		"invalid password",
		nil,
		http.StatusBadRequest,
	)
	ErrDisableUser2FA2FAIsNotEnabled = gonethttpresponse.NewFailResponseError(
		"2fa",
		"2FA is not enabled",
		nil,
		http.StatusBadRequest,
	)
	ErrRegenerateUser2FARecoveryCodesInvalidPassword = gonethttpresponse.NewFailResponseError(
		"password",
		"invalid password",
		nil,
		http.StatusBadRequest,
	)
	ErrRegenerateUser2FARecoveryCodes2FAIsNotEnabled = gonethttpresponse.NewFailResponseError(
		"2fa",
		"2FA is not enabled",
		nil,
		http.StatusBadRequest,
	)
	ErrSendUser2FAEmailCode2FAIsNotEnabled = gonethttpresponse.NewFailResponseError(
		"2fa",
		"2FA is not enabled",
		nil,
		http.StatusBadRequest,
	)
	ErrSendUser2FAEmailCodeInvalidUsername = gonethttpresponse.NewFailResponseError(
		"username",
		"user not found by username",
		nil,
		http.StatusUnauthorized,
	)
	ErrSendUser2FAEmailCodeInvalidPassword = gonethttpresponse.NewFailResponseError(
		"password",
		"invalid password",
		nil,
		http.StatusBadRequest,
	)
)
