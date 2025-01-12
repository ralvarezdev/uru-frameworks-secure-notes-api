package jwt

import (
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
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
}
