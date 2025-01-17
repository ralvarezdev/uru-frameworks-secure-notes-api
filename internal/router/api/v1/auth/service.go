package auth

import (
	"database/sql"
	"errors"
	gocryptobcrypt "github.com/ralvarezdev/go-crypto/bcrypt"
	gocryptototp "github.com/ralvarezdev/go-crypto/otp/totp"
	gojwtcache "github.com/ralvarezdev/go-jwt/cache"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gonethttp "github.com/ralvarezdev/go-net/http"
	internaltotp "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/otp/totp"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresqueries "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/queries"
	internaljwt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	internalapiv1common "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/_common"
	"net/http"
	"strconv"
	"time"
)

type (
	// Service is the structure for the API V1 service for the auth route group
	Service struct {
		jwtIssuer         gojwtissuer.Issuer
		jwtTokenValidator gojwtcache.TokenValidator
		postgresService   *internalpostgres.Service
	}

	// TokenInfo struct for the cache
	TokenInfo struct {
		Type      gojwttoken.Token
		ID        int64
		ExpiresAt time.Time
	}
)

// SetTokenToCache sets the token to the cache
func (s *Service) SetTokenToCache(
	token gojwttoken.Token,
	id int64,
	expiresAt time.Time,
	isValid bool,
) {
	_ = s.jwtTokenValidator.Set(
		token,
		strconv.FormatInt(id, 10),
		isValid,
		expiresAt,
	)
}

// RevokeTokenFromCache revokes the token from the cache
func (s *Service) RevokeTokenFromCache(
	token gojwttoken.Token,
	id int64,
) {
	_ = s.jwtTokenValidator.Revoke(
		token,
		strconv.FormatInt(id, 10),
	)
}

// InsertUserFailedLogInAttempt inserts a failed login attempt for a user
func (s *Service) InsertUserFailedLogInAttempt(
	userID int64,
	ipAddress string,
	badPassword, bad2FACode bool,
) error {
	// Get the database connection
	db := s.postgresService.DB()

	// Insert the failed login attempt
	_, err := db.Exec(
		internalpostgresqueries.InsertUserFailedLogInAttempt,
		userID,
		ipAddress,
		badPassword,
		bad2FACode,
	)
	return err
}

