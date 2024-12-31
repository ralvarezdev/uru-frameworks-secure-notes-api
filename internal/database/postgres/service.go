package postgres

import (
	"database/sql"
	modelpg "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/database/postgres/model"
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
func NewService(database *gorm.DB, connection *sql.DB) (*Service, error) {
	// Check if the database or the connection is nil
	if database == nil {
		return nil, ErrNilDatabase
	}
	if connection == nil {
		return nil, ErrNilConnection
	}

	// Migrate the database
	err := database.AutoMigrate(
		&modelpg.Note{},
		&modelpg.NoteTag{},
		&modelpg.NoteVersion{},
		&modelpg.User{},
		&modelpg.UserAccessToken{},
		&modelpg.UserEmail{},
		&modelpg.UserEmailVerification{},
		&modelpg.UserFailedLogInAttempt{},
		&modelpg.UserHashedPassword{},
		&modelpg.UserPhoneNumber{},
		&modelpg.UserPhoneNumberVerification{},
		&modelpg.UserRefreshToken{},
		&modelpg.UserResetPassword{},
		&modelpg.UserTokenSeed{},
		&modelpg.UserTOTP{},
		&modelpg.UserTOTPRecoveryCode{},
		&modelpg.UserUsername{},
	)
	if err != nil {
		return nil, err
	}

	return &Service{
		database: database,
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
