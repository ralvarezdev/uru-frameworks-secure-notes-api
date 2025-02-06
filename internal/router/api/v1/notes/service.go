package notes

import (
	"database/sql"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	"net/http"
)

type (
	// service is the structure for the API V1 service for the notes route group
	service struct{}
)

// ListUserNotes returns the notes of the user
func (s *service) ListUserNotes(r *http.Request) (
	int64,
	*ListUserNotesResponse,
) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Get the notes
	var userNotesID []sql.NullInt64
	if _, err = internalpostgres.PoolService.Exec(
		&internalpostgresmodel.ListUserNotesProc,
		userID,
		&userNotesID,
	); err != nil {
		panic(err)
	}

	// Parse the notes ID
	parsedUserNotesID := make([]int64, 0, len(userNotesID))
	for _, noteID := range userNotesID {
		if noteID.Valid {
			parsedUserNotesID = append(parsedUserNotesID, noteID.Int64)
		}
	}

	return userID, &ListUserNotesResponse{NotesID: parsedUserNotesID}
}
