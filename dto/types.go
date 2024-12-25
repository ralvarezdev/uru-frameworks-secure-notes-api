package dto

import (
	"time"
)

// BasicRequest is the request DTO for the basic request
type BasicRequest struct {
}

// BasicResponse is the response DTO for the basic response
type BasicResponse struct {
	Message string `json:"message"`
}

// BasicWasUpdatedResponse is the response DTO for the basic was updated response
type BasicWasUpdatedResponse struct {
	Message    string `json:"message"`
	WasUpdated bool   `json:"was_updated"`
}

// BasicWasVerifiedResponse is the response DTO for the basic was verified response
type BasicWasVerifiedResponse struct {
	Message     string `json:"message"`
	WasVerified bool   `json:"was_verified"`
}

// BasicWasRevokedResponse is the response DTO for the basic was revoked response
type BasicWasRevokedResponse struct {
	Message    string `json:"message"`
	WasRevoked bool   `json:"was_revoked"`
}

// BasicWasDeletedResponse is the response DTO for the basic was deleted response
type BasicWasDeletedResponse struct {
	Message    string `json:"message"`
	WasDeleted bool   `json:"was_deleted"`
}

// NoteTag is the response DTO for the note tag
type NoteTag struct {
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NoteTagWithID is the response DTO for the note tag with ID
type NoteTagWithID struct {
	NoteTagID string `json:"note_tag_id"`
	NoteTag
}

// Note is the response DTO for the note
type Note struct {
	Title     string     `json:"title"`
	NoteTags  []string   `json:"note_tags"`
	IsPinned  *bool      `json:"is_pinned,omitempty"`
	Color     *string    `json:"color,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NoteWithID is the response DTO for the note with ID
type NoteWithID struct {
	NoteID string `json:"note_id"`
	Note
}

// NoteVersion is the response DTO for the note version
type NoteVersion struct {
	NoteID           uint       `json:"note_id"`
	EncryptedContent string     `json:"encrypted_content"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
}

// NoteVersionWithID is the response DTO for the note version with ID
type NoteVersionWithID struct {
	NoteVersionID string `json:"note_version_id"`
	NoteVersion
}

// UserRefreshToken is the response DTO for the user refresh token
type UserRefreshToken struct {
	IssuedAt    time.Time  `json:"issued_at"`
	ExpiresAt   time.Time  `json:"expires_at"`
	RevokedAt   *time.Time `json:"revoked_at,omitempty"`
	IPv4Address string     `json:"ipv4_address"`
}

// UserRefreshTokenWithID is the response DTO for the user refresh token with ID
type UserRefreshTokenWithID struct {
	UserRefreshTokenID string `json:"user_refresh_token_id"`
	UserRefreshToken
}

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

// LogInRequest is the request DTO to log in
type LogInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LogInResponse is the response DTO to log in
type LogInResponse struct {
	Message      string  `json:"message"`
	TokenSeed    *string `json:"token_seed,omitempty"`
	Is2FAEnabled bool    `json:"is_2fa_enabled"`
}

// GenerateRefreshTokenRequest is the request DTO to generate refresh token
type GenerateRefreshTokenRequest struct {
	TokenSeed string  `json:"token_seed"`
	TOTPCode  *string `json:"totp_code,omitempty"`
}

// GenerateRefreshTokenResponse is the response DTO to generate refresh token
type GenerateRefreshTokenResponse struct {
	Message      string  `json:"message"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	AccessToken  *string `json:"access_token,omitempty"`
}

// RefreshTokenResponse is the response DTO to refresh token
type RefreshTokenResponse struct {
	Message      string  `json:"message"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	AccessToken  *string `json:"access_token,omitempty"`
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

// HasTOTPResponse is the response DTO to check if user has TOTP enabled
type HasTOTPResponse struct {
	Message string `json:"message"`
	HasTOTP bool   `json:"has_totp"`
}

// GenerateTOTPKeyResponse is the response DTO to generate TOTP key
type GenerateTOTPKeyResponse struct {
	Message string `json:"message"`
	TOTPKey string `json:"totp_key"`
}

// VerifyTOTPRequest is the request DTO to verify TOTP
type VerifyTOTPRequest struct {
	TOTPKey  string `json:"totp_key"`
	TOTPCode string `json:"totp_code"`
}

// VerifyTOTPResponse is the response DTO to verify TOTP
type VerifyTOTPResponse struct {
	Message    string `json:"message"`
	IsVerified *bool  `json:"is_verified,omitempty"`
}

// RevokeTOTPRequest is the request DTO to revoke TOTP
type RevokeTOTPRequest struct {
	Password string `json:"password"`
}

// CreateNoteTagRequest is the request DTO to create a note tag
type CreateNoteTagRequest struct {
	Name string `json:"name"`
}

// UpdateNoteTagRequest is the request DTO to update a note tag
type UpdateNoteTagRequest struct {
	Name *string `json:"name,omitempty"`
}

// DeleteNoteTagRequest is the request DTO to delete a note tag
type DeleteNoteTagRequest struct {
	NoteTagID string `json:"note_tag_id"`
}

// GetNoteTagRequest is the request DTO to get a note tag
type GetNoteTagRequest struct {
	NoteTagID string `json:"note_tag_id"`
}

// GetNoteTagResponse is the response DTO to get a note tag
type GetNoteTagResponse struct {
	Message string  `json:"message"`
	NoteTag NoteTag `json:"note_tag"`
}

// ListNoteTagsResponse is the response DTO to list note tags
type ListNoteTagsResponse struct {
	Message  string          `json:"message"`
	NoteTags []NoteTagWithID `json:"note_tags"`
}

// CreateNoteRequest is the request DTO to create a note
type CreateNoteRequest struct {
	Title    string   `json:"title"`
	NoteTags []string `json:"note_tags"`
	Color    *string  `json:"color,omitempty"`
}

// UpdateNoteRequest is the request DTO to update a note
type UpdateNoteRequest struct {
	NoteID     uint     `json:"note_id"`
	IsPinned   *bool    `json:"is_pinned,omitempty"`
	Title      *string  `json:"title,omitempty"`
	NoteTagsID []string `json:"note_tags_id,omitempty"`
	Color      *string  `json:"color,omitempty"`
}

// DeleteNoteRequest is the request DTO to delete a note
type DeleteNoteRequest struct {
	NoteID uint `json:"note_id"`
}

// GetNoteRequest is the request DTO to get a note
type GetNoteRequest struct {
	NoteID uint `json:"note_id"`
}

// GetNoteResponse is the response DTO to get a note
type GetNoteResponse struct {
	Message string `json:"message"`
	Note    Note   `json:"note"`
}

// ListNotesResponse is the response DTO to list notes
type ListNotesResponse struct {
	Message string       `json:"message"`
	Notes   []NoteWithID `json:"notes"`
}

// CreateNoteVersionRequest is the request DTO to create a note version
type CreateNoteVersionRequest struct {
	NoteID           uint   `json:"note_id"`
	EncryptedContent string `json:"encrypted_content"`
}

// UpdateNoteVersionRequest is the request DTO to update a note version
type UpdateNoteVersionRequest struct {
	NoteVersionID    uint    `json:"note_version_id"`
	EncryptedContent *string `json:"encrypted_content,omitempty"`
}

// DeleteNoteVersionRequest is the request DTO to delete a note version
type DeleteNoteVersionRequest struct {
	NoteVersionID uint `json:"note_version_id"`
}

// GetNoteVersionRequest is the request DTO to get a note version
type GetNoteVersionRequest struct {
	NoteVersionID uint `json:"note_version_id"`
}

// GetNoteVersionResponse is the response DTO to get a note version
type GetNoteVersionResponse struct {
	Message     string      `json:"message"`
	NoteVersion NoteVersion `json:"note_version"`
}

// ListNoteVersionsResponse is the response DTO to list note versions
type ListNoteVersionsResponse struct {
	Message        string   `json:"message"`
	NoteVersionsID []string `json:"note_versions_id"`
}

// ListLastNoteVersionsWithContentResponse is the response DTO to list last note versions with their content
type ListLastNoteVersionsWithContentResponse struct {
	Message      string              `json:"message"`
	NoteVersions []NoteVersionWithID `json:"note_versions"`
}

// DeleteUserRequest is the request DTO to delete a user
type DeleteUserRequest struct {
	Password string `json:"password"`
}

// GetUserRefreshTokenRequest is the request DTO to get a user refresh token that has not been revoked or expired
type GetUserRefreshTokenRequest struct {
	UserRefreshTokenID string `json:"user_refresh_token_id"`
}

// GetUserRefreshTokenResponse is the response DTO to get a user refresh token that has not been revoked or expired
type GetUserRefreshTokenResponse struct {
	Message          string            `json:"message"`
	UserRefreshToken *UserRefreshToken `json:"user_refresh_token,omitempty"`
}

// ListUserRefreshTokensResponse is the response DTO to list user refresh tokens that have not been revoked or expired
type ListUserRefreshTokensResponse struct {
	Message           string                   `json:"message"`
	UserRefreshTokens []UserRefreshTokenWithID `json:"user_refresh_tokens"`
}

// RevokeUserRefreshTokenRequest is the request DTO to revoke a user refresh token
type RevokeUserRefreshTokenRequest struct {
	UserRefreshTokenID string `json:"user_refresh_token_id"`
}
