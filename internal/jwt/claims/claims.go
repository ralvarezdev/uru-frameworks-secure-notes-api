package claims

import (
	"github.com/golang-jwt/jwt/v5"
	gojwtnethttpctx "github.com/ralvarezdev/go-jwt/net/http/context"
	"net/http"
	"strconv"
	"time"
)

type (
	// AccessTokenClaims is the structure for the JWT claims
	AccessTokenClaims struct {
		ParentRefreshTokenID string `json:"prt,omitempty"`
		jwt.RegisteredClaims
	}

	// RefreshTokenClaims is the structure for the JWT claims
	RefreshTokenClaims struct {
		jwt.RegisteredClaims
	}
)

// NewAccessTokenClaims creates a new JWT access token claims
func NewAccessTokenClaims(
	id int64,
	subject string,
	issuedAt, expiresAt time.Time,
	parentRefreshTokenID int64,
) *AccessTokenClaims {
	return &AccessTokenClaims{
		ParentRefreshTokenID: strconv.FormatInt(parentRefreshTokenID, 10),
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
) *RefreshTokenClaims {
	return &RefreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.FormatInt(id, 10),
			Subject:   subject,
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
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

// GetParentRefreshTokenID returns the parent refresh token ID
func GetParentRefreshTokenID(r *http.Request) (int64, error) {
	// Get the claims from the request
	tokenClaims, err := gojwtnethttpctx.GetCtxTokenClaims(r)
	if err != nil {
		return 0, err
	}

	// Get the parent refresh token ID from the token claims
	parentRefreshTokenID, ok := (*tokenClaims)["prt"].(string)
	if !ok {
		return 0, ErrInvalidParentRefreshTokenIDClaim
	}

	// Parse the parent refresh token ID
	return ParseInt64(parentRefreshTokenID)
}
