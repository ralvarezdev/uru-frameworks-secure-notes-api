package _common

import (
	"errors"
)

var (
	UserTOTPSecretNotFoundByUserID = errors.New("user TOTP secret not found by user ID")
)
