package loader

import (
	"github.com/joho/godotenv"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	goloaderenv "github.com/ralvarezdev/go-loader/env"
)

var (
	// Loader is the environment variables loader
	Loader, _ = goloaderenv.NewDefaultLoader(
		func() error {
			// Check if the environment is production
			if goflagsmode.Mode != nil && goflagsmode.Mode.IsProd() {
				return nil
			}

			return godotenv.Load()
		},
	)
)