// ValidatePassword validates a password
func (s *Service) ValidatePassword(
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
	if err := s.InsertUserFailedLogInAttempt(
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
func (s *Service) ValidateTOTPCode(
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
	_ = s.InsertUserFailedLogInAttempt(
		userID,
		ipAddress,
		false,
		true,
	)

	// Register error
	if err != nil {
		return false, err
	}
	return false, nil
}

// ValidateTOTPRecoveryCode validates a TOTP recovery code
func (s *Service) ValidateTOTPRecoveryCode(
	userID int64,
	totpID int64,
	totpRecoveryCode string,
	ipAddress string,
) (bool, error) {
	// Get the database connection
	db := s.postgresService.DB()

	// Revoke the TOTP recovery code
	result, err := db.Exec(
		internalpostgresqueries.UpdateUserTOTPRecoveryCodeRevokedAtByTOTPIDAndCode,
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
	_ = s.InsertUserFailedLogInAttempt(
		userID,
		ipAddress,
		false,
		true,
	)
	return false, nil
}

// GenerateTokens generates user refresh token and user access token
func (s *Service) GenerateTokens(
	userID int64,
	clientIP string,
	time time.Time,
	parentUserRefreshTokenID *int64,
) (*map[gojwttoken.Token]string, error) {
	// Run the transaction
	var userRefreshTokenID, userAccessTokenID int64
	userRefreshTokenInfo := TokenInfo{
		Type:      gojwttoken.RefreshToken,
		ExpiresAt: time.Add(internaljwt.Durations[gojwttoken.RefreshToken]),
	}
	userAccessTokenInfo := TokenInfo{
		Type:      gojwttoken.AccessToken,
		ExpiresAt: time.Add(internaljwt.Durations[gojwttoken.AccessToken]),
	}
	err := s.postgresService.RunTransaction(
		func(tx *sql.Tx) error {
			// Insert the user refresh token
			if err := tx.QueryRow(
				internalpostgresqueries.InsertUserRefreshToken,
				userID,
				parentUserRefreshTokenID,
				clientIP,
				time,
				userRefreshTokenInfo.ExpiresAt,
			).Scan(&userRefreshTokenID); err != nil {
				return err
			}

			// Insert the user access token
			return tx.QueryRow(
				internalpostgresqueries.InsertUserAccessToken,
				userID,
				userRefreshTokenID,
				time,
				userAccessTokenInfo.ExpiresAt,
			).Scan(&userAccessTokenID)
		},
	)
	if err != nil {
		return nil, err
	}

	// Add the tokens ID to the tokens
	userRefreshTokenInfo.ID = userRefreshTokenID
	userAccessTokenInfo.ID = userAccessTokenID

	// Set the tokens in the cache as valid
	go func() {
		for _, token := range []TokenInfo{
			userAccessTokenInfo,
			userRefreshTokenInfo,
		} {
			s.SetTokenToCache(token.Type, token.ID, token.ExpiresAt, true)
		}
	}()

	// Generate the user tokens
	userTokens := make(map[gojwttoken.Token]string)
	for _, token := range []TokenInfo{
		userRefreshTokenInfo,
		userAccessTokenInfo,
	} {
		// Create the user token claims
		userTokenClaims := internaljwtclaims.NewClaims(
			token.Type,
			token.ID,
			strconv.FormatInt(userID, 10),
			time,
			token.ExpiresAt,
		)

		// Issue the user tokens
		rawToken, err := s.jwtIssuer.IssueToken(userTokenClaims)
		if err != nil {
			return nil, err
		}
		userTokens[token.Type] = rawToken
	}

	return &userTokens, nil
}

// LogIn logs in a user
func (s *Service) LogIn(r *http.Request, requestBody *LogInRequest) (
	*int64,
	*map[gojwttoken.Token]string,
	error,
) {
	// Check if the request body is nil
	if requestBody == nil {
		return nil, nil, gonethttp.ErrNilRequestBody
	}

	// Get the database connection
	db := s.postgresService.DB()

	// Get the current time in UTC
	currentTime := time.Now().UTC()

	// Get the client IP
	clientIP := gonethttp.GetClientIP(r)

	// Get the user ID and password hash by the username
	var userID int64
	var passwordHash string
	if err := db.QueryRow(
		internalpostgresqueries.SelectUserIDAndPasswordHashByUsername,
		requestBody.Username,
	).Scan(&userID, &passwordHash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, ErrLogInInvalidUsername
		}
		return nil, nil, err
	}

	// Validate the password
	match, err := s.ValidatePassword(
		userID,
		passwordHash,
		requestBody.Password,
		clientIP,
	)
	if err != nil {
		return nil, nil, err
	}
	if !match {
		return nil, nil, ErrLogInInvalidPassword
	}

	// Get the user TOTP ID and secret
	var userTOTPID int64
	var userTOTPSecret string
	totpIsActive := true
	if err = db.QueryRow(
		internalpostgresqueries.SelectUserTOTPSecretVerifiedByUserID,
		userID,
	).Scan(&userTOTPID, &userTOTPSecret); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, nil, err
		}
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
				userID,
				userTOTPSecret,
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
				userID,
				userTOTPID,
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

	// Generate the user tokens
	userTokens, err := s.GenerateTokens(
		userID,
		clientIP,
		currentTime,
		nil,
	)
	return &userID, userTokens, err
}

// RevokeRefreshToken revokes a user refresh token
func (s *Service) RevokeRefreshToken(
	r *http.Request,
	userRefreshTokenID int64,
) error {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		return err
	}

	// Get the database connection
	db := s.postgresService.DB()

	// Set the tokens in the cache as invalid
	go func() {
		// Get the user access token ID by the user refresh token ID
		var userAccessTokenID int64
		if err := db.QueryRow(
			internalpostgresqueries.SelectUserAccessTokenIDByRefreshTokenID,
			userRefreshTokenID,
		).Scan(&userAccessTokenID); err != nil {
			return
		}

		// Revoke the tokens in the cache
		for token, id := range map[gojwttoken.Token]int64{
			gojwttoken.RefreshToken: userRefreshTokenID,
			gojwttoken.AccessToken:  userAccessTokenID,
		} {
			s.RevokeTokenFromCache(token, id)
		}
	}()

	// Run the transaction
	err = s.postgresService.RunTransaction(
		func(tx *sql.Tx) error {
			// Revoke the user refresh token by the ID and user ID
			if _, err := tx.Exec(
				internalpostgresqueries.UpdateUserRefreshTokenRevokedAtByIDAndUserID,
				userRefreshTokenID,
				userID,
			); err != nil {
				return err
			}

			// Revoke the user access token by the user refresh token ID and user ID
			_, err = tx.Exec(
				internalpostgresqueries.UpdateUserAccessTokenRevokedAtByRefreshTokenIDAndUserID,
				userRefreshTokenID,
				userID,
			)
			return err
		},
	)

	return err
}

