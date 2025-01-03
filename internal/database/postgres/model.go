package postgres

import (
	"time"
)

type (
	// Model is the base model for all Postgres models
	Model struct {
		ID uint `json:"id" gorm:"primaryKey"`
	}

	// User is the Postgres model for the user
	User struct {
		Model
		FirstName string     `json:"first_name" gorm:"not null"`
		LastName  string     `json:"last_name" gorm:"not null"`
		Salt      string     `json:"salt" gorm:"not null"`
		Birthdate *time.Time `json:"birthdate,omitempty"`
		JoinedAt  time.Time  `json:"joined_at" gorm:"not null"`
		DeletedAt *time.Time `json:"deleted_at,omitempty"`
	}

	// UserUsername is the Postgres model for the user username
	UserUsername struct {
		Model
		UserID     uint       `json:"user_id" gorm:"not null"`
		User       User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
		Username   string     `json:"username" gorm:"uniqueIndex:idx_user_username,where:revoked_at IS NULL;not null"`
		AssignedAt time.Time  `json:"assigned_at" gorm:"not null"`
		RevokedAt  *time.Time `json:"revoked_at,omitempty"`
	}

	// UserHashedPassword is the Postgres model for the user hashed password
	UserHashedPassword struct {
		Model
		UserID         uint       `json:"user_id" gorm:"not null"`
		User           User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
		HashedPassword string     `json:"hashed_password" gorm:"not null"`
		AssignedAt     time.Time  `json:"assigned_at" gorm:"not null"`
		RevokedAt      *time.Time `json:"revoked_at,omitempty"`
	}

	// UserResetPassword is the Postgres model for the user password reset
	UserResetPassword struct {
		Model
		UserID     uint       `json:"user_id" gorm:"not null"`
		User       User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
		ResetToken string     `json:"reset_token" gorm:"uniqueIndex;not null"`
		CreatedAt  time.Time  `json:"created_at" gorm:"not null"`
		ExpiresAt  time.Time  `json:"expires_at" gorm:"not null"`
		RevokedAt  *time.Time `json:"revoked_at,omitempty"`
	}

	// UserEmail is the Postgres model for the user email
	UserEmail struct {
		Model
		UserID     uint       `json:"user_id" gorm:"uniqueIndex:idx_user_email,where:revoked_at IS NULL;not null"`
		User       User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
		Email      string     `json:"email" gorm:"uniqueIndex:idx_user_email,where:revoked_at IS NULL;not null"`
		AssignedAt time.Time  `json:"assigned_at" gorm:"not null"`
		RevokedAt  *time.Time `json:"revoked_at,omitempty"`
	}

	// UserEmailVerification is the Postgres model for the user email verification
	UserEmailVerification struct {
		Model
		UserEmailID       uint       `json:"user_email_id" gorm:"not null"`
		UserEmail         UserEmail  `gorm:"foreignKey:UserEmailID;constraint:OnDelete:CASCADE"`
		VerificationToken string     `json:"verification_token" gorm:"not null"`
		CreatedAt         time.Time  `json:"created_at" gorm:"not null"`
		ExpiresAt         time.Time  `json:"expires_at" gorm:"not null"`
		VerifiedAt        *time.Time `json:"verified_at,omitempty"`
		RevokedAt         *time.Time `json:"revoked_at,omitempty"`
	}

	// UserPhoneNumber is the Postgres model for the user phone number
	UserPhoneNumber struct {
		Model
		UserID      uint       `json:"user_id" gorm:"uniqueIndex:idx_user_phone_number,where:revoked_at IS NULL;not null"`
		User        User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
		PhoneNumber string     `json:"phone_number" gorm:"uniqueIndex:idx_user_phone_number,where:revoked_at IS NULL;not null"`
		AssignedAt  time.Time  `json:"assigned_at" gorm:"not null"`
		RevokedAt   *time.Time `json:"revoked_at,omitempty"`
	}

	// UserPhoneNumberVerification is the Postgres model for the user phone number verification
	UserPhoneNumberVerification struct {
		Model
		UserPhoneNumberID uint            `json:"user_phone_number_id" gorm:"not null"`
		UserPhoneNumber   UserPhoneNumber `gorm:"foreignKey:UserPhoneNumberID;constraint:OnDelete:CASCADE"`
		VerificationCode  string          `json:"verification_code" gorm:"not null"`
		CreatedAt         time.Time       `json:"created_at" gorm:"not null"`
		ExpiresAt         time.Time       `json:"expires_at" gorm:"not null"`
		VerifiedAt        *time.Time      `json:"verified_at,omitempty"`
		RevokedAt         *time.Time      `json:"revoked_at,omitempty"`
	}

	// UserFailedLogInAttempt is the Postgres model for the user failed login attempt
	UserFailedLogInAttempt struct {
		Model
		UserID          uint           `json:"user_id" gorm:"not null"`
		User            User           `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
		UserTokenSeedID *uint          `json:"user_token_seed_id,omitempty"`
		UserTokenSeed   *UserTokenSeed `gorm:"foreignKey:UserTokenSeedID"`
		IPv4Address     string         `json:"ipv4_address" gorm:"not null"`
		BadPassword     *bool          `json:"bad_password,omitempty"`
		Bad2FACode      *bool          `json:"bad_2fa_code,omitempty"`
		AttemptedAt     time.Time      `json:"attempted_at" gorm:"not null"`
	}

	// UserTokenSeed is the Postgres model for the user token seed
	UserTokenSeed struct {
		Model
		UserID    uint       `json:"user_id" gorm:"not null"`
		User      User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
		TokenSeed string     `json:"token_seed" gorm:"uniqueIndex;not null"`
		CreatedAt time.Time  `json:"created_at" gorm:"not null"`
		ExpiresAt time.Time  `json:"expires_at" gorm:"not null"`
		RevokedAt *time.Time `json:"revoked_at,omitempty"`
	}

	// UserRefreshToken is the Postgres model for the user refresh token
	UserRefreshToken struct {
		Model
		UserID               uint              `json:"user_id" gorm:"not null"`
		User                 User              `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
		ParentRefreshTokenID *uint             `json:"parent_refresh_token_id,omitempty" gorm:"uniqueIndex:idx_parent_refresh_token,where:parent_refresh_token_id IS NOT NULL"`
		ParentRefreshToken   *UserRefreshToken `gorm:"foreignKey:ParentRefreshTokenID;constraint:OnDelete:CASCADE"`
		UserTokenSeedID      *uint             `json:"user_token_seed_id,omitempty" gorm:"uniqueIndex:idx_user_token_seed,where:user_token_seed_id IS NOT NULL"`
		UserTokenSeed        *UserTokenSeed    `gorm:"foreignKey:UserTokenSeedID;constraint:OnDelete:CASCADE"`
		IPv4Address          string            `json:"ipv4_address" gorm:"not null"`
		IssuedAt             time.Time         `json:"issued_at" gorm:"not null"`
		ExpiresAt            time.Time         `json:"expires_at" gorm:"not null"`
		RevokedAt            *time.Time        `json:"revoked_at,omitempty"`
	}

	// UserAccessToken is the Postgres model for the user access token
	UserAccessToken struct {
		Model
		UserID             uint             `json:"user_id" gorm:"not null"`
		User               User             `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
		UserRefreshTokenID uint             `json:"user_refresh_token_id" gorm:"uniqueIndex;not null"`
		UserRefreshToken   UserRefreshToken `gorm:"foreignKey:UserRefreshTokenID;constraint:OnDelete:CASCADE"`
		IssuedAt           time.Time        `json:"issued_at" gorm:"not null"`
		ExpiresAt          time.Time        `json:"expires_at" gorm:"not null"`
		RevokedAt          *time.Time       `json:"revoked_at,omitempty"`
	}

	// UserTOTP is the Postgres model for the user TOTP
	UserTOTP struct {
		Model
		UserID     uint       `json:"user_id" gorm:"not null"`
		User       User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
		Secret     string     `json:"secret" gorm:"uniqueIndex;not null"`
		CreatedAt  time.Time  `json:"created_at" gorm:"not null"`
		ExpiresAt  time.Time  `json:"expires_at" gorm:"not null"`
		VerifiedAt *time.Time `json:"verified_at,omitempty"`
		RevokedAt  *time.Time `json:"revoked_at,omitempty"`
	}

	// UserTOTPRecoveryCode is the Postgres model for the user TOTP recovery code
	UserTOTPRecoveryCode struct {
		Model
		UserTOTPID uint       `json:"user_totp_id" gorm:"uniqueIndex:idx_user_totp_code;not null"`
		UserTOTP   UserTOTP   `gorm:"foreignKey:UserTOTPID;constraint:OnDelete:CASCADE"`
		Code       string     `json:"code" gorm:"uniqueIndex:idx_user_totp_code" gorm:"not null"`
		RevokedAt  *time.Time `json:"revoked_at,omitempty"`
	}

	// Tag is the Postgres model for the tag
	Tag struct {
		Model
		UserID uint   `json:"user_id" gorm:"uniqueIndex:idx_user_tag;not null"`
		User   User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
		Name   string `json:"name" gorm:"uniqueIndex:idx_user_tag;not null"`
	}

	// Note is the Postgres model for the user note
	Note struct {
		Model
		UserID    uint      `json:"user_id" gorm:"not null"`
		User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
		IsPinned  *bool     `json:"is_pinned,omitempty"`
		Title     string    `json:"title" gorm:"not null"`
		Color     *string   `json:"color,omitempty"`
		CreatedAt time.Time `json:"created_at" gorm:"not null"`
	}

	// NoteTag is the join table for the many-to-many relationship between Note and NoteTag
	NoteTag struct {
		Model
		NoteID     uint      `json:"note_id" gorm:"not null"`
		Note       Note      `gorm:"foreignKey:NoteID;constraint:OnDelete:CASCADE"`
		TagID      uint      `json:"tag_id" gorm:"not null"`
		Tag        Tag       `gorm:"foreignKey:TagID;constraint:OnDelete:CASCADE"`
		AssignedAt time.Time `json:"assigned_at" gorm:"not null"`
	}

	// NoteVersion is the Postgres model for the note version
	NoteVersion struct {
		Model
		NoteID        uint      `json:"note_id" gorm:"not null"`
		Note          Note      `gorm:"foreignKey:NoteID"`
		EncryptedBody string    `json:"encrypted_body" gorm:"not null"`
		CreatedAt     time.Time `json:"created_at" gorm:"not null"`
	}
)
