package version

import (
	"database/sql"
	gonethttp "github.com/ralvarezdev/go-net/http"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
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
	body *GetUserNoteVersionByIDRequest,
) (
	int64,
	*GetUserNoteVersionByIDResponseBody,
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
	var userNoteID sql.NullInt64
	var userNoteVersionEncryptedContent sql.NullString
	var userNoteVersionCreatedAt sql.NullTime
	rows, err := internalpostgres.PoolService.Query(
		&internalpostgresmodel.GetUserNoteVersionByIDFn,
		userID,
		body.NoteVersionID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Check if the note version exists
	if !rows.Next() {
		panic(ErrGetUserNoteVersionByIDNotFound)
	}

	// Scan the row
	if err = rows.Scan(
		&userNoteID,
		&userNoteVersionEncryptedContent,
		&userNoteVersionCreatedAt,
	); err != nil {
		panic(err)
	}

	return userID, &GetUserNoteVersionByIDResponseBody{
		BaseJSendSuccessBody: *gonethttpresponse.NewBaseJSendSuccessBody(),
		Data: GetUserNoteVersionByIDResponseData{
			NoteVersion: &internalpostgresmodel.UserNoteVersion{
				NoteID:           &userNoteID.Int64,
				EncryptedContent: userNoteVersionEncryptedContent.String,
				CreatedAt:        userNoteVersionCreatedAt.Time,
			},
		},
	}
}
