package user

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	"time"
)

type (
	// UpdateProfileRequest is the request DTO to update profile
	UpdateProfileRequest struct {
		FirstName *string    `json:"first_name,omitempty"`
		LastName  *string    `json:"last_name,omitempty"`
		Birthdate *time.Time `json:"birthdate,omitempty"`
	}

	// GetMyProfileResponseData is the response data DTO to get my profile
	GetMyProfileResponseData struct {
		FirstName       string     `json:"first_name"`
		LastName        string     `json:"last_name"`
		Birthdate       *time.Time `json:"birthdate,omitempty"`
		Username        string     `json:"username"`
		Email           string     `json:"email"`
		EmailIsVerified bool       `json:"email_is_verified"`
		Phone           *string    `json:"phone,omitempty"`
		PhoneIsVerified *bool      `json:"phone_is_verified,omitempty"`
		HasTOTPEnabled  bool       `json:"has_totp_enabled"`
	}

	// GetMyProfileResponseBody is the response body DTO to get my profile
	GetMyProfileResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data GetMyProfileResponseData `json:"data"`
	}

	// ChangeUsernameRequest is the request DTO to change username
	ChangeUsernameRequest struct {
		Username string `json:"username"`
	}

	// DeleteUserRequest is the request DTO to delete a user
	DeleteUserRequest struct {
		Password string `json:"password"`
	}
)
