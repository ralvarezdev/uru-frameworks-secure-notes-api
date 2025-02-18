package model

const (
	// CreateGetUserEmailIDProc is the query to create the stored procedure to get user email ID by user ID
	CreateGetUserEmailIDProc = `
CREATE OR REPLACE PROCEDURE get_user_email_id(
	IN in_user_id BIGINT,
	OUT out_user_email_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN	
	-- Select the user email ID by user ID
	SELECT
		user_emails.id
	INTO
		out_user_email_id
	FROM
		user_emails
	WHERE
		user_emails.user_id = in_user_id
	AND
		user_emails.revoked_at IS NULL;
END;	
$$;
`

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

	// CreateRevokeUser2FATOTPProc is the query to create the stored procedure to revoke user 2FA TOTP
	CreateRevokeUser2FATOTPProc = `
CREATE OR REPLACE PROCEDURE revoke_user_2fa_totp(
	IN in_user_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_2fa_totp table
	UPDATE	
		user_2fa_totp
	SET	
		revoked_at = NOW()
	WHERE
		user_2fa_totp.user_id = in_user_id 
	AND	
		user_2fa_totp.revoked_at IS NULL;
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

	// CreateGetUserAccessTokenIDByUserRefreshTokenIDProc is the query to create the stored procedure to get user access token ID by user refresh token ID
	CreateGetUserAccessTokenIDByUserRefreshTokenIDProc = `
CREATE OR REPLACE PROCEDURE get_user_access_token_id_by_user_refresh_token_id(
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

	// CreateGetUser2FATOTPProc is the query to create the stored procedure to get user 2FA TOTP by user ID
	CreateGetUser2FATOTPProc = `
CREATE OR REPLACE PROCEDURE get_user_2fa_totp(
	IN in_user_id BIGINT,
	OUT out_user_2fa_totp_id BIGINT,
	OUT out_user_2fa_totp_secret VARCHAR,
	OUT out_user_2fa_totp_verified_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the TOTP ID, secret, and verified at by user ID
	SELECT
		user_2fa_totp.id,
		user_2fa_totp.secret,
		user_2fa_totp.verified_at
	INTO
		out_user_2fa_totp_id,
		out_user_2fa_totp_secret,
		out_user_2fa_totp_verified_at
	FROM
		user_2fa_totp
	WHERE
		user_2fa_totp.user_id = in_user_id
	AND
		user_2fa_totp.revoked_at IS NULL;
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

	// CreateGenerate2FATOTPUrlProc is the query to create the stored procedure to generate 2FA TOTP URL
	CreateGenerate2FATOTPUrlProc = `
CREATE OR REPLACE PROCEDURE generate_2fa_totp_url(
	IN in_user_id BIGINT,
	IN in_new_user_2fa_totp_secret VARCHAR,
	OUT out_has_user_2fa_enabled BOOLEAN,
	OUT out_user_email VARCHAR,
	OUT out_old_user_2fa_totp_id BIGINT,
	OUT out_old_user_2fa_totp_secret VARCHAR,
	OUT out_old_user_2fa_totp_verified_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Check if the user has 2FA enabled
	call has_user_2fa_enabled(in_user_id, out_has_user_2fa_enabled);

	IF out_has_user_2fa_enabled THEN
		-- Select the user email by user ID
		CALL get_user_email(in_user_id, out_user_email);
	
		-- Select the TOTP ID, secret, and verified at by user ID
		CALL get_user_2fa_totp(in_user_id, out_old_user_2fa_totp_id, out_old_user_2fa_totp_secret, out_old_user_2fa_totp_verified_at);
	
		-- If the TOTP is not verified, revoke it
		IF out_old_user_2fa_totp_id IS NOT NULL AND out_old_user_2fa_totp_verified_at IS NULL THEN
			CALL revoke_user_2fa_totp(in_user_id);
		END IF;
	
		-- If the TOTP wasn't active or verified, insert a new TOTP
		IF out_old_user_2fa_totp_verified_at IS NULL THEN
			-- Insert into user_2fa_totp table
			INSERT INTO user_2fa_totp (
				user_id,
				secret
			)
			VALUES (
				in_user_id,
				in_new_user_2fa_totp_secret
			);
		END IF;
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

	IF out_user_refresh_token_expires_at IS NOT NULL THEN
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
	
	IF out_user_access_token_expires_at IS NOT NULL THEN
		out_user_access_token_found = TRUE;
		out_user_access_token_is_expired = out_user_access_token_expires_at < NOW();
	END IF;
END;	
$$;
`

	// CreateUseUser2FARecoveryCodeProc is the query to create the stored procedure to use user 2FA TOTP recovery code
	CreateUseUser2FARecoveryCodeProc = `
CREATE OR REPLACE PROCEDURE use_user_2fa_recovery_code(
	IN in_user_id BIGINT,
	IN in_user_2fa_recovery_code VARCHAR,
	OUT out_user_2fa_recovery_code_is_valid BOOLEAN,
	OUT out_user_2fa_recovery_code_left INT
)
LANGUAGE plpgsql
AS $$
DECLARE
	out_user_2fa_recovery_code_id BIGINT;
BEGIN
	-- Select if the user 2FA recovery code is valid
	SELECT
		user_2fa_recovery_codes.id
	INTO
		out_user_2fa_recovery_code_id
	FROM
		user_2fa_recovery_codes
	WHERE
		user_2fa_recovery_codes.user_id = in_user_id
	AND
		user_2fa_recovery_codes.recovery_code = in_user_2fa_recovery_code
	AND
		user_2fa_recovery_codes.revoked_at IS NULL
	AND
		user_2fa_recovery_codes.used_at IS NULL;

	IF out_user_2fa_recovery_code_id IS NULL THEN
		out_user_2fa_recovery_code_is_valid = FALSE;
	ELSE
		out_user_2fa_recovery_code_is_valid = TRUE;

		-- Update the user_2fa_recovery_codes table
		UPDATE
			user_2fa_recovery_codes
		SET
			used_at = NOW()
		WHERE
			user_2fa_recovery_codes.id = out_user_2fa_recovery_code_id;
	
		-- Select the count of the user_2fa_recovery_codes table
		SELECT
			COUNT(*)
		INTO
			out_user_2fa_recovery_code_left
		FROM
			user_2fa_recovery_codes
		WHERE
			user_2fa_recovery_codes.user_id = in_user_id
		AND
			user_2fa_recovery_codes.revoked_at IS NULL
		AND
			user_2fa_recovery_codes.used_at IS NULL;
	END IF;
END;
$$;
`

	// CreateRevokeUser2FARecoveryCodesProc is the query to create the stored procedure to revoke user 2FA recovery codes
	CreateRevokeUser2FARecoveryCodesProc = `
CREATE OR REPLACE PROCEDURE revoke_user_2fa_recovery_codes(
	IN in_user_id BIGINT
)	
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_2fa_recovery_codes table
	UPDATE
		user_2fa_recovery_codes
	SET
		revoked_at = NOW()
	WHERE	
		user_2fa_recovery_codes.user_id = in_user_id
	AND	
		user_2fa_recovery_codes.revoked_at IS NULL
	AND
		user_2fa_recovery_codes.used_at IS NULL;
END;
$$;
`

	// CreateVerify2FATOTPProc is the query to create verify 2FA TOTP
	CreateVerify2FATOTPProc = `
CREATE OR REPLACE PROCEDURE verify_2fa_totp(
	IN in_user_2fa_totp_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_2fa_totp table
	UPDATE
		user_2fa_totp
	SET
		verified_at = NOW()
	WHERE
		user_2fa_totp.id = in_user_2fa_totp_id
	AND 
		user_2fa_totp.verified_at IS NULL
	AND
		user_2fa_totp.revoked_at IS NULL;
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
		out_user_email_verification_id,
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
		user_emails.revoked_at IS NULL
	AND
		users.deleted_at IS NULL;
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
		users.id = in_user_id
	AND
		users.deleted_at IS NULL;
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
	CreateForgotPasswordProc = `
CREATE OR REPLACE PROCEDURE forgot_password(
	IN in_user_email VARCHAR,
	IN in_user_reset_password_token VARCHAR,
	IN in_user_reset_password_token_expires_at TIMESTAMP,
	OUT out_user_id BIGINT,
	OUT out_user_first_name VARCHAR,
	OUT out_user_last_name VARCHAR
)
LANGUAGE plpgsql	
AS $$
BEGIN
	-- Select the user ID, first name, last name, and email by username
	SELECT
		users.id,
		users.first_name,
		users.last_name
	INTO
		out_user_id,
		out_user_first_name,
		out_user_last_name
	FROM
		users
	INNER JOIN
		user_emails ON users.id = user_emails.user_id
	WHERE
		user_emails.email = in_user_email
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
	IN in_user_reset_password_token VARCHAR,
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

	-- Revoke the user reset password token
	CALL revoke_user_reset_password_token(in_user_id);

	-- Disable the user 2FA 
	CALL disable_user_2fa(in_user_id);

	-- Update the users table
	UPDATE
		users
	SET
		deleted_at = NOW()
	WHERE
		users.id = in_user_id
	AND
		users.deleted_at IS NULL;

	-- Update the user_tags table
	UPDATE
		user_tags	
	SET
		deleted_at = NOW()
	WHERE
		user_tags.user_id = in_user_id
	AND
		user_tags.deleted_at IS NULL;

	-- Update the user_note_tags table
	UPDATE
		user_note_tags
	SET
		deleted_at = NOW()
	FROM
		user_notes
	WHERE
		user_note_tags.note_id = user_notes.id
	AND
		user_notes.user_id = in_user_id
	AND
		user_note_tags.deleted_at IS NULL;
	
	-- Update the user_note_versions table
	UPDATE
		user_note_versions
	SET
		deleted_at = NOW()
	FROM
		user_notes
	WHERE
		user_note_versions.note_id = user_notes.id
	AND
		user_notes.user_id = in_user_id
	AND
		user_note_versions.deleted_at IS NULL;

	-- Update the user_notes table
	UPDATE
		user_notes
	SET
		deleted_at = NOW()
	WHERE
		user_notes.user_id = in_user_id
	AND
		user_notes.deleted_at IS NULL;
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
		users.id = in_user_id
	AND
		users.deleted_at IS NULL;
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
		first_name = COALESCE(in_user_first_name, out_user_first_name),
		last_name = COALESCE(in_user_last_name, out_user_last_name),
		birthdate = COALESCE(in_user_birthdate, out_user_birthdate),
		update_at = NOW()
	WHERE
		users.id = in_user_id
	AND
		users.deleted_at IS NULL;
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

	// CreateHasUser2FAEnabledProc is the query to create the stored procedure to check if the user has 2FA enabled
	CreateHasUser2FAEnabledProc = `
CREATE OR REPLACE PROCEDURE has_user_2fa_enabled(
	IN in_user_id BIGINT,
	OUT out_has_user_2fa_enabled BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Check if the user is on the user_2fa table
	SELECT
		user_2fa.id IS NOT NULL
	INTO
		out_has_user_2fa_enabled
	FROM
		user_2fa
	WHERE
		user_2fa.user_id = in_user_id
	AND
		user_2fa.revoked_at IS NULL;
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
	OUT out_user_has_2fa_enabled BOOLEAN
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

	-- Check if the user has 2FA enabled
	CALL has_user_2fa_enabled(in_user_id, out_user_has_2fa_enabled);
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
	-- Insert into user_tags table
	INSERT INTO user_tags (
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
	-- Update the user_tags table
	UPDATE
		user_tags
	SET
		name = in_user_tag_name,
		updated_at = NOW()
	WHERE
		user_tags.id = in_user_tag_id
	AND
		user_tags.user_id = in_user_id
	AND
		user_tags.deleted_at IS NULL;
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
	-- Update the user_tags table
	UPDATE
		user_tags
	SET
		deleted_at = NOW()
	WHERE
		user_tags.id = in_user_tag_id
	AND
		user_tags.user_id = in_user_id
	AND
		user_tags.deleted_at IS NULL;
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
	-- Update the user_notes table
	UPDATE
		user_notes
	SET
		pinned_at = CASE
			WHEN in_user_note_pin THEN NOW()
			ELSE NULL
		END,
		updated_at = NOW()
	WHERE
		user_notes.id = in_user_note_id
	AND
		user_notes.user_id = in_user_id
	AND
		user_notes.deleted_at IS NULL;
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
	-- Update the user_notes table
	UPDATE
		user_notes
	SET
		archived_at = CASE
			WHEN in_user_note_archive THEN NOW()
			ELSE NULL
		END,
		updated_at = NOW()
	WHERE
		user_notes.id = in_user_note_id
	AND
		user_notes.user_id = in_user_id
	AND
		user_notes.deleted_at IS NULL;
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
	-- Update the user_notes table
	UPDATE
		user_notes
	SET
		trashed_at = CASE
			WHEN in_user_note_trash THEN NOW()
			ELSE NULL
		END,
		updated_at = NOW()
	WHERE
		user_notes.id = in_user_note_id
	AND
		user_notes.user_id = in_user_id
	AND
		user_notes.deleted_at IS NULL;
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
	-- Update the user_notes table
	UPDATE
		user_notes
	SET
		starred_at = CASE
			WHEN in_user_note_star THEN NOW()
			ELSE NULL
		END,
		updated_at = NOW()
	WHERE
		user_notes.id = in_user_note_id
	AND
		user_notes.user_id = in_user_id
	AND
		user_notes.deleted_at IS NULL;
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
		user_notes.id IS NOT NULL
	INTO
		out_user_note_id_is_valid
	FROM
		user_notes
	WHERE
		user_notes.id = in_user_note_id
	AND
		user_notes.user_id = in_user_id
	AND
		user_notes.deleted_at IS NULL;

	-- If the user note ID is valid, insert into user_note_versions table
	IF out_user_note_id_is_valid THEN
		-- Insert into user_note_versions table
		INSERT INTO user_note_versions (
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
	-- Update the user_note_versions table
	UPDATE
		user_note_versions
	SET
		encrypted_content = NULL,
		deleted_at = NOW()
	FROM
		user_notes 
	WHERE 
		user_note_versions.note_id = user_notes.id
	AND
		user_note_versions.id = in_user_note_version_id
	AND
		user_notes.user_id = in_user_id
	AND
		user_notes.deleted_at IS NULL
	AND
		user_note_versions.deleted_at IS NULL;		
END;
$$;
`

	// CreateValidateUserTagsIDProc is the query to create the stored procedure to validate user tags ID
	CreateValidateUserTagsIDProc = `
CREATE OR REPLACE PROCEDURE validate_user_tags_id(
	IN in_user_id BIGINT,
	IN in_user_tags_id BIGINT[],
	OUT out_valid_user_tags_id BIGINT[]
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Check if the user tags ID are valid
	SELECT
		ARRAY(
			SELECT
				user_tags.id
			FROM
				user_tags
			WHERE
				user_tags.user_id = in_user_id
			AND
				user_tags.id = ANY(in_user_tags_id)
			AND
				user_tags.deleted_at IS NULL
		)
	INTO
		out_valid_user_tags_id;
END;
$$;
`

	// CreateAddUserNoteTagsProc is the query to create the stored procedure to add user note tags
	CreateAddUserNoteTagsProc = `
CREATE OR REPLACE PROCEDURE add_user_note_tags(
	IN in_user_id BIGINT,
	IN in_user_note_id BIGINT,
	IN in_user_note_tags_id BIGINT[]
)
LANGUAGE plpgsql
AS $$
DECLARE
	out_valid_user_note_tag_id BIGINT;
BEGIN
	-- Check if the user tags ID are valid
	CALL validate_user_tags_id(in_user_id, in_user_note_tags_id, out_valid_user_note_tags_id);

	-- Insert into note_tags table
	FOREACH out_valid_user_note_tag_id IN ARRAY out_valid_user_note_tags_id
	LOOP
		INSERT INTO user_note_tags (
			note_id,
			tag_id
		)
		VALUES (
			out_user_note_id,
			out_valid_user_note_tag_id
		);
	END LOOP;
END;
$$;
`

	// CreateCreateUserNoteProc is the query to create the stored procedure to create user note
	CreateCreateUserNoteProc = `
CREATE OR REPLACE PROCEDURE create_user_note(
	IN in_user_id BIGINT,
	IN in_user_note_title VARCHAR,
	IN in_user_note_color VARCHAR,
	IN in_user_note_pinned BOOLEAN,
	IN in_user_note_archived BOOLEAN,
	IN in_user_note_trashed BOOLEAN,
	IN in_user_note_starred BOOLEAN,
	IN in_user_note_encrypted_content TEXT,
	IN in_user_note_tags_id BIGINT[],
	OUT out_user_note_id BIGINT
)
LANGUAGE plpgsql
AS $$
DECLARE
	out_valid_user_note_tags_id BIGINT[];
BEGIN
	-- Insert into user_notes table
	INSERT INTO user_notes (
		user_id,
		title,
		color,
		pinned_at,
		archived_at,
		trashed_at,
		starred_at
	)
	VALUES (
		in_user_id,
		in_user_note_title,
		in_user_note_color,
		CASE
			WHEN in_user_note_pinned THEN NOW()
			ELSE NULL
		END,
		CASE
			WHEN in_user_note_archived THEN NOW()
			ELSE NULL
		END,
		CASE
			WHEN in_user_note_trashed THEN NOW()
			ELSE NULL
		END,
		CASE	
			WHEN in_user_note_starred THEN NOW()	
			ELSE NULL
		END
	)
	RETURNING
		id INTO out_user_note_id;

	-- Insert into user_note_versions table
	INSERT INTO user_note_versions (
		user_id,
		note_id,
		encrypted_content
	)
	VALUES (
		in_user_id,
		out_user_note_id,
		in_user_note_encrypted_content
	);

	-- Add user note tags
	CALL add_user_note_tags(in_user_id, out_user_note_id, in_user_note_tags_id);
END;
$$;
`

	// CreateCreateUser2FARecoveryCodesProc is the query to create the stored procedure to create user TOTP recovery codes
	CreateCreateUser2FARecoveryCodesProc = `
CREATE OR REPLACE PROCEDURE create_user_2fa_recovery_codes(
	IN in_user_id BIGINT,
	IN in_user_2fa_recovery_codes VARCHAR[],
	OUT out_has_user_2fa_enabled BOOLEAN
)
LANGUAGE plpgsql
AS $$
DECLARE
	in_user_2fa_recovery_code VARCHAR;
BEGIN
	-- Check if the user has 2FA enabled
	CALL has_user_2fa_enabled(in_user_id, out_has_user_2fa_enabled);

	IF out_has_user_2fa_enabled THEN
		-- Revoke the user 2FA recovery codes
		CALL revoke_user_2fa_recovery_codes(in_user_id);
	
		-- Insert into user_2fa_recovery_codes table
		FOREACH in_user_2fa_recovery_code IN ARRAY in_user_2fa_recovery_codes
		LOOP
			INSERT INTO user_2fa_recovery_codes (
				user_id,
				recovery_code
			)
			VALUES (
				in_user_id,
				in_user_2fa_recovery_code
			);
		END LOOP;
	END IF;
END;
$$;
`

	// CreateDeleteUserNoteProc is the query to create the stored procedure to delete user note
	CreateDeleteUserNoteProc = `
CREATE OR REPLACE PROCEDURE delete_user_note(
	IN in_user_id BIGINT,
	IN in_user_note_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Update the user_notes table
	UPDATE
		user_notes
	SET
		title = '',
		color = NULL,
		pinned_at = NULL,
		archived_at = NULL,
		trashed_at = NULL,
		starred_at = NULL,
		deleted_at = NOW()
	WHERE
		user_notes.id = in_user_note_id
	AND
		user_notes.user_id = in_user_id
	AND
		user_notes.deleted_at IS NULL;

	-- Update the user_note_versions table
	UPDATE
		user_note_versions
	SET
		encrypted_content = NULL,
		deleted_at = NOW()
	FROM
		user_notes 
	WHERE
		user_note_versions.note_id = user_notes.id
	AND
		user_notes.id = in_user_note_id
	AND
		user_notes.user_id = in_user_id
	AND
		user_note_versions.deleted_at IS NULL;

	-- Update the user_note_tags table
	UPDATE
		user_note_tags
	SET
		assigned_at = NULL,
		deleted_at = NOW()
	FROM
		user_notes 
	WHERE
		user_note_tags.note_id = user_notes.id
	AND
		user_notes.id = in_user_note_id
	AND
		user_notes.user_id = in_user_id
	AND
		user_note_tags.deleted_at IS NULL;
END;
$$;
`

	// CreateRemoveUserNoteTagsProc is the query to create the stored procedure to remove user note tags
	CreateRemoveUserNoteTagsProc = `
CREATE OR REPLACE PROCEDURE remove_user_note_tags(
	IN in_user_id BIGINT,
	IN in_user_note_id BIGINT,
	IN in_user_note_tags_id BIGINT[]
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Check if the user tags ID are valid
	CALL validate_user_tags_id(in_user_id, in_user_note_tags_id, out_valid_user_note_tags_id);

	-- Update the user_note_tags table
	UPDATE
		user_note_tags
	SET
		deleted_at = NOW()
	WHERE
		user_note_tags.note_id = in_user_note_id
	AND
		user_note_tags.tag_id = ANY(out_valid_user_note_tags_id)
	AND
		user_note_tags.deleted_at IS NULL;
END;
$$;
`

	// CreateUpdateUserNoteProc is the query to create the stored procedure to update user note
	CreateUpdateUserNoteProc = `
CREATE OR REPLACE PROCEDURE update_user_note(
	IN in_user_id BIGINT,
	IN in_user_note_id BIGINT,
	IN in_user_note_title VARCHAR,
	IN in_user_note_color VARCHAR
)
LANGUAGE plpgsql
AS $$
DECLARE
	out_user_note_title VARCHAR;
	out_user_note_color VARCHAR;
BEGIN
	-- Select the user note title and color by user ID and user note ID
	SELECT
		user_notes.title,
		user_notes.color
	INTO
		out_user_note_title,
		out_user_note_color
	FROM
		user_notes
	WHERE
		user_notes.id = in_user_note_id
	AND
		user_notes.user_id = in_user_id
	AND
		user_notes.deleted_at IS NULL;

	-- Update the user_notes table
	UPDATE
		user_notes
	SET
		title = COALESCE(in_user_note_title, out_user_note_title),
		color = COALESCE(in_user_note_color, out_user_note_color),
		updated_at = NOW()
	WHERE
		user_notes.id = in_user_note_id
	AND
		user_notes.user_id = in_user_id;
END;
$$;
`

	// CreateListUserNotesProc is the query to create the stored procedure to list user notes
	CreateListUserNotesProc = `
CREATE OR REPLACE PROCEDURE list_user_notes(
	IN in_user_id BIGINT,
	OUT out_user_notes_id BIGINT[]
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user notes ID by user ID
	SELECT
		ARRAY(
			SELECT
				user_notes.id
			FROM
				user_notes	
			WHERE
				user_notes.user_id = in_user_id
			AND
				user_notes.deleted_at IS NULL
		)
	INTO
		out_user_notes_id;
END;
$$;
`

	// CreateRevokeUser2FAEmailCodeProc is the query to create the stored procedure to revoke user 2FA email code
	CreateRevokeUser2FAEmailCodeProc = `
CREATE OR REPLACE PROCEDURE revoke_user_2fa_email_code(
	IN in_user_id BIGINT
)	
LANGUAGE plpgsql	
AS $$
BEGIN
	-- Update the user_2fa_email_codes table
	UPDATE
		user_2fa_email_codes
	SET
		revoked_at = NOW()
	WHERE
		user_2fa_email_codes.user_id = in_user_id
	AND
		user_2fa_email_codes.revoked_at IS NULL
	AND
		user_2fa_email_codes.used_at IS NULL;
END;	
$$;
`

	// CreateCreateUser2FAEmailCodeProc is the query to create the stored procedure to create user 2FA email code
	CreateCreateUser2FAEmailCodeProc = `
CREATE OR REPLACE PROCEDURE create_user_2fa_email_code(	
	IN in_user_id BIGINT,
	IN in_user_2fa_email_code VARCHAR,
	IN in_user_2fa_email_code_expires_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN	
	-- Revoke the user 2FA email code
	CALL revoke_user_2fa_email_code(in_user_id);
	
	-- Insert into user_2fa_email_codes table
	INSERT INTO user_2fa_email_codes (
		user_id,
		code,
		expires_at
	)
	VALUES (
		in_user_id,
		in_user_2fa_email_code,
		in_user_2fa_email_code_expires_at
	);
END;
$$;
`

	// CreateUseUser2FAEmailCodeProc is the query to create the stored procedure to use user 2FA email code
	CreateUseUser2FAEmailCodeProc = `
CREATE OR REPLACE PROCEDURE use_user_2fa_email_code(
	IN in_user_id BIGINT,
	IN in_user_2fa_email_code VARCHAR,
	OUT out_user_2fa_email_code_is_valid BOOLEAN
)
LANGUAGE plpgsql
AS $$
DECLARE
	out_user_2fa_email_code_id BIGINT;
BEGIN
	-- Check if the user 2FA email code is valid
	SELECT
		user_2fa_email_codes.id
	INTO
		out_user_2fa_email_code_id
	FROM
		user_2fa_email_codes
	WHERE
		user_2fa_email_codes.user_id = in_user_id
	AND
		user_2fa_email_codes.code = in_user_2fa_email_code
	AND
		user_2fa_email_codes.expires_at > NOW()
	AND
		user_2fa_email_codes.revoked_at IS NULL
	AND
		user_2fa_email_codes.used_at IS NULL;

	IF out_user_2fa_email_code_id IS NULL THEN
		out_user_2fa_email_code_is_valid = FALSE;
	ELSE
		out_user_2fa_email_code_is_valid = TRUE;

		-- Update the user_2fa_email_codes table
		UPDATE
			user_2fa_email_codes
		SET
			used_at = NOW()
		WHERE
			user_2fa_email_codes.id = out_user_2fa_email_code_id;
	END IF;
END;
$$;
`

	// CreateEnableUser2FAProc is the query to create the stored procedure to enable user 2FA
	CreateEnableUser2FAProc = `
CREATE OR REPLACE PROCEDURE enable_2fa(
	IN in_user_id BIGINT,
	IN in_user_2fa_recovery_codes VARCHAR[],
	OUT out_is_user_email_verified BOOLEAN,
	OUT out_has_user_2fa_enabled BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Check if the user email is verified
	CALL is_user_email_verified(in_user_id, out_is_user_email_verified);

	IF out_is_user_email_verified THEN
		-- Check if the user has 2FA enabled
		CALL has_user_2fa_enabled(in_user_id, out_has_user_2fa_enabled);

		IF NOT out_has_user_2fa_enabled THEN
			-- Create the user 2FA recovery codes
			CALL create_user_2fa_recovery_codes(in_user_id, in_user_2fa_recovery_codes);
		
			-- Insert into user_2fa table
			INSERT INTO user_2fa (
				user_id
			)
			VALUES (
				in_user_id
			);
		END IF;
	END IF;
END;	
$$;
`

	// CreateDisableUser2FAProc is the query to create the stored procedure to disable user 2FA
	CreateDisableUser2FAProc = `
CREATE OR REPLACE PROCEDURE disable_2fa(
	IN in_user_id BIGINT,
	OUT out_has_user_2fa_enabled BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Check if the user has 2FA enabled
	CALL has_user_2fa_enabled(in_user_id, out_has_user_2fa_enabled);

	IF out_has_user_2fa_enabled THEN
		-- Update the user_2fa table
		UPDATE
			user_2fa
		SET
			revoked_at = NOW()
		WHERE
			user_2fa.user_id = in_user_id
		AND
			user_2fa.revoked_at IS NULL;
	
		-- Revoke the user 2FA recovery codes
		CALL revoke_user_2fa_recovery_codes(in_user_id);
	
		-- Revoke the user 2FA email code
		CALL revoke_user_2fa_email_code(in_user_id);
	
		-- Revoke the user 2FA TOTP
		CALL revoke_user_2fa_totp(in_user_id);
	END IF;
END;
$$;
`

	// CreateSendUser2FAEmailCodeProc is the query to create the stored procedure to send user 2FA email code
	CreateSendUser2FAEmailCodeProc = `
CREATE OR REPLACE PROCEDURE send_user_2fa_email_code(
	IN in_user_id BIGINT,
	IN in_user_2fa_email_code VARCHAR,
	IN in_user_2fa_email_code_expires_at TIMESTAMP,
	OUT has_user_2fa_enabled BOOLEAN,
	OUT out_user_first_name VARCHAR,
	OUT out_user_last_name VARCHAR,
	OUT out_user_email VARCHAR
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Check if the user has 2FA enabled
	CALL has_user_2fa_enabled(in_user_id, has_user_2fa_enabled);

	IF has_user_2fa_enabled THEN
		-- Select the user first name, last name, and email by user ID
		SELECT
			users.first_name,
			users.last_name,
			user_emails.email
		INTO	
			out_user_first_name,
			out_user_last_name,
			out_user_email
		FROM	
			users
		INNER JOIN	
			user_emails ON users.id = user_emails.user_id
		WHERE
			users.id = in_user_id
		AND
			users.deleted_at IS NULL
		AND	
			user_emails.revoked_at IS NULL;

		-- Revoke the user 2FA email code
		CALL revoke_user_2fa_email_code(in_user_id);

		-- Create the user 2FA email code
		CALL create_user_2fa_email_code(in_user_id, in_user_2fa_email_code, in_user_2fa_email_code_expires_at);
	END IF;
END;
$$;
`

	// CreateHasUser2FATOTPEnabledProc is the query to create the stored procedure to check if the user has 2FA TOTP enabled
	CreateHasUser2FATOTPEnabledProc = `
CREATE OR REPLACE PROCEDURE has_user_2fa_totp_enabled(
	IN in_user_id BIGINT,
	OUT out_has_user_2fa_totp_enabled BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Check if the user is on the user_2fa_totp table
	SELECT
		user_2fa_totp.id IS NOT NULL
	INTO
		out_has_user_2fa_totp_enabled
	FROM
		user_2fa_totp
	WHERE
		user_2fa_totp.user_id = in_user_id
	AND
		user_2fa_totp.revoked_at IS NULL;
END;	
$$;
`

	// CreateGetUser2FAMethodsProc is the query to create the stored procedure to get user 2FA methods
	CreateGetUser2FAMethodsProc = `
CREATE OR REPLACE PROCEDURE get_user_2fa_methods(
	IN in_user_id BIGINT,
	OUT out_has_user_2fa_enabled BOOLEAN,
	OUT out_has_user_2fa_totp_enabled BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Check if the user has 2FA enabled
	CALL has_user_2fa_enabled(in_user_id, out_has_user_2fa_enabled);

	-- Check if the user has 2FA TOTP enabled
	CALL has_user_2fa_totp_enabled(in_user_id, out_has_user_2fa_totp_enabled);
END;
$$;
`
)
