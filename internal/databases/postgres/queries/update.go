package queries

const (
	// UpdateUserTOTPRecoveryCodeRevokedAtByTOTPIDAndCode is the query to update the TOTP recovery code revoked at field by the TOTP ID and recovery code
	UpdateUserTOTPRecoveryCodeRevokedAtByTOTPIDAndCode = `
UPDATE
	user_totp_recovery_codes
SET
	revoked_at = NOW()
WHERE
	user_totp_recovery_codes.user_totp_id = $1
AND
	user_totp_recovery_codes.recovery_code = $2
AND
	user_totp_recovery_codes.revoked_at IS NULL
`

	// UpdateUserRefreshTokenRevokedAtByIDAndUserID is the query to update the refresh token revoked at field by ID and user ID
	UpdateUserRefreshTokenRevokedAtByIDAndUserID = `
UPDATE
	user_refresh_tokens
SET	
	revoked_at = NOW()
WHERE
	user_refresh_tokens.id = $1
AND
	user_refresh_tokens.user_id = $2
AND 
	user_refresh_tokens.revoked_at IS NULL
`

	// UpdateUserRefreshTokensRevokedAtByUserID is the query to update the refresh tokens revoked at field by the user ID
	UpdateUserRefreshTokensRevokedAtByUserID = `
UPDATE
	user_refresh_tokens
SET	
	revoked_at = NOW()
WHERE
	user_refresh_tokens.user_id = $1
AND
	user_refresh_tokens.revoked_at IS NULL
`

	// UpdateUserAccessTokenRevokedAtByRefreshTokenIDAndUserID is the query to update the access token revoked at field by the user refresh token ID and user ID
	UpdateUserAccessTokenRevokedAtByRefreshTokenIDAndUserID = `
UPDATE	
	user_access_tokens
SET
	revoked_at = NOW()
WHERE
	user_access_tokens.user_refresh_token_id = $1
AND
	user_access_tokens.user_id = $2
AND
	user_access_tokens.revoked_at IS NULL
`

	// UpdateUserAccessTokensRevokedAtByUserID is the query to update the access tokens revoked at field by the user ID
	UpdateUserAccessTokensRevokedAtByUserID = `
UPDATE
	user_access_tokens
SET	
	revoked_at = NOW()
WHERE
	user_access_tokens.user_id = $1
AND
	user_access_tokens.revoked_at IS NULL
`

	// UpdateUserTOTPRevokedAtByID is the query to update the TOTP revoked at field by the ID
	UpdateUserTOTPRevokedAtByID = `
UPDATE
	user_totps
SET
	revoked_at = NOW()
WHERE
	user_totps.id = $1
AND
	user_totps.revoked_at IS NULL
`

	// UpdateUserTOTPRevokedAtByUserID is the query to update the TOTP revoked at field by the user ID
	UpdateUserTOTPRevokedAtByUserID = `
UPDATE
	user_totps
SET
	revoked_at = NOW()
WHERE
	user_totps.user_id = $1
AND
	user_totps.revoked_at IS NULL
`

	// UpdateUserTOTPVerifiedAtByID is the query to update the TOTP verified at field by the ID
	UpdateUserTOTPVerifiedAtByID = `
UPDATE
	user_totps
SET
	verified_at = NOW()
WHERE
	user_totps.id = $1
AND 
	user_totps.verified_at IS NULL
AND
	user_totps.revoked_at IS NULL
`

	// UpdateUserTOTPRecoveryCodeRevokedAtByUserID is the query to update the TOTP recovery code revoked at field by the user ID
	UpdateUserTOTPRecoveryCodeRevokedAtByUserID = `
UPDATE
	user_totp_recovery_codes
SET
	revoked_at = NOW()
WHERE
	user_totp_recovery_codes.user_id = $1
AND
	user_totp_recovery_codes.revoked_at IS NULL
`
)
