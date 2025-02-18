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
		user_tags.user_id = in_user_id
	AND
		user_tags.deleted_at IS NULL;
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
		user_note_versions.user_note_id = in_user_note_id
	AND
		user_note_versions.deleted_at IS NULL;
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
		user_note_tags.user_note_id = in_user_note_id
	AND
		user_note_tags.deleted_at IS NULL;
END;	
$$ LANGUAGE plpgsql;
`

	// CreateSyncUserNoteVersionsByLastSyncedAtFn is the query to create the function to sync user note versions by last synced at
	CreateSyncUserNoteVersionsByLastSyncedAtFn = `
CREATE OR REPLACE FUNCTION sync_user_note_versions_by_last_synced_at(
	in_user_id BIGINT,
	in_user_note_id BIGINT,
	in_last_synced_at TIMESTAMP
) RETURNS
TABLE(
	out_user_note_version_id BIGINT,
	out_user_note_version_encrypted_content TEXT,
	out_user_note_version_created_at TIMESTAMP,
	out_user_note_version_deleted_at TIMESTAMP
)
AS $$
BEGIN	
	-- Return the user note versions
	RETURN QUERY
	SELECT
		user_note_versions.id AS out_user_note_version_id,
		user_note_versions.encrypted_content AS out_user_note_version_encrypted_content,
		user_note_versions.created_at AS out_user_note_version_created_at,
		user_note_versions.deleted_at AS out_user_note_version_deleted_at
	FROM
		user_note_versions
	INNER JOIN
		user_notes ON user_note_versions.user_note_id = user_notes.id
	WHERE
		user_notes.user_id = in_user_id
	AND
		user_note_versions.user_note_id = in_user_note_id
	AND (
		last_synced_at IS NULL
	OR
		user_note_versions.created_at > in_last_synced_at
	OR
		user_note_versions.deleted_at > in_last_synced_at
	);
END;
$$ LANGUAGE plpgsql;
`

	// CreateSyncUserTagsByLastSyncedAtFn is the query to create the function to sync user tags by last synced at
	CreateSyncUserTagsByLastSyncedAtFn = `
CREATE OR REPLACE FUNCTION sync_user_tags_by_last_synced_at(
	in_user_id BIGINT,
	in_last_synced_at TIMESTAMP
) RETURNS
TABLE(	
	out_user_tag_id BIGINT,
	out_user_tag_name VARCHAR,
	out_user_tag_created_at TIMESTAMP,
	out_user_tag_updated_at TIMESTAMP,
	out_user_tag_deleted_at TIMESTAMP
)
AS $$
BEGIN
	-- Return the user tags
	RETURN QUERY
	SELECT
		user_tags.id AS out_user_tag_id,
		user_tags.name AS out_user_tag_name,
		user_tags.created_at AS out_user_tag_created_at,
		user_tags.updated_at AS out_user_tag_updated_at,
		user_tags.deleted_at AS out_user_tag_deleted_at
	FROM
		user_tags
	WHERE
		user_tags.user_id = in_user_id
	AND (
		in_last_synced_at IS NULL
	OR
		user_tags.created_at > in_last_synced_at
	OR
		user_tags.updated_at > in_last_synced_at
	OR
		user_tags.deleted_at > in_last_synced_at
	);
END;
$$ LANGUAGE plpgsql;
`

	// CreateSyncUserNoteTagsByLastSyncedAtFn is the query to create the function to sync user note tags by last synced at
	CreateSyncUserNoteTagsByLastSyncedAtFn = `
CREATE OR REPLACE FUNCTION sync_user_note_tags_by_last_synced_at(
	in_user_id BIGINT,
	in_last_synced_at TIMESTAMP,
	in_user_note_id BIGINT
) RETURNS	
TABLE(	
	out_user_note_tag_user_tag_id BIGINT,
	out_user_note_tag_assigned_at TIMESTAMP,
	out_user_note_tag_deleted_at TIMESTAMP
)
AS $$	
BEGIN	
	-- Return the user note tags
	RETURN QUERY
	SELECT
		user_note_tags.user_tag_id AS out_user_note_tag_user_tag_id,
		user_note_tags.assigned_at AS out_user_note_tag_assigned_at,
		user_note_tags.deleted_at AS out_user_note_tag_deleted_at
	FROM
		user_note_tags
	INNER JOIN
		user_tags ON user_note_tags.user_tag_id = user_tags.id
	INNER JOIN
		user_notes ON user_note_tags.user_note_id = user_notes.id
	WHERE
		user_notes.user_id = in_user_id
	AND
		user_note_tags.user_note_id = in_user_note_id
	AND (
		in_last_synced_at IS NULL
	OR
		user_note_tags.assigned_at > in_last_synced_at
	OR
		user_note_tags.deleted_at > in_last_synced_at
	);
