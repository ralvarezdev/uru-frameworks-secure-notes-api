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
		Username           string  `json:"username"`
		Password           string  `json:"password"`
		TOTPCode           *string `json:"totp_code,omitempty"`
		IsTOTPRecoveryCode *bool   `json:"is_totp_recovery_code,omitempty"`
	}

	// GenerateTOTPUrlResponseData is the response data DTO to generate TOTP URL
	GenerateTOTPUrlResponseData struct {
		TOTPUrl string `json:"totp_url"`
	}

	// GenerateTOTPUrlResponseBody is the response body DTO to generate TOTP URL
	GenerateTOTPUrlResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data GenerateTOTPUrlResponseData `json:"data"`
	}

	// VerifyTOTPRequest is the request DTO to verify TOTP
	VerifyTOTPRequest struct {
		TOTPCode string `json:"totp_code"`
	}

	// VerifyTOTPResponseData is the response data DTO to verify TOTP
	VerifyTOTPResponseData struct {
		RecoveryCodes []string `json:"recovery_codes"`
	}

	// VerifyTOTPResponseBody is the response body DTO to verify TOTP
	VerifyTOTPResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data VerifyTOTPResponseData `json:"data"`
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
		Password string `json:"password"`
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
)
