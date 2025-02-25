package cookie

import (
	"database/sql"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gonethttp "github.com/ralvarezdev/go-net/http"
	gonethttpcookie "github.com/ralvarezdev/go-net/http/cookie"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internaljwt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt"
	internaljwtcache "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/cache"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	internaljwtinfo "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/info"
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

	// UserID is the cookies attributes for the user ID cookie
	UserID = &gonethttpcookie.Attributes{
		Name:     "user_id",
		HTTPOnly: false,
		Secure:   Secure,
		Path:     "/",
	}

	// UserPasswordHash is the cookies attributes for the user password hash cookie
	UserPasswordHash = &gonethttpcookie.Attributes{
		Name:     "user_password_hash",
		HTTPOnly: false,
		Secure:   Secure,
		Path:     "/",
	}

	// SyncNotes is the cookies attributes for the sync notes cookie
	SyncNotes = &gonethttpcookie.Attributes{
		Name:     "sync_notes",
		HTTPOnly: true,
		Secure:   Secure,
		Path:     "/",
	}

	// SyncTags is the cookies attributes for the sync tags cookie
	SyncTags = &gonethttpcookie.Attributes{
		Name:     "sync_tags",
		HTTPOnly: true,
		Secure:   Secure,
		Path:     "/",
	}
)

// GenerateTokensInfo generates the user tokens info
func GenerateTokensInfo() (
	*internaljwtinfo.TokenInfo,
	*internaljwtinfo.TokenInfo,
) {
	// Get the current time
	currentTime := time.Now().UTC()

	// Create the user tokens info
	userRefreshTokenInfo := internaljwtinfo.TokenInfo{
		Type:             gojwttoken.RefreshToken,
		CookieAttributes: RefreshToken,
		IssuedAt:         currentTime,
		ExpiresAt:        currentTime.Add(internaljwt.Durations[gojwttoken.RefreshToken]),
	}
	userAccessTokenInfo := internaljwtinfo.TokenInfo{
		Type:             gojwttoken.AccessToken,
		CookieAttributes: AccessToken,
		IssuedAt:         currentTime,
		ExpiresAt:        currentTime.Add(internaljwt.Durations[gojwttoken.AccessToken]),
	}
	return &userRefreshTokenInfo, &userAccessTokenInfo
}

