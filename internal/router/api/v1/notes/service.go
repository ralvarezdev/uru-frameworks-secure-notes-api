package notes

import (
	"database/sql"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalcookie "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/cookie"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	"net/http"
	"time"
)

type (
	// service is the structure for the API V1 service for the notes route group
	service struct{}
)

// ListUserNotes returns the notes of the user
func (s *service) ListUserNotes(r *http.Request) (
	int64,
	*ListUserNotesResponseBody,
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

	return userID, &ListUserNotesResponseBody{
		BaseJSendSuccessBody: *gonethttpresponse.NewBaseJSendSuccessBody(),
		Data: ListUserNotesResponseData{
			NotesID: parsedUserNotesID,
		},
	}
}

// SyncUserNoteVersionsByLastSyncedAt returns the note versions of the user by the last synced at timestamp
func (s *service) SyncUserNoteVersionsByLastSyncedAt(
	userID int64,
	lastSyncedAt *time.Time,
	userNoteID int64,
) *[]*internalpostgresmodel.UserNoteVersionWithID {
	// Get the user note versions
	var userNoteVersions []*internalpostgresmodel.UserNoteVersionWithID
	rows, err := internalpostgres.PoolService.Query(
		&internalpostgresmodel.SyncUserNoteVersionsByLastSyncedAtFn,
		userID,
		lastSyncedAt,
		userNoteID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate over the user note versions rows
	for rows.Next() {
		var userNoteVersion internalpostgresmodel.UserNoteVersionWithID
		if err = rows.Scan(
			&userNoteVersion.ID,
			&userNoteVersion.EncryptedContent,
			&userNoteVersion.CreatedAt,
			&userNoteVersion.DeletedAt,
		); err != nil {
			panic(err)
		}
		userNoteVersions = append(userNoteVersions, &userNoteVersion)
	}

	return &userNoteVersions
}

// SyncUserNoteTagsByLastSyncedAt returns the note tags of the user by the last synced at timestamp
func (s *service) SyncUserNoteTagsByLastSyncedAt(
	userID int64,
	lastSyncedAt *time.Time,
	userNoteID int64,
) *[]*internalpostgresmodel.UserNoteTag {
	// Get the user note tags
	var userNoteTags []*internalpostgresmodel.UserNoteTag
	rows, err := internalpostgres.PoolService.Query(
		&internalpostgresmodel.SyncUserNoteTagsByLastSyncedAtFn,
		userID,
		lastSyncedAt,
		userNoteID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate over the user note tags rows
	for rows.Next() {
		var userNoteTag internalpostgresmodel.UserNoteTag
		if err = rows.Scan(
			&userNoteTag.TagID,
			&userNoteTag.AssignedAt,
			&userNoteTag.DeletedAt,
		); err != nil {
			panic(err)
		}
		userNoteTags = append(userNoteTags, &userNoteTag)
	}

	return &userNoteTags
}

// SyncUserNotesByLastSyncedAt returns the notes of the user by the last synced at timestamp
func (s *service) SyncUserNotesByLastSyncedAt(
	w http.ResponseWriter,
	r *http.Request,
) (
	int64,
	int64,
	*time.Time,
	*SyncUserNotesByLastSyncedAtResponseBody,
) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Get the user refresh token ID from the request
	userRefreshTokenID, err := internaljwtclaims.GetParentRefreshTokenID(r)
	if err != nil {
		panic(err)
	}

	// Get the last synced at from the cookies
	lastSyncedAt, err := internalcookie.GetSyncNotesCookie(r)
	if err != nil {
		panic(err)
	}

	// Get the user notes
	newLastSyncedAt := time.Now()
	var syncUserNotes []*internalpostgresmodel.SyncUserNoteWithID
	rows, err := internalpostgres.PoolService.Query(
		&internalpostgresmodel.SyncUserNotesByLastSyncedAtFn,
		userID,
		lastSyncedAt,
	)
	if err != nil {
		panic(err)
	}

	// Iterate over the user notes rows
	for rows.Next() {
		var syncUserNote internalpostgresmodel.SyncUserNoteWithID
		var hasToSyncNoteTags, hasToSyncNoteVersions bool
		if err = rows.Scan(
			&syncUserNote.ID,
			&syncUserNote.Title,
			&syncUserNote.Color,
			&syncUserNote.CreatedAt,
			&syncUserNote.UpdatedAt,
			&syncUserNote.PinnedAt,
			&syncUserNote.StarredAt,
			&syncUserNote.ArchivedAt,
			&syncUserNote.TrashedAt,
			&hasToSyncNoteTags,
			&hasToSyncNoteVersions,
		); err != nil {
			panic(err)
		}
		syncUserNotes = append(syncUserNotes, &syncUserNote)

		// Check if the user note versions has to be synced
		if hasToSyncNoteVersions {
			syncUserNote.SyncNoteVersions = *s.SyncUserNoteVersionsByLastSyncedAt(
				userID,
				lastSyncedAt,
				syncUserNote.ID,
			)
		}

		// Check if the user note tags has to be synced
		if hasToSyncNoteTags {
			syncUserNote.SyncNoteTags = *s.SyncUserNoteTagsByLastSyncedAt(
				userID,
				lastSyncedAt,
				syncUserNote.ID,
			)
		}
	}

	// Set the last sync at cookie
	internalcookie.SetSyncNotesCookie(w, newLastSyncedAt)

	return userID, userRefreshTokenID, lastSyncedAt, &SyncUserNotesByLastSyncedAtResponseBody{
		BaseJSendSuccessBody: *gonethttpresponse.NewBaseJSendSuccessBody(),
		Data: SyncUserNotesByLastSyncedAtResponseData{
			SyncNotes: syncUserNotes,
		},
	}
}
