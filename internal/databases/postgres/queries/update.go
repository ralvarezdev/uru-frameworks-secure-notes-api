package queries

const (
	// UpdateUserTOTPRecoveryCodeRevokedAtByID is the query to update the TOTP recovery code revoked at field by ID
	UpdateUserTOTPRecoveryCodeRevokedAtByID = `
UPDATE
	user_totp_recovery_codes
SET
	revoked_at = NOW()
WHERE
	user_totp_recovery_codes.id = $1
`

	// UpdateUserRefreshTokenRevokedAtByID is the query to update the refresh token revoked at field by ID
	UpdateUserRefreshTokenRevokedAtByID = `
UPDATE
	user_refresh_tokens
SET	
	revoked_at = NOW()
WHERE
	user_refresh_tokens.id = $1
`

	// UpdateUserRefreshTokensRevokedAtByUserID is the query to update the refresh tokens revoked at field by the user ID
	UpdateUserRefreshTokensRevokedAtByUserID = `
UPDATE
	user_refresh_tokens
SET	
	revoked_at = NOW()
WHERE
	user_refresh_tokens.user_id = $1
`

	// UpdateUserAccessTokenRevokedAtByUserRefreshTokenID is the query to update the access token revoked at field by the user refresh token ID
	UpdateUserAccessTokenRevokedAtByUserRefreshTokenID = `
UPDATE	
	user_access_tokens
SET
	revoked_at = NOW()
WHERE
	user_access_tokens.user_refresh_token_id = $1
`

	// UpdateUserAccessTokensRevokedAtByUserID is the query to update the access tokens revoked at field by the user ID
	UpdateUserAccessTokensRevokedAtByUserID = `
UPDATE
	user_access_tokens
SET	
	revoked_at = NOW()
WHERE
	user_access_tokens.user_id = $1
`

	// UpdateUserTOTPRevokedAtByID is the query to update the TOTP revoked at field by the ID
	UpdateUserTOTPRevokedAtByID = `
UPDATE
	user_totps
SET
	revoked_at = NOW()
WHERE
	user_totps.id = $1
`

	// UpdateUserTOTPVerifiedAtByID is the query to update the TOTP verified at field by the ID
	UpdateUserTOTPVerifiedAtByID = `
UPDATE
	user_totps
SET
	verified_at = NOW()
WHERE
	user_totps.id = $1
`
)
