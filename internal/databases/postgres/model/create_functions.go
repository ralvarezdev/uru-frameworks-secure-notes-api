package model

const (
	// CreateGetUserRefreshTokenByIDFn is the query to create the function to get user refresh token by ID
	CreateGetUserRefreshTokenByIDFn = `
CREATE OR REPLACE FUNCTION get_user_refresh_token_by_id(
	in_user_refresh_token_id BIGINT,
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
		user_refresh_tokens.id = in_user_refresh_token_id
	AND
		user_refresh_tokens.user_id = in_user_id
	AND
		user_refresh_tokens.revoked_at IS NULL;
END;
$$ LANGUAGE plpgsql;
`

	// CreateListUserRefreshTokensFn is the query to create the function to list user refresh tokens
	CreateListUserRefreshTokensFn = `
CREATE OR REPLACE FUNCTION list_user_refresh_tokens(
	in_user_id BIGINT
) RETURNS
TABLE(
	out_user_refresh_token_id BIGINT,
	out_user_refresh_token_issued_at TIMESTAMP,
	out_user_refresh_token_expires_at TIMESTAMP,
	out_user_refresh_token_ip_address VARCHAR
) AS $$
BEGIN
	-- Return the user refresh tokens details
	RETURN QUERY
	SELECT
		user_refresh_tokens.id AS out_user_refresh_token_id,
		user_refresh_tokens.issued_at AS out_user_refresh_token_issued_at,
		user_refresh_tokens.expires_at AS out_user_refresh_token_expires_at,
		user_refresh_tokens.ip_address AS out_user_refresh_token_ip_address
	FROM
		user_refresh_tokens
	WHERE
		user_refresh_tokens.user_id = in_user_id
	AND
		user_refresh_tokens.revoked_at IS NULL;
END;
$$ LANGUAGE plpgsql;
`

	// CreateListUserTokensFn is the query to create the function to list user tokens
	CreateListUserTokensFn = `
CREATE OR REPLACE FUNCTION list_user_tokens(
	in_user_id BIGINT
) RETURNS
TABLE(
	out_user_refresh_token_id BIGINT,
	out_user_access_token_id BIGINT
)
AS $$
BEGIN
	-- Return the user tokens
	RETURN QUERY
	SELECT
		user_refresh_tokens.id AS out_user_refresh_token_id,
		user_access_tokens.id AS out_user_access_token_id
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

	// CreateListUserTagsFn is the query to create the function to list user tags
	CreateListUserTagsFn = `
CREATE OR REPLACE FUNCTION list_user_tags(
	in_user_id BIGINT
) RETURNS
TABLE(
	out_user_tag_id BIGINT,
	out_user_tag_name VARCHAR,
	out_user_tag_created_at TIMESTAMP,
	out_user_tag_updated_at TIMESTAMP
)
AS $$
BEGIN
	-- Return the user tags
	RETURN QUERY
	SELECT
		user_tags.id AS out_user_tag_id,
		user_tags.name AS out_user_tag_name,
		user_tags.created_at AS out_user_tag_created_at,
		user_tags.updated_at AS out_user_tag_updated_at
	FROM
		user_tags
	WHERE
		user_tags.user_id = in_user_id;
END;
$$ LANGUAGE plpgsql;
`

	// CreateListUserNoteVersionsFn is the query to create the function to list user note versions
	CreateListUserNoteVersionsFn = `
CREATE OR REPLACE FUNCTION list_user_note_versions(
	in_user_id BIGINT,
	in_user_note_id BIGINT
) RETURNS
TABLE(
	out_user_note_version_id BIGINT
)
AS $$
BEGIN
	-- Return the user note versions
	RETURN QUERY
	SELECT
		user_note_versions.id AS out_user_note_version_id
	FROM
		user_note_versions
	INNER JOIN
		user_notes ON user_note_versions.user_note_id = user_notes.id
	WHERE
		user_notes.user_id = in_user_id
	AND
		user_note_versions.user_note_id = in_user_note_id;
END;
$$ LANGUAGE plpgsql;
`

	// CreateSyncUserNoteVersionsFn is the query to create the function to sync user note versions
	CreateSyncUserNoteVersionsFn = `
CREATE OR REPLACE FUNCTION sync_user_note_versions(
	in_user_id BIGINT,
	in_user_note_id BIGINT,
	in_latest_user_note_version_id BIGINT
) RETURNS
TABLE(
	out_user_note_version_id BIGINT,
	out_user_note_version_encrypted_content TEXT,
	out_user_note_version_created_at TIMESTAMP
)
AS $$
BEGIN
	-- Return the user note versions
	RETURN QUERY
	SELECT
		user_note_versions.id AS out_user_note_version_id,
		user_note_versions.encrypted_content AS out_user_note_version_encrypted_content,
		user_note_versions.created_at AS out_user_note_version_created_at
	FROM
		user_note_versions
	INNER JOIN
		user_notes ON user_note_versions.user_note_id = user_notes.id
	WHERE
		user_notes.user_id = in_user_id
	AND
		user_note_versions.user_note_id = in_user_note_id
	AND
		user_note_versions.id > in_latest_user_note_version_id;
END;
$$ LANGUAGE plpgsql;
`

	// CreateListUserNoteTagsFn is the query to create the function to list user note tags
	CreateListUserNoteTagsFn = `
CREATE OR REPLACE FUNCTION list_user_note_tags(
	in_user_id BIGINT,
	in_user_note_id BIGINT
) RETURNS	
TABLE(
	out_user_tag_id BIGINT,
	out_user_note_tag_assigned_at TIMESTAMP
)
AS $$	
BEGIN	
	-- Return the user note tags
	RETURN QUERY
	SELECT
		user_note_tags.user_tag_id AS out_user_tag_id,
		user_note_tags.assigned_at AS out_user_note_tag_assigned_at
	FROM
		user_note_tags
	INNER JOIN
		user_tags ON user_note_tags.user_tag_id = user_tags.id
	INNER JOIN
		user_notes ON user_note_tags.user_note_id = user_notes.id
	WHERE
		user_notes.user_id = in_user_id
	AND
		user_note_tags.user_note_id = in_user_note_id;
END;	
$$ LANGUAGE plpgsql;
`
)
