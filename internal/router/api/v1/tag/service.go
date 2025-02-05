package tag

import (
	"database/sql"
	"errors"
	godatabasespgx "github.com/ralvarezdev/go-databases/sql/pgx"
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	"net/http"
)

type (
	// service is the structure for the API V1 service for the tag route group
	service struct{}
)

// CreateUserTag creates a tag for the authenticated user
func (s *service) CreateUserTag(
	r *http.Request,
	body *CreateUserTagRequest,
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

	// Create the tag
	var tagID sql.NullInt64
	if err = internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.CreateUserTagProc,
		userID,
		body.Name,
	).Scan(
		&tagID,
	); err != nil {
		isUniqueViolation, constraintName := godatabasespgx.IsUniqueViolationError(err)
		if !isUniqueViolation {
			panic(err)
		}
		if constraintName == internalpostgresmodel.UserTagsUniqueName {
			panic(ErrCreateUserTagAlreadyExists)
		}
	}
	return userID, tagID.Int64
}

// UpdateUserTag updates a tag for the authenticated user
func (s *service) UpdateUserTag(
	r *http.Request,
	body *UpdateUserTagRequest,
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

	// Update the tag
	commandTag, err := internalpostgres.PoolService.Exec(
		&internalpostgresmodel.UpdateUserTagProc,
		userID,
		body.TagID,
		body.Name,
	)
	if err != nil {
		panic(err)
	}

	// Check if the tag was updated
	if commandTag.RowsAffected() == 0 {
		panic(ErrUpdateUserTagNotFound)
	}
	return userID
}

// DeleteUserTag deletes a tag for the authenticated user
func (s *service) DeleteUserTag(
	r *http.Request,
	body *DeleteUserTagRequest,
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

	// Delete the tag
	commandTag, err := internalpostgres.PoolService.Exec(
		&internalpostgresmodel.DeleteUserTagProc,
		userID,
		body.TagID,
	)
	if err != nil {
		panic(err)
	}

	// Check if the tag was deleted
	if commandTag.RowsAffected() == 0 {
		panic(ErrDeleteUserTagNotFound)
	}
	return userID
}

// GetUserTagByTagID gets a tag for the authenticated user by tag ID
func (s *service) GetUserTagByTagID(
	r *http.Request,
	body *GetUserTagByTagIDRequest,
) (
	int64,
	*GetUserTagByTagIDResponse,
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

	// Get the tag
	var tag internalpostgresmodel.Tag
	if err = internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.GetUserTagByTagIDProc,
		userID,
		body.TagID,
	).Scan(
		&tag.Name,
		&tag.CreatedAt,
		&tag.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(ErrGetUserTagByTagIDNotFound)
		}
		panic(err)
	}
	return userID, &GetUserTagByTagIDResponse{
		Tag: tag,
	}
}
