package user

import (
	"errors"
	gocryptobcrypt "github.com/ralvarezdev/go-crypto/bcrypt"
	gocryptorandomutf8 "github.com/ralvarezdev/go-crypto/random/strings/utf8"
	"github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal"
	internalbcrypt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/bcrypt"
	internalpbkdf2 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/pbkdf2"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	"gorm.io/gorm"
	"time"
)

type (
	// Service is the structure for the API V1 service for the user route group
	Service struct {
		PostgresService *internalpostgres.Service
	}
)

// SignUp signs up a user
func (s *Service) SignUp(body *SignUpRequest) (*internalpostgres.User, error) {
	if body == nil {
		return nil, internal.ErrNilRequestBody
	}

	// Hash the password
	passwordHash, err := gocryptobcrypt.HashPassword(
		body.Password,
		internalbcrypt.Cost,
	)
	if err != nil {
		return nil, err
	}

	// Generate a random salt
	salt, err := gocryptorandomutf8.Generate(internalpbkdf2.SaltLength)
	if err != nil {
		return nil, err
	}

	// Instantiate the new user
	user := internalpostgres.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Salt:      salt,
		JoinedAt:  time.Now(),
	}

	// Instantiate the new user's password hash
	userPasswordHash := internalpostgres.UserPasswordHash{
		PasswordHash: passwordHash,
		AssignedAt:   time.Now(),
	}

	// Instantiate the new user's email
	userEmail := internalpostgres.UserEmail{
		Email:      body.Email,
		AssignedAt: time.Now(),
	}

	// Instantiate the new user's username
	userUsername := internalpostgres.UserUsername{
		Username:   body.Username,
		AssignedAt: time.Now(),
	}

	// Run the transaction
	err = s.PostgresService.RunTransaction(
		func(tx *gorm.DB) error {
			// Create the new user
			if err = tx.Create(&user).Error; err != nil {
				return err
			}

			// Create the new user's password hash
			userPasswordHash.UserID = user.ID
			if err = tx.Create(&userPasswordHash).Error; err != nil {
				return err
			}

			// Create the new user's email
			userEmail.UserID = user.ID
			if err = tx.Create(&userEmail).Error; err != nil {
				if errors.Is(err, gorm.ErrDuplicatedKey) {
					return ErrEmailAlreadyRegistered
				}
				return err
			}

			// Create the new user's username
			userUsername.UserID = user.ID
			if err = tx.Create(&userUsername).Error; err != nil {
				if errors.Is(err, gorm.ErrDuplicatedKey) {
					return ErrUsernameAlreadyRegistered
				}
				return err
			}

			return nil
		},
	)
	return &user, err
}
