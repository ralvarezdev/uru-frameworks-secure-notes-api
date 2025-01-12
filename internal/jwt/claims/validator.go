package claims

import (
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	gojwt "github.com/ralvarezdev/go-jwt"
	gojwtredisauth "github.com/ralvarezdev/go-jwt/redis/auth"
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresqueries "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/queries"
	"time"
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

// IsRefreshTokenValid checks if the refresh token is valid
func (d *DefaultValidator) IsRefreshTokenValid(id string) (bool, error) {
	// Get the database connection
	db := d.postgresService.DB()

	// Get the refresh token by the ID
	var expiresAt time.Time
	if err := db.QueryRow(
		internalpostgresqueries.SelectUserRefreshTokenExpiresAtByID,
		id,
	).Scan(&expiresAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return time.Now().Before(expiresAt), nil
}

// IsAccessTokenValid checks if the access token is valid
func (d *DefaultValidator) IsAccessTokenValid(id string) (bool, error) {
	// Get the database connection
	db := d.postgresService.DB()

	// Get the access token by the ID
	var expiresAt time.Time
	if err := db.QueryRow(
		internalpostgresqueries.SelectUserAccessTokenExpiresAtByID,
		id,
	).Scan(&expiresAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return time.Now().Before(expiresAt), nil
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
		return d.IsRefreshTokenValid(
			jwtId,
		)
	}

	// Check if the access token is valid
	return d.IsAccessTokenValid(
		jwtId,
	)
}
