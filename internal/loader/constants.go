package loader

import (
	"github.com/joho/godotenv"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	goloaderenv "github.com/ralvarezdev/go-loader/env"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
)

var (
	// Loader is the environment variables loader
	Loader goloaderenv.Loader
)

// Load loads the loader
func Load() {
	// Load the environment variables loader
	loader, _ := goloaderenv.NewDefaultLoader(
		func() error {
			// Check if the environment is production
			if goflagsmode.ModeFlag != nil && goflagsmode.ModeFlag.IsProd() {
				return nil
			}

			return godotenv.Load()
		},
		internallogger.Environment,
	)
	Loader = loader
}
