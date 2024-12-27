package postgresql

import (
	"gorm.io/gorm"
	"time"
)

// User is the PostgreSQL model for the user
type User struct {
	gorm.Model
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Salt      string     `json:"salt"`
	DeletedAt *string    `json:"deleted_at,omitempty"`
	Birthdate *time.Time `json:"birthdate,omitempty"`
}

// UserUsername is the PostgreSQL model for the user username
type UserUsername struct {
	gorm.Model
	UserID    uint       `json:"user_id"`
	User      User       `gorm:"foreignKey:UserID"`
	Username  string     `json:"username" gorm:"uniqueIndex:idx_user_username,where:revoked_at IS NULL"`
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}

// UserHashedPassword is the PostgreSQL model for the user hashed password
type UserHashedPassword struct {
	gorm.Model
	UserID         uint       `json:"user_id"`
	User           User       `gorm:"foreignKey:UserID"`
	HashedPassword string     `json:"hashed_password"`
	RevokedAt      *time.Time `json:"revoked_at,omitempty"`
}

// UserResetPassword is the PostgreSQL model for the user password reset
type UserResetPassword struct {
	gorm.Model
	UserID     uint       `json:"user_id"`
	User       User       `gorm:"foreignKey:UserID"`
	ResetToken string     `json:"reset_token" gorm:"uniqueIndex"`
	RevokedAt  *time.Time `json:"revoked_at,omitempty"`
	ExpiresAt  time.Time  `json:"expires_at"`
}

// UserEmail is the PostgreSQL model for the user email
type UserEmail struct {
	gorm.Model
	UserID         uint                   `json:"user_id" gorm:"uniqueIndex:idx_user_email,where:revoked_at IS NULL"`
	User           User                   `gorm:"foreignKey:UserID"`
	VerificationID *uint                  `json:"verification_id,omitempty" gorm:"uniqueIndex:idx_verification_id,where:verification_id IS NOT NULL"`
	Verification   *UserEmailVerification `gorm:"foreignKey:VerificationID"`
	Email          string                 `json:"email" gorm:"uniqueIndex:idx_user_email,where:revoked_at IS NULL"`
	RevokedAt      *time.Time             `json:"revoked_at,omitempty"`
}

// UserEmailVerification is the PostgreSQL model for the user email verification
type UserEmailVerification struct {
	gorm.Model
	UserEmailID       uint       `json:"user_email_id"`
	UserEmail         UserEmail  `gorm:"foreignKey:UserEmailID"`
	VerificationToken string     `json:"verification_token"`
	ExpiresAt         time.Time  `json:"expires_at"`
	VerifiedAt        *time.Time `json:"verified_at,omitempty"`
	RevokedAt         *time.Time `json:"revoked_at,omitempty"`
}

// UserPhoneNumber is the PostgreSQL model for the user phone number
type UserPhoneNumber struct {
	gorm.Model
	UserID         uint                         `json:"user_id" gorm:"uniqueIndex:idx_user_phone_number,where:revoked_at IS NULL"`
	User           User                         `gorm:"foreignKey:UserID"`
	VerificationID *uint                        `json:"verification_id,omitempty" gorm:"uniqueIndex:idx_verification_id,where:verification_id IS NOT NULL"`
	Verification   *UserPhoneNumberVerification `gorm:"foreignKey:VerificationID"`
	PhoneNumber    string                       `json:"phone_number" gorm:"uniqueIndex:idx_user_phone_number,where:revoked_at IS NULL"`
	RevokedAt      *time.Time                   `json:"revoked_at,omitempty"`
}

// UserPhoneNumberVerification is the PostgreSQL model for the user phone number verification
type UserPhoneNumberVerification struct {
	gorm.Model
	UserPhoneNumberID uint            `json:"user_phone_number_id"`
	UserPhoneNumber   UserPhoneNumber `gorm:"foreignKey:UserPhoneNumberID"`
	VerificationCode  string          `json:"verification_code"`
	ExpiresAt         time.Time       `json:"expires_at"`
	VerifiedAt        *time.Time      `json:"verified_at,omitempty"`
	RevokedAt         *time.Time      `json:"revoked_at,omitempty"`
}

