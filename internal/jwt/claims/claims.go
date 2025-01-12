package claims

import (
	"github.com/golang-jwt/jwt/v5"
	gojwtnethttpctx "github.com/ralvarezdev/go-jwt/net/http/context"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	"net/http"
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

// ParseInt64 parses the string to int64
func ParseInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// GetSubject returns the subject
func GetSubject(r *http.Request) (int64, error) {
	// Get the claims from the request
	tokenClaims, err := gojwtnethttpctx.GetCtxTokenClaims(r)
	if err != nil {
		return 0, err
	}

	// Get the subject from the token claims
	subject, ok := (*tokenClaims)["sub"].(string)
	if !ok {
		return 0, ErrInvalidSubjectClaim
	}

	// Parse the subject
	return ParseInt64(subject)
}

// GetID returns the ID
func GetID(r *http.Request) (int64, error) {
	// Get the claims from the request
	tokenClaims, err := gojwtnethttpctx.GetCtxTokenClaims(r)
	if err != nil {
		return 0, err
	}

	// Get the ID from the token claims
	id, ok := (*tokenClaims)["jti"].(string)
	if !ok {
		return 0, ErrInvalidIDClaim
	}

	// Parse the ID
	return ParseInt64(id)
}
