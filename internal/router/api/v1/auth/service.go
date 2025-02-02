package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	gocryptoaes "github.com/ralvarezdev/go-crypto/aes"
	gocryptobcrypt "github.com/ralvarezdev/go-crypto/bcrypt"
	gocryptototp "github.com/ralvarezdev/go-crypto/otp/totp"
	gocryptopbdkf2 "github.com/ralvarezdev/go-crypto/pbkdf2"
	gocryptorandomutf8 "github.com/ralvarezdev/go-crypto/random/strings/utf8"
	godatabasespgx "github.com/ralvarezdev/go-databases/sql/pgx"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gonethttp "github.com/ralvarezdev/go-net/http"
	gonethttpcookie "github.com/ralvarezdev/go-net/http/cookie"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalcookie "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/cookie"
	internalaes "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/aes"
	internalbcrypt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/bcrypt"
	internaltotp "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/otp/totp"
	internalpbkdf2 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/pbkdf2"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internaljwt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt"
	internaljwtcache "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/cache"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/_common"
	"io"
	"net/http"
	"strconv"
	"time"
)

type (
	// service is the structure for the API V1 service for the auth route group
	service struct{}

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

// SetTokenToCache sets the token to the cache
func (s *service) SetTokenToCache(
	token gojwttoken.Token,
	id int64,
	expiresAt time.Time,
	isValid bool,
) {
	_ = internaljwtcache.TokenValidator.Set(
		token,
		strconv.FormatInt(id, 10),
		isValid,
		expiresAt,
	)
}

// RevokeTokenFromCache revokes the token from the cache
func (s *service) RevokeTokenFromCache(
	token gojwttoken.Token,
	id int64,
) {
	_ = internaljwtcache.TokenValidator.Revoke(
		token,
		strconv.FormatInt(id, 10),
	)
}

// RegisterFailedLoginAttempt registers a failed login attempt
func (s *service) RegisterFailedLoginAttempt(
	userID int64,
	ipAddress string,
	badPassword, bad2FACode bool,
) {
	// Insert the failed login attempt
	_, err := internalpostgres.PoolService.Exec(
		&internalpostgresmodel.RegisterFailedLoginAttemptProc,
		userID,
		ipAddress,
		badPassword,
		bad2FACode,
	)
	if err != nil {
		panic(err)
	}
}

// ValidatePassword validates a password
func (s *service) ValidatePassword(
	userID int64,
	hash, password, ipAddress string,
) bool {
	// Check if the password is correct
	if gocryptobcrypt.CompareHashAndPassword(
		hash,
		password,
	) {
		return true
	}

	// Register the failed login attempt
	s.RegisterFailedLoginAttempt(
		userID,
		ipAddress,
		true,
		false,
	)
	return false
}

// ValidateTOTPCode validates a TOTP code
func (s *service) ValidateTOTPCode(
	userID int64,
	totpSecret,
	totpCode,
	ipAddress string,
	time time.Time,
) bool {
	match, err := gocryptototp.CompareTOTPSha1(
		totpCode,
		totpSecret,
		time,
		uint64(internaltotp.Period),
		internaltotp.Digits,
	)
	if err != nil {
		panic(err)
	}
	if match {
		return true
	}

	// Register the failed login attempt
	s.RegisterFailedLoginAttempt(
		userID,
		ipAddress,
		false,
		true,
	)
	return false
}

// ValidateTOTPRecoveryCode validates a TOTP recovery code
func (s *service) ValidateTOTPRecoveryCode(
	userID int64,
	totpID int64,
	totpRecoveryCode string,
	ipAddress string,
) bool {
	// Revoke the TOTP recovery code
	result, err := internalpostgres.PoolService.Exec(
		&internalpostgresmodel.RevokeTOTPRecoveryCodeProc,
		totpID,
		totpRecoveryCode,
	)
	if err != nil {
		panic(err)
	}

	// Check if the TOTP recovery code was revoked
	if result.RowsAffected() > 0 {
		return true
	}

	// Register the failed login attempt
	s.RegisterFailedLoginAttempt(
		userID,
		ipAddress,
		false,
		true,
	)
	return false
}

// GenerateTokensInfo generates the user tokens info
func (s *service) GenerateTokensInfo() (*TokenInfo, *TokenInfo) {
	// Get the current time
	currentTime := time.Now().UTC()

	// Create the user tokens info
	userRefreshTokenInfo := TokenInfo{
		Type:             gojwttoken.RefreshToken,
		CookieAttributes: internalcookie.RefreshToken,
		IssuedAt:         currentTime,
		ExpiresAt:        currentTime.Add(internaljwt.Durations[gojwttoken.RefreshToken]),
	}
	userAccessTokenInfo := TokenInfo{
		Type:             gojwttoken.AccessToken,
		CookieAttributes: internalcookie.AccessToken,
		IssuedAt:         currentTime,
		ExpiresAt:        currentTime.Add(internaljwt.Durations[gojwttoken.AccessToken]),
	}
	return &userRefreshTokenInfo, &userAccessTokenInfo
}

// SetTokensCookies generates user refresh token and user access token cookies
func (s *service) SetTokensCookies(
	w http.ResponseWriter,
	userID int64,
	userRefreshToken,
	userAccessToken *TokenInfo,
) {
	// Set the tokens in the cache as valid
	go func() {
		for _, token := range []*TokenInfo{userRefreshToken, userAccessToken} {
			s.SetTokenToCache(token.Type, token.ID, token.ExpiresAt, true)
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

	for _, userToken := range []*TokenInfo{userRefreshToken, userAccessToken} {
		// Issue the user tokens
		rawToken, err := internaljwt.Issuer.IssueToken(userToken.Claims)
		if err != nil {
			panic(err)
		}

		// Set the cookies
		gonethttpcookie.SetCookie(
			w,
			userToken.CookieAttributes,
			rawToken,
			userToken.ExpiresAt,
		)
	}
}

// RenovateCookie creates a new cookie with the same value and a new expiration time
func (s *service) RenovateCookie(
	w http.ResponseWriter,
	r *http.Request,
	cookie *gonethttpcookie.Attributes,
	expiresAt time.Time,
) {
	cookieValue, err := r.Cookie(cookie.Name)
	if err != nil {
		// Clear the cookies
		s.ClearCookies(w)

		// An essential cookie is missing, so the user must log in again
		panic(
			gonethttpresponse.NewCookieError(
				cookie.Name,
				"cookie not found, please log in again",
				http.StatusInternalServerError,
				gonethttp.ErrCodeCookieNotFound,
			),
		)
	}
	gonethttpcookie.SetCookie(
		w,
		cookie,
		cookieValue.Value,
		expiresAt,
	)
}

// ClearCookies clears the user cookies
func (s *service) ClearCookies(w http.ResponseWriter) {
	// Remove the cookies
	for _, cookie := range []*gonethttpcookie.Attributes{
		internalcookie.RefreshToken,
		internalcookie.AccessToken,
		internalcookie.Salt,
		internalcookie.EncryptedKey,
	} {
		gonethttpcookie.SetCookie(
			w,
			cookie,
			"",
			time.Now().Add(-time.Hour),
		)
	}
}

// SignUp signs up a user
func (s *service) SignUp(body *SignUpRequest) *int64 {
	if body == nil {
		panic(gonethttp.ErrNilRequestBody)
	}

	// Hash the password
	passwordHash, err := gocryptobcrypt.HashPassword(
		body.Password,
		internalbcrypt.Cost,
	)
	if err != nil {
		panic(err)
	}

	// Generate a random salt
	salt, err := gocryptorandomutf8.Generate(internalpbkdf2.SaltLength)
	if err != nil {
		panic(err)
	}

	// Derive the password with the salt
	derivedKey := gocryptopbdkf2.DeriveKey(
		body.Password,
		[]byte(salt),
		internalpbkdf2.Iterations,
		internalpbkdf2.KeyLength,
		sha256.New,
	)

	// Generate a random key for the AES-256 encryption
	key := make([]byte, internalaes.KeySize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}

	// Encrypt the key with the derived key
	encryptedKey, err := gocryptoaes.Encrypt(key, derivedKey)
	if err != nil {
		panic(err)
	}

	// Run the SQL function to sign up the user
	var userID sql.NullInt64
	if err = internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.SignUpProc,
		body.FirstName,
		body.LastName,
		salt,
		encryptedKey,
		body.Username,
		body.Email,
		passwordHash,
		nil,
	).Scan(
		&userID,
	); err != nil {
		isUniqueViolation, constraintName := godatabasespgx.IsUniqueViolationError(err)
		if !isUniqueViolation {
			panic(err)
		}
		if constraintName == internalpostgresmodel.UserEmailsUniqueEmail {
			panic(ErrSignUpEmailAlreadyRegistered)
		}
		if constraintName == internalpostgresmodel.UserUsernamesUniqueUsername {
			panic(ErrSignUpUsernameAlreadyRegistered)
		}
	}

	// Send email verification

	return &(userID.Int64)
}

// LogIn logs in a user
func (s *service) LogIn(
	w http.ResponseWriter,
	r *http.Request,
	body *LogInRequest,
) *int64 {
	// Check if the request body is nil
	if body == nil {
		panic(gonethttp.ErrNilRequestBody)
	}

	// Get the current time in UTC
	currentTime := time.Now().UTC()

	// Get the client IP
	clientIP := gonethttp.GetClientIP(r)

	// Get the user ID and password hash by the username, and the user TOTP ID and secret if it is active
	var userID, userTOTPID sql.NullInt64
	var userPasswordHash, userSalt, userEncryptedKey, userTOTPSecret sql.NullString
	totpIsActive := true
	if err := internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.PreLogInProc,
		body.Username,
		nil,
		nil,
		nil,
		nil,
		nil, nil,
	).Scan(
		&userID,
		&userPasswordHash,
		&userSalt,
		&userEncryptedKey,
		&userTOTPID,
		&userTOTPSecret,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(ErrLogInInvalidUsername)
		}
		panic(err)
	}

	// Validate the password
	if !s.ValidatePassword(
		userID.Int64,
		userPasswordHash.String,
		body.Password,
		clientIP,
	) {
		panic(ErrLogInInvalidPassword)
	}

	// Check if the user TOTP ID is nil
	if userTOTPID.Int64 == 0 {
		totpIsActive = false
	}

	// Validate the TOTP code, if it is active
	if totpIsActive {
		// Check if the TOTP code-related fields are nil
		if body.TOTPCode == nil {
			panic(ErrLogInRequiredTOTPCode)
		}
		if body.IsTOTPRecoveryCode == nil {
			panic(ErrLogInRequiredIsTOTPRecoveryCode)
		}

		// Validate the TOTP code
		if !(*body.IsTOTPRecoveryCode) {
			if !s.ValidateTOTPCode(
				userID.Int64,
				userTOTPSecret.String,
				*body.TOTPCode,
				clientIP,
				currentTime,
			) {
				panic(ErrLogInInvalidTOTPCode)
			}
		} else {
			// Validate the TOTP recovery code
			if !s.ValidateTOTPRecoveryCode(
				userID.Int64,
				userTOTPID.Int64,
				*body.TOTPCode,
				clientIP,
			) {
				panic(ErrLogInInvalidTOTPRecoveryCode)
			}
		}
	}

	// Create the user tokens info
	userRefreshTokenInfo, userAccessTokenInfo := s.GenerateTokensInfo()

	// Call the refresh token stored procedure
	var userAccessTokenID, userRefreshTokenID sql.NullInt64
	if err := internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.GenerateTokensProc,
		userID,
		nil,
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
	s.SetTokensCookies(
		w,
		userID.Int64,
		userRefreshTokenInfo,
		userAccessTokenInfo,
	)

	// Set the user salt and encrypted key cookies
	gonethttpcookie.SetCookie(
		w,
		internalcookie.Salt,
		userSalt.String,
		userRefreshTokenInfo.ExpiresAt,
	)
	gonethttpcookie.SetCookie(
		w,
		internalcookie.EncryptedKey,
		userEncryptedKey.String,
		userRefreshTokenInfo.ExpiresAt,
	)
	return &(userID.Int64)
}

