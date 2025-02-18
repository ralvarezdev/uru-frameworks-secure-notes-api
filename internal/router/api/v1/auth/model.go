package auth

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
)

type (
	// SignUpRequest is the request DTO to sign up
	SignUpRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
	}

	// LogInRequest is the request DTO to log in
	LogInRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// LogInResponseData is the response data DTO to log in
	LogInResponseData struct {
		TwoFactorAuthenticationMethods *[]string `json:"2fa_methods,omitempty"`
	}

	// EnableUser2FARequest is the request DTO to enable user 2FA
	EnableUser2FARequest struct {
		Password string `json:"password"`
	}

	// EnableUser2FAResponseData is the response data DTO to enable user 2FA
	EnableUser2FAResponseData struct {
		RecoveryCodes []string `json:"recovery_codes"`
	}

	// EnableUser2FAResponseBody is the response body DTO to enable user 2FA
	EnableUser2FAResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data EnableUser2FAResponseData `json:"data"`
	}

	// DisableUser2FARequest is the request DTO to disable user 2FA
	DisableUser2FARequest struct {
		Password string `json:"password"`
	}

	// RegenerateUser2FARecoveryCodesRequest is the request DTO to regenerate user 2FA recovery codes
	RegenerateUser2FARecoveryCodesRequest struct {
		Password string `json:"password"`
	}

	// RegenerateUser2FARecoveryCodesResponseData is the response data DTO to regenerate user 2FA recovery codes
	RegenerateUser2FARecoveryCodesResponseData struct {
		RecoveryCodes []string `json:"recovery_codes"`
	}

	// RegenerateUser2FARecoveryCodesResponseBody is the response body DTO to regenerate user 2FA recovery codes
	RegenerateUser2FARecoveryCodesResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data RegenerateUser2FARecoveryCodesResponseData `json:"data"`
	}

	// Generate2FATOTPUrlResponseData is the response data DTO to generate 2FA TOTP URL
	Generate2FATOTPUrlResponseData struct {
		TOTPUrl string `json:"totp_url"`
	}

	// Generate2FATOTPUrlResponseBody is the response body DTO to generate 2FA TOTP URL
	Generate2FATOTPUrlResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data Generate2FATOTPUrlResponseData `json:"data"`
	}

	// Verify2FATOTPRequest is the request DTO to verify 2FA TOTP
	Verify2FATOTPRequest struct {
		TOTPCode string `json:"totp_code"`
	}

	// RevokeRefreshTokenRequest is the request DTO to revoke a refresh token
	RevokeRefreshTokenRequest struct {
		RefreshTokenID int64 `json:"refresh_token_id"`
	}

	// GetRefreshTokenRequest is the request DTO to get a refresh token
	GetRefreshTokenRequest struct {
		RefreshTokenID int64 `json:"refresh_token_id"`
	}

	// GetRefreshTokenResponseData is the response data DTO to get a refresh token that has not been revoked or expired
	GetRefreshTokenResponseData struct {
		RefreshToken *internalpostgresmodel.UserRefreshToken `json:"refresh_token"`
	}

	// GetRefreshTokenResponseBody is the response body DTO to get a refresh token that has not been revoked or expired
	GetRefreshTokenResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data GetRefreshTokenResponseData `json:"data"`
	}

	// ListRefreshTokensResponseData is the response data DTO to list refresh tokens that have not been revoked or expired
	ListRefreshTokensResponseData struct {
		RefreshTokens []*internalpostgresmodel.UserRefreshTokenWithID `json:"refresh_tokens"`
	}

	// ListRefreshTokensResponseBody is the response body DTO to list refresh tokens that have not been revoked or expired
	ListRefreshTokensResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data ListRefreshTokensResponseData `json:"data"`
	}

	// ChangePasswordRequest is the request DTO to change password
	ChangePasswordRequest struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	// ForgotPasswordRequest is the request DTO to forgot password
	ForgotPasswordRequest struct {
		Email string `json:"email"`
	}

	// ResetPasswordRequest is the request DTO to reset password
	ResetPasswordRequest struct {
		NewPassword string `json:"new_password"`
		Token       string `json:"token"`
	}

	// ChangeEmailRequest is the request DTO to change email
	ChangeEmailRequest struct {
		Email string `json:"email"`
	}

	// ChangePhoneNumberRequest is the request DTO to change phone number
	ChangePhoneNumberRequest struct {
		PhoneNumber string `json:"phone_number"`
	}

	// VerifyPhoneNumberRequest is the request DTO to verify phone number
	VerifyPhoneNumberRequest struct {
		Token string `json:"token"`
	}

	// VerifyEmailRequest is the request DTO to verify email
	VerifyEmailRequest struct {
		Token string `json:"token"`
	}
)
