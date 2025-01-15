package auth

import (
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/_common"
)

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

// GetRefreshTokenResponse is the response DTO to get a refresh token that has not been revoked or expired
type GetRefreshTokenResponse struct {
	RefreshToken *internalapiv1common.UserRefreshToken `json:"refresh_token"`
}

// ListRefreshTokensResponse is the response DTO to list refresh tokens that have not been revoked or expired
type ListRefreshTokensResponse struct {
	RefreshTokens []*internalapiv1common.UserRefreshTokenWithID `json:"refresh_tokens"`
}