// RevokeRefreshToken revokes a user refresh token
func (s *service) RevokeRefreshToken(
	w http.ResponseWriter,
	r *http.Request,
	userRefreshTokenID int64,
) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}
	// Get the parent refresh token ID from the request
	parentRefreshTokenID, err := internaljwtclaims.GetParentRefreshTokenID(r)
	if err != nil {
		panic(err)
	}

	// Set the tokens in the cache as invalid
	go func() {
		// Get the user access token ID by the user refresh token ID
		var userAccessTokenID sql.NullInt64
		if err := internalpostgres.PoolService.QueryRow(
			&internalpostgresmodel.GetAccessTokenByRefreshTokenIDProc,
			userRefreshTokenID,
			nil,
		).Scan(
			&userAccessTokenID,
		); err != nil {
			return
		}

		// Revoke the tokens in the cache
		for token, id := range map[gojwttoken.Token]int64{
			gojwttoken.RefreshToken: userRefreshTokenID,
			gojwttoken.AccessToken:  userAccessTokenID.Int64,
		} {
			s.RevokeTokenFromCache(token, id)
		}
	}()

	// Call the revoke tokens by ID stored procedure
	_, err = internalpostgres.PoolService.Exec(
		&internalpostgresmodel.RevokeTokensByIDProc,
		userID,
		userRefreshTokenID,
	)
	if err != nil {
		panic(err)
	}

	// Clear cookies if the parent refresh token ID is the same as the user refresh token ID
	if parentRefreshTokenID == userRefreshTokenID {
		s.ClearCookies(w)
	}
}

