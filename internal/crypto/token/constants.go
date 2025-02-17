package token

import (
	gostringsconvert "github.com/ralvarezdev/go-strings/convert"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	"time"
)

const (
	// EnvEmailVerificationTokenDuration is the environment variable for the email verification token duration
	EnvEmailVerificationTokenDuration = "URU_FRAMEWORKS_SECURE_NOTES_EMAIL_VERIFICATION_TOKEN_DURATION"

	// EnvResetPasswordTokenDuration is the environment variable for the reset password token duration
	EnvResetPasswordTokenDuration = "URU_FRAMEWORKS_SECURE_NOTES_RESET_PASSWORD_TOKEN_DURATION"
)

var (
	// EmailVerificationTokenDuration is the email verification token duration
	EmailVerificationTokenDuration time.Duration

	// ResetPasswordTokenDuration is the reset password token duration
	ResetPasswordTokenDuration time.Duration

	// PrettyEmailVerificationTokenDuration is the pretty email verification token duration
	PrettyEmailVerificationTokenDuration string

	// PrettyResetPasswordTokenDuration is the pretty reset password token duration
	PrettyResetPasswordTokenDuration string
)

// Load loads the token constants
func Load() {
	// Load the email verification token and reset password token durations
	for env, dest := range map[string]*time.Duration{
		EnvEmailVerificationTokenDuration: &EmailVerificationTokenDuration,
		EnvResetPasswordTokenDuration:     &ResetPasswordTokenDuration,
	} {
		if err := internalloader.Loader.LoadDurationVariable(
			env,
			dest,
		); err != nil {
			panic(err)
		}
	}

	// Set the pretty email verification token duration
	PrettyEmailVerificationTokenDuration, _ = gostringsconvert.PrettyDuration(
		EmailVerificationTokenDuration,
		"s",
	)

	// Set the pretty reset password token duration
	PrettyResetPasswordTokenDuration, _ = gostringsconvert.PrettyDuration(
		ResetPasswordTokenDuration,
		"s",
	)
}