// UserFailedLogInAttempt is the PostgreSQL model for the user failed login attempt
type UserFailedLogInAttempt struct {
	gorm.Model
	UserID          uint           `json:"user_id"`
	User            User           `gorm:"foreignKey:UserID"`
	UserTokenSeedID *uint          `json:"user_token_seed_id,omitempty"`
	UserTokenSeed   *UserTokenSeed `gorm:"foreignKey:UserTokenSeedID"`
	IPv4Address     string         `json:"ipv4_address"`
	BadPassword     *bool          `json:"bad_password,omitempty"`
	Bad2FACode      *bool          `json:"bad_2fa_code,omitempty"`
	AttemptedAt     time.Time      `json:"attempted_at"`
}

// UserTokenSeed is the PostgreSQL model for the user token seed
type UserTokenSeed struct {
	gorm.Model
	UserID    uint       `json:"user_id"`
	User      User       `gorm:"foreignKey:UserID"`
	TokenSeed string     `json:"token_seed" gorm:"uniqueIndex"`
	ExpiresAt time.Time  `json:"expires_at"`
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}

// UserRefreshToken is the PostgreSQL model for the user refresh token
type UserRefreshToken struct {
	gorm.Model
	UserID               uint              `json:"user_id"`
	User                 User              `gorm:"foreignKey:UserID"`
	ParentRefreshTokenID *uint             `json:"parent_refresh_token_id,omitempty" gorm:"uniqueIndex:idx_parent_refresh_token,where:parent_refresh_token_id IS NOT NULL"`
	ParentRefreshToken   *UserRefreshToken `gorm:"foreignKey:ParentRefreshTokenID"`
	UserTokenSeedID      *uint             `json:"user_token_seed_id,omitempty" gorm:"uniqueIndex:idx_user_token_seed,where:user_token_seed_id IS NOT NULL"`
	UserTokenSeed        *UserTokenSeed    `gorm:"foreignKey:UserTokenSeedID"`
	IPv4Address          string            `json:"ipv4_address"`
	ExpiresAt            time.Time         `json:"expires_at"`
	RevokedAt            *time.Time        `json:"revoked_at,omitempty"`
}

// UserAccessToken is the PostgreSQL model for the user access token
type UserAccessToken struct {
	gorm.Model
	UserID             uint             `json:"user_id"`
	User               User             `gorm:"foreignKey:UserID"`
	UserRefreshTokenID uint             `json:"user_refresh_token_id" gorm:"uniqueIndex"`
	UserRefreshToken   UserRefreshToken `gorm:"foreignKey:UserRefreshTokenID"`
	ExpiresAt          time.Time        `json:"expires_at"`
	RevokedAt          *time.Time       `json:"revoked_at,omitempty"`
}

// UserTOTP is the PostgreSQL model for the user TOTP
type UserTOTP struct {
	gorm.Model
	UserID     uint       `json:"user_id"`
	User       User       `gorm:"foreignKey:UserID"`
	Secret     string     `json:"secret" gorm:"uniqueIndex"`
	ExpiresAt  time.Time  `json:"expires_at"`
	VerifiedAt *time.Time `json:"verified_at,omitempty"`
	RevokedAt  *time.Time `json:"revoked_at,omitempty"`
}

// UserTOTPRecoveryCode is the PostgreSQL model for the user TOTP recovery code
type UserTOTPRecoveryCode struct {
	gorm.Model
	UserTOTPID uint       `json:"user_totp_id" gorm:"uniqueIndex:idx_user_totp_code"`
	UserTOTP   UserTOTP   `gorm:"foreignKey:UserTOTPID"`
	Code       string     `json:"code" gorm:"uniqueIndex:idx_user_totp_code"`
	RevokedAt  *time.Time `json:"revoked_at,omitempty"`
}

// NoteTag is the PostgreSQL model for the note tag
type NoteTag struct {
	gorm.Model
	UserID uint   `json:"user_id" gorm:"uniqueIndex:idx_user_note_tag"`
	User   User   `gorm:"foreignKey:UserID"`
	Name   string `json:"name" gorm:"uniqueIndex:idx_user_note_tag"`
}

// Note is the PostgreSQL model for the user note
type Note struct {
	gorm.Model
	UserID   uint      `json:"user_id"`
	User     User      `gorm:"foreignKey:UserID"`
	IsPinned *bool     `json:"is_pinned,omitempty"`
	Title    string    `json:"title"`
	Color    *string   `json:"color,omitempty"`
	NoteTags []NoteTag `gorm:"many2many:notes_tags;"`
}

// NoteVersion is the PostgreSQL model for the note version
type NoteVersion struct {
	gorm.Model
	NoteID        uint   `json:"note_id"`
	Note          Note   `gorm:"foreignKey:NoteID"`
	EncryptedBody string `json:"encrypted_body"`
}
