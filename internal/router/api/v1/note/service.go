package note

import (
	"database/sql"
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

	// Check if the note ID exists
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

	// Check if the note ID exists
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

	// Check if the note ID exists
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

	// Check if the note ID exists
	if commandTag.RowsAffected() == 0 {
		panic(ErrUpdateUserNoteTrashNotFound)
	}
	return userID
}

// GetUserNoteByID gets a note by ID for the authenticated user
func (s *service) GetUserNoteByID(
	r *http.Request,
	body *GetUserNoteByIDRequest,
) (
	int64,
	*GetUserNoteByIDResponse,
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

	// Get the note
	var userNoteLatestNoteVersionID sql.NullInt64
	var userNoteTitle, userNoteColor sql.NullString
	var userNoteCreatedAt, userNoteUpdatedAt, userNotePinnedAt, userNoteArchivedAt, userNoteTrashedAt, userNoteStarredAt sql.NullTime
	rows, err := internalpostgres.PoolService.Query(
		&internalpostgresmodel.GetUserNoteByIDFn,
		userID,
		body.NoteID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Check if the note ID exists
	if !rows.Next() {
		panic(ErrGetUserNoteByIDNotFound)
	}

	// Scan the row
	if err = rows.Scan(
		&userNoteTitle,
		&userNoteColor,
		&userNoteCreatedAt,
		&userNoteUpdatedAt,
		&userNotePinnedAt,
		&userNoteArchivedAt,
		&userNoteTrashedAt,
		&userNoteStarredAt,
		&userNoteLatestNoteVersionID,
	); err != nil {
		panic(err)
	}

	return userID, &GetUserNoteByIDResponse{
		Note: internalpostgresmodel.UserNote{
			Title:               userNoteTitle.String,
			Color:               &userNoteColor.String,
			CreatedAt:           userNoteCreatedAt.Time,
			UpdatedAt:           &userNoteUpdatedAt.Time,
			PinnedAt:            &userNotePinnedAt.Time,
			ArchivedAt:          &userNoteArchivedAt.Time,
			TrashedAt:           &userNoteTrashedAt.Time,
			StarredAt:           &userNoteStarredAt.Time,
			LatestNoteVersionID: &userNoteLatestNoteVersionID.Int64,
		},
	}
}

// CreateUserNote creates a note for the authenticated user
func (s *service) CreateUserNote(
	r *http.Request,
	body *CreateUserNoteRequest,
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

	// Create the note
	var userNoteID sql.NullInt64
	if err = internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.CreateUserNoteProc,
		userID,
		body.Title,
		body.Color,
		body.Pinned,
		body.Archived,
		body.Trashed,
		body.Starred,
		body.EncryptedContent,
		body.NoteTagsID,
		nil,
	).Scan(&userNoteID); err != nil {
		panic(err)
	}
	return userID, userNoteID.Int64
}

// UpdateUserNote updates a note for the authenticated user
func (s *service) UpdateUserNote(
	r *http.Request,
	body *UpdateUserNoteRequest,
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

	// Update the note
	commandTag, err := internalpostgres.PoolService.Exec(
		&internalpostgresmodel.UpdateUserNoteProc,
		userID,
		body.NoteID,
		body.Title,
		body.Color,
	)
	if err != nil {
		panic(err)
	}

	// Check if the note ID exists
	if commandTag.RowsAffected() == 0 {
		panic(ErrUpdateUserNoteNotFound)
	}
	return userID
}

// DeleteUserNote deletes a note for the authenticated user
func (s *service) DeleteUserNote(
	r *http.Request,
	body *DeleteUserNoteRequest,
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

	// Delete the note
	commandTag, err := internalpostgres.PoolService.Exec(
		&internalpostgresmodel.DeleteUserNoteProc,
		userID,
		body.NoteID,
	)
	if err != nil {
		panic(err)
	}

	// Check if the note ID exists
	if commandTag.RowsAffected() == 0 {
		panic(ErrDeleteUserNoteNotFound)
	}
	return userID
}