// SetTokensCookies generates user refresh token and user access token cookies
func SetTokensCookies(
	w http.ResponseWriter,
	userID int64,
	userRefreshToken,
	userAccessToken *internaljwtinfo.TokenInfo,
) (*map[gojwttoken.Token]string, error) {
	// Set the tokens in the cache as valid
	go func() {
		for _, token := range []*internaljwtinfo.TokenInfo{
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

	// Issue the user tokens
	var rawTokens = make(map[gojwttoken.Token]string)
	for _, userToken := range []*internaljwtinfo.TokenInfo{
		userRefreshToken,
		userAccessToken,
	} {
		// Issue the user tokens
		rawToken, err := internaljwt.Issuer.IssueToken(userToken.Claims)
		if err != nil {
			return nil, err
		}

		// Set the raw token
		rawTokens[userToken.Type] = rawToken
	}

	// Set the user tokens cookies
	for _, userToken := range []*internaljwtinfo.TokenInfo{
		userRefreshToken,
		userAccessToken,
	} {
		// Set the cookies
		gonethttpcookie.SetCookie(
			w,
			userToken.CookieAttributes,
			rawTokens[userToken.Type],
			userToken.ExpiresAt,
		)
	}
	return &rawTokens, nil
}

// SetSaltCookie sets the salt cookie
func SetSaltCookie(w http.ResponseWriter, salt string) {
	gonethttpcookie.SetCookie(
		w,
		Salt,
		salt,
		time.Now().Add(internaljwt.Durations[gojwttoken.RefreshToken]),
	)
}

// SetEncryptedKeyCookie sets the encrypted key cookie
func SetEncryptedKeyCookie(w http.ResponseWriter, encryptedKey string) {
	gonethttpcookie.SetCookie(
		w,
		EncryptedKey,
		encryptedKey,
		time.Now().Add(internaljwt.Durations[gojwttoken.RefreshToken]),
	)
}

// SetSyncNotesCookie sets the sync notes cookie
func SetSyncNotesCookie(w http.ResponseWriter, lastSyncedAt time.Time) {
	gonethttpcookie.SetTimestampCookie(
		w,
		SyncNotes,
		lastSyncedAt,
		time.Now().Add(internaljwt.Durations[gojwttoken.RefreshToken]),
	)
}

// SetSyncTagsCookie sets the sync tags cookie
func SetSyncTagsCookie(w http.ResponseWriter, lastSyncedAt time.Time) {
	gonethttpcookie.SetTimestampCookie(
		w,
		SyncTags,
		lastSyncedAt,
		time.Now().Add(internaljwt.Durations[gojwttoken.RefreshToken]),
	)
}

// GetSaltCookie gets the salt cookie
func GetSaltCookie(r *http.Request) (*string, error) {
	// Get the salt cookie
	cookie, err := r.Cookie(Salt.Name)
	if err != nil {
		return nil, err
	}
	return &cookie.Value, nil
}

// GetEncryptedKeyCookie gets the encrypted key cookie
func GetEncryptedKeyCookie(r *http.Request) (*string, error) {
	// Get the encrypted key cookie
	cookie, err := r.Cookie(EncryptedKey.Name)
	if err != nil {
		return nil, err
	}
	return &cookie.Value, nil
}

// GetSyncNotesCookie gets the sync notes cookie
func GetSyncNotesCookie(r *http.Request) (*time.Time, error) {
	return gonethttpcookie.GetTimestampCookie(r, SyncNotes)
}

// GetSyncTagsCookie gets the sync tags cookie
func GetSyncTagsCookie(r *http.Request) (*time.Time, error) {
	return gonethttpcookie.GetTimestampCookie(r, SyncTags)
}

// SetUserIDCookie sets the user ID cookie
func SetUserIDCookie(w http.ResponseWriter, userID int64) {
	gonethttpcookie.SetCookie(
		w,
		UserID,
		strconv.FormatInt(userID, 10),
		time.Now().Add(internaljwt.Durations[gojwttoken.RefreshToken]),
	)
}

// SetUserPasswordHashCookie sets the user password hash cookie
func SetUserPasswordHashCookie(w http.ResponseWriter, userPasswordHash string) {
	gonethttpcookie.SetCookie(
		w,
		UserPasswordHash,
		userPasswordHash,
		time.Now().Add(internaljwt.Durations[gojwttoken.RefreshToken]),
	)
}

// ClearCookies clears the user cookies
func ClearCookies(w http.ResponseWriter) {
	gonethttpcookie.DeleteCookies(
		w, RefreshToken,
		AccessToken,
		Salt,
		EncryptedKey,
		UserID,
		UserPasswordHash,
		SyncTags,
		SyncNotes,
	)
}

// RenovateCookie creates a new cookie with the same value and a new expiration time
func RenovateCookie(
	w http.ResponseWriter,
	r *http.Request,
	attributes *gonethttpcookie.Attributes,
	expiresAt time.Time,
) error {
	return gonethttpcookie.RenovateCookie(
		w, r, attributes, expiresAt,
		func(
			w http.ResponseWriter,
			attributes *gonethttpcookie.Attributes,
			err error,
		) error {
			// Clear the cookies
			ClearCookies(w)

			// An essential cookie is missing, so the user must log in again
			return gonethttpresponse.NewFailResponseError(
				attributes.Name,
				"cookie not found, please log in again",
				gonethttp.ErrCodeCookieNotFound,
				http.StatusInternalServerError,
			)
		},
	)
}

// RefreshTokenFn function to refresh the user tokens
func RefreshTokenFn(token gojwttoken.Token) func(
	w http.ResponseWriter,
	r *http.Request,
) (
	int64,
	*map[gojwttoken.Token]string,
) {
	return func(w http.ResponseWriter, r *http.Request) (
		int64,
		*map[gojwttoken.Token]string,
	) {

		// Get the user ID and the user refresh token ID from the request
		userID, err := internaljwtclaims.GetSubject(r)
		if err != nil {
			panic(err)
		}

		// Check if the token is the access token
		var oldUserRefreshTokenID int64
		if token == gojwttoken.AccessToken {
			// Get the parent refresh token ID from the access token
			oldUserRefreshTokenID, err = internaljwtclaims.GetParentRefreshTokenID(r)
		} else if token == gojwttoken.RefreshToken {
			// Get the user refresh token ID from the request
			oldUserRefreshTokenID, err = internaljwtclaims.GetID(r)
		}
		if err != nil {
			panic(err)
		}

		// Get the client IP
		clientIP := gonethttp.GetClientIP(r)

		// Create the user tokens info
		userRefreshTokenInfo, userAccessTokenInfo := GenerateTokensInfo()

		// Call the refresh token stored procedure
		var userRefreshTokenID, userAccessTokenID sql.NullInt64
		if err = internalpostgres.PoolService.QueryRow(
			&internalpostgresmodel.RefreshTokenProc,
			userID,
			oldUserRefreshTokenID,
			clientIP,
			userRefreshTokenInfo.ExpiresAt,
			userAccessTokenInfo.ExpiresAt,
			nil, nil,
		).Scan(
			&userRefreshTokenID,
			&userAccessTokenID,
		); err != nil {
			panic(err)
		}

		// Set the token ID to its respective token info
		userRefreshTokenInfo.ID = userRefreshTokenID.Int64
		userAccessTokenInfo.ID = userAccessTokenID.Int64

		// Set the user tokens cookies
		rawTokens, err := SetTokensCookies(
			w,
			userID,
			userRefreshTokenInfo,
			userAccessTokenInfo,
		)
		if err != nil {
			panic(err)
		}

		// Renovate the user salt and encrypted key cookies
		for _, cookie := range []*gonethttpcookie.Attributes{
			Salt,
			EncryptedKey,
			SyncTags,
			SyncNotes,
		} {
			if err = RenovateCookie(
				w,
				r,
				cookie,
				userAccessTokenInfo.ExpiresAt,
			); err != nil {
				panic(err)
			}
		}
		return userID, rawTokens
	}
}
