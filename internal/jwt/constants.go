package jwt

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internaljwtcache "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/cache"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	"time"
)

const (
	// EnvPublicKey is the key of the JWT public key
	EnvPublicKey = "URU_FRAMEWORKS_SECURE_NOTES_JWT_PUBLIC_KEY"

	// EnvPrivateKey is the key of the JWT private key
	EnvPrivateKey = "URU_FRAMEWORKS_SECURE_NOTES_JWT_PRIVATE_KEY"

	// EnvAccessTokenDuration is the key of the access token duration
	EnvAccessTokenDuration = "URU_FRAMEWORKS_SECURE_NOTES_ACCESS_TOKEN_DURATION"

	// EnvRefreshTokenDuration is the key of the refresh token duration
	EnvRefreshTokenDuration = "URU_FRAMEWORKS_SECURE_NOTES_REFRESH_TOKEN_DURATION"
)

var (
	// Keys are the JWT keys
	Keys = make(map[string]string)

	// Durations are the JWT tokens duration
	Durations = make(map[gojwttoken.Token]time.Duration)

	// Validator is the JWT validator
	Validator gojwtvalidator.Validator

	// Issuer is the JWT issuer
	Issuer gojwtissuer.Issuer
)

// Load loads the JWT constants
func Load() {
	// Get the JWT keys
	for _, env := range []string{
		EnvPublicKey,
		EnvPrivateKey,
	} {
		var key string
		if err := internalloader.Loader.LoadVariable(env, &key); err != nil {
			panic(err)
		}
		Keys[env] = key
	}

	// Get the JWT tokens duration
	for key, env := range map[gojwttoken.Token]string{
		gojwttoken.AccessToken:  EnvAccessTokenDuration,
		gojwttoken.RefreshToken: EnvRefreshTokenDuration,
	} {
		var tokenDuration time.Duration
		if err := internalloader.Loader.LoadDurationVariable(
			env,
			&tokenDuration,
		); err != nil {
			panic(err)
		}
		Durations[key] = tokenDuration
	}

	// Create the JWT claims validator
	claimsValidator, _ := internaljwtclaims.NewDefaultValidator(
		internalpostgres.PoolService, internaljwtcache.TokenValidator,
	)

	// Create the JWT validator with ED25519 public key
	validator, err := gojwtvalidator.NewEd25519Validator(
		[]byte(Keys[EnvPublicKey]),
		claimsValidator,
		goflagsmode.ModeFlag,
	)
	if err != nil {
		panic(err)
	}
	Validator = validator

	// Create the JWT issuer with ED25519 private key
	issuer, err := gojwtissuer.NewEd25519Issuer(
		[]byte(Keys[EnvPrivateKey]),
	)
	if err != nil {
		panic(err)
	}
	Issuer = issuer
}
