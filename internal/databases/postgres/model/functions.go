package model

var (
	// GetUserRefreshTokenByIDFn is the query to get a refresh token by ID
	GetUserRefreshTokenByIDFn = "SELECT * FROM get_user_refresh_token_by_id($1, $2);"

	// ListUserRefreshTokensFn is the query to list user refresh tokens
	ListUserRefreshTokensFn = "SELECT * FROM list_user_refresh_tokens($1);"

	// ListUserTokensFn is the query to list user tokens
	ListUserTokensFn = "SELECT * FROM list_user_tokens($1);"

	// ListUserTagsFn is the query to list user tags
	ListUserTagsFn = "SELECT * FROM list_user_tags($1);"

	// ListUserNoteVersionsFn is the query to list user note versions
	ListUserNoteVersionsFn = "SELECT * FROM list_user_note_versions($1, $2);"

	// ListUserNoteTagsFn is the query to list user note tags
	ListUserNoteTagsFn = "SELECT * FROM list_user_note_tags($1, $2);"

	// SyncUserNoteVersionsByLastSyncedAtFn is the query to sync user note versions by last synced at
	SyncUserNoteVersionsByLastSyncedAtFn = "SELECT * FROM sync_user_note_versions_by_last_synced_at($1, $2, $3);"

	// SyncUserNotesByLastSyncedAtFn is the query to sync user notes by last synced at
	SyncUserNotesByLastSyncedAtFn = "SELECT * FROM sync_user_notes_by_last_synced_at($1, $2, $3);"

	// SyncUserNoteTagsByLastSyncedAtFn is the query to sync user note tags by last synced at
	SyncUserNoteTagsByLastSyncedAtFn = "SELECT * FROM sync_user_note_tags_by_last_synced_at($1, $2);"

	// SyncUserTagsByLastSyncedAtFn is the query to sync user tags by last synced at
	SyncUserTagsByLastSyncedAtFn = "SELECT * FROM sync_user_tags_by_last_synced_at($1, $2);"

	// GetUserTagByIDFn is the query to get user tag by tag ID
	GetUserTagByIDFn = "SELECT * FROM get_user_tag_by_id($1, $2, $3)"

	// GetUserNoteVersionByIDFn is the query to get user note version by note version ID
	GetUserNoteVersionByIDFn = "SELECT * FROM get_user_note_version_by_id($1, $2)"

	// GetUserNoteByIDFn is the query to get user note by note ID
	GetUserNoteByIDFn = "SELECT * FROM get_user_note_by_id($1, $2)"

	// GetLogInInformationFn is the query to get the login information
	GetLogInInformationFn = "SELECT * FROM get_log_in_information($1);"
)
