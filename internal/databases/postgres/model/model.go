package model

import (
	"database/sql"
	"time"
)

type (
	// User is the Postgres model for the user
	User struct {
		ID        int64
		FirstName string
		LastName  string
		Salt      string
		Birthdate sql.NullTime
		JoinedAt  time.Time
		DeletedAt sql.NullTime
	}

	// UserUsername is the Postgres model for the user username
	UserUsername struct {
		ID         int64
		UserID     int64
		Username   string
		AssignedAt time.Time
		RevokedAt  sql.NullTime
	}

	// UserPasswordHash is the Postgres model for the user password hash
	UserPasswordHash struct {
		ID           int64
		UserID       int64
		PasswordHash string
		AssignedAt   time.Time
		RevokedAt    sql.NullTime
	}

	// UserResetPassword is the Postgres model for the user password reset
	UserResetPassword struct {
		ID         int64
		UserID     int64
		ResetToken string
		CreatedAt  time.Time
		ExpiresAt  time.Time
		RevokedAt  sql.NullTime
	}

	// UserEmail is the Postgres model for the user email
	UserEmail struct {
		ID         int64
		UserID     int64
		Email      string
		AssignedAt time.Time
		RevokedAt  sql.NullTime
	}

	// UserEmailVerification is the Postgres model for the user email verification
	UserEmailVerification struct {
		ID                int64
		UserEmailID       int64
		VerificationToken string
		CreatedAt         time.Time
		ExpiresAt         time.Time
		VerifiedAt        sql.NullTime
		RevokedAt         sql.NullTime
	}

	// UserPhoneNumber is the Postgres model for the user phone number
	UserPhoneNumber struct {
		ID          int64
		UserID      int64
		PhoneNumber string
		AssignedAt  time.Time
		RevokedAt   sql.NullTime
	}

	// UserPhoneNumberVerification is the Postgres model for the user phone number verification
	UserPhoneNumberVerification struct {
		ID                int64
		UserPhoneNumberID int64
		VerificationCode  string
		CreatedAt         time.Time
		ExpiresAt         time.Time
		VerifiedAt        sql.NullTime
		RevokedAt         sql.NullTime
	}

	// UserFailedLogInAttempt is the Postgres model for the user failed login attempt
	UserFailedLogInAttempt struct {
		ID          int64
		UserID      int64
		IPAddress   string
		BadPassword sql.NullBool
		Bad2FACode  sql.NullBool
		AttemptedAt time.Time
	}

	// UserRefreshToken is the Postgres model for the user refresh token
	UserRefreshToken struct {
		ID                   int64
		UserID               int64
		ParentRefreshTokenID sql.NullInt64
		IPAddress            string
		IssuedAt             time.Time
		ExpiresAt            time.Time
		RevokedAt            sql.NullTime
	}

	// UserAccessToken is the Postgres model for the user access token
	UserAccessToken struct {
		ID                 int64
		UserID             int64
		UserRefreshTokenID int64
		IssuedAt           time.Time
		ExpiresAt          time.Time
		RevokedAt          sql.NullTime
	}

	// UserTOTP is the Postgres model for the user TOTP
	UserTOTP struct {
		ID         int64
		UserID     int64
		Secret     string
		CreatedAt  time.Time
		VerifiedAt sql.NullTime
		RevokedAt  sql.NullTime
	}

	// UserTOTPRecoveryCode is the Postgres model for the user TOTP recovery code
	UserTOTPRecoveryCode struct {
		ID         int64
		UserTOTPID int64
		Code       string
		RevokedAt  sql.NullTime
	}

	// Tag is the Postgres model for the tag
	Tag struct {
		ID     int64
		UserID int64
		Name   string
	}

	// Note is the Postgres model for the user note
	Note struct {
		ID        int64
		UserID    int64
		IsPinned  sql.NullBool
		Title     string
		Color     sql.NullString
		CreatedAt time.Time
	}

	// NoteTag is the join table for the many-to-many relationship between Note and NoteTag
	NoteTag struct {
		ID         int64
		NoteID     int64
		TagID      int64
		AssignedAt time.Time
	}

	// NoteVersion is the Postgres model for the note version
	NoteVersion struct {
		ID            string
		NoteID        int64
		EncryptedBody string
		CreatedAt     time.Time
	}
)
