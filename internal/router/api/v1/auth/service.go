package auth

import (
	"database/sql"
	"errors"
	gocryptobcrypt "github.com/ralvarezdev/go-crypto/bcrypt"
	gocryptototp "github.com/ralvarezdev/go-crypto/otp/totp"
	gocryptorandomutf8 "github.com/ralvarezdev/go-crypto/random/strings/utf8"
	godatabasessql "github.com/ralvarezdev/go-databases/sql"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gonethttp "github.com/ralvarezdev/go-net/http"
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
	internalbcrypt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/bcrypt"
	internaltotp "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/otp/totp"
	internalpbkdf2 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/pbkdf2"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresconstraints "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/constraints"
	internalpostgresqueries "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/queries"
	internaljwt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt"
	internaljwtcache "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/cache"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/_common"
	"net/http"
	"strconv"
	"time"
)

type (
	// service is the structure for the API V1 service for the auth route group
	service struct {
		gonethttpfactory.ServiceWrapper
	}

	// TokenInfo struct for the cache
	TokenInfo struct {
		Type      gojwttoken.Token
		ID        int64
		ExpiresAt time.Time
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
) error {
	// Insert the failed login attempt
	_, err := internalpostgres.DB.Exec(
		internalpostgresqueries.RegisterFailedLoginAttemptProc,
		userID,
		ipAddress,
		badPassword,
		bad2FACode,
	)
	return err
}

// ValidatePassword validates a password
func (s *service) ValidatePassword(
	userID int64,
	hash, password, ipAddress string,
) (bool, error) {
	// Check if the password is correct
	if gocryptobcrypt.CompareHashAndPassword(
		hash,
		password,
	) {
		return true, nil
	}

	// Register the failed login attempt
	if err := s.RegisterFailedLoginAttempt(
		userID,
		ipAddress,
		true,
		false,
	); err != nil {
		return false, err
	}
	return false, nil
}

// ValidateTOTPCode validates a TOTP code
func (s *service) ValidateTOTPCode(
	userID int64,
	totpSecret,
	totpCode,
	ipAddress string,
	time time.Time,
) (bool, error) {
	match, err := gocryptototp.CompareTOTPSha1(
		totpCode,
		totpSecret,
		time,
		uint64(internaltotp.Period),
		internaltotp.Digits,
	)
	if match {
		return true, nil
	}

	// Register the failed login attempt
	_ = s.RegisterFailedLoginAttempt(
		userID,
		ipAddress,
		false,
		true,
	)
	if err != nil {
		return false, err
	}
	return false, nil
}

// ValidateTOTPRecoveryCode validates a TOTP recovery code
func (s *service) ValidateTOTPRecoveryCode(
	userID int64,
	totpID int64,
	totpRecoveryCode string,
	ipAddress string,
) (bool, error) {
	// Revoke the TOTP recovery code
	result, err := internalpostgres.DB.Exec(
		internalpostgresqueries.RevokeTOTPRecoveryCodeProc,
		totpID,
		totpRecoveryCode,
	)
	if err != nil {
		return false, err
	}

	// Check if the TOTP recovery code was revoked
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAffected > 0 {
		return true, nil
	}

	// Register the failed login attempt
	_ = s.RegisterFailedLoginAttempt(
		userID,
		ipAddress,
		false,
		true,
	)
	return false, nil
}

// GenerateTokensInfo generates the user tokens info
func (s *service) GenerateTokensInfo() (*TokenInfo, *TokenInfo) {
	// Get the current time
	currentTime := time.Now().UTC()

	// Create the user tokens info
	userRefreshTokenInfo := TokenInfo{
		Type:      gojwttoken.RefreshToken,
		ExpiresAt: currentTime.Add(internaljwt.Durations[gojwttoken.RefreshToken]),
	}
	userAccessTokenInfo := TokenInfo{
		Type:      gojwttoken.AccessToken,
		ExpiresAt: currentTime.Add(internaljwt.Durations[gojwttoken.AccessToken]),
	}
	return &userRefreshTokenInfo, &userAccessTokenInfo
}

// GenerateTokens generates user refresh token and user access token
func (s *service) GenerateTokens(
	userID int64,
	userTokensInfo ...*TokenInfo,
) (*map[gojwttoken.Token]string, error) {
	// Set the tokens in the cache as valid
	go func() {
		for _, token := range userTokensInfo {
			s.SetTokenToCache(token.Type, token.ID, token.ExpiresAt, true)
		}
	}()

	// Generate the user tokens
	userTokens := make(map[gojwttoken.Token]string)
	for _, token := range userTokensInfo {
		// Create the user token claims
		userTokenClaims := internaljwtclaims.NewClaims(
			token.Type,
			token.ID,
			strconv.FormatInt(userID, 10),
			time.Now(),
			token.ExpiresAt,
		)

		// Issue the user tokens
		rawToken, err := internaljwt.Issuer.IssueToken(userTokenClaims)
		if err != nil {
			return nil, err
		}
		userTokens[token.Type] = rawToken
	}

	return &userTokens, nil
}

// SignUp signs up a user
func (s *service) SignUp(r *http.Request, body *SignUpRequest) (
	*int64,
	error,
) {
	if body == nil {
		return nil, gonethttp.ErrNilRequestBody
	}

	// Hash the password
	passwordHash, err := gocryptobcrypt.HashPassword(
		body.Password,
		internalbcrypt.Cost,
	)
	if err != nil {
		return nil, err
	}

	// Generate a random salt
	salt, err := gocryptorandomutf8.Generate(internalpbkdf2.SaltLength)
	if err != nil {
		return nil, err
	}

	// Run the SQL function to sign up the user
	var userID sql.NullInt64
	if err = internalpostgres.DB.QueryRow(
		internalpostgresqueries.SignUpProc,
		body.FirstName,
		body.LastName,
		salt,
		body.Username,
		body.Email,
		passwordHash,
		nil,
	).Scan(
		&userID,
	); err != nil {
		isUniqueViolation, constraintName := godatabasessql.IsUniqueViolationError(err)
		if !isUniqueViolation {
			return nil, err
		}
		if constraintName == internalpostgresconstraints.UserEmailsUniqueEmail {
			return nil, ErrSignUpEmailAlreadyRegistered
		}
		if constraintName == internalpostgresconstraints.UserUsernamesUniqueUsername {
			return nil, ErrSignUpUsernameAlreadyRegistered
		}
	}
	return &(userID.Int64), nil
}

// LogIn logs in a user
func (s *service) LogIn(r *http.Request, requestBody *LogInRequest) (
	*int64,
	*map[gojwttoken.Token]string,
	error,
) {
	// Check if the request body is nil
	if requestBody == nil {
		return nil, nil, gonethttp.ErrNilRequestBody
	}

	// Get the current time in UTC
	currentTime := time.Now().UTC()

	// Get the client IP
	clientIP := gonethttp.GetClientIP(r)

	// Get the user ID and password hash by the username, and the user TOTP ID and secret if it is active
	var userID, userTOTPID sql.NullInt64
	var passwordHash, userTOTPSecret sql.NullString
	totpIsActive := true
	if err := internalpostgres.DB.QueryRow(
		internalpostgresqueries.PreLogInProc,
		requestBody.Username,
		nil,
		nil,
		nil,
		nil,
	).Scan(
		&userID,
		&passwordHash,
		&userTOTPID,
		&userTOTPSecret,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, ErrLogInInvalidUsername
		}
		return nil, nil, err
	}

	// Validate the password
	match, err := s.ValidatePassword(
		userID.Int64,
		passwordHash.String,
		requestBody.Password,
		clientIP,
	)
	if err != nil {
		return nil, nil, err
	}
	if !match {
		return nil, nil, ErrLogInInvalidPassword
	}

	// Check if the user TOTP ID is nil
	if userTOTPID.Int64 == 0 {
		totpIsActive = false
	}

	// Validate the TOTP code, if it is active
	if totpIsActive {
		// Check if the TOTP code-related fields are nil
		if requestBody.TOTPCode == nil {
			return nil, nil, ErrLogInRequiredTOTPCode
		}
		if requestBody.IsTOTPRecoveryCode == nil {
			return nil, nil, ErrLogInRequiredIsTOTPRecoveryCode
		}

		if !(*requestBody.IsTOTPRecoveryCode) {
			// Validate the TOTP code
			match, err = s.ValidateTOTPCode(
				userID.Int64,
				userTOTPSecret.String,
				*requestBody.TOTPCode,
				clientIP,
				currentTime,
			)
			if err != nil {
				return nil, nil, err
			}
			if !match {
				return nil, nil, ErrLogInInvalidTOTPCode
			}
		} else {
			// Validate the TOTP recovery code
			match, err = s.ValidateTOTPRecoveryCode(
				userID.Int64,
				userTOTPID.Int64,
				*requestBody.TOTPCode,
				clientIP,
			)
			if err != nil {
				return nil, nil, err
			}
			if !match {
				return nil, nil, ErrLogInInvalidTOTPRecoveryCode
			}
		}
	}

	// Create the user tokens info
	userRefreshTokenInfo, userAccessTokenInfo := s.GenerateTokensInfo()

	// Call the refresh token stored procedure
	var userAccessTokenID, userRefreshTokenID sql.NullInt64
	if err = internalpostgres.DB.QueryRow(
		internalpostgresqueries.GenerateTokensProc,
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
		return nil, nil, err
	}

	// Set the token ID to its respective token info
	userRefreshTokenInfo.ID = userRefreshTokenID.Int64
	userAccessTokenInfo.ID = userAccessTokenID.Int64

	// Generate the user tokens
	userTokens, err := s.GenerateTokens(
		userID.Int64,
		userRefreshTokenInfo,
		userAccessTokenInfo,
	)
	if err != nil {
		return nil, nil, err
	}
	return &(userID.Int64), userTokens, err
}

