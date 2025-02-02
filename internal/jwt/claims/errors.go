package claims

import (
	"errors"
)

var (
	ErrInvalidSubjectClaim              = errors.New("invalid subject claim")
	ErrInvalidIDClaim                   = errors.New("invalid id claim")
	ErrInvalidParentRefreshTokenIDClaim = errors.New("invalid parent refresh token id claim")
)
