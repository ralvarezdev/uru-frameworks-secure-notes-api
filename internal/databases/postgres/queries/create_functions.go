package queries

const (
	// CreateGetUserRefreshTokenByIDFn is the SQL query to create the get user refresh token by ID function
	CreateGetUserRefreshTokenByIDFn = `
CREATE OR REPLACE FUNCTION get_user_refresh_token_by_id(
	in_refresh_token_id BIGINT,
	in_user_id BIGINT
) RETURNS 
TABLE(
	out_issued_at TIMESTAMP,
	out_expires_at TIMESTAMP,
	out_ip_address VARCHAR
) AS $$
BEGIN
	-- Return the user refresh token details
	RETURN QUERY
	SELECT
		user_refresh_tokens.issued_at AS out_issued_at,
		user_refresh_tokens.expires_at AS out_expires_at,
		user_refresh_tokens.ip_address AS out_ip_address
	FROM
		user_refresh_tokens
	WHERE
		user_refresh_tokens.id = in_refresh_token_id
	AND
		user_refresh_tokens.user_id = in_user_id
	AND
		user_refresh_tokens.revoked_at IS NULL;
END;
$$ LANGUAGE plpgsql;
`

	// CreateListUserRefreshTokensFn is the SQL query to create the list user refresh tokens function
	CreateListUserRefreshTokensFn = `
CREATE OR REPLACE FUNCTION list_user_refresh_tokens(
	in_user_id BIGINT
) RETURNS
TABLE(
	out_id BIGINT,
	out_issued_at TIMESTAMP,
	out_expires_at TIMESTAMP,
	out_ip_address VARCHAR
) AS $$
BEGIN
	-- Return the user refresh tokens details
	RETURN QUERY
	SELECT
		user_refresh_tokens.id AS out_id,
		user_refresh_tokens.issued_at AS out_issued_at,
		user_refresh_tokens.expires_at AS out_expires_at,
		user_refresh_tokens.ip_address AS out_ip_address
	FROM
		user_refresh_tokens
	WHERE
		user_refresh_tokens.user_id = in_user_id
	AND
		user_refresh_tokens.revoked_at IS NULL;
END;
$$ LANGUAGE plpgsql;
`

	// CreateListUserTokensFn is the SQL query to create the list user tokens function
	CreateListUserTokensFn = `
CREATE OR REPLACE FUNCTION list_user_tokens(
	in_user_id BIGINT
) RETURNS
TABLE(
	out_refresh_token_id BIGINT,
	out_access_token_id BIGINT
)
AS $$
BEGIN
	-- Return the user tokens
	RETURN QUERY
	SELECT
		user_refresh_tokens.id AS out_refresh_token_id,
		user_access_tokens.id AS out_access_token_id
	FROM
		user_refresh_tokens
	INNER JOIN
		user_access_tokens ON user_refresh_tokens.id = user_access_tokens.user_refresh_token_id
	WHERE
		user_refresh_tokens.user_id = in_user_id
	AND
		user_refresh_tokens.revoked_at IS NULL;
END;
$$ LANGUAGE plpgsql;
`
)
