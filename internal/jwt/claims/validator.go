package claims

import (
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	gojwt "github.com/ralvarezdev/go-jwt"
	gojwtcache "github.com/ralvarezdev/go-jwt/cache"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresqueries "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/queries"
	"time"
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
		postgresService: postgresService,
		tokenValidator:  tokenValidator,
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

	// Check if it is before the expiration time
	isValid := time.Now().Before(expiresAt)

	// Check if the token validator is not nil
	if d.tokenValidator != nil {
		// Set the refresh token in the cache
		if err := d.tokenValidator.Set(
			gojwttoken.RefreshToken,
			id,
			isValid,
			expiresAt,
		); err != nil {
			return false, err
		}
	}
	return isValid, nil
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

	// Check if it is before the expiration time
	isValid := time.Now().Before(expiresAt)

	// Check if the token validator is not nil
	if d.tokenValidator != nil {
		// Set the access token in the cache
		if err := d.tokenValidator.Set(
			gojwttoken.AccessToken,
			id,
			isValid,
			expiresAt,
		); err != nil {
			return false, err
		}
	}
	return isValid, nil
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
	jti, ok := (*claims)[gojwt.IdClaim].(string)
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

	// Check if the token validator is not nil
	if d.tokenValidator != nil {
		// Check if it is a refresh token
		var token gojwttoken.Token
		if isRefreshToken {
			token = gojwttoken.RefreshToken
		} else {
			token = gojwttoken.AccessToken
		}

		// Check if the token is valid
		isValid, err := d.tokenValidator.IsValid(token, jti)
		if err == nil {
			return isValid, nil
		}
	}

	// Validate the token
	if isRefreshToken {
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
