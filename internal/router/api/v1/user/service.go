package user

import (
	"database/sql"
	gocryptobcrypt "github.com/ralvarezdev/go-crypto/bcrypt"
	gocryptorandomutf8 "github.com/ralvarezdev/go-crypto/random/strings/utf8"
	godatabasessql "github.com/ralvarezdev/go-databases/sql"
	"github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal"
	internalbcrypt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/bcrypt"
	internalpbkdf2 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/pbkdf2"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresconstraints "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/constraints"
	internalpostgresqueries "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/queries"
	"net/http"
)

type (
	// Service is the structure for the API V1 service for the user route group
	Service struct {
		PostgresService *internalpostgres.Service
	}
)

// SignUp signs up a user
func (s *Service) SignUp(r *http.Request, body *SignUpRequest) (
	*int64,
	error,
) {
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

	// Run the transaction
	var userID int64
	err = s.PostgresService.RunTransaction(
		func(tx *sql.Tx) error {
			// Create the new user
			if err = tx.QueryRow(
				internalpostgresqueries.InsertUser,
				body.FirstName,
				body.LastName,
				salt,
			).Scan(&userID); err != nil {
				return err
			}

			// Create the new user's email
			if _, err = tx.Exec(
				internalpostgresqueries.InsertUserEmail,
				userID,
				body.Email,
			); err != nil {
				isUniqueViolation, constraintName := godatabasessql.IsUniqueViolationError(err)
				if isUniqueViolation {
					if constraintName == internalpostgresconstraints.UserEmailsUniqueEmail {
						return ErrEmailAlreadyRegistered
					}
				}
				return err
			}

			// Create the new user's username
			if _, err = tx.Exec(
				internalpostgresqueries.InsertUserUsername,
				userID,
				body.Username,
			); err != nil {
				isUniqueViolation, constraintName := godatabasessql.IsUniqueViolationError(err)
				if isUniqueViolation {
					if constraintName == internalpostgresconstraints.UserUsernamesUniqueUsername {
						return ErrUsernameAlreadyRegistered
					}
				}
				return err
			}

			// Create the new user's password hash
			_, err = tx.Exec(
				internalpostgresqueries.InsertUserPasswordHash,
				userID,
				passwordHash,
			)
			return err
		},
	)
	return &userID, err
}
