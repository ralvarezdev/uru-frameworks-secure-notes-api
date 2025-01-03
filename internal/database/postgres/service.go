package postgres

import (
	"database/sql"
	"gorm.io/gorm"
)

type (
	// Service is the Postgres service struct
	Service struct {
		database   *gorm.DB
		connection *sql.DB
	}
)

// NewService creates a new Postgres service
func NewService(database *gorm.DB, connection *sql.DB) (
	instance *Service,
	err error,
) {
	// Check if the database or the connection is nil
	if database == nil {
		return nil, ErrNilDatabase
	}
	if connection == nil {
		return nil, ErrNilConnection
	}

	/*
		// Migrate the database without the join tables and foreign keys
		err = database.AutoMigrate(
			&User{},
			&UserTokenSeed{},
			&UserFailedLogInAttempt{},
			&UserRefreshToken{},
			&UserAccessToken{},
			&UserTOTP{},
			&UserTOTPRecoveryCode{},
			&UserHashedPassword{},
			&UserUsername{},
			&UserResetPassword{},
			&UserEmail{},
			&UserEmailVerification{},
			&UserPhoneNumber{},
			&UserPhoneNumberVerification{},
			&Note{},
			&NoteVersion{},
			&Tag{},
			&NoteTag{},
		)
		if err != nil {
			return nil, err
		}
	*/

	return &Service{
		database:   database,
		connection: connection,
	}, nil
}

// Database returns the Postgres database
func (s *Service) Database() *gorm.DB {
	return s.database
}

// Close closes the Postgres service
func (s *Service) Close() error {
	// Close the connection
	return s.connection.Close()
}
