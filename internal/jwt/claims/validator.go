package claims

import (
	"github.com/golang-jwt/jwt/v5"
	gojwt "github.com/ralvarezdev/go-jwt"
	gojwtredisauth "github.com/ralvarezdev/go-jwt/redis/auth"
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/database/postgres"
)

// DefaultValidator struct
type DefaultValidator struct {
	postgresService     *internalpostgres.Service
	redisTokenValidator gojwtredisauth.TokenValidator
}

// NewDefaultValidator creates a new default validator
func NewDefaultValidator(
	postgresService *internalpostgres.Service,
	redisTokenValidator gojwtredisauth.TokenValidator,
) (*DefaultValidator, error) {
	return &DefaultValidator{
		postgresService:     postgresService,
		redisTokenValidator: redisTokenValidator,
	}, nil
}

// IsAccessTokenValid checks if the access token is valid
func (d *DefaultValidator) IsAccessTokenValid(jwtId string) (bool, error) {
	return false, nil
}

// IsRefreshTokenValid checks if the refresh token is valid
func (d *DefaultValidator) IsRefreshTokenValid(jwtId string) (bool, error) {
	return false, nil
}

// ValidateClaims validates the claims
func (d *DefaultValidator) ValidateClaims(
	claims *jwt.MapClaims,
	interception gojwtinterception.Interception,
) (bool, error) {
	// Check if is a refresh token
	isRefreshToken, ok := (*claims)[gojwt.IsRefreshTokenClaim].(bool)
	if !ok {
		return false, ErrIsRefreshTokenClaimNotValid
	}

	// Get the JWT Identifier
	jwtId, ok := (*claims)[gojwt.IdClaim].(string)
	if !ok {
		return false, ErrIdClaimNotValid
	}

	// Check if it must be refresh token
	if !isRefreshToken && interception == gojwtinterception.RefreshToken {
		return false, ErrMustBeRefreshToken
	}

	// Check if it must be access token
	if isRefreshToken && interception == gojwtinterception.AccessToken {
		return false, ErrMustBeAccessToken
	}

	// Check if redis is enabled. If it is, check if the token is valid
	if d.redisTokenValidator != nil {
		return d.redisTokenValidator.IsTokenValid(jwtId)
	}

	// Validate the token
	if isRefreshToken {
		// Check if the refresh token is valid
		isValid, err := d.IsRefreshTokenValid(
			jwtId,
		)
		if err != nil {
			return false, err
		}
		return isValid, nil
	}

	// Check if the access token is valid
	isValid, err := d.IsAccessTokenValid(
		jwtId,
	)
	if err != nil {
		return false, err
	}
	return isValid, nil
}