package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"github.com/google/uuid"
	gocryptoaes "github.com/ralvarezdev/go-crypto/aes"
	gocryptobcrypt "github.com/ralvarezdev/go-crypto/bcrypt"
	gocryptototp "github.com/ralvarezdev/go-crypto/otp/totp"
	gocryptopbdkf2 "github.com/ralvarezdev/go-crypto/pbkdf2"
	gocryptorandomutf8 "github.com/ralvarezdev/go-crypto/random/strings/utf8"
	godatabasespgx "github.com/ralvarezdev/go-databases/sql/pgx"
	gonethttp "github.com/ralvarezdev/go-net/http"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	"github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal"
	internalcookie "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/cookie"
	internalaes "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/aes"
	internalbcrypt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/bcrypt"
	internaltotp "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/otp/totp"
	internalpbkdf2 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/pbkdf2"
	internaltoken "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/token"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internaljwtcache "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/cache"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	internalmailersend "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/mailersend"
	"io"
	"net/http"
	"time"
)

type (
	// service is the structure for the API V1 service for the auth route group
	service struct{}
)

// RegisterFailedLoginAttempt registers a failed login attempt
func (s *service) RegisterFailedLoginAttempt(
	userID int64,
	ipAddress string,
	badPassword, bad2FACode bool,
) {
	// Insert the failed login attempt
	_, err := internalpostgres.PoolService.Exec(
		&internalpostgresmodel.RegisterFailedLogInAttemptProc,
		userID,
		ipAddress,
		badPassword,
		bad2FACode,
	)
	if err != nil {
		panic(err)
	}
}

