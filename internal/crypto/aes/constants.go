package aes

import (
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
)

const (
	// EnvKeySize is the key for the AES encryption key size
	EnvKeySize = "URU_FRAMEWORKS_SECURE_NOTES_AES_KEY_SIZE"
)

var (
	// KeySize is the size of the AES encryption key
	KeySize = 32
)

// Load loads the AES constants
func Load() {
	// Get the key size for the AES encryption
	if err := internalloader.Loader.LoadIntVariable(
		EnvKeySize,
		&KeySize,
	); err != nil {
		panic(err)
	}
}
