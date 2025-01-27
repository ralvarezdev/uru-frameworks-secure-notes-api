package model

var (
	// InsertUserTOTPRecoveryCodes is the SQL query to insert the user TOTP recovery codes
	// This is dynamically set based on the recovery codes count
	InsertUserTOTPRecoveryCodes string
)
