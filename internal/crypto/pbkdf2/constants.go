package pbkdf2

import (
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
)

const (
	// SaltLengthKey is the key for the salt length
	SaltLengthKey = "URU_FRAMEWORKS_SECURE_NOTES_PBKDF2_SALT_LENGTH"
)

var (
	// SaltLength is the length of the salt
	SaltLength int
)

// init initializes the salt length
func init() {
	// Get the salt length
	saltLength, err := internalloader.Loader.LoadIntVariable(SaltLengthKey)
	if err != nil {
		panic(err)
	}
	internallogger.Environment.EnvironmentVariableLoaded(SaltLengthKey)
	SaltLength = saltLength
}