// ComparePassword compares a password with its hash
func (s *service) ComparePassword(
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
		&internalpostgresmodel.RevokeUserTOTPRecoveryCodeProc,
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

// GenerateEmailVerificationToken generates an email verification token
func (s *service) GenerateEmailVerificationToken() (uuid.UUID, time.Time) {
	return uuid.New(), time.Now().UTC().Add(internaltoken.EmailVerificationTokenDuration)
}

// GenerateResetPasswordToken generates a reset password token
func (s *service) GenerateResetPasswordToken() (uuid.UUID, time.Time) {
	return uuid.New(), time.Now().UTC().Add(internaltoken.ResetPasswordTokenDuration)
}

// SignUp signs up a user
func (s *service) SignUp(body *SignUpRequest) int64 {
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

	// Generate the email verification token and its expiration time
	emailVerificationToken, emailVerificationTokenExpiresAt := s.GenerateEmailVerificationToken()

	// Run the SQL stored procedure to sign up the user
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
		emailVerificationToken,
		emailVerificationTokenExpiresAt,
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

	// Send welcome email
	fullName := body.FirstName + " " + body.LastName
	go internalmailersend.SendWelcomeEmail(
		fullName,
		body.Email,
	)

	// Send email verification
	go internalmailersend.SendVerificationEmail(
		fullName,
		body.Email,
		emailVerificationToken,
	)

	return userID.Int64
}

// LogIn logs in a user
func (s *service) LogIn(
	w http.ResponseWriter,
	r *http.Request,
	body *LogInRequest,
) int64 {
	// Check if the request body is nil
	if body == nil {
		panic(gonethttp.ErrNilRequestBody)
	}

	// Get the current time in UTC
	currentTime := time.Now().UTC()

	// Get the client IP
	clientIP := gonethttp.GetClientIP(r)

	// Get the user ID and password hash by the username
	var userID sql.NullInt64
	var userPasswordHash, userSalt, userEncryptedKey sql.NullString
	var tooManyFailedAttempts sql.NullBool
	rows, err := internalpostgres.PoolService.Query(
		&internalpostgresmodel.GetLogInInformationFn,
		body.Username,
		clientIP,
		internal.MaximumFailedAttemptsCount,
		internal.MaximumFailedAttemptsPeriodSeconds,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Check if the user exists
	if !rows.Next() {
		panic(ErrLogInInvalidUsername)
	}

	// Scan the row
	if err = rows.Scan(
		&userID,
		&userSalt,
		&userEncryptedKey,
		&userPasswordHash,
		&tooManyFailedAttempts,
	); err != nil {
		panic(err)
	}

	// Check if there are too many failed attempts
	if tooManyFailedAttempts.Valid && tooManyFailedAttempts.Bool {
		panic(ErrLogInTooManyFailedAttempts)
	}

	// Compare the password with its hash
	if !s.ComparePassword(
		userID.Int64,
		userPasswordHash.String,
		body.Password,
		clientIP,
	) {
		panic(ErrLogInInvalidPassword)
	}

	// Get the user TOTP ID and secret by the user ID if it is active
	var userTOTPID sql.NullInt64
	var userTOTPSecret sql.NullString
	var userTOTPVerifiedAt sql.NullTime
	if err := internalpostgres.PoolService.QueryRow(
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

	// Check if the user TOTP ID is not nil and the TOTP is verified
	if userTOTPID.Int64 != 0 && userTOTPVerifiedAt.Valid {
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
	userRefreshTokenInfo, userAccessTokenInfo := internalcookie.GenerateTokensInfo()

	// Call the refresh token stored procedure
	var userAccessTokenID, userRefreshTokenID sql.NullInt64
	if err = internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.GenerateUserTokensProc,
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
	if _, err = internalcookie.SetTokensCookies(
		w,
		userID.Int64,
		userRefreshTokenInfo,
		userAccessTokenInfo,
	); err != nil {
		panic(err)
	}

	// Set the user salt, encrypted key and user ID cookies
	internalcookie.SetSaltCookie(w, userSalt.String)
	internalcookie.SetEncryptedKeyCookie(w, userEncryptedKey.String)
	internalcookie.SetUserIDCookie(w, userID.Int64)
	return userID.Int64
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
	go internaljwtcache.RevokeRefreshTokenFromCache(userRefreshTokenID)

	// Call the revoke tokens by ID stored procedure
	_, err = internalpostgres.PoolService.Exec(
		&internalpostgresmodel.RevokeUserTokensByIDProc,
		userID,
		userRefreshTokenID,
	)
	if err != nil {
		panic(err)
	}

	// Clear cookies if the parent refresh token ID is the same as the user refresh token ID
	if parentRefreshTokenID == userRefreshTokenID {
		internalcookie.ClearCookies(w)
	}
}

// LogOut logs out a user
func (s *service) LogOut(w http.ResponseWriter, r *http.Request) int64 {
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

	return userRefreshTokenID
}

// RevokeRefreshTokens revokes all user refresh tokens
func (s *service) RevokeRefreshTokens(
	w http.ResponseWriter,
	r *http.Request,
) int64 {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Set the tokens in the cache as invalid
	go internaljwtcache.RevokeUserRefreshTokensFromCache(userID)

	// Call the revoke tokens stored procedure
	_, err = internalpostgres.PoolService.Exec(
		&internalpostgresmodel.RevokeUserTokensProc,
		userID,
	)
	if err != nil {
		panic(err)
	}

	// Clear cookies
	internalcookie.ClearCookies(w)

	return userID
}

// GenerateTOTPUrl generates a TOTP URL
func (s *service) GenerateTOTPUrl(r *http.Request) (
	int64,
	*GenerateTOTPUrlResponseBody,
) {
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
	return userID, &GenerateTOTPUrlResponseBody{
		BaseJSendSuccessBody: *gonethttpresponse.NewBaseJSendSuccessBody(),
		Data: GenerateTOTPUrlResponseData{
			TOTPUrl: totpUrl,
		},
	}
}

// VerifyTOTP verifies a TOTP secret
func (s *service) VerifyTOTP(
	r *http.Request,
	body *VerifyTOTPRequest,
) (int64, *VerifyTOTPResponseBody) {
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
	userTOTPRecoveryCodes, err := gocryptototp.GenerateRecoveryCodes(
		internaltotp.RecoveryCodesCount,
		internaltotp.RecoveryCodesLength,
	)
	if err != nil {
		panic(err)
	}

	// Verify the TOTP and insert the recovery codes
	if _, err = internalpostgres.PoolService.Exec(
		&internalpostgresmodel.VerifyTOTPProc,
		userTOTPID,
		userTOTPRecoveryCodes,
	); err != nil {
		panic(err)
	}

	return userID, &VerifyTOTPResponseBody{
		BaseJSendSuccessBody: *gonethttpresponse.NewBaseJSendSuccessBody(),
		Data: VerifyTOTPResponseData{
			RecoveryCodes: *userTOTPRecoveryCodes,
		},
	}
}

// RevokeTOTP revokes a TOTP secret
func (s *service) RevokeTOTP(r *http.Request) int64 {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Run the SQL stored procedure to get the user TOTP ID by the user ID
	_, err = internalpostgres.PoolService.Exec(
		&internalpostgresmodel.RevokeUserTOTPProc,
		userID,
	)
	if err != nil {
		panic(err)
	}
	return userID
}

// ListRefreshTokens lists all user refresh tokens
func (s *service) ListRefreshTokens(r *http.Request) (
	int64,
	*ListRefreshTokensResponseBody,
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
	var userRefreshTokens []*internalpostgresmodel.UserRefreshTokenWithID
	for rows.Next() {
		var userRefreshToken internalpostgresmodel.UserRefreshTokenWithID
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

	return userID, &ListRefreshTokensResponseBody{
		BaseJSendSuccessBody: *gonethttpresponse.NewBaseJSendSuccessBody(),
		Data: ListRefreshTokensResponseData{
			RefreshTokens: userRefreshTokens,
		},
	}
}

// GetRefreshToken gets a user refresh token
func (s *service) GetRefreshToken(
	r *http.Request,
	userRefreshTokenID int64,
) (int64, *GetRefreshTokenResponseBody) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Run the SQL function to get the user refresh token by the ID and user ID
	var userRefreshToken internalpostgresmodel.UserRefreshToken
	rows, err := internalpostgres.PoolService.Query(
		&internalpostgresmodel.GetUserRefreshTokenByIDFn,
		userRefreshTokenID,
		userID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Check if the user refresh token exists
	if !rows.Next() {
		panic(ErrGetRefreshTokenNotFound)
	}

	// Scan the row
	if err = rows.Scan(
		&userRefreshToken.IssuedAt,
		&userRefreshToken.ExpiresAt,
		&userRefreshToken.IPAddress,
	); err != nil {
		panic(err)
	}

	return userID, &GetRefreshTokenResponseBody{
		BaseJSendSuccessBody: *gonethttpresponse.NewBaseJSendSuccessBody(),
		Data: GetRefreshTokenResponseData{
			RefreshToken: &userRefreshToken,
		},
	}
}

// VerifyEmail verifies a user email
func (s *service) VerifyEmail(
	emailVerificationToken string,
) int64 {
	// Run the SQL stored procedure to verify the user email by the email verification token
	var userID sql.NullInt64
	var userInvalidToken sql.NullBool
	if err := internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.VerifyEmailProc,
		emailVerificationToken,
		nil,
		nil,
	).Scan(
		&userID,
		&userInvalidToken,
	); err != nil {
		panic(err)
	}

	// Check if the email verification token is invalid
	if !userInvalidToken.Valid || userInvalidToken.Bool {
		panic(ErrVerifyEmailInvalidToken)
	}
	return userID.Int64
}

// SendEmailVerificationToken sends an email verification token
func (s *service) SendEmailVerificationToken(
	r *http.Request,
) int64 {
	// Get the subject from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Run the SQL stored procedure to check if the user email is verified
	var userFirstName, userLastName, userEmail sql.NullString
	var userEmailIsVerified sql.NullBool
	if err = internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.PreSendEmailVerificationTokenProc,
		userID,
		nil,
		nil, nil,
		nil,
	).Scan(
		&userFirstName,
		&userLastName,
		&userEmail,
		&userEmailIsVerified,
	); err != nil {
		panic(err)
	}

	// Check if the user email is already verified
	if userEmailIsVerified.Bool {
		panic(ErrSendEmailVerificationTokenAlreadyVerified)
	}

	// Generate the email verification token and its expiration time
	emailVerificationToken, emailVerificationTokenExpiresAt := s.GenerateEmailVerificationToken()

	// Run the SQL function to send the email verification
	if _, err = internalpostgres.PoolService.Exec(
		&internalpostgresmodel.SendEmailVerificationTokenProc,
		userID,
		emailVerificationToken,
		emailVerificationTokenExpiresAt,
	); err != nil {
		panic(err)
	}

	// Send email verification
	go internalmailersend.SendVerificationEmail(
		userFirstName.String+" "+userLastName.String,
		userEmail.String,
		emailVerificationToken,
	)

	return userID
}

// ChangeEmail changes a user email
func (s *service) ChangeEmail(
	r *http.Request,
	body *ChangeEmailRequest,
) int64 {
	// Check if the request body is nil
	if body == nil {
		panic(gonethttp.ErrNilRequestBody)
	}

	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Generate the email verification token and its expiration time
	emailVerificationToken, emailVerificationTokenExpiresAt := s.GenerateEmailVerificationToken()

	// Run the SQL stored procedure to change the user email
	var userFirstName, userLastName sql.NullString
	if err = internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.ChangeEmailProc,
		userID,
		body.Email,
		emailVerificationToken,
		emailVerificationTokenExpiresAt,
		nil, nil,
	).Scan(
		&userFirstName,
		&userLastName,
	); err != nil {
		isUniqueViolation, constraintName := godatabasespgx.IsUniqueViolationError(err)
		if !isUniqueViolation {
			panic(err)
		}
		if constraintName == internalpostgresmodel.UserEmailsUniqueEmail {
			panic(ErrChangeEmailAlreadyRegistered)
		}
	}

	// Send email verification
	go internalmailersend.SendVerificationEmail(
		userFirstName.String+" "+userLastName.String,
		body.Email,
		emailVerificationToken,
	)

	return userID
}

// ForgotPassword sends a forgot password email
func (s *service) ForgotPassword(
	body *ForgotPasswordRequest,
) int64 {
	// Check if the request body is nil
	if body == nil {
		panic(gonethttp.ErrNilRequestBody)
	}

	// Generate the reset password token and its expiration time
	resetPasswordToken, resetPasswordTokenExpiresAt := s.GenerateResetPasswordToken()

	// Run the SQL stored procedure to send the forgot password email
	var userID sql.NullInt64
	var userFirstName, userLastName sql.NullString
	if err := internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.ForgotPasswordProc,
		body.Email,
		resetPasswordToken,
		resetPasswordTokenExpiresAt,
		nil,
		nil, nil,
		nil,
	).Scan(
		&userID,
		&userFirstName,
		&userLastName,
	); err != nil {
		panic(err)
	}

	// Send reset password email
	go internalmailersend.SendResetPasswordEmail(
		userFirstName.String+" "+userLastName.String,
		body.Email,
		resetPasswordToken,
	)

	return userID.Int64
}

// ResetPassword resets a user password
func (s *service) ResetPassword(
	resetPasswordToken string,
	body *ResetPasswordRequest,
) int64 {
	// Check if the request body is nil
	if body == nil {
		panic(gonethttp.ErrNilRequestBody)
	}

	// Hash the password
	passwordHash, err := gocryptobcrypt.HashPassword(
		body.NewPassword,
		internalbcrypt.Cost,
	)
	if err != nil {
		panic(err)
	}

	// Run the SQL stored procedure to reset the user password
	var userID sql.NullInt64
	var userInvalidToken sql.NullBool
	if err = internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.ResetPasswordProc,
		resetPasswordToken,
		passwordHash,
		nil, nil,
	).Scan(
		&userID,
		&userInvalidToken,
	); err != nil {
		panic(err)
	}

	// Check if the reset password token is invalid
	if userInvalidToken.Bool {
		panic(ErrResetPasswordInvalidToken)
	}

	return userID.Int64
}

