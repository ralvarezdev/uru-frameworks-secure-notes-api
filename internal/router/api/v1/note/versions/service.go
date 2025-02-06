package versions

import (
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	"net/http"
)

type (
	// service is the structure for the API V1 service for the note versions route group
	service struct{}
)

// ListUserNoteVersions lists the note versions for the authenticated user
func (s *service) ListUserNoteVersions(
	r *http.Request,
	body *ListUserNoteVersionsRequest,
) (int64, *ListUserNoteVersionsResponse) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// List the note versions
	rows, err := internalpostgres.PoolService.Query(
		&internalpostgresmodel.ListUserNoteVersionsFn,
		userID,
		body.NoteID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate through the note versions
	userNoteVersionsID := make([]int64, 0)
	for rows.Next() {
		var noteVersionID int64
		if err = rows.Scan(
			&noteVersionID,
		); err != nil {
			panic(err)
		}
		userNoteVersionsID = append(userNoteVersionsID, noteVersionID)
	}

	// Check if the note ID exists
	if len(userNoteVersionsID) == 0 {
		panic(ErrListUserNoteVersionsNotFound)
	}

	return userID, &ListUserNoteVersionsResponse{
		NoteVersionsID: userNoteVersionsID,
	}
}
