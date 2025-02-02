package cookie

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gonethttpcookie "github.com/ralvarezdev/go-net/http/cookie"
)

var (
	// Secure is the flag that sets the cookie 'secure' field
	Secure = goflagsmode.ModeFlag.IsProd()

	// AccessToken is the cookies attributes for the access token cookie
	AccessToken = &gonethttpcookie.Attributes{
		Name:     gojwttoken.AccessToken.String(),
		HTTPOnly: true,
		Secure:   Secure,
		Path:     "/",
	}

	// RefreshToken is the cookies attributes for the refresh token cookie
	RefreshToken = &gonethttpcookie.Attributes{
		Name:     gojwttoken.RefreshToken.String(),
		HTTPOnly: true,
		Secure:   Secure,
		Path:     "/",
	}

	// Salt is the cookies attributes for the salt cookie
	Salt = &gonethttpcookie.Attributes{
		Name:     "salt",
		HTTPOnly: false,
		Secure:   Secure,
		Path:     "/",
	}

	// EncryptedKey is the cookies attributes for the encrypted key cookie
	EncryptedKey = &gonethttpcookie.Attributes{
		Name:     "encrypted_key",
		HTTPOnly: false,
		Secure:   Secure,
		Path:     "/",
	}
)
