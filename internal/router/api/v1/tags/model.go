package tags

import (
	interalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
)

type (
	// ListUserTagsResponse is the response DTO to list tags
	ListUserTagsResponse struct {
		Tags []*interalpostgresmodel.TagWithID `json:"tags"`
	}
)
