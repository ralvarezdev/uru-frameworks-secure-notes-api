package postgres

import (
	"database/sql"
	godatabasessqlservice "github.com/ralvarezdev/go-databases/sql/service"
	internalpostgresqueries "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/queries"
)

type (
	// Service is the Postgres service struct
	Service struct {
		godatabasessqlservice.Service
	}
)

// NewService creates a new Postgres service
func NewService(db *sql.DB) (
	*Service,
	error,
) {
	// Create the default service
	defaultService, err := godatabasessqlservice.NewDefaultService(db)
	if err != nil {
		return nil, err
	}

	// Create the instance
	instance := &Service{
		Service: defaultService,
	}

	// Migrate the database
	/**/
	err = instance.Migrate()
	if err != nil {
		return nil, err
	}
	/**/

	return instance, nil
}

// Migrate migrates the database
func (s *Service) Migrate() error {
	return s.Service.Migrate(
		internalpostgresqueries.CreateUsers,
		internalpostgresqueries.CreateUserUsernames,
		internalpostgresqueries.CreateUserPasswordHashes,
		internalpostgresqueries.CreateUserResetPasswords,
		internalpostgresqueries.CreateUserEmails,
		internalpostgresqueries.CreateUserEmailVerifications,
		internalpostgresqueries.CreateUserPhoneNumbers,
		internalpostgresqueries.CreateUserPhoneNumberVerifications,
		internalpostgresqueries.CreateUserTokenSeeds,
		internalpostgresqueries.UserFailedLogInAttemptsCreate,
		internalpostgresqueries.CreateUserRefreshTokens,
		internalpostgresqueries.CreateUserAccessTokens,
		internalpostgresqueries.CreateUserTOTPs,
		internalpostgresqueries.CreateUserTOTPRecoveryCodes,
		internalpostgresqueries.CreateNotes,
		internalpostgresqueries.CreateTags,
		internalpostgresqueries.CreateNoteTags,
		internalpostgresqueries.CreateNoteVersions,
	)
}
