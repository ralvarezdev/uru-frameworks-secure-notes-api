package auth

import (
	"errors"
)

var (
	ErrInvalidPassword           = errors.New("invalid password")
	ErrInvalidTOTPCode           = errors.New("invalid TOTP code")
	ErrInvalidTOTPRecoveryCode   = errors.New("invalid TOTP recovery code")
	ErrMissingTOTPCode           = errors.New("missing TOTP code")
	ErrMissingIsTOTPRecoveryCode = errors.New("missing is TOTP recovery code")
)