END;
$$ LANGUAGE plpgsql;
`

	// CreateSyncUserNotesByLastSyncedAtFn is the query to create the function to sync user notes by last synced at
	CreateSyncUserNotesByLastSyncedAtFn = `
CREATE OR REPLACE FUNCTION sync_user_notes_by_last_synced_at(
	in_user_id BIGINT,
	in_last_synced_at TIMESTAMP
) RETURNS
TABLE(
	out_user_note_id BIGINT,
	out_user_note_title VARCHAR,
	out_user_note_color VARCHAR,
	out_user_note_created_at TIMESTAMP,
	out_user_note_updated_at TIMESTAMP,
	out_user_note_pinned_at TIMESTAMP,
	out_user_note_starred_at TIMESTAMP,
	out_user_note_archived_at TIMESTAMP,
	out_user_note_trashed_at TIMESTAMP,
	out_user_note_deleted_at TIMESTAMP,
	out_user_note_has_to_sync_note_tags BOOLEAN,
	out_user_note_has_to_sync_note_versions BOOLEAN
)
AS $$
BEGIN
	-- Return the user notes
	RETURN QUERY
	SELECT *
	FROM (
		SELECT
			user_notes.id AS out_user_note_id,
			user_notes.title AS out_user_note_title,
			user_notes.color AS out_user_note_color,
			user_notes.created_at AS out_user_note_created_at,
			user_notes.updated_at AS out_user_note_updated_at,
			user_notes.pinned_at AS out_user_note_pinned_at,
			user_notes.starred_at AS out_user_note_starred_at,
			user_notes.archived_at AS out_user_note_archived_at,
			user_notes.trashed_at AS out_user_note_trashed_at,
			user_notes.deleted_at AS out_user_note_deleted_at,
			CASE
				WHEN EXISTS (
					SELECT
						1
					FROM
						user_note_tags
					WHERE
						user_note_tags.user_note_id = user_notes.id
					AND	 (
						user_note_tags.assigned_at > in_last_synced_at
					OR
						user_note_tags.deleted_at > in_last_synced_at
					)	
				) THEN TRUE
				ELSE FALSE
			END AS out_user_note_has_to_sync_note_tags,
			CASE
				WHEN EXISTS (
					SELECT
						1
					FROM
						user_note_versions
					WHERE
						user_note_versions.user_note_id = user_notes.id
					AND	(
						user_note_versions.created_at > in_last_synced_at
					OR	
						user_note_versions.deleted_at > in_last_synced_at
					)	
				) THEN TRUE	
				ELSE FALSE
			END AS out_user_note_has_to_sync_note_versions
		FROM
			user_notes
		WHERE
			user_notes.user_id = in_user_id
	) AS user_notes
	WHERE
		in_last_synced_at IS NULL
	OR
		user_notes.out_user_note_has_to_sync_note_tags = TRUE
	OR
		user_notes.out_user_note_has_to_sync_note_versions = TRUE
	OR
		user_notes.out_user_note_created_at > in_last_synced_at
	OR
		user_notes.out_user_note_updated_at > in_last_synced_at
	OR
		user_notes.out_user_note_deleted_at > in_last_synced_at;
END;
$$ LANGUAGE plpgsql;
`

	// CreateGetUserTagByIDFn is the query to create the function to get user tag by tag ID
	CreateGetUserTagByIDFn = `
