package jwt

import (
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
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
		key, err := internalloader.Loader.LoadVariable(env)
		if err != nil {
			panic(err)
		}
		internallogger.Environment.EnvironmentVariableLoaded(env)
		Keys[env] = key
	}

	// Get the JWT tokens duration
	for key, env := range map[gojwttoken.Token]string{
		gojwttoken.AccessToken:  EnvAccessTokenDuration,
		gojwttoken.RefreshToken: EnvRefreshTokenDuration,
	} {
		tokenDuration, err := internalloader.Loader.LoadDurationVariable(env)
		if err != nil {
			panic(err)
		}
		internallogger.Environment.EnvironmentVariableLoaded(env)
		Durations[key] = tokenDuration
	}
}