// LogOut logs out a user
func (s *service) LogOut(w http.ResponseWriter, r *http.Request) *int64 {
	// Get the user refresh token ID from the request
	userRefreshTokenID, err := internaljwtclaims.GetID(r)
	if err != nil {
		panic(err)
	}

	// Revoke the user refresh token
	s.RevokeRefreshToken(
		w,
		r,
		userRefreshTokenID,
	)

	return &userRefreshTokenID
}

// RevokeRefreshTokens revokes all user refresh tokens
func (s *service) RevokeRefreshTokens(
	w http.ResponseWriter,
	r *http.Request,
) *int64 {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Set the tokens in the cache as invalid
	go func() {
		// Get the user refresh tokens and user access tokens ID by user ID
		var userRefreshTokenID, userAccessTokenID int64
		rows, err := internalpostgres.PoolService.Query(
			&internalpostgresmodel.ListUserTokensFn,
			userID,
		)
		if err != nil {
			return
		}
		defer rows.Close()

		// Parse the user refresh tokens ID
		for rows.Next() {
			if err := rows.Scan(
				&userRefreshTokenID,
				&userAccessTokenID,
			); err != nil {
				return
			}

			// Revoke the user refresh token and user access token from the cache
			s.RevokeTokenFromCache(gojwttoken.RefreshToken, userRefreshTokenID)
			s.RevokeTokenFromCache(gojwttoken.AccessToken, userAccessTokenID)
		}
	}()

	// Call the revoke tokens stored procedure
	_, err = internalpostgres.PoolService.Exec(
		&internalpostgresmodel.RevokeTokensProc,
		userID,
	)
	if err != nil {
		panic(err)
	}

	// Clear cookies
	s.ClearCookies(w)

	return &userID
}

