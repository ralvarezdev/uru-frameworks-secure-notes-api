package token

import (
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	"time"
)

const (
	// EnvEmailVerificationTokenDuration is the environment variable for the email verification token duration
	EnvEmailVerificationTokenDuration = "URU_FRAMEWORKS_SECURE_NOTES_EMAIL_VERIFICATION_TOKEN_DURATION"
)

var (
	// EmailVerificationTokenDuration is the email verification token duration
	EmailVerificationTokenDuration time.Duration
)

// Load loads the token constants
func Load() {
	// Load the email verification token duration
	if err := internalloader.Loader.LoadDurationVariable(
		EnvEmailVerificationTokenDuration,
		&EmailVerificationTokenDuration,
	); err != nil {
		panic(err)
	}
}