// LogOut logs out a user
func (s *Service) LogOut(r *http.Request) (*int64, error) {
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
func (s *Service) RevokeRefreshTokens(r *http.Request) (*int64, error) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		return nil, err
	}

	// Get the database connection
	db := s.postgresService.DB()

	// Set the tokens in the cache as invalid
	go func() {
		// Get the user refresh tokens ID by user ID
		var userRefreshTokenID int64
		rows, err := db.Query(
			internalpostgresqueries.SelectUserRefreshTokensIDByUserID,
			userID,
		)
		if err != nil {
			return
		}
		defer rows.Close()

		// Parse the user refresh tokens ID
		for rows.Next() {
			if err := rows.Scan(&userRefreshTokenID); err != nil {
				return
			}

			// Revoke the user refresh token from the cache
			s.RevokeTokenFromCache(gojwttoken.RefreshToken, userRefreshTokenID)
		}

		// Get the user access tokens ID by user ID
		var userAccessTokenID int64
		rows, err = db.Query(
			internalpostgresqueries.SelectUserAccessTokensIDByUserID,
			userID,
		)
		if err != nil {
			return
		}

		// Parse the user access tokens ID
		for rows.Next() {
			if err := rows.Scan(&userAccessTokenID); err != nil {
				return
			}

			// Revoke the user access token from the cache
			s.RevokeTokenFromCache(gojwttoken.AccessToken, userAccessTokenID)
		}
	}()

	// Run the transaction
	return &userID, s.postgresService.RunTransaction(
		func(tx *sql.Tx) error {
			// Revoke the user refresh tokens by the user ID
			if _, err = tx.Exec(
				internalpostgresqueries.UpdateUserRefreshTokensRevokedAtByUserID,
				userID,
			); err != nil {
				return err
			}

			// Revoke the user access tokens by the user refresh token ID
			_, err = tx.Exec(
				internalpostgresqueries.UpdateUserAccessTokensRevokedAtByUserID,
				userID,
			)
			return err
		},
	)
}

// RefreshToken refreshes a user token
func (s *Service) RefreshToken(r *http.Request) (
	*int64,
	*map[gojwttoken.Token]string,
	error,
) {
	// Get the user ID and the user refresh token ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		return nil, nil, err
	}
	userRefreshTokenID, err := internaljwtclaims.GetID(r)
	if err != nil {
		return nil, nil, err
	}

	// Get the current time in UTC
	currentTime := time.Now().UTC()

	// Get the client IP
	clientIP := gonethttp.GetClientIP(r)

	// Revoke the user refresh token
	err = s.RevokeRefreshToken(
		r,
		userRefreshTokenID,
	)
	if err != nil {
		return nil, nil, err
	}

	// Generate the user tokens
	userTokens, err := s.GenerateTokens(
		userID,
		clientIP,
		currentTime,
		&userRefreshTokenID,
	)
	return &userID, userTokens, err
}

