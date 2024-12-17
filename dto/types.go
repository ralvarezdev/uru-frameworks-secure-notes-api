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

// NoteTag is the response DTO for the note tag
type NoteTag struct {
	NoteTagID string     `json:"note_tag_id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Note is the response DTO for the note
type Note struct {
	NoteID    string     `json:"note_id"`
	Title     string     `json:"title"`
	NoteTags  []string   `json:"note_tags"`
	Color     *string    `json:"color,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NoteVersion is the response DTO for the note version
type NoteVersion struct {
	NoteVersionID    string     `json:"note_version_id"`
	EncryptedContent string     `json:"encrypted_content"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
}

// SignUpRequest is the request DTO for the sign up
type SignUpRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

// LogInRequest is the request DTO for the log in
type LogInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LogInResponse is the response DTO for the log in
type LogInResponse struct {
	Message      string  `json:"message"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	AccessToken  *string `json:"access_token,omitempty"`
}

// RefreshTokenResponse is the response DTO for the refresh token
type RefreshTokenResponse struct {
	Message      string  `json:"message"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	AccessToken  *string `json:"access_token,omitempty"`
}

// ChangeUsernameRequest is the request DTO for the change username
type ChangeUsernameRequest struct {
	Username string `json:"username"`
}

// ChangePasswordRequest is the request DTO for the change password
type ChangePasswordRequest struct {
	Password string `json:"password"`
}

// ForgotPasswordRequest is the request DTO for the forgot password
type ForgotPasswordRequest struct {
	Username string `json:"username"`
}

// ResetPasswordRequest is the request DTO for the reset password
type ResetPasswordRequest struct {
	ResetToken string `json:"reset_token"`
	Password   string `json:"password"`
}

// ChangeEmailRequest is the request DTO for the change email
type ChangeEmailRequest struct {
	Email string `json:"email"`
}

// SendEmailVerificationTokenRequest is the request DTO for the send email verification token
type SendEmailVerificationTokenRequest struct {
	Email string `json:"email"`
}

// VerifyEmailRequest is the request DTO for the verify email
type VerifyEmailRequest struct {
	Email             string `json:"email"`
	VerificationToken string `json:"verification_token"`
}

// ChangePhoneNumberRequest is the request DTO for the change phone number
type ChangePhoneNumberRequest struct {
	PhoneNumber string `json:"phone_number"`
}

// SendPhoneNumberVerificationCodeRequest is the request DTO for the send phone number verification code
type SendPhoneNumberVerificationCodeRequest struct {
	PhoneNumber string `json:"phone_number"`
}

// VerifyPhoneNumberRequest is the request DTO for the verify phone number
type VerifyPhoneNumberRequest struct {
	PhoneNumber      string `json:"phone_number"`
	VerificationCode string `json:"verification_code"`
}

// CreateNoteTagRequest is the request DTO for the create note tag
type CreateNoteTagRequest struct {
	Name string `json:"name"`
}

// UpdateNoteTagRequest is the request DTO for the update note tag
type UpdateNoteTagRequest struct {
	Name *string `json:"name,omitempty"`
}

// DeleteNoteTagRequest is the request DTO for the delete note tag
type DeleteNoteTagRequest struct {
	NoteTagID string `json:"note_tag_id"`
}

// GetNoteTagRequest is the request DTO for the get note tag
type GetNoteTagRequest struct {
	NoteTagID string `json:"note_tag_id"`
}

// GetNoteTagResponse is the response DTO for the get note tag
type GetNoteTagResponse struct {
	Message string  `json:"message"`
	Name    *string `json:"name,omitempty"`
}

// ListNoteTagsResponse is the response DTO for the list note tags
type ListNoteTagsResponse struct {
	Message  string    `json:"message"`
	NoteTags []NoteTag `json:"note_tags"`
}

// CreateNoteRequest is the request DTO for the create note
type CreateNoteRequest struct {
	Title    string   `json:"title"`
	NoteTags []string `json:"note_tags"`
	Color    *string  `json:"color,omitempty"`
}

// UpdateNoteRequest is the request DTO for the update note
type UpdateNoteRequest struct {
	NoteID     uint     `json:"note_id"`
	Title      *string  `json:"title,omitempty"`
	NoteTagsID []string `json:"note_tags_id,omitempty"`
	Color      *string  `json:"color,omitempty"`
}

// DeleteNoteRequest is the request DTO for the delete note
type DeleteNoteRequest struct {
	NoteID uint `json:"note_id"`
}

// GetNoteRequest is the request DTO for the get note
type GetNoteRequest struct {
	NoteID uint `json:"note_id"`
}

// GetNoteResponse is the response DTO for the get note
type GetNoteResponse struct {
	Message    string   `json:"message"`
	Title      string   `json:"title"`
	NoteTagsID []string `json:"note_tags_id"`
	Color      *string  `json:"color,omitempty"`
}

// ListNotesResponse is the response DTO for the list notes
type ListNotesResponse struct {
	Message string `json:"message"`
	Notes   []Note `json:"notes"`
}

// CreateNoteVersionRequest is the request DTO for the create note version
type CreateNoteVersionRequest struct {
	NoteID           uint   `json:"note_id"`
	EncryptedContent string `json:"encrypted_content"`
}

// UpdateNoteVersionRequest is the request DTO for the update note version
type UpdateNoteVersionRequest struct {
	NoteVersionID    uint    `json:"note_version_id"`
	EncryptedContent *string `json:"encrypted_content,omitempty"`
}

// DeleteNoteVersionRequest is the request DTO for the delete note version
type DeleteNoteVersionRequest struct {
	NoteVersionID uint `json:"note_version_id"`
}

// GetNoteVersionRequest is the request DTO for the get note version
type GetNoteVersionRequest struct {
	NoteVersionID uint `json:"note_version_id"`
}

// GetNoteVersionResponse is the response DTO for the get note version
type GetNoteVersionResponse struct {
	Message          string `json:"message"`
	EncryptedContent string `json:"encrypted_content"`
}

// ListLast5NoteVersionsResponse is the response DTO for the list last 5 note versions
type ListLast5NoteVersionsResponse struct {
	Message      string        `json:"message"`
	NoteVersions []NoteVersion `json:"note_versions"`
}
