package info

import (
	"github.com/golang-jwt/jwt/v5"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gonethttpcookie "github.com/ralvarezdev/go-net/http/cookie"
	"time"
)

type (
	// TokenInfo struct with the token information and the cookie attributes
	TokenInfo struct {
		Type             gojwttoken.Token
		ID               int64
		CookieAttributes *gonethttpcookie.Attributes
		IssuedAt         time.Time
		ExpiresAt        time.Time
		Claims           jwt.Claims
	}
)
