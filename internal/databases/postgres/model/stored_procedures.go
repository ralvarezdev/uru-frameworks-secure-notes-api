package model

var (
	// SignUpProc is the query to call the sign-up stored procedure
	SignUpProc = "CALL sign_up($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	// RevokeTOTPProc is the query to call the revoke TOTP stored procedure
	RevokeTOTPProc = "CALL revoke_totp($1)"

	// GenerateTokensProc is the query to call the generate tokens stored procedure
	GenerateTokensProc = "CALL generate_tokens($1, $2, $3, $4, $5, $6, $7)"

	// RevokeTokensByIDProc is the query to call the revoke tokens by ID stored procedure
	RevokeTokensByIDProc = "CALL revoke_tokens_by_id($1, $2)"

	// RefreshTokenProc is the query to call the refresh token stored procedure
	RefreshTokenProc = "CALL refresh_token($1, $2, $3, $4, $5, $6, $7)"

	// RevokeTokensProc is the query to call the revoke tokens stored procedure
	RevokeTokensProc = "CALL revoke_tokens($1)"

	// GetAccessTokenByRefreshTokenIDProc is the query to call the get access token by refresh token ID stored procedure
	GetAccessTokenByRefreshTokenIDProc = "CALL get_access_token_by_refresh_token_id($1, $2)"

	// PreLogInProc is the query to call the pre-login stored procedure
	PreLogInProc = "CALL pre_log_in($1, $2, $3, $4, $5, $6, $7);"

	// RegisterFailedLoginAttemptProc is the query to call the register failed login attempt stored procedure
	RegisterFailedLoginAttemptProc = "CALL register_failed_login_attempt($1, $2, $3, $4)"

	// GetUserTOTPProc is the query to call the get user TOTP by user ID stored procedure
	GetUserTOTPProc = "CALL get_user_totp($1, $2, $3, $4)"

	// GetUserEmailProc is the query to call the get user email by user ID stored procedure
	GetUserEmailProc = "CALL get_user_email($1, $2)"

	// GenerateTOTPUrlProc is the query to call the generate TOTP URL stored procedure
	GenerateTOTPUrlProc = "CALL pre_generate_totp_url($1, $2, $3, $4, $5, $6)"

	// IsRefreshTokenValidProc is the query to call the is refresh token valid stored procedure
	IsRefreshTokenValidProc = "CALL is_refresh_token_valid($1, $2, $3, $4)"

	// IsAccessTokenValidProc is the query to call the is access token valid stored procedure
	IsAccessTokenValidProc = "CALL is_access_token_valid($1, $2, $3, $4)"

	// RevokeTOTPRecoveryCodeProc is the query to call the revoke TOTP recovery code stored procedure
	RevokeTOTPRecoveryCodeProc = "CALL revoke_totp_recovery_code($1, $2)"

	// VerifyTOTPProc is the query to call verify TOTP stored procedure
	VerifyTOTPProc = "CALL verify_totp($1)"

	// SendEmailVerificationTokenProc is the query to call the send email verification token stored procedure
	SendEmailVerificationTokenProc = "CALL send_email_verification_token($1, $2, $3)"

	// GetUserEmailIDProc is the query to call the get user email ID stored procedure
	GetUserEmailIDProc = "CALL get_user_email_id($1, $2)"

	// VerifyEmailProc is the query to call the verify email stored procedure
	VerifyEmailProc = "CALL verify_email($1, $2)"

	// IsUserEmailVerifiedProc is the query to call the is user email verified stored procedure
	IsUserEmailVerifiedProc = "CALL is_user_email_verified($1, $2, $3, $4, $5)"

	// RevokeEmailProc is the query to call the revoke email stored procedure
	RevokeEmailProc = "CALL revoke_email($1)"

	// ChangeEmailProc is the query to call the change email stored procedure
	ChangeEmailProc = "CALL change_email($1, $2, $3, $4, $5, $6)"
)
