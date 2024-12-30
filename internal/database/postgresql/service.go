package postgresql

import (
	"database/sql"
	modelpg "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/database/postgresql/model"
	"gorm.io/gorm"
)

type (
	// Service is the PostgreSQL service struct
	Service struct {
		database   *gorm.DB
		connection *sql.DB
	}
)

// NewService creates a new PostgreSQL service
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

// Database returns the PostgreSQL database
func (s *Service) Database() *gorm.DB {
	return s.database
}

// Close closes the PostgreSQL service
func (s *Service) Close() error {
	// Close the connection
	return s.connection.Close()
}

// IsAccessTokenValid checks if the access token is valid
func (s *Service) IsAccessTokenValid(jwtId string) (bool, error) {
	return false, nil
}

// IsRefreshTokenValid checks if the refresh token is valid
func (s *Service) IsRefreshTokenValid(jwtId string) (bool, error) {
	return false, nil
}
