package cookie

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gonethttp "github.com/ralvarezdev/go-net/http"
	gonethttpcookie "github.com/ralvarezdev/go-net/http/cookie"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internaljwt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt"
	internaljwtcache "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/cache"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	"net/http"
	"strconv"
	"time"
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

// SetTokensCookies generates user refresh token and user access token cookies
func SetTokensCookies(
	w http.ResponseWriter,
	userID int64,
	userRefreshToken,
	userAccessToken *internaljwt.TokenInfo,
) error {
	// Set the tokens in the cache as valid
	go func() {
		for _, token := range []*internaljwt.TokenInfo{
			userRefreshToken,
			userAccessToken,
		} {
			internaljwtcache.SetTokenToCache(
				token.Type,
				token.ID,
				token.ExpiresAt,
				true,
			)
		}
	}()

	// Generate the user tokens claims
	userRefreshToken.Claims = internaljwtclaims.NewRefreshTokenClaims(
		userRefreshToken.ID,
		strconv.FormatInt(userID, 10),
		userRefreshToken.IssuedAt,
		userRefreshToken.ExpiresAt,
	)
	userAccessToken.Claims = internaljwtclaims.NewAccessTokenClaims(
		userAccessToken.ID,
		strconv.FormatInt(userID, 10),
		userAccessToken.IssuedAt,
		userAccessToken.ExpiresAt,
		userRefreshToken.ID,
	)

	// Create the user token claims and set the cookies
	for _, userToken := range []*internaljwt.TokenInfo{
		userRefreshToken,
		userAccessToken,
	} {
		// Issue the user tokens
		rawToken, err := internaljwt.Issuer.IssueToken(userToken.Claims)
		if err != nil {
			return err
		}

		// Set the cookies
		gonethttpcookie.SetCookie(
			w,
			userToken.CookieAttributes,
			rawToken,
			userToken.ExpiresAt,
		)
	}
	return nil
}

// ClearCookies clears the user cookies
func ClearCookies(w http.ResponseWriter) {
	// Remove the cookies
	for _, cookie := range []*gonethttpcookie.Attributes{
		RefreshToken,
		AccessToken,
		Salt,
		EncryptedKey,
	} {
		gonethttpcookie.SetCookie(
			w,
			cookie,
			"",
			time.Now().Add(-time.Hour),
		)
	}
}

// RenovateCookie creates a new cookie with the same value and a new expiration time
func RenovateCookie(
	w http.ResponseWriter,
	r *http.Request,
	cookie *gonethttpcookie.Attributes,
	expiresAt time.Time,
) error {
	cookieValue, err := r.Cookie(cookie.Name)
	if err != nil {
		// Clear the cookies
		ClearCookies(w)

		// An essential cookie is missing, so the user must log in again
		return gonethttpresponse.NewCookieError(
			cookie.Name,
			"cookie not found, please log in again",
			gonethttp.ErrCodeCookieNotFound,
			http.StatusInternalServerError,
		)
	}
	gonethttpcookie.SetCookie(
		w,
		cookie,
		cookieValue.Value,
		expiresAt,
	)
}
