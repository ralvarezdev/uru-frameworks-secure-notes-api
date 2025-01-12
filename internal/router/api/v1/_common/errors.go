package _common

import (
	"errors"
)

var (
	UserNotFoundByUsername         = errors.New("user not found by username")
	UserTOTPSecretNotFoundByUserID = errors.New("user TOTP secret not found by user ID")
)
