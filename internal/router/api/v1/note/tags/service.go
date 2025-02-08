package tags

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	"net/http"
)

type (
	// service is the structure for the API V1 service for the note tags route group
	service struct{}
)

// ListUserNoteTags lists the note tags for the authenticated user
func (s *service) ListUserNoteTags(
	r *http.Request,
	body *ListUserNoteTagsRequest,
) (int64, *ListUserNoteTagsResponseBody) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// List the user note tags
	rows, err := internalpostgres.PoolService.Query(
		&internalpostgresmodel.ListUserNoteTagsFn,
		userID,
		body.NoteID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate over the rows
	var userNoteTags []*internalpostgresmodel.UserNoteTagWithID
	for rows.Next() {
		var userNoteTag internalpostgresmodel.UserNoteTagWithID
		if err = rows.Scan(
			&userNoteTag.TagID,
			&userNoteTag.AssignedAt,
		); err != nil {
			panic(err)
		}
		userNoteTags = append(userNoteTags, &userNoteTag)
	}

	// Check if the note ID exists
	if len(userNoteTags) == 0 {
		panic(ErrListUserNoteTagsNotFound)
	}

	return userID, &ListUserNoteTagsResponseBody{
		BaseJSendSuccessBody: *gonethttpresponse.NewBaseJSendSuccessBody(),
		Data: ListUserNoteTagsResponseData{
			NoteTags: userNoteTags,
		},
	}
}

// AddUserNoteTags adds note tags for the authenticated user
func (s *service) AddUserNoteTags(
	r *http.Request,
	body *AddUserNoteTagsRequest,
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

	// Add the note tags
	if _, err = internalpostgres.PoolService.Exec(
		&internalpostgresmodel.AddUserNoteTagsProc,
		userID,
		body.NoteID,
		body.TagsID,
	); err != nil {
		panic(err)
	}

	return userID
}

// RemoveUserNoteTags removes note tags for the authenticated user
func (s *service) RemoveUserNoteTags(
	r *http.Request,
	body *RemoveUserNoteTagsRequest,
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

	// Remove the note tags
	if _, err = internalpostgres.PoolService.Exec(
		&internalpostgresmodel.RemoveUserNoteTagsProc,
		userID,
		body.NoteID,
		body.TagsID,
	); err != nil {
		panic(err)
	}

	return userID
}