// GenerateTOTPUrl generates a TOTP URL
func (s *Service) GenerateTOTPUrl(r *http.Request) (*int64, *string, error) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		return nil, nil, err
	}

	// Get the database connection
	db := s.postgresService.DB()

	// Run transaction
	var userTOTPID int64
	var userEmail, userTOTPSecret string
	var userTOTPVerifiedAt *time.Time
	err = s.postgresService.RunTransaction(
		func(tx *sql.Tx) error {
			// Get the user email by the user ID
			if err = tx.QueryRow(
				internalpostgresqueries.SelectUserEmailByUserID,
				userID,
			).Scan(&userEmail); err != nil {
				return err
			}

			// Get the user TOTP ID by the user ID
			if err = tx.QueryRow(
				internalpostgresqueries.SelectUserTOTPByUserID,
				userID,
			).Scan(
				&userTOTPID,
				&userTOTPSecret,
				&userTOTPVerifiedAt,
			); err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return nil
				}
				return err
			}

			// Check if the TOTP is already verified
			if userTOTPVerifiedAt != nil {
				return ErrGenerateTOTPUrlAlreadyVerified
			}

			// Revoke the existing TOTP secret
			_, err = tx.Exec(
				internalpostgresqueries.UpdateUserTOTPRevokedAtByID,
				userTOTPID,
			)
			return err
		},
	)
	if err != nil {
		return nil, nil, err
	}

	// Generate the TOTP secret
	totpSecret, err := gocryptototp.NewSecret(internaltotp.SecretLength)
	if err != nil {
		return nil, nil, err
	}

	// Insert the TOTP secret
	if _, err = db.Exec(
		internalpostgresqueries.InsertUserTOTP,
		userID,
		totpSecret,
	); err != nil {
		return nil, nil, err
	}

	// Generate the TOTP URL
	totpUrl, err := internaltotp.Url.Generate(
		totpSecret,
		userEmail,
	)
	if err != nil {
		return nil, nil, err
	}

	return &userID, &totpUrl, nil
}

// VerifyTOTP verifies a TOTP secret
func (s *Service) VerifyTOTP(
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

	// Get the database connection
	db := s.postgresService.DB()

	// Get the user TOTP ID and secret
	var userTOTPID int64
	var userTOTPSecret string
	var userTOTPVerifiedAt *time.Time
	if err = db.QueryRow(
		internalpostgresqueries.SelectUserTOTPByUserID,
		userID,
	).Scan(&userTOTPID, &userTOTPSecret, &userTOTPVerifiedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, ErrVerifyTOTPNotGenerated
		}
		return nil, nil, err
	}

	// Check if the TOTP is already verified
	if userTOTPVerifiedAt != nil {
		return nil, nil, ErrVerifyTOTPAlreadyVerified
	}

	// Verify the TOTP code with the secret
	match, err := gocryptototp.CompareTOTPSha1(
		requestBody.TOTPCode,
		userTOTPSecret,
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
	err = s.postgresService.RunTransaction(
		func(tx *sql.Tx) error {
			// Update the user TOTP verified at
			if _, err = tx.Exec(
				internalpostgresqueries.UpdateUserTOTPVerifiedAtByID,
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
func (s *Service) RevokeTOTP(r *http.Request) (*int64, error) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		return nil, err
	}

	// Run transaction
	err = s.postgresService.RunTransaction(
		func(tx *sql.Tx) error {
			// Revoke the user TOTP by the user ID
			if _, err = tx.Exec(
				internalpostgresqueries.UpdateUserTOTPRevokedAtByUserID,
				userID,
			); err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return internalapiv1common.UserTOTPSecretNotFoundByUserID
				}
				return err
			}

			// Revoke the user TOTP recovery codes by the user ID
			_, err = tx.Exec(
				internalpostgresqueries.UpdateUserTOTPRecoveryCodeRevokedAtByUserID,
				userID,
			)
			return err
		},
	)
	if err != nil {
		return nil, err
	}

	return &userID, nil
}

// ListRefreshTokens lists all user refresh tokens
func (s *Service) ListRefreshTokens(r *http.Request) (
	*int64,
	*[]*internalapiv1common.UserRefreshTokenWithID,
	error,
) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		return nil, nil, err
	}

	// Get the database connection
	db := s.postgresService.DB()

	// Get the user refresh tokens
	rows, err := db.Query(
		internalpostgresqueries.SelectUserRefreshTokensByUserID,
		userID,
	)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	// Parse the user refresh tokens
	userRefreshTokens := make([]*internalapiv1common.UserRefreshTokenWithID, 0)
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
func (s *Service) GetRefreshToken(
	r *http.Request,
	userRefreshTokenID int64,
) (*int64, *internalapiv1common.UserRefreshToken, error) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		return nil, nil, err
	}

	// Get the database connection
	db := s.postgresService.DB()

	// Get the user refresh token
	var userRefreshToken internalapiv1common.UserRefreshToken
	if err = db.QueryRow(
		internalpostgresqueries.SelectUserRefreshTokenByIDAndUserID,
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
