package internal

import (
	"errors"
	govalidatormappervalidations "github.com/ralvarezdev/go-validator/struct/mapper/validator"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	"time"
)

const (
	// EnvVerifyEmailURL is the environment variable for the verify email URL
	EnvVerifyEmailURL = "URU_FRAMEWORKS_SECURE_NOTES_VERIFY_EMAIL_URL"

	// EnvResetPasswordURL is the environment variable for the reset password URL
	EnvResetPasswordURL = "URU_FRAMEWORKS_SECURE_NOTES_RESET_PASSWORD_URL"

	// EnvMinimumPasswordLength is the environment for the minimum password length
	EnvMinimumPasswordLength = "URU_FRAMEWORKS_SECURE_NOTES_MINIMUM_PASSWORD_LENGTH"

	// EnvMinimumPasswordSpecialCount is the environment for the minimum password special characters count
	EnvMinimumPasswordSpecialCount = "URU_FRAMEWORKS_SECURE_NOTES_MINIMUM_PASSWORD_SPECIAL_COUNT"

	// EnvMinimumPasswordNumbersCount is the environment for the minimum password numbers characters count
	EnvMinimumPasswordNumbersCount = "URU_FRAMEWORKS_SECURE_NOTES_MINIMUM_PASSWORD_NUMBERS_COUNT"

	// EnvMinimumPasswordCapsCount is the environment for the minimum password caps characters count
	EnvMinimumPasswordCapsCount = "URU_FRAMEWORKS_SECURE_NOTES_MINIMUM_PASSWORD_CAPS_COUNT"

	// EnvMaximumFailedAttemptsCount is the environment for the maximum failed attempts count
	EnvMaximumFailedAttemptsCount = "URU_FRAMEWORKS_SECURE_NOTES_MAXIMUM_FAILED_ATTEMPTS_COUNT"

	// EnvMaximumFailedAttemptsPeriod is the environment for the maximum failed attempts period
	EnvMaximumFailedAttemptsPeriod = "URU_FRAMEWORKS_SECURE_NOTES_MAXIMUM_FAILED_ATTEMPTS_PERIOD"

	// EnvMinimumAge is the environment for the minimum age
	EnvMinimumAge = "URU_FRAMEWORKS_SECURE_NOTES_MINIMUM_AGE"

	// EnvMaximumAge is the environment for the maximum age
	EnvMaximumAge = "URU_FRAMEWORKS_SECURE_NOTES_MAXIMUM_AGE"

	// EnvTwoFactorAuthenticationDuration is the environment for the 2FA duration
	EnvTwoFactorAuthenticationDuration = "URU_FRAMEWORKS_SECURE_NOTES_2FA_DURATION"

	// EnvTwoFactorAuthenticationEmailDuration is the environment for the 2FA email duration
	EnvTwoFactorAuthenticationEmailDuration = "URU_FRAMEWORKS_SECURE_NOTES_2FA_EMAIL_DURATION"

	// EmailCodeType is the type of email code
	EmailCodeType = "email-code"

	// TOTPCodeType is the type of TOTP code
	TOTPCodeType = "totp-code"

	// RecoveryCodeType is the type of recovery code
	RecoveryCodeType = "recovery-code"
)

var (
	// VerifyEmailURL is the verify email URL
	VerifyEmailURL string

	// ResetPasswordURL is the reset password URL
	ResetPasswordURL string

	// MinimumPasswordLength is the minimum password length
	MinimumPasswordLength int

	// MinimumPasswordSpecialCount is the minimum password special characters count
	MinimumPasswordSpecialCount int

	// MinimumPasswordNumbersCount is the minimum password numbers characters count
	MinimumPasswordNumbersCount int

	// MinimumPasswordCapsCount is the minimum password caps characters count
	MinimumPasswordCapsCount int

	// PasswordOptions is the password options
	PasswordOptions *govalidatormappervalidations.PasswordOptions

	// MinimumAge is the minimum age
	MinimumAge int

	// MaximumAge is the maximum age
	MaximumAge int

	// BirthdateOptions is the birthdate options
	BirthdateOptions *govalidatormappervalidations.BirthdateOptions

	// MaximumFailedAttemptsCount is the maximum failed attempts count
	MaximumFailedAttemptsCount int

	// MaximumFailedAttemptsPeriod is the maximum failed attempts period
	MaximumFailedAttemptsPeriod time.Duration

	// MaximumFailedAttemptsPeriodSeconds is the maximum failed attempts period in seconds
	MaximumFailedAttemptsPeriodSeconds int64

	// TwoFactorAuthenticationDuration is the 2FA duration
	TwoFactorAuthenticationDuration time.Duration

	// TwoFactorAuthenticationEmailDuration is the 2FA email duration
	TwoFactorAuthenticationEmailDuration time.Duration
)

var (
	// InDevelopment is the response when a request is made to an endpoint that is not implemented yet
	InDevelopment = errors.New("this endpoint is not implemented yet")
)

// Load loads the URL constants
func Load() {
	// Get the verify email and reset password URL
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

	// Get the password-related counts and length, and age-related counts
	for env, dest := range map[string]*int{
		EnvMinimumPasswordLength:       &MinimumPasswordLength,
		EnvMinimumPasswordSpecialCount: &MinimumPasswordSpecialCount,
		EnvMinimumPasswordNumbersCount: &MinimumPasswordNumbersCount,
		EnvMinimumPasswordCapsCount:    &MinimumPasswordCapsCount,
		EnvMaximumFailedAttemptsCount:  &MaximumFailedAttemptsCount,
		EnvMinimumAge:                  &MinimumAge,
		EnvMaximumAge:                  &MaximumAge,
	} {
		if err := internalloader.Loader.LoadIntVariable(
			env,
			dest,
		); err != nil {
			panic(err)
		}
	}

	// Get the maximum failed attempts period duration and 2FA duration
	for env, dest := range map[string]*time.Duration{
		EnvMaximumFailedAttemptsPeriod:          &MaximumFailedAttemptsPeriod,
		EnvTwoFactorAuthenticationDuration:      &TwoFactorAuthenticationDuration,
		EnvTwoFactorAuthenticationEmailDuration: &TwoFactorAuthenticationEmailDuration,
	} {
		if err := internalloader.Loader.LoadDurationVariable(
			env,
			dest,
		); err != nil {
			panic(err)
		}
	}

	// Get the maximum failed attempts period in seconds
	MaximumFailedAttemptsPeriodSeconds = int64(MaximumFailedAttemptsPeriod.Seconds())

	// Create the password options
	PasswordOptions = &govalidatormappervalidations.PasswordOptions{
		MinimumLength:       MinimumPasswordLength,
		MinimumSpecialCount: MinimumPasswordSpecialCount,
		MinimumNumbersCount: MinimumPasswordNumbersCount,
		MinimumCapsCount:    MinimumPasswordCapsCount,
	}

	// Create the birthdate options
	BirthdateOptions = &govalidatormappervalidations.BirthdateOptions{
		MinimumAge: MinimumAge,
		MaximumAge: MaximumAge,
	}
}