// RevokeRefreshToken revokes a user refresh token
func (s *service) RevokeRefreshToken(
	r *http.Request,
	userRefreshTokenID int64,
) error {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		return err
	}

	// Set the tokens in the cache as invalid
	go func() {
		// Get the user access token ID by the user refresh token ID
		var userAccessTokenID sql.NullInt64
		if err := internalpostgres.DB.QueryRow(
			internalpostgresqueries.GetAccessTokenByRefreshTokenIDProc,
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
	_, err = internalpostgres.DB.Exec(
		internalpostgresqueries.RevokeTokensByIDProc,
		userID,
		userRefreshTokenID,
	)
	return err
}

// LogOut logs out a user
func (s *service) LogOut(r *http.Request) (*int64, error) {
	// Get the user refresh token ID from the request
	userRefreshTokenID, err := internaljwtclaims.GetID(r)
	if err != nil {
		return nil, err
	}

	// Revoke the user refresh token
	return &userRefreshTokenID, s.RevokeRefreshToken(
		r,
		userRefreshTokenID,
	)
}

// RevokeRefreshTokens revokes all user refresh tokens
func (s *service) RevokeRefreshTokens(r *http.Request) (*int64, error) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		return nil, err
	}

	// Set the tokens in the cache as invalid
	go func() {
		// Get the user refresh tokens and user access tokens ID by user ID
		var userRefreshTokenID, userAccessTokenID int64
		rows, err := internalpostgres.DB.Query(
			internalpostgresqueries.ListUserTokensFn,
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
	_, err = internalpostgres.DB.Exec(
		internalpostgresqueries.RevokeTokensProc,
		userID,
	)
	if err != nil {
		return nil, err
	}
	return &userID, err
}

// RefreshToken refreshes a user token
func (s *service) RefreshToken(r *http.Request) (
	*int64,
	*map[gojwttoken.Token]string,
	error,
) {
	// Get the user ID and the user refresh token ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		return nil, nil, err
	}
	oldUserRefreshTokenID, err := internaljwtclaims.GetID(r)
	if err != nil {
		return nil, nil, err
	}

	// Get the client IP
	clientIP := gonethttp.GetClientIP(r)

	// Create the user tokens info
	userRefreshTokenInfo, userAccessTokenInfo := s.GenerateTokensInfo()

	// Call the refresh token stored procedure
	var userRefreshTokenID, userAccessTokenID sql.NullInt64
	if err = internalpostgres.DB.QueryRow(
		internalpostgresqueries.RefreshTokenProc,
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
		return nil, nil, err
	}

	// Set the token ID to its respective token info
	userRefreshTokenInfo.ID = userRefreshTokenID.Int64
	userAccessTokenInfo.ID = userAccessTokenID.Int64

	// Generate the user tokens
	userTokens, err := s.GenerateTokens(
		userID,
		userRefreshTokenInfo,
		userAccessTokenInfo,
	)
	if err != nil {
		return nil, nil, err
	}
	return &userID, userTokens, err
}

