package auth

import (
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/_common"
)

// SignUpRequest is the request DTO to sign up
type SignUpRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

// LogInRequest is the request DTO to log in
type LogInRequest struct {
	Username           string  `json:"username"`
	Password           string  `json:"password"`
	TOTPCode           *string `json:"totp_code,omitempty"`
	IsTOTPRecoveryCode *bool   `json:"is_totp_recovery_code,omitempty"`
}

// RefreshTokenResponse is the response DTO to refresh token
type RefreshTokenResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

// GenerateTOTPUrlResponse is the response DTO to generate TOTP URL
type GenerateTOTPUrlResponse struct {
	TOTPUrl string `json:"totp_url"`
}

// VerifyTOTPRequest is the request DTO to verify TOTP
type VerifyTOTPRequest struct {
	TOTPCode string `json:"totp_code"`
}

// VerifyTOTPResponse is the response DTO to verify TOTP
type VerifyTOTPResponse struct {
	RecoveryCodes []string `json:"recovery_codes"`
}

// RevokeRefreshTokenRequest is the request DTO to revoke a refresh token
type RevokeRefreshTokenRequest struct {
	RefreshTokenID int64 `json:"refresh_token_id"`
}

// GetRefreshTokenRequest is the request DTO to get a refresh token
type GetRefreshTokenRequest struct {
	RefreshTokenID int64 `json:"refresh_token_id"`
}

// GetRefreshTokenResponse is the response DTO to get a refresh token that has not been revoked or expired
type GetRefreshTokenResponse struct {
	RefreshToken *internalapiv1common.UserRefreshToken `json:"refresh_token"`
}

// ListRefreshTokensResponse is the response DTO to list refresh tokens that have not been revoked or expired
type ListRefreshTokensResponse struct {
	RefreshTokens []*internalapiv1common.UserRefreshTokenWithID `json:"refresh_tokens"`
}

// ChangePasswordRequest is the request DTO to change password
type ChangePasswordRequest struct {
	Password string `json:"password"`
}

// ForgotPasswordRequest is the request DTO to forgot password
type ForgotPasswordRequest struct {
	Username string `json:"username"`
}

// ResetPasswordRequest is the request DTO to reset password
type ResetPasswordRequest struct {
	ResetToken string `json:"reset_token"`
	Password   string `json:"password"`
}
