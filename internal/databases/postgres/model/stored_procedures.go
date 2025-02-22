package model

var (
	// SignUpProc is the query to call the stored procedure to sign-up
	SignUpProc = "CALL sign_up($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	// RevokeUser2FATOTPProc is the query to call the stored procedure to revoke user 2FA TOTP
	RevokeUser2FATOTPProc = "CALL revoke_user_2fa_totp($1)"

	// GenerateUserTokensProc is the query to call the stored procedure to generate user tokens
	GenerateUserTokensProc = "CALL generate_user_tokens($1, $2, $3, $4, $5, $6, $7)"

	// RevokeUserTokensByIDProc is the query to call the stored procedure to revoke user tokens by ID
	RevokeUserTokensByIDProc = "CALL revoke_user_tokens_by_id($1, $2)"

	// RefreshTokenProc is the query to call the stored procedure to refresh token
	RefreshTokenProc = "CALL refresh_token($1, $2, $3, $4, $5, $6, $7)"

	// RevokeUserTokensProc is the query to call the stored procedure to revoke user tokens
	RevokeUserTokensProc = "CALL revoke_user_tokens($1)"

	// GetUserAccessTokenByUserRefreshTokenIDProc is the query to call the stored procedure to get user access token by user refresh token ID
	GetUserAccessTokenByUserRefreshTokenIDProc = "CALL get_user_access_token_by_user_refresh_token_id($1, $2)"

	// RegisterFailedLogInAttemptProc is the query to call the stored procedure to register failed login attempt
	RegisterFailedLogInAttemptProc = "CALL register_failed_log_in_attempt($1, $2, $3, $4)"

	// GetUser2FATOTPProc is the query to call the stored procedure to get user 2FA TOTP
	GetUser2FATOTPProc = "CALL get_user_2fa_totp($1, $2, $3, $4)"

	// GetUserEmailProc is the query to call the stored procedure to get user email
	GetUserEmailProc = "CALL get_user_email($1, $2)"

	// Generate2FATOTPUrlProc is the query to call the stored procedure to generate 2FA TOTP URL
	Generate2FATOTPUrlProc = "CALL generate_2fa_totp_url($1, $2, $3, $4, $5, $6, $7)"

	// IsRefreshTokenValidProc is the query to call the stored procedure to check if the refresh token is valid
	IsRefreshTokenValidProc = "CALL is_refresh_token_valid($1, $2, $3, $4)"

	// IsAccessTokenValidProc is the query to call the stored procedure to check if the access token is valid
	IsAccessTokenValidProc = "CALL is_access_token_valid($1, $2, $3, $4)"

	// RevokeUser2FARecoveryCodesProc is the query to call the stored procedure to revoke user 2FA recovery codes
	RevokeUser2FARecoveryCodesProc = "CALL revoke_user_2fa_recovery_codes($1)"

	// UseUser2FARecoveryCodeProc is the query to call the stored procedure to use user 2FA recovery code
	UseUser2FARecoveryCodeProc = "CALL use_user_2fa_recovery_code($1, $2, $3, $4)"

	// Verify2FATOTPProc is the query to call the stored procedure to verify 2FA TOTP
	Verify2FATOTPProc = "CALL verify_2fa_totp($1)"

	// SendEmailVerificationTokenProc is the query to call the stored procedure to send email verification token
	SendEmailVerificationTokenProc = "CALL send_email_verification_token($1, $2, $3)"

	// GetUserEmailIDProc is the query to call the stored procedure to get user email ID
	GetUserEmailIDProc = "CALL get_user_email_id($1, $2)"

	// VerifyEmailProc is the query to call the stored procedure to verify email
	VerifyEmailProc = "CALL verify_email($1, $2, $3)"

	// IsUserEmailVerifiedProc is the query to call the stored procedure to check if the user email is verified
	IsUserEmailVerifiedProc = "CALL is_user_email_verified($1, $2)"

	// PreSendEmailVerificationTokenProc is the query to call the stored procedure to pre-send email verification token
	PreSendEmailVerificationTokenProc = "CALL pre_send_email_verification_token($1, $2, $3, $4, $5, $6)"

	// RevokeUserEmailProc is the query to call the stored procedure to revoke user email
	RevokeUserEmailProc = "CALL revoke_user_email($1)"

	// ChangeEmailProc is the query to call the stored procedure to change email
	ChangeEmailProc = "CALL change_email($1, $2, $3, $4, $5, $6)"

	// ForgotPasswordProc is the query to call the stored procedure to forgot password
	ForgotPasswordProc = "CALL forgot_password($1, $2, $3, $4, $5, $6)"

	// RevokeUserResetPasswordTokenProc is the query to call the stored procedure to revoke user reset password token
	RevokeUserResetPasswordTokenProc = "CALL revoke_user_reset_password_token($1)"

	// RevokeUserPasswordHashProc is the query to call the stored procedure to revoke user password hash
	RevokeUserPasswordHashProc = "CALL revoke_user_password_hash($1)"

	// ResetPasswordProc is the query to call the stored procedure to reset password
	ResetPasswordProc = "CALL reset_password($1, $2, $3, $4, $5)"

	// ChangePasswordProc is the query to call the stored procedure to change password
	ChangePasswordProc = "CALL change_password($1, $2, $3)"

	// GetUserPasswordHashProc is the query to call the stored procedure to get user password hash
	GetUserPasswordHashProc = "CALL get_user_password_hash($1, $2)"

	// RevokeUserUsernameProc is the query to call the stored procedure to revoke username
	RevokeUserUsernameProc = "CALL revoke_user_username($1)"

	// RevokeUserPhoneNumberProc is the query to call the stored procedure to revoke user phone number
	RevokeUserPhoneNumberProc = "CALL revoke_user_phone_number($1)"

	// DeleteUserProc is the query to call the stored procedure to delete user
	DeleteUserProc = "CALL delete_user($1)"

	// ChangeUsernameProc is the query to call the stored procedure to change username
	ChangeUsernameProc = "CALL change_username($1, $2)"

	// UpdateProfileProc is the query to call the stored procedure to update profile
	UpdateProfileProc = "CALL update_profile($1, $2, $3, $4)"

	// RevokeUserEmailVerificationTokenProc is the query to call the stored procedure to revoke user email verification token
	RevokeUserEmailVerificationTokenProc = "CALL revoke_user_email_verification_token($1)"

	// GetUserPhoneNumberProc is the query to call the stored procedure to get user phone number
	GetUserPhoneNumberProc = "CALL get_user_phone_number($1, $2)"

	// GetUserUsernameProc is the query to call the stored procedure to get user username
	GetUserUsernameProc = "CALL get_user_username($1, $2)"

	// IsUserPhoneNumberVerifiedProc is the query to call the stored procedure to check if the user phone number is verified
	IsUserPhoneNumberVerifiedProc = "CALL is_user_phone_number_verified($1, $2)"

	// HasUser2FAEnabledProc is the query to call the stored procedure to check if the user has 2FA enabled
	HasUser2FAEnabledProc = "CALL has_user_2fa_enabled($1, $2)"

	// GetMyProfileProc is the query to call the stored procedure to get my profile
	GetMyProfileProc = "CALL get_my_profile($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	// GetUserBasicInfoProc is the query to call the stored procedure to get user basic info
	GetUserBasicInfoProc = "CALL get_user_basic_info($1, $2, $3, $4)"

	// CreateUserTagProc is the query to call the stored procedure to create user tag
	CreateUserTagProc = "CALL create_user_tag($1, $2, $3)"

	// DeleteUserTagProc is the query to call the stored procedure to delete user tag
	DeleteUserTagProc = "CALL delete_user_tag($1, $2)"

	// UpdateUserTagProc is the query to call the stored procedure to update user tag
	UpdateUserTagProc = "CALL update_user_tag($1, $2, $3)"

	// UpdateUserNoteTrashProc is the query to call the stored procedure to update user note trash
	UpdateUserNoteTrashProc = "CALL update_user_note_trash($1, $2, $3)"

	// UpdateUserNoteStarProc is the query to call the stored procedure to update user note star
	UpdateUserNoteStarProc = "CALL update_user_note_star($1, $2, $3)"

	// UpdateUserNoteArchiveProc is the query to call the stored procedure to update user note archive
	UpdateUserNoteArchiveProc = "CALL update_user_note_archive($1, $2, $3)"

	// UpdateUserNotePinProc is the query to call the stored procedure to update user note pin
	UpdateUserNotePinProc = "CALL update_user_note_pin($1, $2, $3)"

	// CreateUserNoteVersionProc is the query to call the stored procedure to create user note version
	CreateUserNoteVersionProc = "CALL create_user_note_version($1, $2, $3, $4, $5)"

	// DeleteUserNoteVersionProc is the query to call the stored procedure to delete user note version
	DeleteUserNoteVersionProc = "CALL delete_user_note_version($1, $2)"

	// CreateUserNoteProc is the query to call the stored procedure to create user note
	CreateUserNoteProc = "CALL create_user_note($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	// AddUserNoteTagsProc is the query to call the stored procedure to add user note tags
	AddUserNoteTagsProc = "CALL add_user_note_tags($1, $2, $3)"

	// ValidateUserTagsIDProc is the query to call the stored procedure to validate user tags ID
	ValidateUserTagsIDProc = "CALL validate_user_tags_id($1, $2, $3)"

	// CreateUser2FARecoveryCodesProc is the query to call the stored procedure to create user 2FA recovery codes
	CreateUser2FARecoveryCodesProc = "CALL create_user_2fa_recovery_codes($1, $2, $3)"

	// DeleteUserNoteProc is the query to call the stored procedure to delete user note
	DeleteUserNoteProc = "CALL delete_user_note($1, $2)"

	// RemoveUserNoteTagsProc is the query to call the stored procedure to remove user note tags
	RemoveUserNoteTagsProc = "CALL remove_user_note_tags($1, $2, $3)"

	// UpdateUserNoteProc is the query to call the stored procedure to update user note
	UpdateUserNoteProc = "CALL update_user_note($1, $2, $3, $4)"

	// ListUserNotesProc is the query to call the stored procedure to list user notes
	ListUserNotesProc = "CALL list_user_notes($1, $2)"

	// RevokeUser2FAEmailCodeProc is the query to call the stored procedure to revoke user 2FA email code
	RevokeUser2FAEmailCodeProc = "CALL revoke_user_2fa_email_code($1)"

	// CreateUser2FAEmailCodeProc is the query to call the stored procedure to create user 2FA email code
	CreateUser2FAEmailCodeProc = "CALL create_user_2fa_email_code($1, $2)"

	// UseUser2FAEmailCodeProc is the query to call the stored procedure to use user 2FA email code
	UseUser2FAEmailCodeProc = "CALL use_user_2fa_email_code($1, $2, $3)"

	// EnableUser2FAProc is the query to call the stored procedure to enable user 2FA
	EnableUser2FAProc = "CALL enable_user_2fa($1, $2, $3, $4)"

	// DisableUser2FAProc is the query to call the stored procedure to disable user 2FA
	DisableUser2FAProc = "CALL disable_user_2fa($1, $2)"

	// SendUser2FAEmailCodeProc is the query to call the stored procedure to send user 2FA email code
	SendUser2FAEmailCodeProc = "CALL send_user_2fa_email_code($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	// HasUser2FATOTPEnabledProc is the query to call the stored procedure to check if the user has 2FA TOTP enabled
	HasUser2FATOTPEnabledProc = "CALL has_user_2fa_totp_enabled($1, $2)"

	// GetUser2FAMethodsProc is the query to call the stored procedure to get user 2FA methods
	GetUser2FAMethodsProc = "CALL get_user_2fa_methods($1, $2, $3)"

	// GetUserIDByResetPasswordTokenProc is the query to call the stored procedure to get user ID by reset password token
	GetUserIDByResetPasswordTokenProc = "CALL get_user_id_by_reset_password_token($1, $2)"
)
