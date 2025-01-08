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
			&UserPasswordHash{},
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

// Close closes the Postgres service
func (s *Service) Close() error {
	return s.connection.Close()
}

// RunTransaction runs a transaction
func (s *Service) RunTransaction(transaction func(tx *gorm.DB) error) error {
	return s.database.Transaction(transaction)
}

// UserUsernameWasFound checks if a user username was found
func (s *Service) UserUsernameWasFound(userUsername *UserUsername) (
	*UserUsername,
	error,
) {
	if userUsername.ID != 0 {
		return userUsername, nil
	}
	return nil, ErrUserNotFound
}

// GetUserUsernameByUsername gets a user username by username preloaded with the user
func (s *Service) GetUserUsernameByUsername(username string) (
	userUsername *UserUsername,
	err error,
) {
	userUsername = &UserUsername{Username: username}
	err = s.database.Preload(
		"User.UserPasswordHash",
		"revoked_at IS NULL",
	).Where(
		"username = ? AND revoked_at IS NULL",
		username,
	).First(&userUsername).Error
	if err != nil {
		return nil, err
	}
	return s.UserUsernameWasFound(userUsername)
}
