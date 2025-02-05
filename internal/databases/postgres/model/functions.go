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
)
