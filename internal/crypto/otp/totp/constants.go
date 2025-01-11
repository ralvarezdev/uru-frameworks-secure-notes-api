package totp

import (
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
)

const (
	// EnvPeriod is the key to the TOTP period
	EnvPeriod = "URU_FRAMEWORKS_SECURE_NOTES_TOTP_PERIOD"

	// EnvDigits is the key to the TOTP digits
	EnvDigits = "URU_FRAMEWORKS_SECURE_NOTES_TOTP_DIGITS"
)

var (
	// Period is the TOTP period
	Period int

	// Digits is the TOTP digits
	Digits int
)

// Load loads the TOTP constants
func Load() {
	// Get the TOTP period
	totpPeriod, err := internalloader.Loader.LoadIntVariable(EnvPeriod)
	if err != nil {
		panic(err)
	}
	internallogger.Environment.EnvironmentVariableLoaded(EnvPeriod)
	Period = totpPeriod

	// Get the TOTP digits
	totpDigits, err := internalloader.Loader.LoadIntVariable(EnvDigits)
	if err != nil {
		panic(err)
	}
	internallogger.Environment.EnvironmentVariableLoaded(EnvDigits)
	Digits = totpDigits
}
