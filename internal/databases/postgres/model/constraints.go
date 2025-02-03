package model

const (
	// UserUsernamesUniqueUsername is the constraint name for the unique username constraint on the user_usernames table
	UserUsernamesUniqueUsername = "user_usernames_unique_username"

	// UserEmailsUniqueEmail is the constraint name for the unique email constraint on the user_emails table
	UserEmailsUniqueEmail = "user_emails_unique_email"

	// UserPhoneNumbersUniquePhoneNumber is the constraint name for the unique phone number constraint on the user_phone_numbers table
	UserPhoneNumbersUniquePhoneNumber = "user_phone_numbers_unique_phone_number"

	// UserTagsUniqueName is the constraint name for the unique name constraint on the user_tags table
	UserTagsUniqueName = "user_tags_unique_name"
)
