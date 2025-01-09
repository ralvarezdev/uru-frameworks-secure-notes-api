package queries

const (
	// UsersInsert is the SQL query to insert a user
	UsersInsert = `
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
	$4
)
RETURNING
	id
`

	// UserUsernamesInsert is the SQL query to insert a user username
	UserUsernamesInsert = `
INSERT INTO user_usernames (
	user_id,
	username,
	assigned_at
)
VALUES (
	$1,
	$2,
	$3
)
RETURNING
	id
`

	// UserPasswordHashesInsert is the SQL query to insert a user password hash
	UserPasswordHashesInsert = `
INSERT INTO user_password_hashes (
	user_id,
	password_hash,
	assigned_at
)
VALUES (
	$1,
	$2,
	$3
)
RETURNING
	id
`

	// UserEmailsInsert is the SQL query to insert a user email
	UserEmailsInsert = `
INSERT INTO user_emails (
	user_id,
	email,
	assigned_at
)	
VALUES (
	$1,
	$2,
	$3
)
RETURNING
	id
`

	// UserPhoneNumbersInsert is the SQL query to insert a user phone number
	UserPhoneNumbersInsert = `
INSERT INTO user_phone_numbers (
	user_id,
	phone_number,
	assigned_at
)
VALUES (
	$1,
	$2,
	$3
)	
RETURNING
	id
`
)
