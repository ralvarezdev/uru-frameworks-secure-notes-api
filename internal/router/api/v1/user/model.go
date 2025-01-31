package user

import (
	"time"
)

// UpdateProfileRequest is the request DTO to update profile
type UpdateProfileRequest struct {
	FirstName *string    `json:"first_name,omitempty"`
	LastName  *string    `json:"last_name,omitempty"`
	Birthdate *time.Time `json:"birthdate,omitempty"`
}

// GetMyProfileResponse is the response DTO to get my profile
type GetMyProfileResponse struct {
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	Username        string  `json:"username"`
	Email           string  `json:"email"`
	EmailIsVerified bool    `json:"email_is_verified"`
	Phone           *string `json:"phone,omitempty"`
	PhoneIsVerified *bool   `json:"phone_is_verified,omitempty"`
	HasTOTP         bool    `json:"has_totp"`
	NumberNotes     int     `json:"number_notes"`
}

// ChangeUsernameRequest is the request DTO to change username
type ChangeUsernameRequest struct {
	Username string `json:"username"`
}

// ChangeEmailRequest is the request DTO to change email
type ChangeEmailRequest struct {
	Email string `json:"email"`
}

// VerifyEmailRequest is the request DTO to verify email
type VerifyEmailRequest struct {
	Token string `json:"token"`
}

// ChangePhoneNumberRequest is the request DTO to change phone number
type ChangePhoneNumberRequest struct {
	PhoneNumber string `json:"phone_number"`
}

// VerifyPhoneNumberRequest is the request DTO to verify phone number
type VerifyPhoneNumberRequest struct {
	Token string `json:"token"`
}

// DeleteUserRequest is the request DTO to delete a user
type DeleteUserRequest struct {
	Password string `json:"password"`
}
