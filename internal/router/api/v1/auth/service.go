package auth

import (
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
)

type (
	// Service is the structure for the API V1 service for the auth route group
	Service struct {
		JwtIssuer       gojwtissuer.Issuer
		PostgresService *internalpostgres.Service
	}
)

// LogIn logs in a user
/*
func (s *Service) LogIn(body *LogInRequest) (*LogInResponse, error) {
	// Check if the body is nil
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

	// Run the transaction
	err = s.PostgresService.RunTransaction(
		func(tx *gorm.DB) error {
			// Get the user username preloaded with the user
			userUsername, err := s.PostgresService.GetUserUsernameByUsername(body.Username)
			if err != nil {
				return err
			}

			fmt.Println(userUsername, passwordHash)

				// Check if the password is correct
				if !gocryptobcrypt.CheckPasswordHash(userUsername.) {
					return nil, ErrInvalidPassword
				}

				// Create the JWT token
				token, err := s.JwtIssuer.IssueToken(user.ID)
				if err != nil {
					return nil, err
				}


			return nil
		},
	)

	return &LogInResponse{}, nil

}
*/
