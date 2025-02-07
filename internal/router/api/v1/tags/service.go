package tags

import (
	internalcookie "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/cookie"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	"net/http"
	"time"
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
	var userTags []*internalpostgresmodel.UserTagWithID
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
		var tag internalpostgresmodel.UserTagWithID
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

// SyncUserTagsByLastSyncedAt syncs the tags of the authenticated user by the last synced at timestamp
func (s *service) SyncUserTagsByLastSyncedAt(
	w http.ResponseWriter,
	r *http.Request,
) (
	int64,
	int64,
	*time.Time,
	*SyncUserTagsResponse,
) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Get the user refresh token ID from the request
	userRefreshTokenID, err := internaljwtclaims.GetParentRefreshTokenID(r)
	if err != nil {
		panic(err)
	}

	// Get the last synced at from the cookies
	lastSyncedAt, err := internalcookie.GetSyncTagsCookie(r)
	if err != nil {
		panic(err)
	}

	// Get the user tags
	newLastSyncedAt := time.Now()
	var userTags []*internalpostgresmodel.UserTagWithID
	rows, err := internalpostgres.PoolService.Query(
		&internalpostgresmodel.SyncUserTagsByLastSyncedAtFn,
		userID,
		lastSyncedAt,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		var userTag internalpostgresmodel.UserTagWithID
		if err = rows.Scan(
			&userTag.ID,
			&userTag.Name,
			&userTag.CreatedAt,
			&userTag.UpdatedAt,
			&userTag.DeletedAt,
		); err != nil {
			panic(err)
		}
		userTags = append(userTags, &userTag)
	}

	// Set the last sync at cookie
	internalcookie.SetSyncTagsCookie(w, newLastSyncedAt)

	return userID, userRefreshTokenID, lastSyncedAt, &SyncUserTagsResponse{
		SyncTags: userTags,
	}
}
