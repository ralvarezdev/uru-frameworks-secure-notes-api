package auth

import (
	"database/sql"
	"errors"
	gocryptobcrypt "github.com/ralvarezdev/go-crypto/bcrypt"
	gocryptototp "github.com/ralvarezdev/go-crypto/otp/totp"
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
		JwtIssuer       gojwtissuer.Issuer
		PostgresService *internalpostgres.Service
	}
)

// InsertUserFailedLogInAttempt inserts a failed login attempt for a user
func (s *Service) InsertUserFailedLogInAttempt(
	userID int64,
	ipAddress string,
	badPassword, bad2FACode bool,
) error {
	// Get the database connection
	db := s.PostgresService.DB()

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
	db := s.PostgresService.DB()

	// Get the TOTP recovery code by the user TOTP ID
	var totpRecoveryCodeID string
	if err := db.QueryRow(
		internalpostgresqueries.SelectUserTOTPRecoveryCodeByCode,
		totpID,
		totpRecoveryCode,
	).Scan(&totpRecoveryCode); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Register the failed login attempt
			_ = s.InsertUserFailedLogInAttempt(
				userID,
				ipAddress,
				false,
				true,
			)
		}
		return false, err
	}

	// Revoke the TOTP recovery code
	if _, err := db.Exec(
		internalpostgresqueries.UpdateUserTOTPRecoveryCodeRevokedAtByID,
		totpRecoveryCodeID,
	); err != nil {
		return false, err
	}

	return true, nil
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
	refreshTokenExpiresAt := time.Add(internaljwt.Durations[gojwttoken.RefreshToken])
	accessTokenExpiresAt := time.Add(internaljwt.Durations[gojwttoken.AccessToken])
	err := s.PostgresService.RunTransaction(
		func(tx *sql.Tx) error {
			// Insert the user refresh token
			if err := tx.QueryRow(
				internalpostgresqueries.InsertUserRefreshToken,
				userID,
				parentUserRefreshTokenID,
				clientIP,
				time,
				refreshTokenExpiresAt,
			).Scan(&userRefreshTokenID); err != nil {
				return err
			}

			// Insert the user access token
			return tx.QueryRow(
				internalpostgresqueries.InsertUserAccessToken,
				userID,
				userRefreshTokenID,
				time,
				accessTokenExpiresAt,
			).Scan(&userAccessTokenID)
		},
	)
	if err != nil {
		return nil, err
	}

	// Create the user tokens claims
	userTokensClaims := make(map[gojwttoken.Token]*internaljwtclaims.Claims)
	userTokensClaims[gojwttoken.RefreshToken] = internaljwtclaims.NewRefreshTokenClaims(
		userRefreshTokenID,
		strconv.FormatInt(userID, 10),
		time,
		refreshTokenExpiresAt,
	)
	userTokensClaims[gojwttoken.AccessToken] = internaljwtclaims.NewAccessTokenClaims(
		userAccessTokenID,
		strconv.FormatInt(userID, 10),
		time,
		accessTokenExpiresAt,
	)

	// Issue the user tokens
	userTokens := make(map[gojwttoken.Token]string)
	for token, claims := range userTokensClaims {
		rawToken, err := s.JwtIssuer.IssueToken(claims)
		if err != nil {
			return nil, err
		}
		userTokens[token] = rawToken
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
	db := s.PostgresService.DB()

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
			return nil, nil, internalapiv1common.UserNotFoundByUsername
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
			return nil, nil, ErrLogInMissingTOTPCode
		}
		if requestBody.IsTOTPRecoveryCode == nil {
			return nil, nil, ErrLogInMissingIsTOTPRecoveryCode
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
	// Run the transaction
	return s.PostgresService.RunTransaction(
		func(tx *sql.Tx) error {
			// Revoke the user refresh token by the ID
			if _, err := tx.Exec(
				internalpostgresqueries.UpdateUserRefreshTokenRevokedAtByID,
				userRefreshTokenID,
			); err != nil {
				return err
			}

			// Revoke the user access token by the user refresh token ID
			_, err := tx.Exec(
				internalpostgresqueries.UpdateUserAccessTokenRevokedAtByUserRefreshTokenID,
				userRefreshTokenID,
			)
			return err
		},
	)
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

	// Run the transaction
	return &userID, s.PostgresService.RunTransaction(
		func(tx *sql.Tx) error {
			// Revoke the user refresh tokens by the user ID
			if _, err = tx.Exec(
				internalpostgresqueries.UpdateUserRefreshTokensRevokedAtByUserID,
				userID,
			); err != nil {
				return err
			}

			// Delete the user access token by the user refresh token ID
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
	db := s.PostgresService.DB()

	// Run transaction
	var userTOTPID int64
	var userEmail, userTOTPSecret string
	err = s.PostgresService.RunTransaction(
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
			).Scan(&userTOTPID, &userTOTPSecret); err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return nil
				}
				return err
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
	db := s.PostgresService.DB()

	// Get the user TOTP ID and secret
	var userTOTPID int64
	var userTOTPSecret string
	if err = db.QueryRow(
		internalpostgresqueries.SelectUserTOTPByUserID,
		userID,
	).Scan(&userTOTPID, &userTOTPSecret); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, internalapiv1common.UserTOTPSecretNotFoundByUserID
		}
		return nil, nil, err
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
	err = s.PostgresService.RunTransaction(
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