// RefreshToken refreshes a user token
func (s *service) RefreshToken(w http.ResponseWriter, r *http.Request) *int64 {
	// Get the user ID and the user refresh token ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}
	oldUserRefreshTokenID, err := internaljwtclaims.GetID(r)
	if err != nil {
		panic(err)
	}

	// Get the client IP
	clientIP := gonethttp.GetClientIP(r)

	// Create the user tokens info
	userRefreshTokenInfo, userAccessTokenInfo := s.GenerateTokensInfo()

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
	s.SetTokensCookies(
		w,
		userID,
		userRefreshTokenInfo,
		userAccessTokenInfo,
	)

	// Renovate the user salt and encrypted key cookies
	for _, cookie := range []*gonethttpcookie.Attributes{
		internalcookie.Salt,
		internalcookie.EncryptedKey,
	} {
		s.RenovateCookie(
			w,
			r,
			cookie,
			userAccessTokenInfo.ExpiresAt,
		)
	}
	return &userID
}

// GenerateTOTPUrl generates a TOTP URL
func (s *service) GenerateTOTPUrl(r *http.Request) (*int64, *string) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Generate the TOTP secret
	totpSecret, err := gocryptototp.NewSecret(internaltotp.SecretLength)
	if err != nil {
		panic(err)
	}

	// Call the generate TOTP URL stored procedure
	var userTOTPID sql.NullInt64
	var userEmail, userTOTPSecret sql.NullString
	var userTOTPVerifiedAt sql.NullTime
	if err = internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.GenerateTOTPUrlProc,
		userID,
		totpSecret,
		nil, nil, nil, nil,
	).Scan(
		&userEmail,
		&userTOTPID,
		&userTOTPSecret,
		&userTOTPVerifiedAt,
	); err != nil {
		panic(err)
	}

	// Check if the TOTP is already verified
	if userTOTPVerifiedAt.Valid {
		panic(ErrGenerateTOTPUrlAlreadyVerified)
	}

	// Generate the TOTP URL
	totpUrl, err := internaltotp.Url.Generate(
		totpSecret,
		userEmail.String,
	)
	if err != nil {
		panic(err)
	}
	return &userID, &totpUrl
}

