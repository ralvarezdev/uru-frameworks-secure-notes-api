package model

const (
	// CreateSendEmailVerificationTokenProc is the query to create the send email verification token stored procedure
	CreateSendEmailVerificationTokenProc = `
CREATE OR REPLACE PROCEDURE send_email_verification_token(
	IN in_user_id BIGINT,
	IN in_email_verification_token VARCHAR,
	IN in_email_verification_token_expires_at TIMESTAMP,
)
LANGUAGE plpgsql
AS $$
DECLARE
	out_user_email_id BIGINT;
BEGIN
	-- Select the user email by user ID
	CALL get_user_email_id(in_user_id, out_user_email_id);

	-- Insert into user_email_verifications table
	INSERT INTO user_email_verifications (
		user_email_id,
		verification_token,
		expires_at
	)
	VALUES (
		out_user_email_id,
		in_email_verification_token,
		in_email_verification_token_expires_at
	);	
END;
$$;
`

	// CreateSignUpProc is the query to create the sign-up stored procedure
	CreateSignUpProc = `
CREATE OR REPLACE PROCEDURE sign_up(
	IN in_first_name VARCHAR,
	IN in_last_name VARCHAR,
	IN in_salt VARCHAR,
	IN in_encrypted_key TEXT, 
	IN in_username VARCHAR,
	IN in_email VARCHAR,
	IN in_password_hash VARCHAR,
	IN in_email_verification_token VARCHAR,
	IN in_email_verification_token_expires_at TIMESTAMP,
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
		in_first_name, 
		in_last_name, 
		in_salt,
		in_encrypted_key
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
		in_username
	);

	-- Insert into user_emails table
	INSERT INTO user_emails (
		user_id, 
		email
	)
	VALUES (
		out_user_id, 
		in_email
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
		in_password_hash
	);

	-- Insert into user_email_verifications table
	call send_email_verification_token(out_user_id, in_email_verification_token, in_email_verification_token_expires_at);
EXCEPTION 
	WHEN OTHERS THEN 
		RAISE;
END;
$$;
`

	CreateRevokeTOTPProc = `
CREATE OR REPLACE PROCEDURE revoke_totp(
	IN in_user_id BIGINT
)
LANGUAGE plpgsql
AS $$
DECLARE
	out_user_totp_id BIGINT;
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

	// CreateGenerateTokensProc is the query to create the generate tokens stored procedure
	CreateGenerateTokensProc = `
CREATE OR REPLACE PROCEDURE generate_tokens(
	IN in_user_id BIGINT,
	IN in_parent_refresh_token_id BIGINT,
	IN in_ip_address VARCHAR,
	IN in_refresh_expires_at TIMESTAMP,
	IN in_access_expires_at TIMESTAMP,
	OUT out_refresh_token_id BIGINT,
	OUT out_access_token_id BIGINT
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
		in_parent_refresh_token_id,
		in_ip_address,
		in_refresh_expires_at
	)
	RETURNING
		id INTO out_refresh_token_id;

	-- Insert into user_access_tokens table
	INSERT INTO user_access_tokens (
		user_id,
		user_refresh_token_id,
		expires_at
	)
	VALUES (
  		in_user_id,
  		out_refresh_token_id,
  		in_access_expires_at
	)
	RETURNING
  		id INTO out_access_token_id;
END;
$$;
`

	// CreateRevokeTokensByIDProc is the query to create the revoke tokens by ID stored procedure
	CreateRevokeTokensByIDProc = `
CREATE OR REPLACE PROCEDURE revoke_tokens_by_id(
	IN in_user_id BIGINT,
	IN in_refresh_token_id BIGINT
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
		user_refresh_tokens.id = in_refresh_token_id
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
		user_access_tokens.user_refresh_token_id = in_refresh_token_id
	AND
		user_access_tokens.user_id = user_id
	AND
		user_access_tokens.revoked_at IS NULL;
END;
$$;
`

	// CreateRefreshTokenProc is the query to create the refresh token stored procedure
	CreateRefreshTokenProc = `
CREATE OR REPLACE PROCEDURE refresh_token(
	IN in_user_id BIGINT,
	IN in_old_refresh_token_id BIGINT,
	IN in_ip_address VARCHAR,
	IN in_new_refresh_expires_at TIMESTAMP,
	IN in_new_access_expires_at TIMESTAMP,
	OUT out_new_refresh_token_id BIGINT,
	OUT out_new_access_token_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Revoke the old user refresh token and access token
	CALL revoke_tokens_by_id(in_user_id, in_old_refresh_token_id);

	-- Generate new tokens
	CALL generate_tokens(in_user_id, in_old_refresh_token_id, in_ip_address, in_new_refresh_expires_at, in_new_access_expires_at, out_new_refresh_token_id, out_new_access_token_id);
END;
$$;
`

	// CreateRevokeTokensProc is the query to create the revoke tokens stored procedure
	CreateRevokeTokensProc = `
CREATE OR REPLACE PROCEDURE revoke_tokens(
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

	// CreateGetAccessTokenIDByRefreshTokenIDProc is the query to create the get access token ID by refresh token ID stored procedure
	CreateGetAccessTokenIDByRefreshTokenIDProc = `
CREATE OR REPLACE PROCEDURE get_access_token_id_by_refresh_token_id(
	IN in_refresh_token_id BIGINT,
	OUT out_access_token_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user_access_tokens table
	SELECT
		user_access_tokens.id
	INTO
		out_access_token_id
	FROM
		user_access_tokens
	WHERE
		user_access_tokens.user_refresh_token_id = in_refresh_token_id
	AND
		user_access_tokens.revoked_at IS NULL;
END;
$$;
`

	// CreatePreLogInProc is the query to create the pre-log in stored procedure
	CreatePreLogInProc = `
CREATE OR REPLACE PROCEDURE pre_log_in(
	IN in_username VARCHAR,
	OUT out_user_id BIGINT,
	OUT out_password_hash VARCHAR,
	OUT out_salt VARCHAR,
	OUT out_encrypted_key TEXT,
	OUT out_totp_id BIGINT,
	OUT out_totp_secret VARCHAR
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
		out_salt,	
		out_encrypted_key,	
		out_password_hash
	FROM
		users
	INNER JOIN
		user_usernames ON users.id = user_usernames.user_id
	INNER JOIN
		user_password_hashes ON users.id = user_password_hashes.user_id
	WHERE
		user_usernames.username = in_username
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
		out_totp_id,
		out_totp_secret
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
		out_totp_id = NULL;
		out_totp_secret = NULL;
	END IF;
END;
$$;
`

	// CreateRegisterFailedLogInAttemptProc is the query to create the register failed log in attempt stored procedure
	CreateRegisterFailedLogInAttemptProc = `
CREATE OR REPLACE PROCEDURE register_failed_login_attempt(
	IN in_user_id BIGINT,
	IN in_ip_address VARCHAR,
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
		in_ip_address,
		in_bad_password,
		in_bad_2fa_code
	);
END;
$$;
`

	// CreateGetUserTOTPProc is the query to create the get user TOTP by user ID stored procedure
	CreateGetUserTOTPProc = `
CREATE OR REPLACE PROCEDURE get_user_totp(
	IN in_user_id BIGINT,
	OUT out_totp_id BIGINT,
	OUT out_totp_secret VARCHAR,
	OUT out_totp_verified_at TIMESTAMP
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
		out_totp_id,
		out_totp_secret,
		out_totp_verified_at
	FROM
		user_totps
	WHERE
		user_totps.user_id = in_user_id
	AND
		user_totps.revoked_at IS NULL;
END;
$$;
`

	// CreateGetUserEmailProc is the query to create the get user email by user ID stored procedure
	CreateGetUserEmailProc = `
CREATE OR REPLACE PROCEDURE get_user_email(
	IN in_user_id BIGINT,
	OUT out_email VARCHAR
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user email by user ID
	SELECT
		user_emails.email
	INTO
		out_email
	FROM
		user_emails
	WHERE
		user_emails.user_id = in_user_id
	AND
		user_emails.revoked_at IS NULL;
END;
$$;
`

	// CreateGetUserEmailIDProc is the query to create the get user email ID by user ID stored procedure
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

	// CreateGenerateTOTPUrlProc is the query to create the generate TOTP URL stored procedure
	CreateGenerateTOTPUrlProc = `
CREATE OR REPLACE PROCEDURE pre_generate_totp_url(
	IN in_user_id BIGINT,
	IN in_new_totp_secret VARCHAR,
	OUT out_email VARCHAR,
	OUT out_old_totp_id BIGINT,
	OUT out_old_totp_secret VARCHAR,
	OUT out_old_totp_verified_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user email by user ID
	CALL get_user_email(in_user_id, out_email);

	-- Select the TOTP ID, secret, and verified at by user ID
	CALL get_user_totp(in_user_id, out_old_totp_id, out_old_totp_secret, out_old_totp_verified_at);

	-- If the TOTP is not verified, revoke it
	IF out_old_totp_verified_at IS NULL THEN
		CALL revoke_totp(in_user_id);
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
			in_new_totp_secret
		);
	END IF;
END;
$$;
`

	// CreateIsRefreshTokenValidProc is the query to create the is refresh token valid stored procedure
	CreateIsRefreshTokenValidProc = `
CREATE OR REPLACE PROCEDURE is_refresh_token_valid(
	IN in_refresh_token_id BIGINT,
	OUT out_expires_at TIMESTAMP,
	OUT out_found BOOLEAN,
	OUT out_is_expired BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user refresh token ID by user ID
	SELECT
		user_refresh_tokens.expires_at
	INTO
		out_expires_at
	FROM
		user_refresh_tokens
	WHERE
		user_refresh_tokens.id = in_refresh_token_id
	AND
		user_refresh_tokens.revoked_at IS NULL;

	IF out_expires_at IS NOT NULL THEN
		out_found = TRUE;
		out_is_expired = out_expires_at < NOW();
	END IF;
END;
$$;
`

	// CreateIsAccessTokenValidProc is the query to create the is access token valid stored procedure
	CreateIsAccessTokenValidProc = `
CREATE OR REPLACE PROCEDURE is_access_token_valid(
	IN in_access_token_id BIGINT,
	OUT out_expires_at TIMESTAMP,
	OUT out_found BOOLEAN,
	OUT out_is_expired BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user access token ID by user ID
	SELECT
		user_access_tokens.expires_at
	INTO
		out_expires_at
	FROM
		user_access_tokens
	WHERE
		user_access_tokens.id = in_access_token_id
	AND
		user_access_tokens.revoked_at IS NULL;
	
	IF out_expires_at IS NOT NULL THEN
		out_found = TRUE;
		out_is_expired = out_expires_at < NOW();
	END IF;
END;	
$$;
`

	// CreateRevokeTOTPRecoveryCodeProc is the query to create the revoke TOTP recovery code stored procedure
	CreateRevokeTOTPRecoveryCodeProc = `
CREATE OR REPLACE PROCEDURE revoke_totp_recovery_code(
	IN out_user_totp_id BIGINT,
	IN in_recovery_code VARCHAR
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
		user_totp_recovery_codes.user_totp_id = out_user_totp_id
	AND
		user_totp_recovery_codes.recovery_code = in_recovery_code
	AND
		user_totp_recovery_codes.revoked_at IS NULL;
END;
$$;
`

	// CreateVerifyTOTPProc is the query to create verify TOTP stored procedure
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

	// CreateVerifyEmailProc is the query to create the verify email stored procedure
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
		user_email_verifications.revoked_at IS NULL;
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

	// CreateIsUserEmailVerifiedProc is the query to create the is user email verified stored procedure
	CreateIsUserEmailVerifiedProc = `
CREATE OR REPLACE PROCEDURE is_user_email_verified(
	IN in_user_id BIGINT,
	OUT out_user_first_name VARCHAR,
	OUT out_user_last_name VARCHAR,
	OUT out_user_email VARCHAR
	OUT out_is_verified BOOLEAN
)
LANGUAGE plpgsql
AS $$
DECLARE 
	out_user_email_id BIGINT;
BEGIN
	-- Select the user email by user ID
	SELECT
		users.first_name,
		users.last_name,
		user_emails.id,
		user_emails.email,
		user_emails.verified_at IS NOT NULL
	INTO
		out_user_first_name,
		out_user_last_name,
		out_user_email_id,
		out_user_email,
		out_is_verified
	FROM
		user_emails
	INNER JOIN
		users ON user_emails.user_id = users.id
	WHERE
		user_emails.user_id = in_user_id
	AND
		user_emails.revoked_at IS NULL;

	-- Revoke the user email verification token, if it exists and hasn't been verified
	IF NOT out_is_verified THEN
		-- Update the user_email_verifications table
		UPDATE
			user_email_verifications
		SET
			revoked_at = NOW()
		WHERE
			user_email_verifications.user_email_id = user_email_id
	END IF;
END;
$$;
`

	// CreateRevokeUserEmailProc is the query to create the revoke user email stored procedure
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

	// CreateChangeEmailProc is the query to create the change email stored procedure
	CreateChangeEmailProc = `
CREATE OR REPLACE PROCEDURE change_email(
	IN in_user_id BIGINT,
	IN in_new_email VARCHAR,
	IN in_email_verification_token VARCHAR,
	IN in_email_verification_token_expires_at TIMESTAMP
	OUT out_user_first_name VARCHAR,
	OUT out_user_last_name VARCHAR,
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
		in_new_email
	);

	-- Insert into user_email_verifications table
	CALL send_email_verification_token(in_user_id, in_email_verification_token, in_email_verification_token_expires_at);

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

	// CreateRevokeResetPasswordTokenProc is the query to create the revoke reset password token stored procedure
	CreateRevokeResetPasswordTokenProc = `
CREATE OR REPLACE PROCEDURE revoke_reset_password_token(
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
		user_reset_passwords.user_id = out_user_id
	AND
		user_reset_passwords.revoked_at IS NULL;
END;
$$;
`

	// CreateForgotPasswordProc is the query to create the forgot password stored procedure
	CreateForgotPasswordProc = `
CREATE OR REPLACE PROCEDURE forgot_password(
	IN in_username VARCHAR,
	IN in_reset_token VARCHAR,
	IN in_reset_token_expires_at TIMESTAMP,
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
		user_usernames.username = in_username
	AND
		user_usernames.revoked_at IS NULL
	AND
		user_emails.revoked_at IS NULL;
	AND
		users.deleted_at IS NULL;

	-- Revoke the user reset password token, if it exists and hasn't been revoked
	CALL revoke_reset_password_token(out_user_id);

	-- Insert into user_reset_passwords table
	INSERT INTO user_reset_passwords (
		user_id,
		reset_token,
		expires_at
	)
	VALUES (
		out_user_id,
		in_reset_token,
		in_reset_token_expires_at
	);
END;	
$$;
`

	// CreateRevokePasswordHashProc is the query to create the revoke password hash stored procedure
	CreateRevokePasswordHashProc = `
CREATE OR REPLACE PROCEDURE revoke_password_hash(
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

	// CreateResetPasswordProc is the query to create the reset password stored procedure
	CreateResetPasswordProc = `
CREATE OR REPLACE PROCEDURE reset_password(
	IN in_reset_password_token BIGINT,
	IN in_new_password_hash VARCHAR,
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
		user_reset_passwords.reset_token = in_reset_password_token
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
		call revoke_password_hash(out_user_id);
	
		-- Insert into user_password_hashes table
		INSERT INTO user_password_hashes (
			user_id,
			password_hash
		)
		VALUES (
			in_user_id,
			in_new_password_hash
		);
	
		-- Revoke the user reset password token
		CALL revoke_reset_password_token(out_user_id);
	END IF;
END;
$$;
`
)
