package tag

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
)

type (
	// CreateUserTagRequest is the request DTO to create a user tag
	CreateUserTagRequest struct {
		Name string `json:"name"`
	}

	// UpdateUserTagRequest is the request DTO to update a user tag
	UpdateUserTagRequest struct {
		TagID int64   `json:"tag_id"`
		Name  *string `json:"name,omitempty"`
	}

	// DeleteUserTagRequest is the request DTO to delete a user tag
	DeleteUserTagRequest struct {
		TagID int64 `json:"tag_id"`
	}

	// GetUserTagByIDRequest is the request DTO to get a user tag by tag ID
	GetUserTagByIDRequest struct {
		TagID int64 `json:"tag_id"`
	}

	// GetUserTagByIDResponseData is the response data DTO to get a user tag by tag ID
	GetUserTagByIDResponseData struct {
		Tag internalpostgresmodel.UserTag `json:"tag"`
	}

	// GetUserTagByIDResponseBody is the response body DTO to get a user tag by tag ID
	GetUserTagByIDResponseBody struct {
		gonethttpresponse.BaseJSendSuccessBody
		Data GetUserTagByIDResponseData `json:"data"`
	}
)
