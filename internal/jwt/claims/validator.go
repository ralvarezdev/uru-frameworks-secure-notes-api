package claims

import (
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	gojwt "github.com/ralvarezdev/go-jwt"
	gojwtcache "github.com/ralvarezdev/go-jwt/cache"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresqueries "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
)

// DefaultValidator struct
type DefaultValidator struct {
	postgresService *internalpostgres.Service
	tokenValidator  gojwtcache.TokenValidator
}

// NewDefaultValidator creates a new default validator
func NewDefaultValidator(
	postgresService *internalpostgres.Service,
	tokenValidator gojwtcache.TokenValidator,
) (*DefaultValidator, error) {
	return &DefaultValidator{
		postgresService,
		tokenValidator,
	}, nil
}

// IsRefreshTokenValid checks if the refresh token is valid
func (d *DefaultValidator) IsRefreshTokenValid(id string) (bool, error) {
	// Get the refresh token by the ID
	var expiresAt sql.NullTime
	var found, isExpired sql.NullBool
	if err := d.postgresService.QueryRow(
		&internalpostgresqueries.IsRefreshTokenValidProc,
		id,
		nil, nil, nil,
	).Scan(
		&expiresAt,
		&found,
		&isExpired,
	); err != nil {
		return false, err
	}

	// Check if it was found
	if !found.Bool {
		return false, nil
	}

	// Check if the token validator is not nil
	isValid := !isExpired.Bool
	if d.tokenValidator != nil {
		// Set the refresh token in the cache
		if err := d.tokenValidator.Set(
			gojwttoken.RefreshToken,
			id,
			isValid,
			expiresAt.Time,
		); err != nil {
			return false, err
		}
	}
	return isValid, nil
}

// IsAccessTokenValid checks if the access token is valid
func (d *DefaultValidator) IsAccessTokenValid(id string) (bool, error) {
	// Get the access token by the ID
	var expiresAt sql.NullTime
	var found, isExpired sql.NullBool
	if err := d.postgresService.QueryRow(
		&internalpostgresqueries.IsAccessTokenValidProc,
		id, nil, nil, nil,
	).Scan(
		&expiresAt,
		&found,
		&isExpired,
	); err != nil {
		return false, err
	}

	// Check if it was found
	if !found.Bool {
		return false, nil
	}

	// Check if the token validator is not nil
	isValid := !isExpired.Bool
	if d.tokenValidator != nil {
		// Set the access token in the cache
		if err := d.tokenValidator.Set(
			gojwttoken.AccessToken,
			id,
			isValid,
			expiresAt.Time,
		); err != nil {
			return false, err
		}
	}
	return isValid, nil
}

// ValidateClaims validates the claims
func (d *DefaultValidator) ValidateClaims(
	claims *jwt.MapClaims,
	token gojwttoken.Token,
) (bool, error) {
	// Get the JWT Identifier
	jti, ok := (*claims)[gojwt.IdClaim].(string)
	if !ok {
		return false, ErrInvalidIDClaim
	}

	// Check if the token validator is not nil
	if d.tokenValidator != nil {
		// Check if the token is valid
		isValid, err := d.tokenValidator.IsValid(token, jti)
		if err == nil {
			return isValid, nil
		}
	}

	// Validate the token
	if token == gojwttoken.RefreshToken {
		// Check if the refresh token is valid
		return d.IsRefreshTokenValid(
			jti,
		)
	}

	// Check if the access token is valid
	return d.IsAccessTokenValid(
		jti,
	)
}
