package postgres

import (
	"database/sql"
	godatabasessqlservice "github.com/ralvarezdev/go-databases/sql/service"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
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
	err = instance.Migrate()
	if err != nil {
		return nil, err
	}

	return instance, nil
}

// Migrate migrates the database
func (s *Service) Migrate() error {
	return s.Service.Migrate(
		internalpostgresqueries.UsersMigrate,
		internalpostgresqueries.UserUsernamesMigrate,
		internalpostgresqueries.UserPasswordHashesMigrate,
		internalpostgresqueries.UserResetPasswordMigrate,
		internalpostgresqueries.UserEmailsMigrate,
		internalpostgresqueries.UserEmailVerificationsMigrate,
		internalpostgresqueries.UserPhoneNumbersMigrate,
		internalpostgresqueries.UserPhoneNumberVerificationsMigrate,
		internalpostgresqueries.UserTokenSeedsMigrate,
		internalpostgresqueries.UserFailedLogInAttemptsMigrate,
		internalpostgresqueries.UserRefreshTokensMigrate,
		internalpostgresqueries.UserAccessTokensMigrate,
		internalpostgresqueries.UserTOTPsMigrate,
		internalpostgresqueries.UserTOTPRecoveryCodesMigrate,
		internalpostgresqueries.NotesMigrate,
		internalpostgresqueries.TagsMigrate,
		internalpostgresqueries.NoteTagsMigrate,
		internalpostgresqueries.NoteVersionsMigrate,
	)
}

// WasUserFound checks if a user was found
func (s *Service) WasUserFound(user *internalpostgresmodel.User) (
	*internalpostgresmodel.User,
	error,
) {
	if user.ID != 0 {
		return user, nil
	}
	return nil, ErrUserNotFound
}
