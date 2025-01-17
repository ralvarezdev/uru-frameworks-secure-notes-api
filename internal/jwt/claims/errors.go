package claims

import (
	"errors"
)

var (
	ErrIsRefreshTokenClaimNotValid = errors.New("irt not valid")
	ErrIdClaimNotValid             = errors.New("jwt_id not valid")
	ErrMustBeAccessToken           = errors.New("must be access token")
	ErrMustBeRefreshToken          = errors.New("must be refresh token")
	ErrInvalidSubjectClaim         = errors.New("invalid subject claim")
	ErrInvalidIDClaim              = errors.New("invalid id claim")
	ErrInvalidTokenValueType       = "invalid token value type, expected: %s"
)
