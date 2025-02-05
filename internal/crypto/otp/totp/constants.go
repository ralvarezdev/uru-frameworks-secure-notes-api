package totp

import (
	gocryptototp "github.com/ralvarezdev/go-crypto/otp/totp"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
)

const (
	// EnvPeriod is the key to the TOTP period
	EnvPeriod = "URU_FRAMEWORKS_SECURE_NOTES_TOTP_PERIOD"

	// EnvDigits is the key to the TOTP digits
	EnvDigits = "URU_FRAMEWORKS_SECURE_NOTES_TOTP_DIGITS"

	// EnvSecretLength is the key to the TOTP secret length
	EnvSecretLength = "URU_FRAMEWORKS_SECURE_NOTES_TOTP_SECRET_LENGTH"

	// EnvRecoveryCodesLength is the key to the TOTP recovery codes length
	EnvRecoveryCodesLength = "URU_FRAMEWORKS_SECURE_NOTES_TOTP_RECOVERY_CODES_LENGTH"

	// EnvRecoveryCodesCount is the key to the TOTP recovery codes count
	EnvRecoveryCodesCount = "URU_FRAMEWORKS_SECURE_NOTES_TOTP_RECOVERY_CODES_COUNT"
)

var (
	// Period is the TOTP period
	Period int

	// Digits is the TOTP digits
	Digits int

	// SecretLength is the TOTP secret length
	SecretLength int

	// RecoveryCodesLength is the TOTP recovery codes length
	RecoveryCodesLength int

	// RecoveryCodesCount is the TOTP recovery codes count
	RecoveryCodesCount int

	// Url is the TOTP URL
	Url *gocryptototp.Url
)

// Load loads the TOTP constants
func Load() {
	// Load the TOTP period, digits, secret length, recovery codes length, and recovery codes count
	for env, dest := range map[string]*int{
		EnvPeriod:              &Period,
		EnvDigits:              &Digits,
		EnvSecretLength:        &SecretLength,
		EnvRecoveryCodesLength: &RecoveryCodesLength,
		EnvRecoveryCodesCount:  &RecoveryCodesCount,
	} {
		if err := internalloader.Loader.LoadIntVariable(
			env,
			dest,
		); err != nil {
			panic(err)
		}
	}

	// Set the TOTP URL
	Url = gocryptototp.NewUrl(
		"Secure Notes",
		"SHA1",
		Digits,
		Period,
	)
}
