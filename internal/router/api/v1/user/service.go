package user

import (
	"database/sql"
	gocryptobcrypt "github.com/ralvarezdev/go-crypto/bcrypt"
	godatabasespgx "github.com/ralvarezdev/go-databases/sql/pgx"
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
	internaljwtcache "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/cache"
	internaljwtclaims "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/claims"
	"net/http"
)

type (
	// service is the structure for the API V1 service for the user route group
	service struct{}
)

// DeleteUser deletes a user
func (s *service) DeleteUser(r *http.Request, body *DeleteUserRequest) int64 {
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
		body.Password,
	) {
		panic(ErrDeleteUserInvalidPassword)
	}

	// Revoke the refresh tokens from cache
	internaljwtcache.RevokeUserRefreshTokensFromCache(userID)

	// Delete the user
	if _, err = internalpostgres.PoolService.Exec(
		&internalpostgresmodel.DeleteUserProc,
		userID,
	); err != nil {
		panic(err)
	}
	return userID
}

// ChangeUsername changes the username of the authenticated user
func (s *service) ChangeUsername(
	r *http.Request,
	body *ChangeUsernameRequest,
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

	// Update the username
	if _, err = internalpostgres.PoolService.Exec(
		&internalpostgresmodel.ChangeUsernameProc,
		userID,
		body.Username,
	); err != nil {
		isUniqueViolation, constraintName := godatabasespgx.IsUniqueViolationError(err)
		if !isUniqueViolation {
			panic(err)
		}
		if constraintName == internalpostgresmodel.UserUsernamesUniqueUsername {
			panic(ErrChangeUsernameAlreadyRegistered)
		}
	}
	return userID
}

// UpdateProfile updates the profile of the authenticated user
func (s *service) UpdateProfile(
	r *http.Request,
	body *UpdateProfileRequest,
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

	// Update the profile
	if _, err = internalpostgres.PoolService.Exec(
		&internalpostgresmodel.UpdateProfileProc,
		userID,
		body.FirstName,
		body.LastName,
		body.Birthdate,
	); err != nil {
		panic(err)
	}
	return userID
}

// GetMyProfile gets the profile of the authenticated user
func (s *service) GetMyProfile(r *http.Request) (int64, *GetMyProfileResponse) {
	// Get the user ID from the request
	userID, err := internaljwtclaims.GetSubject(r)
	if err != nil {
		panic(err)
	}

	// Get the user profile
	var (
		firstName       string
		lastName        string
		birthdate       sql.NullTime
		username        string
		email           string
		emailIsVerified bool
		phone           sql.NullString
		phoneIsVerified sql.NullBool
		hasTOTPEnabled  bool
	)
	if err = internalpostgres.PoolService.QueryRow(
		&internalpostgresmodel.GetMyProfileProc,
		userID,
	).Scan(
		&firstName,
		&lastName,
		&birthdate,
		&username,
		&email,
		&emailIsVerified,
		&phone,
		&phoneIsVerified,
		&hasTOTPEnabled,
	); err != nil {
		panic(err)
	}

	// Return the user profile
	return userID, &GetMyProfileResponse{
		FirstName:       firstName,
		LastName:        lastName,
		Birthdate:       &birthdate.Time,
		Username:        username,
		Email:           email,
		EmailIsVerified: emailIsVerified,
		Phone:           &phone.String,
		PhoneIsVerified: &phoneIsVerified.Bool,
		HasTOTPEnabled:  hasTOTPEnabled,
	}
}
