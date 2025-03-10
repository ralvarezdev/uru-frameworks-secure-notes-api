package tag

import (
	"database/sql"
	godatabasespgx "github.com/ralvarezdev/go-databases/sql/pgx"
	gonethttp "github.com/ralvarezdev/go-net/http"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
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
		if constraintName == internalpostgresmodel.UserTagsUniqueUserIDName {
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

// GetUserTagByID gets a tag for the authenticated user by tag ID
func (s *service) GetUserTagByID(
	r *http.Request,
	body *GetUserTagByIDRequest,
) (
	int64,
	*GetUserTagByIDResponseBody,
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
	var userTagName sql.NullString
	var userTagCreatedAt, userTagUpdatedAt sql.NullTime
	rows, err := internalpostgres.PoolService.Query(
		&internalpostgresmodel.GetUserTagByIDFn,
		userID,
		body.TagID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Check if the tag ID exists
	if !rows.Next() {
		panic(ErrGetUserTagByIDNotFound)
	}

	// Scan the row
	if err = rows.Scan(
		&userTagName,
		&userTagCreatedAt,
		&userTagUpdatedAt,
	); err != nil {
		panic(err)
	}

	return userID, &GetUserTagByIDResponseBody{
		BaseJSendSuccessBody: *gonethttpresponse.NewBaseJSendSuccessBody(),
		Data: GetUserTagByIDResponseData{
			Tag: internalpostgresmodel.UserTag{
				Name:      userTagName.String,
				CreatedAt: userTagCreatedAt.Time,
				UpdatedAt: userTagUpdatedAt.Time,
			},
		},
	}
}
