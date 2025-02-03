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
    deleted_at TIMESTAMP
);
`

	// CreateUserPasswordHashes is the SQL query to create the user_password_hashes table
	CreateUserPasswordHashes = `
CREATE TABLE IF NOT EXISTS user_password_hashes (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    assigned_at TIMESTAMP NOT NULL DEFAULT NOW(),
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// CreateUserResetPasswords is the SQL query to create the user_reset_passwords table
	CreateUserResetPasswords = `
CREATE TABLE IF NOT EXISTS user_reset_passwords (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    reset_token VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// CreateUserEmailVerifications is the SQL query to create the user_email_verifications table
	CreateUserEmailVerifications = `
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
`

	// CreateUserPhoneNumberVerifications is the SQL query to create the user_phone_number_verifications table
	CreateUserPhoneNumberVerifications = `
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

	// CreateUserRefreshTokens is the SQL query to create the user_refresh_tokens table
	CreateUserRefreshTokens = `
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
`

	// CreateUserAccessTokens is the SQL query to create the user_access_tokens table
	CreateUserAccessTokens = `
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
`

	// CreateUserTOTPs is the SQL query to create the user_totps table
	CreateUserTOTPs = `
CREATE TABLE IF NOT EXISTS user_totps (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    secret VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    verified_at TIMESTAMP,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// CreateUserTOTPRecoveryCodes is the SQL query to create the user_totp_recovery_codes table
	CreateUserTOTPRecoveryCodes = `
CREATE TABLE IF NOT EXISTS user_totp_recovery_codes (
    id BIGSERIAL PRIMARY KEY,
    user_totp_id BIGINT NOT NULL,
    code VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_totp_id) REFERENCES user_totps(id) ON DELETE CASCADE
);
`

	// CreateNotes is the SQL query to create the notes table
	CreateNotes = `
CREATE TABLE IF NOT EXISTS notes (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    is_pinned BOOLEAN,
    title VARCHAR(255) NOT NULL,
    color VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// CreateNoteTags is the SQL query to create the note_tags table
	CreateNoteTags = `
CREATE TABLE IF NOT EXISTS note_tags (
    id BIGSERIAL PRIMARY KEY,
    note_id BIGINT NOT NULL,
    tag_id BIGINT NOT NULL,
    assigned_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (note_id) REFERENCES notes(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);
`

	// CreateNoteVersions is the SQL query to create the note_versions table
	CreateNoteVersions = `
CREATE TABLE IF NOT EXISTS note_versions (
    id SERIAL PRIMARY KEY,
    note_id BIGINT NOT NULL,
    encrypted_body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (note_id) REFERENCES notes(id) ON DELETE CASCADE
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
`, UserUsernamesUniqueUsername,
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
`, UserEmailsUniqueEmail,
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
`, UserPhoneNumbersUniquePhoneNumber,
	)

	// CreateTags is the SQL query to create the tags table
	CreateTags = fmt.Sprintf(
		`
CREATE TABLE IF NOT EXISTS tags (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name VARCHAR(50) UNIQUE NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS %s ON tags (user_id, name);
`, UserTagsUniqueName,
	)
)
