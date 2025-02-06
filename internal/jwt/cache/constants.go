package cache

import (
	"database/sql"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwtcache "github.com/ralvarezdev/go-jwt/cache"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"strconv"
	"time"
)

var (
	// TokenValidator is the cache token validator
	TokenValidator gojwtcache.TokenValidator

	// SetTokenToCache sets the token to the cache
	SetTokenToCache func(
		token gojwttoken.Token,
		id int64,
		expiresAt time.Time,
		isValid bool,
	)

	// RevokeTokenFromCache revokes the token from the cache
	RevokeTokenFromCache func(
		token gojwttoken.Token,
		id int64,
	)

	// RevokeRefreshTokenFromCache revokes the refresh token from the cache
	RevokeRefreshTokenFromCache func(
		userRefreshTokenID int64,
	)

	// RevokeUserRefreshTokensFromCache revokes the user refresh tokens from the cache
	RevokeUserRefreshTokensFromCache func(userID int64)
)

// Load initializes the cache
func Load(mode *goflagsmode.Flag) {
	// Check if the mode is debug
	if mode != nil && mode.IsDebug() {
		TokenValidator = gojwtcache.NewTokenValidatorService(internallogger.CacheTokenValidator)
	} else {
		TokenValidator = gojwtcache.NewTokenValidatorService(nil)
	}

	// Set the function to set the token to the cache
	SetTokenToCache = func(
		token gojwttoken.Token,
		id int64,
		expiresAt time.Time,
		isValid bool,
	) {
		_ = TokenValidator.Set(
			token,
			strconv.FormatInt(id, 10),
			isValid,
			expiresAt,
		)
	}

	// Set the function to revoke the token from the cache
	RevokeTokenFromCache = func(
		token gojwttoken.Token,
		id int64,
	) {
		_ = TokenValidator.Revoke(
			token,
			strconv.FormatInt(id, 10),
		)
	}

	// Set the function to revoke the refresh token from the cache
	RevokeRefreshTokenFromCache = func(userRefreshTokenID int64) {
		// Get the user access token ID by the user refresh token ID
		var userAccessTokenID sql.NullInt64
		if err := internalpostgres.PoolService.QueryRow(
			&internalpostgresmodel.GetUserAccessTokenByUserRefreshTokenIDProc,
			userRefreshTokenID,
			nil,
		).Scan(
			&userAccessTokenID,
		); err != nil {
			return
		}

		// Revoke the tokens in the cache
		RevokeTokenFromCache(
			gojwttoken.RefreshToken,
			userRefreshTokenID,
		)
		RevokeTokenFromCache(
			gojwttoken.AccessToken,
			userAccessTokenID.Int64,
		)
	}

	// Set the function to revoke the user refresh tokens from the cache
	RevokeUserRefreshTokensFromCache = func(userID int64) {
		// Get the user refresh tokens and user access tokens ID by user ID
		var userRefreshTokenID, userAccessTokenID int64
		rows, err := internalpostgres.PoolService.Query(
			&internalpostgresmodel.ListUserTokensFn,
			userID,
		)
		if err != nil {
			return
		}
		defer rows.Close()

		// Parse the user refresh tokens ID
		for rows.Next() {
			if err = rows.Scan(
				&userRefreshTokenID,
				&userAccessTokenID,
			); err != nil {
				return
			}

			// Revoke the user refresh token and user access token from the cache
			RevokeTokenFromCache(
				gojwttoken.RefreshToken,
				userRefreshTokenID,
			)
			RevokeTokenFromCache(
				gojwttoken.AccessToken,
				userAccessTokenID,
			)
		}
	}
}
