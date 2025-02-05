package note

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	"net/http"
)

type (
	// service is the structure for the API V1 service for the note route group
	service struct{}
)

// UpdateUserNoteStar updates a note star for the authenticated user
func (s *service) UpdateUserNoteStar(
	r *http.Request,
	body *UpdateUserNoteStarRequest,
) int64 {
	// Check if the request body is nil
	if body == nil {
		panic(gonethttp.ErrNilRequestBody)
	}

	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Update the note star
	commandTag, err := internalpostgres.PoolService.Exec(
		&internalpostgresmodel.UpdateUserNoteStarProc,
		userID,
		body.NoteID,
		body.Star,
	)
	if err != nil {
		panic(err)
	}

	// Check if the note exists
	if commandTag.RowsAffected() == 0 {
		panic(ErrUpdateUserNoteStarNotFound)
	}
	return userID
}

// UpdateUserNoteArchive updates a note archive for the authenticated user
func (s *service) UpdateUserNoteArchive(
	r *http.Request,
	body *UpdateUserNoteArchiveRequest,
) int64 {
	// Check if the request body is nil
	if body == nil {
		panic(gonethttp.ErrNilRequestBody)
	}

	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Update the note archive
	commandTag, err := internalpostgres.PoolService.Exec(
		&internalpostgresmodel.UpdateUserNoteArchiveProc,
		userID,
		body.NoteID,
		body.Archive,
	)
	if err != nil {
		panic(err)
	}

	// Check if the note exists
	if commandTag.RowsAffected() == 0 {
		panic(ErrUpdateUserNoteArchiveNotFound)
	}
	return userID
}

// UpdateUserNotePin updates a note pin for the authenticated user
func (s *service) UpdateUserNotePin(
	r *http.Request,
	body *UpdateUserNotePinRequest,
) int64 {
	// Check if the request body is nil
	if body == nil {
		panic(gonethttp.ErrNilRequestBody)
	}

	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Update the note pin
	commandTag, err := internalpostgres.PoolService.Exec(
		&internalpostgresmodel.UpdateUserNotePinProc,
		userID,
		body.NoteID,
		body.Pin,
	)
	if err != nil {
		panic(err)
	}

	// Check if the note exists
	if commandTag.RowsAffected() == 0 {
		panic(ErrUpdateUserNotePinNotFound)
	}
	return userID
}

// UpdateUserNoteTrash updates a note trash for the authenticated user
func (s *service) UpdateUserNoteTrash(
	r *http.Request,
	body *UpdateUserNoteTrashRequest,
) int64 {
	// Check if the request body is nil
	if body == nil {
		panic(gonethttp.ErrNilRequestBody)
	}

	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Update the note trash
	commandTag, err := internalpostgres.PoolService.Exec(
		&internalpostgresmodel.UpdateUserNoteTrashProc,
		userID,
		body.NoteID,
		body.Trash,
	)
	if err != nil {
		panic(err)
	}

	// Check if the note exists
	if commandTag.RowsAffected() == 0 {
		panic(ErrUpdateUserNoteTrashNotFound)
	}
	return userID
}
