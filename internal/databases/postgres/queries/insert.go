package queries

const (
	// InsertUser is the SQL query to insert a user
	InsertUser = `
INSERT INTO users (
	first_name,
	last_name,
	salt,
	joined_at
)
VALUES (
	$1,
	$2,
	$3,
	NOW()
)
RETURNING
	id
`

	// InsertUserUsername is the SQL query to insert a user username
	InsertUserUsername = `
INSERT INTO user_usernames (
	user_id,
	username,
	assigned_at
)
VALUES (
	$1,
	$2,
	NOW()
)
RETURNING
	id
`

	// InsertUserPasswordHash is the SQL query to insert a user password hash
	InsertUserPasswordHash = `
INSERT INTO user_password_hashes (
	user_id,
	password_hash,
	assigned_at
)
VALUES (
	$1,
	$2,
	NOW()
)
RETURNING
	id
`

	// InsertUserEmail is the SQL query to insert a user email
	InsertUserEmail = `
INSERT INTO user_emails (
	user_id,
	email,
	assigned_at
)	
VALUES (
	$1,
	$2,
	NOW()
)
RETURNING
	id
`

	// InsertUserPhoneNumber is the SQL query to insert a user phone number
	InsertUserPhoneNumber = `
INSERT INTO user_phone_numbers (
	user_id,
	phone_number,
	assigned_at
)
VALUES (
	$1,
	$2,
	NOW()
)	
RETURNING
	id
`

	// InsertUserFailedLogInAttempt is the SQL query to insert a user failed login attempt
	InsertUserFailedLogInAttempt = `
INSERT INTO user_failed_log_in_attempts (
	user_id,
	ip_address,
	bad_password,
	bad_2fa_code,
	attempted_at
)
VALUES (
	$1,
	$2,
	$3,
	$4,
	NOW()
)
`

	// InsertUserRefreshToken is the SQL query to insert a user refresh token
	InsertUserRefreshToken = `
INSERT INTO user_refresh_tokens (
	user_id,
	parent_user_refresh_token_id,
	ip_address,
	issued_at,
	expires_at
)
VALUES (
	$1,
	$2,
	$3,
	$4,
	$5
	)
RETURNING
	id
`

	// InsertUserAccessToken is the SQL query to insert a user access token
	InsertUserAccessToken = `
INSERT INTO user_access_tokens (
	user_id,
	user_refresh_token_id,
	issued_at,
	expires_at
)
VALUES (
	$1,
	$2,
	$3,
	$4
)
RETURNING
	id
`

	// InsertUserTOTP is the SQL query to insert a user TOTP key
	InsertUserTOTP = `
INSERT INTO user_totps (
	user_id,
	secret,
	created_at
)
VALUES (
	$1,
	$2,
	NOW()
)
RETURNING
	id
`
)

var (
	// InsertUserTOTPRecoveryCodes is the SQL query to insert the user TOTP recovery codes
	// This is dynamically set based on the recovery codes count
	InsertUserTOTPRecoveryCodes string
)
