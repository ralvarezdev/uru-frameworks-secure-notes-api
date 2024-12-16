package model

import (
	"gorm.io/gorm"
	"time"
)

// User is the PostgreSQL model for the user
type User struct {
	gorm.Model
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	DeletedAt string    `json:"deleted_at,omitempty"`
	Birthdate time.Time `json:"birthdate,omitempty"`
	Salt      string    `json:"salt"`
}

// UserHashedPassword is the PostgreSQL model for the user hashed password
type UserHashedPassword struct {
	gorm.Model
	UserID         uint   `json:"user_id"`
	User           User   `gorm:"foreignKey:UserID"`
	HashedPassword string `json:"hashed_password"`
}

// UserEmail is the PostgreSQL model for the user email
type UserEmail struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID"`
	Email     string    `json:"email"`
	RevokedAt time.Time `json:"revoked_at,omitempty"`
}

// UserEmailVerification is the PostgreSQL model for the user email verification
type UserEmailVerification struct {
	gorm.Model
	UserEmailID uint      `json:"user_email_id"`
	UserEmail   UserEmail `gorm:"foreignKey:UserEmailID"`
	UUID        string    `json:"uuid"`
	VerifiedAt  time.Time `json:"verified_at,omitempty"`
	RevokedAt   time.Time `json:"revoked_at,omitempty"`
}

// UserResetPassword is the PostgreSQL model for the user password reset
type UserResetPassword struct {
	gorm.Model
	UserID     uint      `json:"user_id"`
	User       User      `gorm:"foreignKey:UserID"`
	ResetToken string    `json:"reset_token"`
	RevokedAt  time.Time `json:"revoked_at,omitempty"`
}

// UserPhoneNumber is the PostgreSQL model for the user phone number
type UserPhoneNumber struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	User        User      `gorm:"foreignKey:UserID"`
	PhoneNumber string    `json:"phone_number"`
	RevokedAt   time.Time `json:"revoked_at,omitempty"`
}

// UserPhoneNumberVerification is the PostgreSQL model for the user phone number verification
type UserPhoneNumberVerification struct {
	gorm.Model
	UserPhoneNumberID uint            `json:"user_phone_number_id"`
	UserPhoneNumber   UserPhoneNumber `gorm:"foreignKey:UserPhoneNumberID"`
	VerificationCode  string          `json:"verification_code"`
	VerifiedAt        time.Time       `json:"verified_at,omitempty"`
	RevokedAt         time.Time       `json:"revoked_at,omitempty"`
}

// UserRefreshToken is the PostgreSQL model for the user refresh token
type UserRefreshToken struct {
	gorm.Model
	UserID               uint              `json:"user_id"`
	User                 User              `gorm:"foreignKey:UserID"`
	ParentRefreshTokenID *uint             `json:"parent_refresh_token_id,omitempty"`
	ParentRefreshToken   *UserRefreshToken `gorm:"foreignKey:ParentRefreshTokenID"`
	IPv4Address          string            `json:"ipv4_address"`
	IssuedAt             time.Time         `json:"issued_at"`
	RevokedAt            time.Time         `json:"revoked_at,omitempty"`
}

// UserAccessToken is the PostgreSQL model for the user access token
type UserAccessToken struct {
	gorm.Model
	UserID             uint             `json:"user_id"`
	User               User             `gorm:"foreignKey:UserID"`
	UserRefreshTokenID uint             `json:"user_refresh_token_id"`
	UserRefreshToken   UserRefreshToken `gorm:"foreignKey:UserRefreshTokenID"`
	IssuedAt           time.Time        `json:"issued_at"`
	RevokedAt          time.Time        `json:"revoked_at,omitempty"`
}

// NoteTag is the PostgreSQL model for the note tag
type NoteTag struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	User   User   `gorm:"foreignKey:UserID"`
	Name   string `json:"name"`
}

// Note is the PostgreSQL model for the note
type Note struct {
	gorm.Model
	UserID   uint      `json:"user_id"`
	User     User      `gorm:"foreignKey:UserID"`
	Title    string    `json:"title"`
	Color    string    `json:"color"`
	NoteTags []NoteTag `gorm:"many2many:notes_tags;"`
}

// NoteVersion is the PostgreSQL model for the note version
type NoteVersion struct {
	gorm.Model
	NoteID        uint   `json:"note_id"`
	Note          Note   `gorm:"foreignKey:NoteID"`
	EncryptedBody string `json:"encrypted_body"`
}
