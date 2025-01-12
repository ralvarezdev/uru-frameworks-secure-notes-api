package totp

import (
	gocryptototp "github.com/ralvarezdev/go-crypto/otp/totp"
	internalpostgresqueries "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/queries"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	"strconv"
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
	// Get the TOTP period
	if err := internalloader.Loader.LoadIntVariable(
		EnvPeriod,
		&Period,
	); err != nil {
		panic(err)
	}

	// Get the TOTP digits
	if err := internalloader.Loader.LoadIntVariable(
		EnvDigits,
		&Digits,
	); err != nil {
		panic(err)
	}

	// Get the TOTP secret length
	if err := internalloader.Loader.LoadIntVariable(
		EnvSecretLength,
		&SecretLength,
	); err != nil {
		panic(err)
	}

	// Get the TOTP recovery codes length
	if err := internalloader.Loader.LoadIntVariable(
		EnvRecoveryCodesLength,
		&RecoveryCodesLength,
	); err != nil {
		panic(err)
	}

	// Get the TOTP recovery codes count
	if err := internalloader.Loader.LoadIntVariable(
		EnvRecoveryCodesCount,
		&RecoveryCodesCount,
	); err != nil {
		panic(err)
	}

	// Set the InsertUserTOTPRecoveryCodes query based on the recovery codes count
	internalpostgresqueries.InsertUserTOTPRecoveryCodes = `
INSERT INTO user_totp_recovery_codes (
	user_totp_id,
	code,
	created_at
)
VALUES
`

	// Set the InsertUserTOTPRecoveryCodes query based on the recovery codes count
	var j string
	for i := 0; i < RecoveryCodesCount; i++ {
		j = strconv.Itoa(i + 2)
		if i == RecoveryCodesCount-1 {
			internalpostgresqueries.InsertUserTOTPRecoveryCodes += "($1, $" + j + ", NOW());"
		} else {
			internalpostgresqueries.InsertUserTOTPRecoveryCodes += "($1, $" + j + ", NOW()),"
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
