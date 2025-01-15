package queries

const (
	// SelectUserIDAndPasswordHashByUsername is the query to select the user ID and password hash by username
	SelectUserIDAndPasswordHashByUsername = `
SELECT
	users.id,
	user_password_hashes.password_hash
FROM
	users
JOIN 
	user_usernames ON users.id = user_usernames.user_id
JOIN
	user_password_hashes ON users.id = user_password_hashes.user_id
WHERE
	user_usernames.username = $1
AND
 	user_usernames.revoked_at IS NULL
AND
	users.deleted_at IS NULL
AND
	user_password_hashes.revoked_at IS NULL
`

	// SelectUserEmailByUserID is the query to select the user email by user ID
	SelectUserEmailByUserID = `
SELECT
	user_emails.email
FROM
	user_emails
WHERE
	user_emails.user_id = $1
AND
	user_emails.revoked_at IS NULL
`

	// SelectUserTOTPSecretVerifiedByUserID is the query to select the user TOTP secret that has been verified by user ID
	SelectUserTOTPSecretVerifiedByUserID = `
SELECT
	user_totps.id,
	user_totps.secret
FROM
	user_totps
WHERE
	user_totps.user_id = $1
AND
	user_totps.revoked_at IS NULL
AND
	user_totps.verified_at IS NOT NULL
`

	// SelectUserRefreshTokenExpiresAtByID is the query to select the user refresh token expires at by ID
	SelectUserRefreshTokenExpiresAtByID = `
SELECT
	user_refresh_tokens.expires_at
FROM
	user_refresh_tokens
WHERE
	user_refresh_tokens.id = $1
AND
	user_refresh_tokens.revoked_at IS NULL
`

	// SelectUserAccessTokenExpiresAtByID is the query to select the user access token expires at by ID
	SelectUserAccessTokenExpiresAtByID = `
SELECT
	user_access_tokens.expires_at
FROM
	user_access_tokens
WHERE
	user_access_tokens.id = $1
AND	
	user_access_tokens.revoked_at IS NULL
`

	// SelectUserTOTPByUserID is the query to select the user TOTP by user ID
	SelectUserTOTPByUserID = `
SELECT
	user_totps.id,
	user_totps.secret
	user_totps.verified_at
FROM
	user_totps
WHERE
	user_totps.user_id = $1
AND
	user_totps.revoked_at IS NULL
`

	// SelectUserRefreshTokensByUserID is the query to select the user refresh tokens by user ID
	SelectUserRefreshTokensByUserID = `
SELECT
	user_refresh_tokens.id,
	user_refresh_tokens.issued_at,
	user_refresh_tokens.expires_at,
	user_refresh_tokens.ip_address
FROM
	user_refresh_tokens
WHERE
	user_refresh_tokens.user_id = $1
AND
	user_refresh_tokens.revoked_at IS NULL
`

	// SelectUserRefreshTokenByIDAndUserID is the query to select the user refresh token by ID and user ID
	SelectUserRefreshTokenByIDAndUserID = `
SELECT
	user_refresh_tokens.issued_at,
	user_refresh_tokens.expires_at,
	user_refresh_tokens.ip_address
FROM
	user_refresh_tokens
WHERE	
	user_refresh_tokens.id = $1
AND
	user_refresh_tokens.user_id = $2
AND
	user_refresh_tokens.revoked_at IS NULL
`
)
