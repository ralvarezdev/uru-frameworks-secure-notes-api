package cache

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwtcache "github.com/ralvarezdev/go-jwt/cache"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
)

var (
	// TokenValidator is the cache token validator
	TokenValidator gojwtcache.TokenValidator
)

// Load initializes the cache
func Load(mode *goflagsmode.Flag) {
	// Check if the mode is debug
	if mode != nil && mode.IsDebug() {
		TokenValidator = gojwtcache.NewTokenValidatorService(internallogger.CacheTokenValidator)
	} else {
		TokenValidator = gojwtcache.NewTokenValidatorService(nil)
	}
}