// ChangePassword changes a user password
func (s *service) ChangePassword(
	r *http.Request,
	body *ChangePasswordRequest,
) int64 {
	// Check if the request body is nil
	if body == nil {
		panic(gonethttp.ErrNilRequestBody)
	}

	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Get the user password hash
	var userPasswordHash sql.NullString
	if err = internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.GetUserPasswordHashProc,
		userID,
		nil,
	).Scan(
		&userPasswordHash,
	); err != nil {
		panic(err)
	}

	// Validate the old password
	if !gocryptobcrypt.CompareHashAndPassword(
		userPasswordHash.String,
		body.OldPassword,
	) {
		panic(ErrChangePasswordInvalidOldPassword)
	}

	// Check if the new password is the same as the old password
	if gocryptobcrypt.CompareHashAndPassword(
		userPasswordHash.String,
		body.NewPassword,
	) {
		panic(ErrChangePasswordSamePassword)
	}

	// Hash the new password
	newPasswordHash, err := gocryptobcrypt.HashPassword(
		body.NewPassword,
		internalbcrypt.Cost,
	)
	if err != nil {
		panic(err)
	}

	// Run the SQL stored procedure to change the user password
	if _, err = internalpostgres.PoolService.Exec(
		&internalpostgresmodel.ChangePasswordProc,
		userID,
		newPasswordHash,
	); err != nil {
		panic(err)
	}

	return userID
}