CREATE OR REPLACE FUNCTION get_user_tag_by_id(
	in_user_id BIGINT,
	in_user_tag_id BIGINT
) RETURNS
TABLE(
	out_user_tag_name VARCHAR,
	out_user_tag_created_at TIMESTAMP,
	out_user_tag_updated_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user tag name, created at, and updated at by user ID and tag ID
	RETURN QUERY
	SELECT
		user_tags.name AS out_user_tag_name,
		user_tags.created_at AS out_user_tag_created_at,
		user_tags.updated_at AS out_user_tag_updated_at
	FROM
		user_tags
	WHERE
		user_tags.id = in_user_tag_id
	AND
		user_tags.user_id = in_user_id
	AND
		user_tags.deleted_at IS NULL;
END;
$$;
`

	// CreateGetUserNoteVersionByIDFn is the query to create the function to get user note version by note version ID
	CreateGetUserNoteVersionByIDFn = `
CREATE OR REPLACE FUNCTION get_user_note_version_by_id(
	in_user_id BIGINT,
	in_user_note_version_id BIGINT
) RETURNS
TABLE(
	out_user_note_version_encrypted_content TEXT,
	out_user_note_version_created_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user note version encrypted content and created at by user ID and user note version ID
	RETURN QUERY
	SELECT
		user_note_versions.encrypted_content AS out_user_note_version_encrypted_content,
		user_note_versions.created_at AS out_user_note_version_created_at
	FROM
		user_note_versions
	INNER JOIN
		user_notes ON user_note_versions.note_id = user_notes.id
	WHERE
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

	// CreateGetUserNoteByIDFn is the query to create the function to get user note by note ID
	CreateGetUserNoteByIDFn = `
CREATE OR REPLACE FUNCTION get_user_note_by_id(
	in_user_id BIGINT,
	in_user_note_id BIGINT
) RETURNS
TABLE(
	out_user_note_title VARCHAR,
	out_user_note_color VARCHAR,
	out_user_note_created_at TIMESTAMP,
	out_user_note_updated_at TIMESTAMP,
	out_user_note_pinned_at TIMESTAMP,
	out_user_note_archived_at TIMESTAMP,
	out_user_note_trashed_at TIMESTAMP,
	out_user_note_starred_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user note title, color, created at, updated at, pinned at, archived at, trashed at, and starred at by user ID and user note ID
	RETURN QUERY
	SELECT
		user_notes.title AS out_user_note_title,
		user_notes.color AS out_user_note_color,
		user_notes.created_at AS out_user_note_created_at,
		user_notes.updated_at AS out_user_note_updated_at,
		user_notes.pinned_at AS out_user_note_pinned_at,
		user_notes.archived_at AS out_user_note_archived_at,
		user_notes.trashed_at AS out_user_note_trashed_at,
		user_notes.starred_at AS out_user_note_starred_at
	FROM
		user_notes
	INNER JOIN
		user_note_versions ON user_notes.id = user_note_versions.note_id
	WHERE
		user_notes.id = in_user_note_id
	AND
		user_notes.user_id = in_user_id
	AND
		user_notes.deleted_at IS NULL
	AND
		user_note_versions.deleted_at IS NULL
	ORDER BY
		user_note_versions.created_at DESC
	LIMIT 1;
END;
$$;
`

	// CreateGetLogInInformationFn is the query to create the function to get log in information
	CreateGetLogInInformationFn = `
CREATE OR REPLACE FUNCTION get_log_in_information(
	in_user_username VARCHAR,
	in_ip_address VARCHAR,
	in_maximum_failed_attempts_count INT,
	in_maximum_failed_attempts_period_seconds INT
) RETURNS
TABLE(
	out_user_id BIGINT,
	out_user_salt VARCHAR,
	out_user_encrypted_key TEXT,
	out_user_password_hash VARCHAR,
	out_too_many_failed_attempts BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
	-- Select the user ID and password hash by username
	RETURN QUERY
	SELECT
		users.id AS out_user_id,
		users.salt AS out_user_salt,
		users.encrypted_key AS out_user_encrypted_key,
		user_password_hashes.password_hash AS out_user_password_hash,
		CASE
			WHEN
				(	
					SELECT
						COUNT(*)
					FROM
						user_failed_log_in_attempts
					WHERE	
						user_failed_log_in_attempts.user_id = users.id	
					AND
						user_failed_log_in_attempts.attempted_at >= NOW() - INTERVAL '1 second' * in_maximum_failed_attempts_period_seconds
					AND
						user_failed_log_in_attempts.ip_address = in_ip_address
				) >= in_maximum_failed_attempts_count THEN TRUE
			ELSE FALSE
		END AS out_too_many_failed_attempts
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
END;
$$;
`
)
