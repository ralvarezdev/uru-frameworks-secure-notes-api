package model

const (
	// CreateSignUpProc is the query to create the sign-up stored procedure
	CreateSignUpProc = `
CREATE OR REPLACE PROCEDURE sign_up(
	IN in_first_name VARCHAR,
	IN in_last_name VARCHAR,
	IN in_salt VARCHAR,
	IN in_username VARCHAR,
	IN in_email VARCHAR,
	IN in_password_hash VARCHAR,
	OUT out_user_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Insert into users table
	INSERT INTO users (
		first_name,
		last_name, 
		salt
	) 
	VALUES (
		in_first_name, 
		in_last_name, 
		in_salt
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
	);

	-- Insert into user_password_hashes table
	INSERT INTO user_password_hashes (
		user_id, 
		password_hash
	) 
	VALUES (
		out_user_id, 
		in_password_hash
	);
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
	// CreateRevokeTOTPProc is the query to create the revoke TOTP stored procedure

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
CREATE OR REPLACE PROCEDURE pre_login(
	IN in_username VARCHAR,
	OUT out_user_id BIGINT,
	OUT out_password_hash VARCHAR,
	OUT out_totp_id BIGINT,
	OUT out_totp_secret VARCHAR
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user ID and password hash by username
	SELECT
		users.id,
		user_password_hashes.password_hash
	INTO
		out_user_id,
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
)
