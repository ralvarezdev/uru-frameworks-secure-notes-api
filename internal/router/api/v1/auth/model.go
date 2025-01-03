package auth

import (
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/common"
)

// LogInRequest is the request DTO to log in
type LogInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LogInResponse is the response DTO to log in
type LogInResponse struct {
	Message      string  `json:"message"`
	TokenSeed    *string `json:"token_seed,omitempty"`
	Is2FAEnabled bool    `json:"is_2fa_enabled"`
}

// GenerateRefreshTokenRequest is the request DTO to generate refresh token
type GenerateRefreshTokenRequest struct {
	TokenSeed string  `json:"token_seed"`
	TOTPCode  *string `json:"totp_code,omitempty"`
}

// GenerateRefreshTokenResponse is the response DTO to generate refresh token
type GenerateRefreshTokenResponse struct {
	Message      string  `json:"message"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	AccessToken  *string `json:"access_token,omitempty"`
}

// RefreshTokenResponse is the response DTO to refresh token
type RefreshTokenResponse struct {
	Message      string  `json:"message"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	AccessToken  *string `json:"access_token,omitempty"`
}

// HasTOTPResponse is the response DTO to check if user has TOTP enabled
type HasTOTPResponse struct {
	Message string `json:"message"`
	HasTOTP bool   `json:"has_totp"`
}

// GenerateTOTPKeyResponse is the response DTO to generate TOTP key
type GenerateTOTPKeyResponse struct {
	Message string `json:"message"`
	TOTPKey string `json:"totp_key"`
}

// VerifyTOTPRequest is the request DTO to verify TOTP
type VerifyTOTPRequest struct {
	TOTPKey  string `json:"totp_key"`
	TOTPCode string `json:"totp_code"`
}

// VerifyTOTPResponse is the response DTO to verify TOTP
type VerifyTOTPResponse struct {
	Message    string `json:"message"`
	IsVerified *bool  `json:"is_verified,omitempty"`
}

// RevokeTOTPRequest is the request DTO to revoke TOTP
type RevokeTOTPRequest struct {
	Password string `json:"password"`
}

// GetUserRefreshTokenRequest is the request DTO to get a user refresh token that has not been revoked or expired
type GetUserRefreshTokenRequest struct {
	UserRefreshTokenID string `json:"user_refresh_token_id"`
}

// GetUserRefreshTokenResponse is the response DTO to get a user refresh token that has not been revoked or expired
type GetUserRefreshTokenResponse struct {
	Message          string                                `json:"message"`
	UserRefreshToken *internalapiv1common.UserRefreshToken `json:"user_refresh_token,omitempty"`
}

// ListUserRefreshTokensResponse is the response DTO to list user refresh tokens that have not been revoked or expired
type ListUserRefreshTokensResponse struct {
	Message           string                                       `json:"message"`
	UserRefreshTokens []internalapiv1common.UserRefreshTokenWithID `json:"user_refresh_tokens"`
}

// RevokeUserRefreshTokenRequest is the request DTO to revoke a user refresh token
type RevokeUserRefreshTokenRequest struct {
	UserRefreshTokenID string `json:"user_refresh_token_id"`
}

// RevokeUserRefreshTokens