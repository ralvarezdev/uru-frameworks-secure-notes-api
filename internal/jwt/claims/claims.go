package claims

import (
	"github.com/golang-jwt/jwt/v5"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	"strconv"
	"time"
)

type (
	// Claims is the structure for the JWT claims
	Claims struct {
		IsRefreshToken bool `json:"irt"`
		jwt.RegisteredClaims
	}
)

// NewClaims creates a new JWT claims
func NewClaims(
	token gojwttoken.Token,
	id int64,
	subject string,
	issuedAt, expiresAt time.Time,
) *Claims {
	return &Claims{
		IsRefreshToken: token == gojwttoken.RefreshToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.FormatInt(id, 10),
			Subject:   subject,
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
}

// NewRefreshTokenClaims creates a new JWT refresh token claims
func NewRefreshTokenClaims(
	id int64,
	subject string,
	issuedAt, expiresAt time.Time,
) *Claims {
	return NewClaims(gojwttoken.RefreshToken, id, subject, issuedAt, expiresAt)
}

// NewAccessTokenClaims creates a new JWT access token claims
func NewAccessTokenClaims(
	id int64,
	subject string,
	issuedAt, expiresAt time.Time,
) *Claims {
	return NewClaims(gojwttoken.AccessToken, id, subject, issuedAt, expiresAt)
}
