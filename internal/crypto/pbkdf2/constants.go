package pbkdf2

import (
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
)

const (
	// EnvSaltLength is the key for the salt length
	EnvSaltLength = "URU_FRAMEWORKS_SECURE_NOTES_PBKDF2_SALT_LENGTH"

	// EnvIterations is the key for the iterations
	EnvIterations = "URU_FRAMEWORKS_SECURE_NOTES_PBKDF2_ITERATIONS"

	// EnvKeyLength is the key for the key length
	EnvKeyLength = "URU_FRAMEWORKS_SECURE_NOTES_PBKDF2_KEY_LENGTH"
)

var (
	// SaltLength is the length of the salt
	SaltLength int

	// Iterations is the number of iterations
	Iterations int

	// KeyLength is the length of the key
	KeyLength int
)

// Load loads the PBKDF2 constants
func Load() {
	// Get the salt length, iterations, and key length
	for env, dest := range map[string]*int{
		EnvSaltLength: &SaltLength,
		EnvIterations: &Iterations,
		EnvKeyLength:  &KeyLength,
	} {
		if err := internalloader.Loader.LoadIntVariable(
			env,
			dest,
		); err != nil {
			panic(err)
		}
	}
}
