package user

// SignUpRequest is the request DTO to sign up
type SignUpRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

// GetMyProfileResponse is the response DTO to get my profile
type GetMyProfileResponse struct {
	Message         string  `json:"message"`
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

// ChangePasswordRequest is the request DTO to change password
type ChangePasswordRequest struct {
	Password string `json:"password"`
}

// ForgotPasswordRequest is the request DTO to forgot password
type ForgotPasswordRequest struct {
	Username string `json:"username"`
}

// ResetPasswordRequest is the request DTO to reset password
type ResetPasswordRequest struct {
	ResetToken string `json:"reset_token"`
	Password   string `json:"password"`
}

// ChangeEmailRequest is the request DTO to change email
type ChangeEmailRequest struct {
	Email string `json:"email"`
}

// SendEmailVerificationTokenRequest is the request DTO to send email verification token
type SendEmailVerificationTokenRequest struct {
	Email string `json:"email"`
}

// VerifyEmailRequest is the request DTO to verify email
type VerifyEmailRequest struct {
	Email             string `json:"email"`
	VerificationToken string `json:"verification_token"`
}

// ChangePhoneNumberRequest is the request DTO to change phone number
type ChangePhoneNumberRequest struct {
	PhoneNumber string `json:"phone_number"`
}

// SendPhoneNumberVerificationCodeRequest is the request DTO to send phone number verification code
type SendPhoneNumberVerificationCodeRequest struct {
	PhoneNumber string `json:"phone_number"`
}

// VerifyPhoneNumberRequest is the request DTO to verify phone number
type VerifyPhoneNumberRequest struct {
	PhoneNumber      string `json:"phone_number"`
	VerificationCode string `json:"verification_code"`
}

// DeleteUserRequest is the request DTO to delete a user
type DeleteUserRequest struct {
	Password string `json:"password"`
}
