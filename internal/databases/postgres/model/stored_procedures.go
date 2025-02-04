package model

var (
	// SignUpProc is the query to call the sign-up stored procedure
	SignUpProc = "CALL sign_up($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	// RevokeUserTOTPProc is the query to call the revoke user TOTP stored procedure
	RevokeUserTOTPProc = "CALL revoke_user_totp($1)"

	// GenerateTokensProc is the query to call the generate tokens stored procedure
	GenerateTokensProc = "CALL generate_tokens($1, $2, $3, $4, $5, $6, $7)"

	// RevokeUserTokensByIDProc is the query to call the revoke user tokens by ID stored procedure
	RevokeUserTokensByIDProc = "CALL revoke_user_tokens_by_id($1, $2)"

	// RefreshTokenProc is the query to call the refresh token stored procedure
	RefreshTokenProc = "CALL refresh_token($1, $2, $3, $4, $5, $6, $7)"

	// RevokeUserTokensProc is the query to call the revoke user tokens stored procedure
	RevokeUserTokensProc = "CALL revoke_user_tokens($1)"

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

	// RevokeUserTOTPRecoveryCodeProc is the query to call the revoke user TOTP recovery code stored procedure
	RevokeUserTOTPRecoveryCodeProc = "CALL revoke_user_totp_recovery_code($1, $2)"

	// VerifyTOTPProc is the query to call verify TOTP stored procedure
	VerifyTOTPProc = "CALL verify_totp($1)"

	// SendEmailVerificationTokenProc is the query to call the send email verification token stored procedure
	SendEmailVerificationTokenProc = "CALL send_email_verification_token($1, $2, $3)"

	// GetUserEmailIDProc is the query to call the get user email ID stored procedure
	GetUserEmailIDProc = "CALL get_user_email_id($1, $2)"

	// VerifyEmailProc is the query to call the verify email stored procedure
	VerifyEmailProc = "CALL verify_email($1, $2, $3)"

	// IsUserEmailVerifiedProc is the query to call the is user email verified stored procedure
	IsUserEmailVerifiedProc = "CALL is_user_email_verified($1, $2)"

	// PreSendEmailVerificationTokenProc is the query to call the pre-send email verification token stored procedure
	PreSendEmailVerificationTokenProc = "CALL pre_send_email_verification_token($1, $2, $3, $4, $5, $6)"

	// RevokeUserEmailProc is the query to call the revoke user email stored procedure
	RevokeUserEmailProc = "CALL revoke_user_email($1)"

	// ChangeEmailProc is the query to call the change email stored procedure
	ChangeEmailProc = "CALL change_email($1, $2, $3, $4, $5, $6)"

	// ForgotPasswordProc is the query to call the forgot password stored procedure
	ForgotPasswordProc = "CALL forgot_password($1, $2, $3, $4, $5, $6, $7)"

	// RevokeUserResetPasswordTokenProc is the query to call the revoke user reset password token stored procedure
	RevokeUserResetPasswordTokenProc = "CALL revoke_user_reset_password_token($1)"

	// RevokeUserPasswordHashProc is the query to call the revoke user password hash stored procedure
	RevokeUserPasswordHashProc = "CALL revoke_user_password_hash($1)"

	// ResetPasswordProc is the query to call the reset password stored procedure
	ResetPasswordProc = "CALL reset_password($1, $2, $3, $4)"

	// RevokeUserTokensExceptRefreshTokenIDProc is the query to call the revoke user tokens except refresh token ID stored procedure
	RevokeUserTokensExceptRefreshTokenIDProc = "CALL revoke_user_tokens_except_refresh_token_id($1, $2)"

	// ChangePasswordProc is the query to call the change password stored procedure
	ChangePasswordProc = "CALL change_password($1, $2, $3)"

	// GetUserPasswordHashProc is the query to call the get user password hash stored procedure
	GetUserPasswordHashProc = "CALL get_user_password_hash($1, $2)"

	// RevokeUserUsernameProc is the query to call the revoke username stored procedure
	RevokeUserUsernameProc = "CALL revoke_user_username($1)"

	// RevokeUserPhoneNumberProc is the query to call the revoke user phone number stored procedure
	RevokeUserPhoneNumberProc = "CALL revoke_user_phone_number($1)"

	// DeleteUserProc is the query to call the delete user stored procedure
	DeleteUserProc = "CALL delete_user($1)"

	// ChangeUsernameProc is the query to call the change username stored procedure
	ChangeUsernameProc = "CALL change_username($1, $2)"

	// UpdateProfileProc is the query to call the update profile stored procedure
	UpdateProfileProc = "CALL update_profile($1, $2, $3, $4)"

	// RevokeUserEmailVerificationTokenProc is the query to call the revoke user email verification token stored procedure
	RevokeUserEmailVerificationTokenProc = "CALL revoke_user_email_verification_token($1)"

	// GetUserPhoneNumberProc is the query to call the get user phone number stored procedure
	GetUserPhoneNumberProc = "CALL get_user_phone_number($1, $2)"

	// GetUserUsernameProc is the query to call the get user username stored procedure
	GetUserUsernameProc = "CALL get_user_username($1, $2)"

	// IsUserPhoneNumberVerifiedProc is the query to call the is user phone number verified stored procedure
	IsUserPhoneNumberVerifiedProc = "CALL is_user_phone_number_verified($1, $2)"

	// HasUserTOTPEnabledProc is the query to call the has user TOTP enabled stored procedure
	HasUserTOTPEnabledProc = "CALL has_user_totp_enabled($1, $2)"

	// GetMyProfileProc is the query to call the get my profile stored procedure
	GetMyProfileProc = "CALL get_my_profile($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
)
