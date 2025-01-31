package tag

import (
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/_common"
)

// CreateTagRequest is the request DTO to create a tag
type CreateTagRequest struct {
	Name string `json:"name"`
}

// UpdateTagRequest is the request DTO to update a tag
type UpdateTagRequest struct {
	TagID int64   `json:"tag_id"`
	Name  *string `json:"name,omitempty"`
}

// DeleteTagRequest is the request DTO to delete a tag
type DeleteTagRequest struct {
	TagID int64 `json:"tag_id"`
}

// GetTagRequest is the request DTO to get a tag
type GetTagRequest struct {
	TagID int64 `json:"tag_id"`
}

// GetTagResponse is the response DTO to get a tag
type GetTagResponse struct {
	Tag internalapiv1common.Tag `json:"tag"`
}