// GenerateTOTPUrl generates a TOTP URL
func (s *service) GenerateTOTPUrl(r *http.Request) (*int64, *string, error) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		return nil, nil, err
	}

	// Generate the TOTP secret
	totpSecret, err := gocryptototp.NewSecret(internaltotp.SecretLength)
	if err != nil {
		return nil, nil, err
	}

	// Call the generate TOTP URL stored procedure
	var userTOTPID sql.NullInt64
	var userEmail, userTOTPSecret sql.NullString
	var userTOTPVerifiedAt sql.NullTime
	if err = internalpostgres.DB.QueryRow(
		internalpostgresqueries.GenerateTOTPUrlProc,
		userID,
		totpSecret,
		nil, nil, nil, nil,
	).Scan(
		&userEmail,
		&userTOTPID,
		&userTOTPSecret,
		&userTOTPVerifiedAt,
	); err != nil {
		return nil, nil, err
	}

	// Check if the TOTP is already verified
	if userTOTPVerifiedAt.Valid {
		return nil, nil, ErrGenerateTOTPUrlAlreadyVerified
	}

	// Generate the TOTP URL
	totpUrl, err := internaltotp.Url.Generate(
		totpSecret,
		userEmail.String,
	)
	if err != nil {
		return nil, nil, err
	}

	return &userID, &totpUrl, nil
}

