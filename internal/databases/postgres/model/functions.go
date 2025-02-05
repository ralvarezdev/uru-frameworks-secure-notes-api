package model

var (
	// GetUserRefreshTokenByIDFn is the SQL query to get a refresh token by ID
	GetUserRefreshTokenByIDFn = "SELECT * FROM get_user_refresh_token_by_id($1, $2);"

	// ListUserRefreshTokensFn is the SQL query to list user refresh tokens
	ListUserRefreshTokensFn = "SELECT * FROM list_user_refresh_tokens($1);"

	// ListUserTokensFn is the SQL query to list user tokens
	ListUserTokensFn = "SELECT * FROM list_user_tokens($1);"

	// ListUserTagsFn is the SQL query to list user tags
	ListUserTagsFn = "SELECT * FROM list_user_tags($1);"

	// ListUserNoteVersionsFn is the SQL query to list user note versions
	ListUserNoteVersionsFn = "SELECT * FROM list_user_note_versions($1, $2);"

	// SyncUserNoteVersionsFn is the SQL query to sync user note versions
	SyncUserNoteVersionsFn = "SELECT * FROM sync_user_note_versions($1, $2, $3);"

	// ListUserNotesTagsFn is the SQL query to list user notes tags
	ListUserNotesTagsFn = "SELECT * FROM list_user_notes_tags($1, $2);"
)
