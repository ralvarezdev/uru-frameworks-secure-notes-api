package version

import (
	"database/sql"
	"errors"
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	"net/http"
)

type (
	// service is the structure for the API V1 service for the note version route group
	service struct{}
)

// CreateUserNoteVersion creates a note version for the authenticated user
func (s *service) CreateUserNoteVersion(
	r *http.Request,
	body *CreateUserNoteVersionRequest,
) (int64, int64) {
	// Check if the request body is nil
	if body == nil {
		panic(gonethttp.ErrNilRequestBody)
	}

	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Create the note version
	var noteIDIsValid sql.NullBool
	var noteVersionID sql.NullInt64
	if err = internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.CreateUserNoteVersionProc,
		userID,
		body.NoteID,
		body.EncryptedContent,
	).Scan(
		&noteIDIsValid,
		&noteVersionID,
	); err != nil {
		panic(err)
	}

	// Check if the note ID is valid
	if !noteIDIsValid.Bool {
		panic(ErrCreateUserNoteVersionNoteIDIsNotValid)
	}
	return userID, noteVersionID.Int64
}

// DeleteUserNoteVersion deletes a note version for the authenticated user
func (s *service) DeleteUserNoteVersion(
	r *http.Request,
	body *DeleteUserNoteVersionRequest,
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

	// Delete the note version
	commandTag, err := internalpostgres.PoolService.Exec(
		&internalpostgresmodel.DeleteUserNoteVersionProc,
		userID,
		body.NoteVersionID,
	)
	if err != nil {
		panic(err)
	}

	// Check if the note version was deleted
	if commandTag.RowsAffected() == 0 {
		panic(ErrDeleteUserNoteVersionNotFound)
	}
	return userID
}

// GetUserNoteVersionByNoteVersionID gets a note version by note version ID
func (s *service) GetUserNoteVersionByNoteVersionID(
	r *http.Request,
	body *GetUserNoteVersionByNoteVersionIDRequest,
) (
	int64,
	*GetUserNoteVersionByNoteVersionIDResponse,
) {
	// Check if the request body is nil
	if body == nil {
		panic(gonethttp.ErrNilRequestBody)
	}

	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Get the note version
	var noteVersion internalpostgresmodel.UserNoteVersion
	if err = internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.GetUserNoteVersionByNoteVersionIDProc,
		userID,
		body.NoteVersionID,
	).Scan(
		&noteVersion.NoteID,
		&noteVersion.EncryptedContent,
		&noteVersion.CreatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(ErrGetUserNoteVersionByNoteVersionIDNotFound)
		}
		panic(err)
	}
	return userID, &GetUserNoteVersionByNoteVersionIDResponse{
		NoteVersion: &noteVersion,
	}
}
