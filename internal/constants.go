package internal

import (
	"errors"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
)

const (
	// EnvVerifyEmailURL is the environment variable for the verify email URL
	EnvVerifyEmailURL = "URU_FRAMEWORKS_SECURE_NOTES_VERIFY_EMAIL_URL"

	// EnvResetPasswordURL is the environment variable for the reset password URL
	EnvResetPasswordURL = "URU_FRAMEWORKS_SECURE_NOTES_RESET_PASSWORD_URL"
)

var (
	// VerifyEmailURL is the verify email URL
	VerifyEmailURL string

	// ResetPasswordURL is the reset password URL
	ResetPasswordURL string
)

var (
	// InDevelopment is the response when a request is made to an endpoint that is not implemented yet
	InDevelopment = errors.New("this endpoint is not implemented yet")
)

// Load loads the URL constants
func Load() {
	// Load the environment variables
	for env, dest := range map[string]*string{
		EnvVerifyEmailURL:   &VerifyEmailURL,
		EnvResetPasswordURL: &ResetPasswordURL,
	} {
		if err := internalloader.Loader.LoadVariable(
			env,
			dest,
		); err != nil {
			panic(err)
		}
	}
}
