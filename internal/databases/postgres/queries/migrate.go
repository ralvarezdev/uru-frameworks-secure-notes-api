package queries

import (
	"fmt"
	internalpostgresconstraints "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/constraints"
)

var (
	// UsersMigrate is the SQL query to create the users table
	UsersMigrate = `
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    salt VARCHAR(255) NOT NULL,
    birthdate TIMESTAMP,
    joined_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
`

	// UserUsernamesMigrate is the SQL query to create the user_usernames table
	UserUsernamesMigrate = fmt.Sprintf(
		`
CREATE TABLE IF NOT EXISTS user_usernames (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    username VARCHAR(50) NOT NULL,
    assigned_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_usernames (username) WHERE revoked_at IS NULL;
`, internalpostgresconstraints.UserUsernamesUniqueUsername,
	)

	// UserPasswordHashesMigrate is the SQL query to create the user_password_hashes table
	UserPasswordHashesMigrate = `
CREATE TABLE IF NOT EXISTS user_password_hashes (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    assigned_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// UserResetPasswordMigrate is the SQL query to create the UserResetPassword table
	UserResetPasswordMigrate = `
CREATE TABLE IF NOT EXISTS user_reset_passwords (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    reset_token VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// UserEmailsMigrate is the SQL query to create the user_emails table
	UserEmailsMigrate = fmt.Sprintf(
		`
CREATE TABLE IF NOT EXISTS user_emails (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    email VARCHAR(100) NOT NULL,
    assigned_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS %s ON user_emails (email) WHERE revoked_at IS NULL;
`, internalpostgresconstraints.UserEmailsUniqueEmail,
	)

	// UserEmailVerificationsMigrate is the SQL query to create the user_email_verifications table
	UserEmailVerificationsMigrate = `
CREATE TABLE IF NOT EXISTS user_email_verifications (
    id BIGSERIAL PRIMARY KEY,
    user_email_id BIGINT NOT NULL,
    verification_token VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    verified_at TIMESTAMP,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_email_id) REFERENCES user_emails(id) ON DELETE CASCADE
);
`

	// UserPhoneNumbersMigrate is the SQL query to create the user_phone_numbers table
	UserPhoneNumbersMigrate = `
CREATE TABLE IF NOT EXISTS user_phone_numbers (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    assigned_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// UserPhoneNumberVerificationsMigrate is the SQL query to create the user_phone_number_verifications table
	UserPhoneNumberVerificationsMigrate = `
CREATE TABLE IF NOT EXISTS user_phone_number_verifications (
    id BIGSERIAL PRIMARY KEY,
    user_phone_number_id BIGINT NOT NULL,
    verification_code VARCHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    verified_at TIMESTAMP,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_phone_number_id) REFERENCES user_phone_numbers(id) ON DELETE CASCADE
);
`

	// UserTokenSeedsMigrate is the SQL query to create the user_token_seeds table
	UserTokenSeedsMigrate = `
CREATE TABLE IF NOT EXISTS user_token_seeds (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    token_seed VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// UserFailedLogInAttemptsMigrate is the SQL query to create the user_failed_log_in_attempts table
	UserFailedLogInAttemptsMigrate = `
CREATE TABLE IF NOT EXISTS user_failed_log_in_attempts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    user_token_seed_id BIGINT,
    ipv4_address VARCHAR(15) NOT NULL,
    bad_password BOOLEAN,
    bad_2fa_code BOOLEAN,
    attempted_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (user_token_seed_id) REFERENCES user_token_seeds(id)
);
`

	// UserRefreshTokensMigrate is the SQL query to create the user_refresh_tokens table
	UserRefreshTokensMigrate = `
CREATE TABLE IF NOT EXISTS user_refresh_tokens (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    parent_refresh_token_id BIGINT,
    user_token_seed_id BIGINT,
    ipv4_address VARCHAR(15) NOT NULL,
    issued_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (parent_refresh_token_id) REFERENCES user_refresh_tokens(id) ON DELETE CASCADE,
    FOREIGN KEY (user_token_seed_id) REFERENCES user_token_seeds(id) ON DELETE CASCADE
);
`

	// UserAccessTokensMigrate is the SQL query to create the user_access_tokens table
	UserAccessTokensMigrate = `
CREATE TABLE IF NOT EXISTS user_access_tokens (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    user_refresh_token_id BIGINT NOT NULL,
    issued_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (user_refresh_token_id) REFERENCES user_refresh_tokens(id) ON DELETE CASCADE
);
`

	// UserTOTPsMigrate is the SQL query to create the user_totps table
	UserTOTPsMigrate = `
CREATE TABLE IF NOT EXISTS user_totps (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    secret VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    verified_at TIMESTAMP,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// UserTOTPRecoveryCodesMigrate is the SQL query to create the user_totp_recovery_codes table
	UserTOTPRecoveryCodesMigrate = `
CREATE TABLE IF NOT EXISTS user_totp_recovery_codes (
    id BIGSERIAL PRIMARY KEY,
    user_totp_id BIGINT NOT NULL,
    code VARCHAR(255) UNIQUE NOT NULL,
    revoked_at TIMESTAMP,
    FOREIGN KEY (user_totp_id) REFERENCES user_totps(id) ON DELETE CASCADE
);
`

	// TagsMigrate is the SQL query to create the tags table
	TagsMigrate = `
CREATE TABLE IF NOT EXISTS tags (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name VARCHAR(50) UNIQUE NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// NotesMigrate is the SQL query to create the notes table
	NotesMigrate = `
CREATE TABLE IF NOT EXISTS notes (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    is_pinned BOOLEAN,
    title VARCHAR(255) NOT NULL,
    color VARCHAR(50),
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	// NoteTagsMigrate is the SQL query to create the note_tags table
	NoteTagsMigrate = `
CREATE TABLE IF NOT EXISTS note_tags (
    id BIGSERIAL PRIMARY KEY,
    note_id BIGINT NOT NULL,
    tag_id BIGINT NOT NULL,
    assigned_at TIMESTAMP NOT NULL,
    FOREIGN KEY (note_id) REFERENCES notes(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);
`

	// NoteVersionsMigrate is the SQL query to create the note_versions table
	NoteVersionsMigrate = `
CREATE TABLE IF NOT EXISTS note_versions (
    id SERIAL PRIMARY KEY,
    note_id BIGINT NOT NULL,
    encrypted_body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (note_id) REFERENCES notes(id) ON DELETE CASCADE
);
`
)