// VerifyTOTP verifies a TOTP secret
func (s *service) VerifyTOTP(
	r *http.Request,
	requestBody *VerifyTOTPRequest,
) (*int64, *[]string, error) {
	// Check if the request body is nil
	if requestBody == nil {
		return nil, nil, gonethttp.ErrNilRequestBody
	}

	// Get the current time in UTC
	currentTime := time.Now().UTC()

	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)

	// Get the user TOTP ID and secret
	var userTOTPID sql.NullInt64
	var userTOTPSecret sql.NullString
	var userTOTPVerifiedAt sql.NullTime
	if err = internalpostgres.DB.QueryRow(
		internalpostgresqueries.GetUserTOTPProc,
		userID,
		nil, nil, nil,
	).Scan(
		&userTOTPID,
		&userTOTPSecret,
		&userTOTPVerifiedAt,
	); err != nil {
		return nil, nil, err
	}

	// Check if the TOTP is already verified
	if userTOTPVerifiedAt.Valid {
		return nil, nil, ErrVerifyTOTPAlreadyVerified
	}

	// Check if the user TOTP ID is nil
	if !userTOTPID.Valid {
		return nil, nil, ErrVerifyTOTPNotGenerated
	}

	// Verify the TOTP code with the secret
	match, err := gocryptototp.CompareTOTPSha1(
		requestBody.TOTPCode,
		userTOTPSecret.String,
		currentTime,
		uint64(internaltotp.Period),
		internaltotp.Digits,
	)
	if err != nil {
		return nil, nil, err
	}
	if !match {
		return nil, nil, ErrVerifyTOTPInvalidTOTPCode
	}

	// Generate the TOTP recovery codes
	totpRecoveryCodes, err := gocryptototp.GenerateRecoveryCodes(
		internaltotp.RecoveryCodesCount,
		internaltotp.RecoveryCodesLength,
	)
	if err != nil {
		return nil, nil, err
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
	err = internalpostgres.DBService.RunTransaction(
		func(tx *sql.Tx) error {
			// Update the user TOTP verified at
			if _, err = tx.Exec(
				internalpostgresqueries.VerifyTOTPProc,
				userTOTPID,
			); err != nil {
				return err
			}

			// Insert the user TOTP recovery codes
			_, err = tx.Exec(
				internalpostgresqueries.InsertUserTOTPRecoveryCodes,
				insertTOTPRecoveryCodesArgs...,
			)
			return err
		},
	)
	if err != nil {
		return nil, nil, err
	}

	return &userID, totpRecoveryCodes, err
}

// RevokeTOTP revokes a TOTP secret
func (s *service) RevokeTOTP(r *http.Request) (*int64, error) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		return nil, err
	}

	// Run the SQL function to get the user TOTP ID by the user ID
	_, err = internalpostgres.DB.Exec(
		internalpostgresqueries.RevokeTOTPProc,
		userID,
	)
	if err != nil {
		return nil, err
	}
	return &userID, nil
}

// ListRefreshTokens lists all user refresh tokens
func (s *service) ListRefreshTokens(r *http.Request) (
	*int64,
	*[]*internalapiv1common.UserRefreshTokenWithID,
	error,
) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		return nil, nil, err
	}

	// Run the SQL function to list the user refresh tokens by the user ID
	rows, err := internalpostgres.DB.Query(
		internalpostgresqueries.ListUserRefreshTokensFn,
		userID,
	)
	if err != nil {
		return nil, nil, err
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
			return nil, nil, err
		}
		userRefreshTokens = append(userRefreshTokens, &userRefreshToken)
	}

	return &userID, &userRefreshTokens, nil
}

// GetRefreshToken gets a user refresh token
func (s *service) GetRefreshToken(
	r *http.Request,
	userRefreshTokenID int64,
) (*int64, *internalapiv1common.UserRefreshToken, error) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		return nil, nil, err
	}

	// Run the SQL function to get the user refresh token by the ID and user ID
	var userRefreshToken internalapiv1common.UserRefreshToken
	if err = internalpostgres.DB.QueryRow(
		internalpostgresqueries.GetUserRefreshTokenByIDFn,
		userRefreshTokenID,
		userID,
	).Scan(
		&userRefreshToken.IssuedAt,
		&userRefreshToken.ExpiresAt,
		&userRefreshToken.IPAddress,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, ErrGetRefreshTokenNotFound
		}
		return nil, nil, err
	}

	return &userID, &userRefreshToken, nil
}
