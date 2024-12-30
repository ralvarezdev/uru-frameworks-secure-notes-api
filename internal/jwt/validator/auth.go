package validator

import (
	gojwtredisauth "github.com/ralvarezdev/go-jwt/redis/auth"
	internalapi "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/api/v1"
	internalapiv1 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/api/v1"
)

type (
	// DefaultTokenValidator struct
	DefaultTokenValidator struct {
		apiV1Service        *internalapiv1.Service
		redisTokenValidator gojwtredisauth.TokenValidator
	}
)

// NewDefaultTokenValidator creates a new default token validator
func NewDefaultTokenValidator(
	apiV1Service *internalapiv1.Database,
	redisTokenValidator gojwtredisauth.TokenValidator,
) (*DefaultTokenValidator, error) {
	// Check if the PostgreSQL service is nil
	if apiV1Service == nil {
		return nil, internalapi.ErrNilService
	}

	return &DefaultTokenValidator{
		apiV1Service:        apiV1Service,
		redisTokenValidator: redisTokenValidator,
	}, nil
}

// IsTokenValid checks if the token is valid
func (d *DefaultTokenValidator) IsTokenValid(
	token string, jwtId string, isRefreshToken bool,
) (bool, error) {
	// Check if Redis is enabled. If it is, use the Redis cache
	if d.redisTokenValidator != nil {
		return d.redisTokenValidator.IsTokenValid(jwtId)
	}

	// Validate the token
	if isRefreshToken {
		// Check if the refresh token is valid
		isValid, err := d.apiV1Service.IsRefreshTokenValid(
			jwtId,
		)
		if err != nil {
			return false, err
		}
		return isValid, nil
	}

	// Check if the access token is valid
	isValid, err := d.apiV1Service.IsAccessTokenValid(
		jwtId,
	)
	if err != nil {
		return false, err
	}
	return isValid, nil
}
