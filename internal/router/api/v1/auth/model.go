package auth

import (
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/_common"
)

// RefreshTokenResponse is the response DTO to refresh token
type RefreshTokenResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

// LogInRequest is the request DTO to log in
type LogInRequest struct {
	Username           string  `json:"username"`
	Password           string  `json:"password"`
	TOTPCode           *string `json:"totp_code,omitempty"`
	IsTOTPRecoveryCode *bool   `json:"is_totp_recovery_code,omitempty"`
}

// RevokeRefreshTokenRequest is the request DTO to revoke refresh token
type RevokeRefreshTokenRequest struct {
	UserRefreshTokenID int64 `json:"user_refresh_token_id"`
}

// VerifyTOTPRequest is the request DTO to verify TOTP
type VerifyTOTPRequest struct {
	TOTPKey  string `json:"totp_key"`
	TOTPCode string `json:"totp_code"`
}

// VerifyTOTPResponse is the response DTO to verify TOTP
type VerifyTOTPResponse struct {
	IsVerified *bool `json:"is_verified,omitempty"`
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
	UserRefreshToken *internalapiv1common.UserRefreshToken `json:"user_refresh_token,omitempty"`
}

// ListUserRefreshTokensResponse is the response DTO to list user refresh tokens that have not been revoked or expired
type ListUserRefreshTokensResponse struct {
	UserRefreshTokens []*internalapiv1common.UserRefreshTokenWithID `json:"user_refresh_tokens"`
}
