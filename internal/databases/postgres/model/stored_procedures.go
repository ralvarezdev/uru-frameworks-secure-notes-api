package model

var (
	// SignUpProc is the query to call the stored procedure to sign-up
	SignUpProc = "CALL sign_up($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	// RevokeUserTOTPProc is the query to call the stored procedure to revoke user TOTP
	RevokeUserTOTPProc = "CALL revoke_user_totp($1)"

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

	// PreLogInProc is the query to call the stored procedure to pre-login
	PreLogInProc = "CALL pre_log_in($1, $2, $3, $4, $5, $6, $7);"

	// RegisterFailedLogInAttemptProc is the query to call the stored procedure to register failed login attempt
	RegisterFailedLogInAttemptProc = "CALL register_failed_log_in_attempt($1, $2, $3, $4)"

	// GetUserTOTPProc is the query to call the stored procedure to get user TOTP
	GetUserTOTPProc = "CALL get_user_totp($1, $2, $3, $4)"

	// GetUserEmailProc is the query to call the stored procedure to get user email
	GetUserEmailProc = "CALL get_user_email($1, $2)"

	// GenerateTOTPUrlProc is the query to call the stored procedure to generate TOTP URL
	GenerateTOTPUrlProc = "CALL generate_totp_url($1, $2, $3, $4, $5, $6)"

	// IsRefreshTokenValidProc is the query to call the stored procedure to check if the refresh token is valid
	IsRefreshTokenValidProc = "CALL is_refresh_token_valid($1, $2, $3, $4)"

	// IsAccessTokenValidProc is the query to call the stored procedure to check if the access token is valid
	IsAccessTokenValidProc = "CALL is_access_token_valid($1, $2, $3, $4)"

	// RevokeUserTOTPRecoveryCodeProc is the query to call the stored procedure to revoke user TOTP recovery code
	RevokeUserTOTPRecoveryCodeProc = "CALL revoke_user_totp_recovery_code($1, $2)"

	// VerifyTOTPProc is the query to call the stored procedure to verify TOTP
	VerifyTOTPProc = "CALL verify_totp($1, $2)"

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
	ResetPasswordProc = "CALL reset_password($1, $2, $3, $4)"

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

	// HasUserTOTPEnabledProc is the query to call the stored procedure to check if the user has TOTP enabled
	HasUserTOTPEnabledProc = "CALL has_user_totp_enabled($1, $2)"

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

	// GetUserTagByIDProc is the query to call the stored procedure to get user tag by tag ID
	GetUserTagByIDProc = "CALL get_user_tag_by_id($1, $2, $3, $4, $5)"

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

	// GetUserNoteVersionByIDProc is the query to call the stored procedure to get user note version by note version ID
	GetUserNoteVersionByIDProc = "CALL get_user_note_version_by_id($1, $2, $3, $4)"

	// CreateUserNoteProc is the query to call the stored procedure to create user note
	CreateUserNoteProc = "CALL create_user_note($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	// AddUserNoteTagsProc is the query to call the stored procedure to add user note tags
	AddUserNoteTagsProc = "CALL add_user_note_tags($1, $2, $3)"

	// ValidateUserTagsIDProc is the query to call the stored procedure to validate user tags ID
	ValidateUserTagsIDProc = "CALL validate_user_tags_id($1, $2, $3)"

	// CreateUserTOTPRecoveryCodesProc is the query to call the stored procedure to create user TOTP recovery codes
	CreateUserTOTPRecoveryCodesProc = "CALL create_user_totp_recovery_codes($1, $2)"

	// DeleteUserNoteProc is the query to call the stored procedure to delete user note
	DeleteUserNoteProc = "CALL delete_user_note($1, $2)"

	// RemoveUserNoteTagsProc is the query to call the stored procedure to remove user note tags
	RemoveUserNoteTagsProc = "CALL remove_user_note_tags($1, $2, $3)"

	// UpdateUserNoteProc is the query to call the stored procedure to update user note
	UpdateUserNoteProc = "CALL update_user_note($1, $2, $3, $4)"

	// GetUserNoteByIDProc is the query to call the stored procedure to get user note by note ID
	GetUserNoteByIDProc = "CALL get_user_note_by_id($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"

	// ListUserNotesProc is the query to call the stored procedure to list user notes
	ListUserNotesProc = "CALL list_user_notes($1, $2)"
)
