package model

const (
	// CreateRevokeUserEmailVerificationTokenProc is the query to create the stored procedure to revoke user email verification token
	CreateRevokeUserEmailVerificationTokenProc = `
CREATE OR REPLACE PROCEDURE revoke_user_email_verification_token(
	IN in_user_email_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_email_verifications table
	UPDATE
		user_email_verifications
	SET
		revoked_at = NOW()
	WHERE
		user_email_verifications.user_email_id = in_user_email_id
	AND
		user_email_verifications.revoked_at IS NULL;
END;	
$$;
`

	// CreateSendEmailVerificationTokenProc is the query to create the stored procedure to send email verification token
	CreateSendEmailVerificationTokenProc = `
CREATE OR REPLACE PROCEDURE send_email_verification_token(
	IN in_user_id BIGINT,
	IN in_user_email_verification_token VARCHAR,
	IN in_user_email_verification_token_expires_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
DECLARE
	out_user_email_id BIGINT;
BEGIN
	-- Select the user email by user ID
	CALL get_user_email_id(in_user_id, out_user_email_id);

	-- Revoke the user email verification token
	CALL revoke_user_email_verification_token(out_user_email_id);

	-- Insert into user_email_verifications table
	INSERT INTO user_email_verifications (
		user_email_id,
		verification_token,
		expires_at
	)
	VALUES (
		out_user_email_id,
		in_user_email_verification_token,
		in_user_email_verification_token_expires_at
	);	
END;
$$;
`

	// CreateSignUpProc is the query to create the stored procedure to sign-up
	CreateSignUpProc = `
CREATE OR REPLACE PROCEDURE sign_up(
	IN in_user_first_name VARCHAR,
	IN in_user_last_name VARCHAR,
	IN in_user_salt VARCHAR,
	IN in_user_encrypted_key TEXT, 
	IN in_user_username VARCHAR,
	IN in_user_email VARCHAR,
	IN in_user_password_hash VARCHAR,
	IN in_user_email_verification_token VARCHAR,
	IN in_user_email_verification_token_expires_at TIMESTAMP,
	OUT out_user_id BIGINT
)
LANGUAGE plpgsql
AS $$
DECLARE
	out_user_email_id BIGINT;
BEGIN
	-- Insert into users table
	INSERT INTO users (
		first_name,
		last_name,
		salt,
		encrypted_key
	) 
	VALUES (
		in_user_first_name, 
		in_user_last_name, 
		in_user_salt,
		in_user_encrypted_key
	)
	RETURNING 
		id INTO out_user_id;

	-- Insert into user_usernames table
	INSERT INTO user_usernames (
		user_id, 
		username
	)
	VALUES (
		out_user_id, 
		in_user_username
	);

	-- Insert into user_emails table
	INSERT INTO user_emails (
		user_id, 
		email
	)
	VALUES (
		out_user_id, 
		in_user_email
	) 
	RETURNING
		id INTO out_user_email_id;

	-- Insert into user_password_hashes table
	INSERT INTO user_password_hashes (
		user_id, 
		password_hash
	) 
	VALUES (
		out_user_id, 
		in_user_password_hash
	);

	-- Insert into user_email_verifications table
	call send_email_verification_token(out_user_id, in_user_email_verification_token, in_user_email_verification_token_expires_at);
END;
$$;
`

	// CreateRevokeUserTOTPProc is the query to create the stored procedure to revoke user TOTP
	CreateRevokeUserTOTPProc = `
CREATE OR REPLACE PROCEDURE revoke_user_totp(
	IN in_user_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_totp_recovery_codes table
	UPDATE
		user_totp_recovery_codes
	SET
		revoked_at = NOW()
	FROM
		user_totps 
	WHERE
		user_totps.id = user_totp_recovery_codes.user_totp_id
	AND	
		user_totps.user_id = in_user_id
	AND	
		user_totp_recovery_codes.revoked_at IS NULL;

	-- Update the user_totps table
	UPDATE	
		user_totps
	SET	
		revoked_at = NOW()
	WHERE
		user_totps.user_id = in_user_id 
	AND	
		user_totps.revoked_at IS NULL;
END;
$$;
`

	// CreateGenerateUserTokensProc is the query to create the stored procedure to generate user tokens
	CreateGenerateUserTokensProc = `
CREATE OR REPLACE PROCEDURE generate_user_tokens(
	IN in_user_id BIGINT,
	IN in_user_parent_refresh_token_id BIGINT,
	IN in_user_ip_address VARCHAR,
	IN in_user_refresh_token_expires_at TIMESTAMP,
	IN in_user_access_token_expires_at TIMESTAMP,
	OUT out_user_refresh_token_id BIGINT,
	OUT out_user_access_token_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Insert into user_refresh_tokens table
 	INSERT INTO user_refresh_tokens (
  		user_id,
  		parent_user_refresh_token_id,
  		ip_address,
		expires_at
	)
	VALUES (
		in_user_id,
		in_user_parent_refresh_token_id,
		in_user_ip_address,
		in_user_refresh_token_expires_at
	)
	RETURNING
		id INTO out_user_refresh_token_id;

	-- Insert into user_access_tokens table
	INSERT INTO user_access_tokens (
		user_id,
		user_refresh_token_id,
		expires_at
	)
	VALUES (
  		in_user_id,
  		out_user_refresh_token_id,
  		in_user_access_token_expires_at
	)
	RETURNING
  		id INTO out_user_access_token_id;
END;
$$;
`

	// CreateRevokeUserTokensByIDProc is the query to create the stored procedure to revoke user tokens by ID
	CreateRevokeUserTokensByIDProc = `
CREATE OR REPLACE PROCEDURE revoke_user_tokens_by_id(
	IN in_user_id BIGINT,
	IN in_user_refresh_token_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_refresh_tokens table
	UPDATE
		user_refresh_tokens
	SET	
		revoked_at = NOW()
	WHERE
		user_refresh_tokens.id = in_user_refresh_token_id
	AND
		user_refresh_tokens.user_id = in_user_id
	AND 
		user_refresh_tokens.revoked_at IS NULL;

	-- Update the user_access_tokens table
	UPDATE
		user_access_tokens
	SET
		revoked_at = NOW()
	WHERE
		user_access_tokens.user_refresh_token_id = in_user_refresh_token_id
	AND
		user_access_tokens.user_id = user_id
	AND
		user_access_tokens.revoked_at IS NULL;
END;
$$;
`

	// CreateRefreshTokenProc is the query to create the stored procedure to refresh token
	CreateRefreshTokenProc = `
CREATE OR REPLACE PROCEDURE refresh_token(
	IN in_user_id BIGINT,
	IN in_old_user_refresh_token_id BIGINT,
	IN in_user_ip_address VARCHAR,
	IN in_new_user_refresh_token_expires_at TIMESTAMP,
	IN in_new_user_access_token_expires_at TIMESTAMP,
	OUT out_new_user_refresh_token_id BIGINT,
	OUT out_new_user_access_token_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Revoke the old user refresh token and access token
	CALL revoke_user_tokens_by_id(in_user_id, in_old_user_refresh_token_id);

	-- Generate new tokens
	CALL generate_user_tokens(in_user_id, in_old_user_refresh_token_id, in_user_ip_address, in_new_user_refresh_token_expires_at, in_new_user_access_token_expires_at, out_new_user_refresh_token_id, out_new_user_access_token_id);
END;
$$;
`

	// CreateRevokeUserTokensProc is the query to create the stored procedure to revoke user tokens
	CreateRevokeUserTokensProc = `
CREATE OR REPLACE PROCEDURE revoke_user_tokens(
	IN in_user_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_refresh_tokens table
	UPDATE
		user_refresh_tokens
	SET	
		revoked_at = NOW()
	WHERE
		user_refresh_tokens.user_id = in_user_id
	AND
		user_refresh_tokens.revoked_at IS NULL;
	
	-- Update the user_access_tokens table
	UPDATE
		user_access_tokens
	SET	
		revoked_at = NOW()
	WHERE
		user_access_tokens.user_id = in_user_id
	AND
		user_access_tokens.revoked_at IS NULL;
END;
$$;
`

	// CreateGetAccessTokenIDByRefreshTokenIDProc is the query to create the stored procedure to get access token ID by refresh token ID
	CreateGetAccessTokenIDByRefreshTokenIDProc = `
CREATE OR REPLACE PROCEDURE get_access_token_id_by_refresh_token_id(
	IN in_user_refresh_token_id BIGINT,
	OUT out_user_access_token_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user_access_tokens table
	SELECT
		user_access_tokens.id
	INTO
		out_user_access_token_id
	FROM
		user_access_tokens
	WHERE
		user_access_tokens.user_refresh_token_id = in_user_refresh_token_id
	AND
		user_access_tokens.revoked_at IS NULL;
END;
$$;
`

	// CreatePreLogInProc is the query to create the stored procedure to pre-log in
	CreatePreLogInProc = `
CREATE OR REPLACE PROCEDURE pre_log_in(
	IN in_user_username VARCHAR,
	OUT out_user_id BIGINT,
	OUT out_user_password_hash VARCHAR,
	OUT out_user_salt VARCHAR,
	OUT out_user_encrypted_key TEXT,
	OUT out_user_totp_id BIGINT,
	OUT out_user_totp_secret VARCHAR
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user ID and password hash by username
	SELECT
		users.id,
		users.salt,
		users.encrypted_key,
		user_password_hashes.password_hash
	INTO
		out_user_id,
		out_user_salt,	
		out_user_encrypted_key,	
		out_user_password_hash
	FROM
		users
	INNER JOIN
		user_usernames ON users.id = user_usernames.user_id
	INNER JOIN
		user_password_hashes ON users.id = user_password_hashes.user_id
	WHERE
		user_usernames.username = in_user_username
	AND
		user_usernames.revoked_at IS NULL
	AND
		users.deleted_at IS NULL
	AND
		user_password_hashes.revoked_at IS NULL;

	-- Select the TOTP ID and secret by user ID
	SELECT
		user_totps.id,
		user_totps.secret
	INTO
		out_user_totp_id,
		out_user_totp_secret
	FROM
		user_totps
	WHERE
		user_totps.user_id = out_user_id
	AND
		user_totps.revoked_at IS NULL
	AND
		user_totps.verified_at IS NOT NULL;

	-- If the user doesn't have a TOTP, set the TOTP ID and secret to NULL
	IF NOT FOUND THEN
		out_user_totp_id = NULL;
		out_user_totp_secret = NULL;
	END IF;
END;
$$;
`

	// CreateRegisterFailedLogInAttemptProc is the query to create the stored procedure to register failed log in attempt
	CreateRegisterFailedLogInAttemptProc = `
CREATE OR REPLACE PROCEDURE register_failed_log_in_attempt(
	IN in_user_id BIGINT,
	IN in_user_ip_address VARCHAR,
	IN in_bad_password BOOLEAN,
	IN in_bad_2fa_code BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Insert into user_failed_log_in_attempts table
	INSERT INTO user_failed_log_in_attempts (
		user_id,
		ip_address,
		bad_password,
		bad_2fa_code
	)
	VALUES (
		in_user_id,
		in_user_ip_address,
		in_bad_password,
		in_bad_2fa_code
	);
END;
$$;
`

	// CreateGetUserTOTPProc is the query to create the stored procedure to get user TOTP by user ID
	CreateGetUserTOTPProc = `
CREATE OR REPLACE PROCEDURE get_user_totp(
	IN in_user_id BIGINT,
	OUT out_user_totp_id BIGINT,
	OUT out_user_totp_secret VARCHAR,
	OUT out_user_totp_verified_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the TOTP ID, secret, and verified at by user ID
	SELECT
		user_totps.id,
		user_totps.secret,
		user_totps.verified_at
	INTO
		out_user_totp_id,
		out_user_totp_secret,
		out_user_totp_verified_at
	FROM
		user_totps
	WHERE
		user_totps.user_id = in_user_id
	AND
		user_totps.revoked_at IS NULL;
END;
$$;
`

	// CreateGetUserEmailProc is the query to create the stored procedure to get user email by user ID
	CreateGetUserEmailProc = `
CREATE OR REPLACE PROCEDURE get_user_email(
	IN in_user_id BIGINT,
	OUT out_user_email VARCHAR
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user email by user ID
	SELECT
		user_emails.email
	INTO
		out_user_email
	FROM
		user_emails
	WHERE
		user_emails.user_id = in_user_id
	AND
		user_emails.revoked_at IS NULL;
END;
$$;
`

	// CreateGetUserEmailIDProc is the query to create the stored procedure to get user email ID by user ID
	CreateGetUserEmailIDProc = `
CREATE OR REPLACE PROCEDURE get_user_email_id(
	IN in_user_id BIGINT,
	OUT out_email_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN	
	-- Select the user email ID by user ID
	SELECT
		user_emails.id
	INTO
		out_email_id
	FROM
		user_emails
	WHERE
		user_emails.user_id = in_user_id
	AND
		user_emails.revoked_at IS NULL;
END;	
$$;
`

	// CreateGenerateTOTPUrlProc is the query to create the stored procedure to generate TOTP URL
	CreateGenerateTOTPUrlProc = `
CREATE OR REPLACE PROCEDURE generate_totp_url(
	IN in_user_id BIGINT,
	IN in_new_user_totp_secret VARCHAR,
	OUT out_user_email VARCHAR,
	OUT out_old_user_totp_id BIGINT,
	OUT out_old_user_totp_secret VARCHAR,
	OUT out_old_user_totp_verified_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user email by user ID
	CALL get_user_email(in_user_id, out_user_email);

	-- Select the TOTP ID, secret, and verified at by user ID
	CALL get_user_totp(in_user_id, out_old_user_totp_id, out_old_user_totp_secret, out_old_user_totp_verified_at);

	-- If the TOTP is not verified, revoke it
	IF out_old_totp_verified_at IS NULL THEN
		CALL revoke_user_totp(in_user_id);
	END IF;

	-- If the TOTP wasn't active or verified, insert a new TOTP
	IF out_old_totp_id IS NULL OR out_old_totp_verified_at IS NULL THEN
		-- Insert into user_totps table
		INSERT INTO user_totps (
			user_id,
			secret
		)
		VALUES (
			in_user_id,
			in_new_user_totp_secret
		);
	END IF;
END;
$$;
`

	// CreateIsRefreshTokenValidProc is the query to create the stored procedure to check if the refresh token is valid
	CreateIsRefreshTokenValidProc = `
CREATE OR REPLACE PROCEDURE is_refresh_token_valid(
	IN in_user_refresh_token_id BIGINT,
	OUT out_user_refresh_token_expires_at TIMESTAMP,
	OUT out_user_refresh_token_found BOOLEAN,
	OUT out_user_refresh_token_is_expired BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user refresh token ID by user ID
	SELECT
		user_refresh_tokens.expires_at
	INTO
		out_user_refresh_token_expires_at
	FROM
		user_refresh_tokens
	WHERE
		user_refresh_tokens.id = in_user_refresh_token_id
	AND
		user_refresh_tokens.revoked_at IS NULL;

	IF out_expires_at IS NOT NULL THEN
		out_user_refresh_token_found = TRUE;
		out_user_refresh_token_is_expired = out_user_refresh_token_expires_at < NOW();
	END IF;
END;
$$;
`

	// CreateIsAccessTokenValidProc is the query to create the stored procedure to check if the access token is valid
	CreateIsAccessTokenValidProc = `
CREATE OR REPLACE PROCEDURE is_access_token_valid(
	IN in_user_access_token_id BIGINT,
	OUT out_user_access_token_expires_at TIMESTAMP,
	OUT out_user_access_token_found BOOLEAN,
	OUT out_user_access_token_is_expired BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user access token ID by user ID
	SELECT
		user_access_tokens.expires_at
	INTO
		out_user_access_token_expires_at
	FROM
		user_access_tokens
	WHERE
		user_access_tokens.id = in_user_access_token_id
	AND
		user_access_tokens.revoked_at IS NULL;
	
	IF out_expires_at IS NOT NULL THEN
		out_user_access_token_found = TRUE;
		out_user_access_token_is_expired = out_user_access_token_expires_at < NOW();
	END IF;
END;	
$$;
`

	// CreateRevokeUserTOTPRecoveryCodeProc is the query to create the stored procedure to revoke user TOTP recovery code
	CreateRevokeUserTOTPRecoveryCodeProc = `
CREATE OR REPLACE PROCEDURE revoke_user_totp_recovery_code(
	IN in_user_totp_id BIGINT,
	IN in_user_totp_recovery_code VARCHAR
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_totp_recovery_codes table
	UPDATE
		user_totp_recovery_codes
	SET
		revoked_at = NOW()
	WHERE
		user_totp_recovery_codes.user_totp_id = in_user_totp_id
	AND
		user_totp_recovery_codes.recovery_code = in_user_totp_recovery_code
	AND
		user_totp_recovery_codes.revoked_at IS NULL;
END;
$$;
`

	// CreateVerifyTOTPProc is the query to create verify TOTP
	CreateVerifyTOTPProc = `
CREATE OR REPLACE PROCEDURE verify_totp(
	IN in_user_totp_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_totps table
	UPDATE
		user_totps
	SET
		verified_at = NOW()
	WHERE
		user_totps.id = in_user_totp_id
	AND 
		user_totps.verified_at IS NULL
	AND
		user_totps.revoked_at IS NULL;
END;
$$;
`

	// CreateVerifyEmailProc is the query to create the stored procedure to verify email
	CreateVerifyEmailProc = `
CREATE OR REPLACE PROCEDURE verify_email(
	IN in_user_email_verification_token VARCHAR,
	OUT out_user_id BIGINT,
	OUT out_invalid_token BOOLEAN
)
LANGUAGE plpgsql
AS $$
DECLARE 
	out_user_email_verification_id BIGINT;
BEGIN
	-- Select the user email ID by verification token
	SELECT
		user_email_verifications.id,
		user_emails.user_id
	INTO
		out_user_email_verification_id
		out_user_id
	FROM
		user_email_verifications
	INNER JOIN
		user_emails ON user_email_verifications.user_email_id = user_emails.id
	WHERE
		user_email_verifications.verification_token = in_user_email_verification_token
	AND
		user_email_verifications.expires_at > NOW()
	AND
		user_email_verifications.revoked_at IS NULL
	AND
		user_email_verifications.verified_at IS NULL;

	-- Check if the user email verification token is invalid
	IF out_user_email_verification_id IS NULL THEN
		out_invalid_token = TRUE;
	ELSE
		out_invalid_token = FALSE;

		-- Update the user_email_verifications table
		UPDATE
			user_email_verifications
		SET
			verified_at = NOW()
		WHERE
			user_email_verifications.id = out_user_email_verification_id;
	END IF;
END;
$$;
`

	// CreateIsUserEmailVerifiedProc is the query to create the stored procedure to check if the user email is verified
	CreateIsUserEmailVerifiedProc = `
CREATE OR REPLACE PROCEDURE is_user_email_verified(
	IN in_user_id BIGINT,
	OUT out_user_email_is_verified BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user email by user ID
	SELECT
		user_emails.verified_at IS NOT NULL
	INTO
		out_user_email_is_verified
	FROM
		user_emails
	WHERE
		user_emails.user_id = in_user_id
	AND
		user_emails.revoked_at IS NULL;	
END;
$$;
`

	// CreatePreSendEmailVerificationTokenProc is the query to create the stored procedure to pre-send email verification token
	CreatePreSendEmailVerificationTokenProc = `
CREATE OR REPLACE PROCEDURE pre_send_email_verification_token(
	IN in_user_id BIGINT,
	OUT out_user_first_name VARCHAR,
	OUT out_user_last_name VARCHAR,
	OUT out_user_email VARCHAR,
	OUT out_user_email_is_verified BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user email by user ID
	SELECT
		users.first_name,
		users.last_name,
		user_emails.email,
		user_emails.verified_at IS NOT NULL
	INTO
		out_user_first_name,
		out_user_last_name,
		out_user_email,
		out_user_email_is_verified
	FROM
		user_emails
	INNER JOIN
		users ON user_emails.user_id = users.id
	WHERE
		user_emails.user_id = in_user_id
	AND
		user_emails.revoked_at IS NULL;
END;
$$;
`

	// CreateRevokeUserEmailProc is the query to create the stored procedure to revoke user email
	CreateRevokeUserEmailProc = `
CREATE OR REPLACE PROCEDURE revoke_user_email(
	IN in_user_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_emails table
	UPDATE
		user_emails
	SET
		revoked_at = NOW()
	WHERE
		user_emails.user_id = in_user_id
	AND
		user_emails.revoked_at IS NULL;

	-- Update the user_email_verifications table
	UPDATE
		user_email_verifications
	SET
		revoked_at = NOW()
	FROM 
		user_emails
	WHERE
		user_email_verifications.user_email_id = user_emails.id
	AND
		user_emails.user_id = in_user_id
	AND
		user_email_verifications.revoked_at IS NULL;
END;
$$;
`

	// CreateChangeEmailProc is the query to create the stored procedure to change email
	CreateChangeEmailProc = `
CREATE OR REPLACE PROCEDURE change_email(
	IN in_user_id BIGINT,
	IN in_new_user_email VARCHAR,
	IN in_user_email_verification_token VARCHAR,
	IN in_user_email_verification_token_expires_at TIMESTAMP,
	OUT out_user_first_name VARCHAR,
	OUT out_user_last_name VARCHAR
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Revoke the user email
	CALL revoke_user_email(in_user_id);

	-- Insert into user_emails table
	INSERT INTO user_emails (
		user_id,
		email
	)
	VALUES (
		in_user_id,
		in_new_user_email
	);

	-- Insert into user_email_verifications table
	CALL send_email_verification_token(in_user_id, in_user_email_verification_token, in_user_email_verification_token_expires_at);

	-- Select the user first name and last name by user ID
	SELECT
		users.first_name,
		users.last_name
	INTO
		out_user_first_name,
		out_user_last_name
	FROM
		users
	WHERE
		users.id = in_user_id;
END;
$$;
`

	// CreateRevokeUserResetPasswordTokenProc is the query to create the stored procedure to revoke user reset password token
	CreateRevokeUserResetPasswordTokenProc = `
CREATE OR REPLACE PROCEDURE revoke_user_reset_password_token(
	IN in_user_id BIGINT
)	
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_reset_passwords table
	UPDATE
		user_reset_passwords
	SET
		revoked_at = NOW()
	WHERE
		user_reset_passwords.user_id = in_user_id
	AND
		user_reset_passwords.revoked_at IS NULL;
END;
$$;
`

	// CreateForgotPasswordProc is the query to create the stored procedure to forgot password
	CreateForgotPasswordProc = `
CREATE OR REPLACE PROCEDURE forgot_password(
	IN in_user_username VARCHAR,
	IN in_user_reset_password_token VARCHAR,
	IN in_user_reset_password_token_expires_at TIMESTAMP,
	OUT out_user_id BIGINT,
	OUT out_user_first_name VARCHAR,
	OUT out_user_last_name VARCHAR,
	OUT out_user_email VARCHAR
)
LANGUAGE plpgsql	
AS $$
BEGIN
	-- Select the user ID, first name, last name, and email by username
	SELECT
		users.id,
		users.first_name,
		users.last_name,
		user_emails.email
	INTO
		out_user_id
		out_user_first_name,
		out_user_last_name,
		out_user_email
	FROM
		users
	INNER JOIN
		user_usernames ON users.id = user_usernames.user_id
	INNER JOIN
		user_emails ON users.id = user_emails.user_id
	WHERE
		user_usernames.username = in_user_username
	AND
		user_usernames.revoked_at IS NULL
	AND
		user_emails.revoked_at IS NULL
	AND
		users.deleted_at IS NULL;

	-- Revoke the user reset password token, if it exists and hasn't been revoked
	CALL revoke_user_reset_password_token(out_user_id);

	-- Insert into user_reset_passwords table
	INSERT INTO user_reset_passwords (
		user_id,
		reset_token,
		expires_at
	)
	VALUES (
		out_user_id,
		in_user_reset_password_token,
		in_user_reset_password_token_expires_at
	);
END;	
$$;
`

	// CreateRevokeUserPasswordHashProc is the query to create the stored procedure to revoke user password hash
	CreateRevokeUserPasswordHashProc = `
CREATE OR REPLACE PROCEDURE revoke_user_password_hash(
	IN in_user_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_password_hashes table
	UPDATE
		user_password_hashes
	SET
		revoked_at = NOW()
	WHERE
		user_password_hashes.user_id = in_user_id
	AND
		user_password_hashes.revoked_at IS NULL;
END;
$$;
`

	// CreateResetPasswordProc is the query to create the stored procedure to reset password
	CreateResetPasswordProc = `
CREATE OR REPLACE PROCEDURE reset_password(
	IN in_user_reset_password_token BIGINT,
	IN in_new_user_password_hash VARCHAR,
	OUT out_user_id BIGINT,
	OUT out_invalid_token BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user ID by reset password token
	SELECT
		user_reset_passwords.user_id
	INTO
		out_user_id
	FROM
		user_reset_passwords
	WHERE
		user_reset_passwords.reset_token = in_user_reset_password_token
	AND
		user_reset_passwords.expires_at > NOW()
	AND
		user_reset_passwords.revoked_at IS NULL;

	-- Check if the user ID exists
	IF out_user_id IS NULL THEN
		out_invalid_token = TRUE;
	ELSE
		out_invalid_token = FALSE;

		-- Revoke the user password hash
		call revoke_user_password_hash(out_user_id);
	
		-- Insert into user_password_hashes table
		INSERT INTO user_password_hashes (
			user_id,
			password_hash
		)
		VALUES (
			in_user_id,
			in_new_user_password_hash
		);
	
		-- Revoke the user reset password token
		CALL revoke_user_reset_password_token(out_user_id);

		-- Revoke the user tokens
		CALL revoke_user_tokens(out_user_id);
	END IF;
END;
$$;
`

	// CreateRevokeUserTokensExceptRefreshTokenIDProc is the query to create the stored procedure to revoke user tokens except refresh token ID
	CreateRevokeUserTokensExceptRefreshTokenIDProc = `
CREATE OR REPLACE PROCEDURE revoke_user_tokens_except_refresh_token_id(
	IN in_user_id BIGINT,
	IN in_user_refresh_token_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_refresh_tokens table
	UPDATE
		user_refresh_tokens
	SET	
		revoked_at = NOW()
	WHERE
		user_refresh_tokens.user_id = in_user_id
	AND 
		user_refresh_tokens.revoked_at IS NULL
	AND
		user_refresh_tokens.id != in_user_refresh_token_id;

	-- Update the user_access_tokens table
	UPDATE
		user_access_tokens
	SET
		revoked_at = NOW()
	WHERE
		user_access_tokens.user_id = in_user_id
	AND
		user_access_tokens.revoked_at IS NULL
	AND
		user_access_tokens.user_refresh_token_id != in_user_refresh_token_id;
END;
$$;
`

	// CreateGetUserPasswordHashProc is the query to create the stored procedure to get user password hash
	CreateGetUserPasswordHashProc = `
CREATE OR REPLACE PROCEDURE get_user_password_hash(
	IN in_user_id BIGINT,
	OUT out_user_password_hash VARCHAR
)	
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user password hash by user ID
	SELECT
		user_password_hashes.password_hash
	INTO
		out_user_password_hash
	FROM
		user_password_hashes
	WHERE
		user_password_hashes.user_id = in_user_id
	AND
		user_password_hashes.revoked_at IS NULL;
END;	
$$;
`

	// CreateChangePasswordProc is the query to create the stored procedure to change password
	CreateChangePasswordProc = `
CREATE OR REPLACE PROCEDURE change_password(
	IN in_user_id BIGINT,
	IN in_new_user_password_hash VARCHAR,
	IN in_user_refresh_token_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Revoke the user password hash
	CALL revoke_user_password_hash(in_user_id);

	-- Insert into user_password_hashes table
	INSERT INTO user_password_hashes (
		user_id,	
		password_hash
	)
	VALUES (
		in_user_id,
		in_new_user_password_hash
	);

	-- Revoke the user tokens, except the current access token and refresh token
	CALL revoke_user_tokens_except_refresh_token_id(in_user_id, in_user_refresh_token_id);
END;
$$;
`

	// CreateRevokeUserUsernameProc is the query to create the stored procedure to revoke user username
	CreateRevokeUserUsernameProc = `
CREATE OR REPLACE PROCEDURE revoke_user_username(
	IN in_user_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_usernames table
	UPDATE
		user_usernames
	SET
		revoked_at = NOW()	
	WHERE
		user_usernames.user_id = in_user_id
	AND
		user_usernames.revoked_at IS NULL;
END;
$$;
`

	// CreateRevokeUserPhoneNumberProc is the query to create the stored procedure to revoke user phone number
	CreateRevokeUserPhoneNumberProc = `
CREATE OR REPLACE PROCEDURE revoke_user_phone_number(
	IN in_user_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_phone_numbers table
	UPDATE
		user_phone_numbers
	SET
		revoked_at = NOW()
	WHERE
		user_phone_numbers.user_id = in_user_id
	AND
		user_phone_numbers.revoked_at IS NULL;

	-- Update the user_phone_number_verifications table
	UPDATE
		user_phone_number_verifications
	SET
		revoked_at = NOW()
	FROM
		user_phone_numbers
	WHERE
		user_phone_number_verifications.user_phone_number_id = user_phone_numbers.id
	AND
		user_phone_numbers.user_id = in_user_id
	AND
		user_phone_number_verifications.revoked_at IS NULL;
END;
$$;
`

	// CreateDeleteUserProc is the query to create the stored procedure to delete user
	CreateDeleteUserProc = `
CREATE OR REPLACE PROCEDURE delete_user(
	IN in_user_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Revoke the user username
	CALL revoke_user_username(in_user_id);

	-- Revoke the user email 
	CALL revoke_user_email(in_user_id);

	-- Revoke the user phone number
	CALL revoke_user_phone_number(in_user_id);

	-- Revoke the user password hash
	CALL revoke_user_password_hash(in_user_id);

	-- Revoke the user tokens
	CALL revoke_user_tokens(in_user_id);

	-- Revoke the user TOTP
	CALL revoke_user_totp(in_user_id);

	-- Revoke the user reset password token
	CALL revoke_user_reset_password_token(in_user_id);

	-- Update the users table
	UPDATE
		users
	SET
		deleted_at = NOW()
	WHERE
		users.id = in_user_id;

	-- Delete the user tags
	DELETE FROM
		tags
	WHERE
		tags.user_id = in_user_id;

	-- Delete the user note tags
	DELETE FROM
		note_tags
	INNER JOIN
		notes ON note_tags.note_id = notes.id
	WHERE
		notes.user_id = in_user_id;
	
	-- Delete the user note versions
	DELETE FROM
		note_versions
	INNER JOIN
		notes ON note_versions.note_id = notes.id
	WHERE
		notes.user_id = in_user_id;

	-- Delete the user notes
	DELETE FROM
		notes
	WHERE
		notes.user_id = in_user_id;
END;
$$;
`

	// CreateChangeUsernameProc is the query to create the stored procedure to change username
	CreateChangeUsernameProc = `
CREATE OR REPLACE PROCEDURE change_username(
	IN in_user_id BIGINT,
	IN in_new_user_username VARCHAR
)
LANGUAGE plpgsql	
AS $$
BEGIN
	-- Revoke the user username
	CALL revoke_user_username(in_user_id);

	-- Insert into user_usernames table
	INSERT INTO user_usernames (
		user_id,
		username
	)
	VALUES (
		in_user_id,
		in_new_user_username
	);
END;
$$;
`

	// CreateGetUserBasicInfoProc is the query to create the stored procedure to get user basic info
	CreateGetUserBasicInfoProc = `
CREATE OR REPLACE PROCEDURE get_user_basic_info(
	IN in_user_id BIGINT,
	OUT out_user_first_name VARCHAR,
	OUT out_user_last_name VARCHAR,
	OUT out_user_birthdate TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user first name, last name, and birthdate by user ID
	SELECT
		users.first_name,
		users.last_name,
		users.birthdate
	INTO
		out_user_first_name,
		out_user_last_name,
		out_user_birthdate
	FROM
		users
	WHERE
		users.id = in_user_id;
END;
$$;
`

	// CreateUpdateProfileProc is the query to create the stored procedure to update profile
	CreateUpdateProfileProc = `
CREATE OR REPLACE PROCEDURE update_profile(
	IN in_user_id BIGINT,
	IN in_user_first_name VARCHAR,
	IN in_user_last_name VARCHAR,
	IN in_user_birthdate TIMESTAMP
)
LANGUAGE plpgsql
AS $$
DECLARE 
	out_user_first_name VARCHAR;
	out_user_last_name VARCHAR;
	out_user_birthdate TIMESTAMP;
BEGIN
	-- Select the user first name, last name, and birthdate by user ID
	CALL get_user_basic_info(in_user_id, out_user_first_name, out_user_last_name, out_user_birthdate);

	-- Update the users table conditionally
	UPDATE
		users
	SET
		first_name = COALESCE(in_user_first_name, out_first_name),
		last_name = COALESCE(in_user_last_name, out_last_name),
		birthdate = COALESCE(in_user_birthdate, out_birthdate)
	WHERE
		users.id = in_user_id;
END;
$$;
`

	// CreateGetUserPhoneNumberProc is the query to create the stored procedure to get user phone number by user ID
	CreateGetUserPhoneNumberProc = `
CREATE OR REPLACE PROCEDURE get_user_phone_number(
	IN in_user_id BIGINT,
	OUT out_user_phone_number VARCHAR
)	
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user phone number by user ID
	SELECT
		user_phone_numbers.phone_number
	INTO
		out_user_phone_number
	FROM
		user_phone_numbers
	WHERE
		user_phone_numbers.user_id = in_user_id
	AND
		user_phone_numbers.revoked_at IS NULL;
END;
$$;
`

	// CreateGetUserUsernameProc is the query to create the stored procedure to get user username by user ID
	CreateGetUserUsernameProc = `
CREATE OR REPLACE PROCEDURE get_user_username(
	IN in_user_id BIGINT,
	OUT out_user_username VARCHAR
)	
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user username by user ID
	SELECT
		user_usernames.username
	INTO
		out_user_username
	FROM
		user_usernames
	WHERE
		user_usernames.user_id = in_user_id
	AND
		user_usernames.revoked_at IS NULL;
END;
$$;
`

	// CreateHasUserTOTPEnabledProc is the query to create the stored procedure to check if the user has TOTP enabled
	CreateHasUserTOTPEnabledProc = `
CREATE OR REPLACE PROCEDURE has_user_totp_enabled(
	IN in_user_id BIGINT,
	OUT out_user_has_totp_enabled BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user TOTP ID by user ID
	SELECT
		user_totps.verified_at IS NOT NULL
	INTO
		out_user_has_totp_enabled
	FROM
		user_totps
	WHERE
		user_totps.user_id = in_user_id
	AND
		user_totps.revoked_at IS NULL;
END;
$$;
`

	// CreateIsUserPhoneNumberVerifiedProc is the query to create the stored procedure to check if the user phone number is verified
	CreateIsUserPhoneNumberVerifiedProc = `
CREATE OR REPLACE PROCEDURE is_user_phone_number_verified(
	IN in_user_id BIGINT,
	OUT out_user_phone_number_is_verified BOOLEAN
)	
LANGUAGE plpgsql	
AS $$
BEGIN
	-- Select the user phone number by user ID
	SELECT
		user_phone_numbers.verified_at IS NOT NULL
	INTO
		out_user_phone_number_is_verified
	FROM
		user_phone_numbers
	WHERE
		user_phone_numbers.user_id = in_user_id
	AND
		user_phone_numbers.revoked_at IS NULL;
END;	
$$;
`

	// CreateGetMyProfileProc is the query to create the stored procedure to get my profile
	CreateGetMyProfileProc = `
CREATE OR REPLACE PROCEDURE get_my_profile(
	IN in_user_id BIGINT,
	OUT out_user_first_name VARCHAR,
	OUT out_user_last_name VARCHAR,
	OUT out_user_birthdate TIMESTAMP,
	OUT out_user_username VARCHAR,
	OUT out_user_email VARCHAR,
	OUT out_user_email_is_verified BOOLEAN,
	OUT out_user_phone_number VARCHAR,
	OUT out_user_phone_number_is_verified BOOLEAN,
	OUT out_user_has_totp_enabled BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user first name, last name, and birthdate by user ID
	CALL get_user_basic_info(in_user_id, out_user_first_name, out_user_last_name, out_user_birthdate);

	-- Select the user username by user ID
	CALL get_user_username(in_user_id, out_username);

	-- Get the user email
	CALL get_user_email(in_user_id, out_user_email);

	-- Select the user phone number by user ID
	CALL get_user_phone_number(in_user_id, out_user_phone_number);

	-- Check if the user email is verified
	CALL is_user_email_verified(in_user_id, out_user_email_is_verified);

	-- Check if the user phone number is verified
	CALL is_user_phone_number_verified(in_user_id, out_user_phone_number_is_verified);

	-- Check if the user has TOTP enabled
	CALL has_user_totp_enabled(in_user_id, out_user_has_totp_enabled);
END;
$$;
`

	// CreateCreateUserTagProc is the query to create the stored procedure to create user tag
	CreateCreateUserTagProc = `
CREATE OR REPLACE PROCEDURE create_user_tag(
	IN in_user_id BIGINT,
	IN in_user_tag_name VARCHAR,
	OUT out_user_tag_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Insert into tags table
	INSERT INTO tags (
		user_id,
		name
	)
	VALUES (
		in_user_id,
		in_user_tag_name
	)
	RETURNING
		id INTO out_user_tag_id;
END;
$$;
`

	// CreateUpdateUserTagProc is the query to create the stored procedure to update user tag
	CreateUpdateUserTagProc = `
CREATE OR REPLACE PROCEDURE update_user_tag(
	IN in_user_id BIGINT,
	IN in_user_tag_id BIGINT,
	IN in_user_tag_name VARCHAR
)
LANGUAGE plpgsql
AS $$
BEGIN	
	-- Update the tags table
	UPDATE
		tags
	SET
		name = in_user_tag_name
	WHERE
		tags.id = in_user_tag_id
	AND
		tags.user_id = in_user_id;
END;	
$$;
`

	// CreateDeleteUserTagProc is the query to create the stored procedure to delete user tag
	CreateDeleteUserTagProc = `
CREATE OR REPLACE PROCEDURE delete_user_tag(
	IN in_user_id BIGINT,
	IN in_user_tag_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Delete the user tag
	DELETE FROM
		tags
	WHERE
		tags.id = in_user_tag_id
	AND
		tags.user_id = in_user_id;
END;
$$;
`

	// CreateGetUserTagByTagIDProc is the query to create the stored procedure to get user tag by tag ID
	CreateGetUserTagByTagIDProc = `
CREATE OR REPLACE PROCEDURE get_user_tag_by_tag_id(
	IN in_user_id BIGINT,
	IN in_user_tag_id BIGINT,
	OUT out_user_tag_name VARCHAR
	OUT out_user_tag_created_at TIMESTAMP,
	OUT out_user_tag_updated_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user tag name, created at, and updated at by user ID and tag ID
	SELECT
		tags.name,
		tags.created_at,
		tags.updated_at
	INTO
		out_user_tag_name,
		out_user_tag_created_at,
		out_user_tag_updated_at
	FROM
		tags
	WHERE
		tags.id = in_user_tag_id
	AND
		tags.user_id = in_user_id;
END;
$$;
`

	// CreateUpdateUserNotePinProc is the query to create the stored procedure to update user note pin
	CreateUpdateUserNotePinProc = `
CREATE OR REPLACE PROCEDURE update_user_note_pin(
	IN in_user_id BIGINT,
	IN in_user_note_id BIGINT,
	IN in_user_note_pin BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the notes table
	UPDATE
		notes
	SET
		pinned_at = CASE
			WHEN in_user_note_pin THEN NOW()
			ELSE NULL
	WHERE
		notes.id = in_user_note_id
	AND
		notes.user_id = in_user_id;
END;
$$;
`

	// CreateUpdateUserNoteArchiveProc is the query to create the stored procedure to update user note archive
	CreateUpdateUserNoteArchiveProc = `
CREATE OR REPLACE PROCEDURE update_user_note_archive(
	IN in_user_id BIGINT,
	IN in_user_note_id BIGINT,
	IN in_user_note_archive BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the notes table
	UPDATE
		notes
	SET
		archived_at = CASE
			WHEN in_user_note_archive THEN NOW()
			ELSE NULL
	WHERE
		notes.id = in_user_note_id
	AND
		notes.user_id = in_user_id;
END;	
$$;
`

	// CreateUpdateUserNoteTrashProc is the query to create the stored procedure to update user note trash
	CreateUpdateUserNoteTrashProc = `
CREATE OR REPLACE PROCEDURE update_user_note_trash(
	IN in_user_id BIGINT,
	IN in_user_note_id BIGINT,
	IN in_user_note_trash BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the notes table
	UPDATE
		notes
	SET
		trashed_at = CASE
			WHEN in_user_note_trash THEN NOW()
			ELSE NULL
	WHERE
		notes.id = in_user_note_id
	AND
		notes.user_id = in_user_id;
END;
$$;
`

	// CreateUpdateUserNoteStarProc is the query to create the stored procedure to update user note star
	CreateUpdateUserNoteStarProc = `
CREATE OR REPLACE PROCEDURE update_user_note_star(
	IN in_user_id BIGINT,
	IN in_user_note_id BIGINT,
	IN in_user_note_star BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the notes table
	UPDATE
		notes
	SET
		starred_at = CASE
			WHEN in_user_note_star THEN NOW()
			ELSE NULL
	WHERE
		notes.id = in_user_note_id
	AND
		notes.user_id = in_user_id;
END;
$$;
`

	// CreateCreateUserNoteVersionProc is the query to create the stored procedure to create user note version
	CreateCreateUserNoteVersionProc = `
CREATE OR REPLACE PROCEDURE create_user_note_version(
	IN in_user_id BIGINT,
	IN in_user_note_id BIGINT,
	IN in_user_note_version_encrypted_content TEXT,
	OUT out_user_note_id_is_valid BOOLEAN,
	OUT out_user_note_version_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Check if the user note ID is valid
	SELECT
		notes.id IS NOT NULL
	INTO
		out_user_note_id_is_valid
	FROM
		notes
	WHERE
		notes.id = in_user_note_id
	AND
		notes.user_id = in_user_id;

	-- If the user note ID is valid, insert into note_versions table
	IF out_user_note_id_is_valid THEN
		-- Insert into note_versions table
		INSERT INTO note_versions (
			user_id,
			note_id,
			encrypted_content
		)
		VALUES (
			in_user_id,
			in_user_note_id,
			in_user_note_version_encrypted_content
		)
		RETURNING
			id INTO out_user_note_version_id;
	END IF;
END;
$$;
`

	// CreateDeleteUserNoteVersionProc is the query to create the stored procedure to delete user note version
	CreateDeleteUserNoteVersionProc = `
CREATE OR REPLACE PROCEDURE delete_user_note_version(
	IN in_user_id BIGINT,
	IN in_user_note_version_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Delete the user note version
	DELETE FROM
		note_versions
	INNER JOIN
		notes ON note_versions.note_id = notes.id
	WHERE
		note_versions.id = in_user_note_version_id
	AND
		note_versions.user_id = in_user_id;
END;
$$;
`

	// CreateGetUserNoteVersionByNoteVersionIDProc is the query to create the stored procedure to get user note version by note version ID
	CreateGetUserNoteVersionByNoteVersionIDProc = `
CREATE OR REPLACE PROCEDURE get_user_note_version_by_note_version_id(
	IN in_user_id BIGINT,
	IN in_user_note_version_id BIGINT,	
	OUT out_user_note_version_encrypted_content TEXT,
	OUT out_user_note_version_created_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user note version encrypted content and created at by user ID and note version ID
	SELECT
		note_versions.encrypted_content,
		note_versions.created_at
	INTO
		out_user_note_version_encrypted_content,
		out_user_note_version_created_at
	FROM
		note_versions
	INNER JOIN
		notes ON note_versions.note_id = notes.id
	WHERE
		note_versions.id = in_user_note_version_id
	AND
		notes.user_id = in_user_id;
END;
$$;
`
)
