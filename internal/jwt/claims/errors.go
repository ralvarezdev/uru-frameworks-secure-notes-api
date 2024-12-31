package claims

import (
	"errors"
)

var (
	ErrIsRefreshTokenClaimNotValid = errors.New("irt not valid")
	ErrIdClaimNotValid             = errors.New("jwt_id not valid")
	ErrMustBeAccessToken           = errors.New("must be access token")
	ErrMustBeRefreshToken          = errors.New("must be refresh token")
	ErrTokenExpired                = errors.New("token expired")
)
