package versions

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
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
	noteVersionsID := make([]int64, 0)
	for rows.Next() {
		var noteVersionID int64
		if err = rows.Scan(
			&noteVersionID,
		); err != nil {
			panic(err)
		}
		noteVersionsID = append(noteVersionsID, noteVersionID)
	}

	// Check if the note ID exists
	if len(noteVersionsID) == 0 {
		panic(ErrListUserNoteVersionsNotFound)
	}

	return userID, &ListUserNoteVersionsResponse{
		NoteVersionsID: noteVersionsID,
	}
}

// SyncUserNoteVersions syncs the note versions for the authenticated user
func (s *service) SyncUserNoteVersions(
	r *http.Request,
	body *SyncUserNoteVersionsRequest,
) (int64, *SyncUserNoteVersionsResponse) {
	// Check if the request body is nil
	if body == nil {
		panic(gonethttp.ErrNilRequestBody)
	}

	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Check if the latest note version ID is nil
	if body.LatestNoteVersionID == nil {
		// Set the latest note version ID to 0
		latestNoteVersionID := int64(0)
		body.LatestNoteVersionID = &latestNoteVersionID
	}

	// Sync the note versions
	var userNoteVersions []*internalpostgresmodel.UserNoteVersionWithID
	rows, err := internalpostgres.PoolService.Query(
		&internalpostgresmodel.SyncUserNoteVersionsFn,
		userID,
		body.NoteID,
		body.LatestNoteVersionID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate through the note versions
	for rows.Next() {
		var userNoteVersion internalpostgresmodel.UserNoteVersionWithID
		if err = rows.Scan(
			&userNoteVersion.ID,
			&userNoteVersion.EncryptedContent,
			&userNoteVersion.CreatedAt,
		); err != nil {
			panic(err)
		}
		userNoteVersions = append(userNoteVersions, &userNoteVersion)
	}

	// Check if the note versions exist
	if len(userNoteVersions) == 0 {
		panic(ErrSyncUserNoteVersionsNotFound)
	}
	return userID, &SyncUserNoteVersionsResponse{
		NoteVersions: userNoteVersions,
	}
}
