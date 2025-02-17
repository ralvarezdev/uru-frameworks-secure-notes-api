package model

const (
	// UserUsernamesUniqueUsername is the constraint name for the unique username constraint on the user_usernames table
	UserUsernamesUniqueUsername = "user_usernames_unique_username"

	// UserUsernamesUniqueUserID is the constraint name for the unique user ID constraint on the user_usernames table
	UserUsernamesUniqueUserID = "user_usernames_unique_user_id"

	// UserEmailsUniqueEmail is the constraint name for the unique email constraint on the user_emails table
	UserEmailsUniqueEmail = "user_emails_unique_email"

	// UserEmailsUniqueUserID is the constraint name for the unique user ID constraint on the user_emails table
	UserEmailsUniqueUserID = "user_emails_unique_user_id"

	// UserPhoneNumbersUniquePhoneNumber is the constraint name for the unique phone number constraint on the user_phone_numbers table
	UserPhoneNumbersUniquePhoneNumber = "user_phone_numbers_unique_phone_number"

	// UserPhoneNumbersUniqueUserID is the constraint name for the unique user ID constraint on the user_phone_numbers table
	UserPhoneNumbersUniqueUserID = "user_phone_numbers_unique_user_id"

	// UserTagsUniqueUserIDName is the constraint name for the unique user ID and name constraint on the user_tags table
	UserTagsUniqueUserIDName = "user_tags_unique_user_id_name"

	// UserPasswordHashesUniqueUserID is the constraint name for the unique user ID constraint on the user_password_hashes table
	UserPasswordHashesUniqueUserID = "user_password_hashes_unique_user_id"

	// UserResetPasswordsUniqueUserID is the constraint name for the unique user ID constraint on the user_reset_passwords table
	UserResetPasswordsUniqueUserID = "user_reset_passwords_unique_user_id"

	// UserEmailVerificationsUniqueUserEmailID is the constraint name for the unique user email ID constraint on the user_email_verifications table
	UserEmailVerificationsUniqueUserEmailID = "user_email_verifications_unique_user_email_id"

	// UserPhoneNumberVerificationsUniqueUserPhoneNumberID is the constraint name for the unique user phone number ID constraint on the user_phone_number_verifications table
	UserPhoneNumberVerificationsUniqueUserPhoneNumberID = "user_phone_number_verifications_unique_user_phone_number_id"

	// UserRefreshTokensUniqueParentUserRefreshTokenID is the constraint name for the unique parent user refresh token ID constraint on the user_refresh_tokens table
	UserRefreshTokensUniqueParentUserRefreshTokenID = "user_refresh_tokens_unique_parent_user_refresh_token_id"

	// UserAccessTokensUniqueUserRefreshTokenID is the constraint name for the unique user refresh token ID constraint on the user_access_tokens table
	UserAccessTokensUniqueUserRefreshTokenID = "user_access_tokens_unique_user_refresh_token_id"

	// User2FATOTPUniqueUserID is the constraint name for the unique user ID constraint on the user_2fa_totp table
	User2FATOTPUniqueUserID = "user_2fa_totp_unique_user_id"

	// UserNoteTagsUniqueUserNoteIDUserTagID is the constraint name for the unique user note ID and user tag ID constraint on the user_note_tags table
	UserNoteTagsUniqueUserNoteIDUserTagID = "user_note_tags_unique_user_note_id_user_tag_id"
)
