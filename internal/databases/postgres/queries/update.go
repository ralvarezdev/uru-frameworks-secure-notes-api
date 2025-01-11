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
)
