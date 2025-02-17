package model

import (
	"fmt"
)

const (
	// CreateUsers is the SQL query to create the users table
	CreateUsers = `
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    salt VARCHAR(255) NOT NULL,
	encrypted_key TEXT NOT NULL,
    birthdate TIMESTAMP,
    joined_at TIMESTAMP NOT NULL DEFAULT NOW(),
	update_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);
`

	// CreateUserFailedLogInAttempts is the SQL query to create the user_failed_log_in_attempts table
	CreateUserFailedLogInAttempts = `
CREATE TABLE IF NOT EXISTS user_failed_log_in_attempts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    ip_address VARCHAR(15) NOT NULL,
    bad_password BOOLEAN,
    bad_2fa_code BOOLEAN,
    attempted_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// CreateUser2FA is the SQL query to create the user_2fa table
	CreateUser2FA = `
CREATE TABLE IF NOT EXISTS user_2fa (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	revoked_at TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// CreateUser2FARecoveryCodes is the SQL query to create the user_2fa_recovery_codes table
	CreateUser2FARecoveryCodes = `
CREATE TABLE IF NOT EXISTS user_2fa_recovery_codes (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    code VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	used_at TIMESTAMP,
    revoked_at TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// CreateUser2FAEmailCodes is the SQL query to create the user_2fa_email_codes table
	CreateUser2FAEmailCodes = `
CREATE TABLE IF NOT EXISTS user_2fa_email_codes (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	code VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	expires_at TIMESTAMP NOT NULL,
	used_at TIMESTAMP,
	revoked_at TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// CreateUserNotes is the SQL query to create the user_notes table
	CreateUserNotes = `
CREATE TABLE IF NOT EXISTS user_notes (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL,
    color VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	pinned_at TIMESTAMP,
	starred_at TIMESTAMP,
	archived_at TIMESTAMP,
	trashed_at TIMESTAMP,
	deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// CreateUserNoteVersions is the SQL query to create the user_note_versions table
	CreateUserNoteVersions = `
CREATE TABLE IF NOT EXISTS user_note_versions (
    id SERIAL PRIMARY KEY,
    user_note_id BIGINT NOT NULL,
    encrypted_content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	deleted_at TIMESTAMP,
    FOREIGN KEY (user_note_id) REFERENCES user_notes(id) ON DELETE CASCADE
);
`
)

var (
	// CreateUserUsernames is the SQL query to create the user_usernames table
	CreateUserUsernames = fmt.Sprintf(
		`
CREATE TABLE IF NOT EXISTS user_usernames (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    username VARCHAR(50) NOT NULL,
    assigned_at TIMESTAMP NOT NULL DEFAULT NOW(),
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_usernames (username) WHERE revoked_at IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_usernames (user_id) WHERE revoked_at IS NULL;
`, UserUsernamesUniqueUsername, UserUsernamesUniqueUserID,
	)

	// CreateUserEmails is the SQL query to create the user_emails table
	CreateUserEmails = fmt.Sprintf(
		`
CREATE TABLE IF NOT EXISTS user_emails (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    email VARCHAR(100) NOT NULL,
    assigned_at TIMESTAMP NOT NULL DEFAULT NOW(),
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_emails (email) WHERE revoked_at IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_emails (user_id) WHERE revoked_at IS NULL;
`, UserEmailsUniqueEmail, UserEmailsUniqueUserID,
	)

	// CreateUserPasswordHashes is the SQL query to create the user_password_hashes table
	CreateUserPasswordHashes = fmt.Sprintf(
		`
CREATE TABLE IF NOT EXISTS user_password_hashes (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    assigned_at TIMESTAMP NOT NULL DEFAULT NOW(),
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_password_hashes (user_id) WHERE revoked_at IS NULL;
`, UserPasswordHashesUniqueUserID,
	)

	// CreateUserPhoneNumbers is the SQL query to create the user_phone_numbers table
	CreateUserPhoneNumbers = fmt.Sprintf(
		`
CREATE TABLE IF NOT EXISTS user_phone_numbers (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    assigned_at TIMESTAMP NOT NULL DEFAULT NOW(),
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_phone_numbers (phone_number) WHERE revoked_at IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_phone_numbers (user_id) WHERE revoked_at IS NULL;
`, UserPhoneNumbersUniquePhoneNumber, UserPhoneNumbersUniqueUserID,
	)

	// CreateUserTags is the SQL query to create the user_tags table
	CreateUserTags = fmt.Sprintf(
		`
CREATE TABLE IF NOT EXISTS user_tags (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name VARCHAR(50) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_tags (user_id, name) WHERE deleted_at IS NULL;
`, UserTagsUniqueUserIDName,
	)

	// CreateUserResetPasswords is the SQL query to create the user_reset_passwords table
	CreateUserResetPasswords = fmt.Sprintf(
		`
CREATE TABLE IF NOT EXISTS user_reset_passwords (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    reset_token VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_reset_passwords (user_id) WHERE revoked_at IS NULL;
`, UserResetPasswordsUniqueUserID,
	)

	// CreateUserEmailVerifications is the SQL query to create the user_email_verifications table
	CreateUserEmailVerifications = fmt.Sprintf(
		`
CREATE TABLE IF NOT EXISTS user_email_verifications (
    id BIGSERIAL PRIMARY KEY,
    user_email_id BIGINT NOT NULL,
    verification_token VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    verified_at TIMESTAMP,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_email_id) REFERENCES user_emails(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_email_verifications (user_email_id) WHERE revoked_at IS NULL;
`, UserEmailVerificationsUniqueUserEmailID,
	)

	// CreateUserPhoneNumberVerifications is the SQL query to create the user_phone_number_verifications table
	CreateUserPhoneNumberVerifications = fmt.Sprintf(
		`
CREATE TABLE IF NOT EXISTS user_phone_number_verifications (
    id BIGSERIAL PRIMARY KEY,
    user_phone_number_id BIGINT NOT NULL,
    verification_code VARCHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    verified_at TIMESTAMP,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_phone_number_id) REFERENCES user_phone_numbers(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_phone_number_verifications (user_phone_number_id) WHERE revoked_at IS NULL;
`, UserPhoneNumberVerificationsUniqueUserPhoneNumberID,
	)

	// CreateUserRefreshTokens is the SQL query to create the user_refresh_tokens table
	CreateUserRefreshTokens = fmt.Sprintf(
		`
CREATE TABLE IF NOT EXISTS user_refresh_tokens (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    parent_user_refresh_token_id BIGINT,
    ip_address VARCHAR(15) NOT NULL,
    issued_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (parent_user_refresh_token_id) REFERENCES user_refresh_tokens(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_refresh_tokens (parent_user_refresh_token_id) WHERE parent_user_refresh_token_id IS NOT NULL;
`, UserRefreshTokensUniqueParentUserRefreshTokenID,
	)

	// CreateUserAccessTokens is the SQL query to create the user_access_tokens table
	CreateUserAccessTokens = fmt.Sprintf(
		`
CREATE TABLE IF NOT EXISTS user_access_tokens (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    user_refresh_token_id BIGINT NOT NULL,
    issued_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (user_refresh_token_id) REFERENCES user_refresh_tokens(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_access_tokens (user_refresh_token_id);
`, UserAccessTokensUniqueUserRefreshTokenID,
	)

	// CreateUser2FATOTP is the SQL query to create the user_2fa_totp table
	CreateUser2FATOTP = fmt.Sprintf(
		`
CREATE TABLE IF NOT EXISTS user_2fa_totp (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    secret VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    verified_at TIMESTAMP,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_2fa_totp (user_id) WHERE revoked_at IS NULL;
`, User2FATOTPUniqueUserID,
	)

	// CreateUserNoteTags is the SQL query to create the user_note_tags table
	CreateUserNoteTags = fmt.Sprintf(
		`
CREATE TABLE IF NOT EXISTS user_note_tags (
    id BIGSERIAL PRIMARY KEY,
    user_note_id BIGINT NOT NULL,
    user_tag_id BIGINT NOT NULL,
    assigned_at TIMESTAMP NOT NULL DEFAULT NOW(),
	deleted_at TIMESTAMP,
    FOREIGN KEY (user_note_id) REFERENCES user_notes(id) ON DELETE CASCADE,
    FOREIGN KEY (user_tag_id) REFERENCES user_tags(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_note_tags (user_note_id, user_tag_id) WHERE deleted_at IS NULL;
`, UserNoteTagsUniqueUserNoteIDUserTagID,
	)
)