// VerifyTOTP verifies a TOTP secret
func (s *service) VerifyTOTP(
	r *http.Request,
	body *VerifyTOTPRequest,
) (*int64, *[]string) {
	// Check if the request body is nil
	if body == nil {
		panic(gonethttp.ErrNilRequestBody)
	}

	// Get the current time in UTC
	currentTime := time.Now().UTC()

	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)

	// Get the user TOTP ID and secret
	var userTOTPID sql.NullInt64
	var userTOTPSecret sql.NullString
	var userTOTPVerifiedAt sql.NullTime
	if err = internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.GetUserTOTPProc,
		userID,
		nil, nil, nil,
	).Scan(
		&userTOTPID,
		&userTOTPSecret,
		&userTOTPVerifiedAt,
	); err != nil {
		panic(err)
	}

	// Check if the TOTP is already verified
	if userTOTPVerifiedAt.Valid {
		panic(ErrVerifyTOTPAlreadyVerified)
	}

	// Check if the user TOTP ID is nil
	if !userTOTPID.Valid {
		panic(ErrVerifyTOTPNotGenerated)
	}

	// Verify the TOTP code with the secret
	match, err := gocryptototp.CompareTOTPSha1(
		body.TOTPCode,
		userTOTPSecret.String,
		currentTime,
		uint64(internaltotp.Period),
		internaltotp.Digits,
	)
	if err != nil {
		panic(err)
	}
	if !match {
		panic(ErrVerifyTOTPInvalidTOTPCode)
	}

	// Generate the TOTP recovery codes
	totpRecoveryCodes, err := gocryptototp.GenerateRecoveryCodes(
		internaltotp.RecoveryCodesCount,
		internaltotp.RecoveryCodesLength,
	)
	if err != nil {
		panic(err)
	}

	// Create the insert TOTP recovery codes arguments
	insertTOTPRecoveryCodesArgs := make(
		[]interface{},
		len(*totpRecoveryCodes)+1,
	)
	insertTOTPRecoveryCodesArgs[0] = &userTOTPID
	for i, code := range *totpRecoveryCodes {
		insertTOTPRecoveryCodesArgs[i+1] = &code
	}

	// Run transaction
	err = internalpostgres.PoolService.CreateTransaction(
		func(ctx context.Context, tx pgx.Tx) error {
			// Update the user TOTP verified at
			if _, err = tx.Exec(
				ctx,
				internalpostgresmodel.VerifyTOTPProc,
				userTOTPID,
			); err != nil {
				return err
			}

			// Insert the user TOTP recovery codes
			_, err = tx.Exec(
				ctx,
				internalpostgresmodel.InsertUserTOTPRecoveryCodes,
				insertTOTPRecoveryCodesArgs...,
			)
			return err
		},
	)
	if err != nil {
		panic(err)
	}

	return &userID, totpRecoveryCodes
}

// RevokeTOTP revokes a TOTP secret
func (s *service) RevokeTOTP(r *http.Request) *int64 {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Run the SQL function to get the user TOTP ID by the user ID
	_, err = internalpostgres.PoolService.Exec(
		&internalpostgresmodel.RevokeTOTPProc,
		userID,
	)
	if err != nil {
		panic(err)
	}
	return &userID
}

// ListRefreshTokens lists all user refresh tokens
func (s *service) ListRefreshTokens(r *http.Request) (
	*int64,
	*[]*internalapiv1common.UserRefreshTokenWithID,
) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Run the SQL function to list the user refresh tokens by the user ID
	rows, err := internalpostgres.PoolService.Query(
		&internalpostgresmodel.ListUserRefreshTokensFn,
		userID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Parse the user refresh tokens
	var userRefreshTokens []*internalapiv1common.UserRefreshTokenWithID
	for rows.Next() {
		var userRefreshToken internalapiv1common.UserRefreshTokenWithID
		if err = rows.Scan(
			&userRefreshToken.ID,
			&userRefreshToken.IssuedAt,
			&userRefreshToken.ExpiresAt,
			&userRefreshToken.IPAddress,
		); err != nil {
			panic(err)
		}
		userRefreshTokens = append(userRefreshTokens, &userRefreshToken)
	}

	return &userID, &userRefreshTokens
}

// GetRefreshToken gets a user refresh token
func (s *service) GetRefreshToken(
	r *http.Request,
	userRefreshTokenID int64,
) (*int64, *internalapiv1common.UserRefreshToken) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Run the SQL function to get the user refresh token by the ID and user ID
	var userRefreshToken internalapiv1common.UserRefreshToken
	if err = internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.GetUserRefreshTokenByIDFn,
		userRefreshTokenID,
		userID,
	).Scan(
		&userRefreshToken.IssuedAt,
		&userRefreshToken.ExpiresAt,
		&userRefreshToken.IPAddress,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(ErrGetRefreshTokenNotFound)
		}
		panic(err)
	}

	return &userID, &userRefreshToken
}
