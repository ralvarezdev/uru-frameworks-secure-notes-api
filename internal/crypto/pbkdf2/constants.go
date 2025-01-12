package pbkdf2

import (
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
)

const (
	// EnvSaltLength is the key for the salt length
	EnvSaltLength = "URU_FRAMEWORKS_SECURE_NOTES_PBKDF2_SALT_LENGTH"
)

var (
	// SaltLength is the length of the salt
	SaltLength int
)

// Load loads the PBKDF2 constants
func Load() {
	// Get the salt length
	if err := internalloader.Loader.LoadIntVariable(
		EnvSaltLength,
		&SaltLength,
	); err != nil {
		panic(err)
	}
}
