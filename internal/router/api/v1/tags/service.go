package tags

import (
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	"net/http"
)

type (
	// service is the structure for the API V1 service for the tags route group
	service struct{}
)

// ListUserTags lists the tags of the authenticated user
func (s *service) ListUserTags(r *http.Request) (int64, *ListUserTagsResponse) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Get the user tags
	var userTags []*internalpostgresmodel.TagWithID
	rows, err := internalpostgres.PoolService.Query(
		&internalpostgresmodel.ListUserTagsFn,
		userID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		var tag internalpostgresmodel.TagWithID
		if err = rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.CreatedAt,
			&tag.UpdatedAt,
		); err != nil {
			panic(err)
		}
		userTags = append(userTags, &tag)
	}

	return userID, &ListUserTagsResponse{
		Tags: userTags,
	}
}
