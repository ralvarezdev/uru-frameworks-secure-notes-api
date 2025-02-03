package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	godatabasespgxpool "github.com/ralvarezdev/go-databases/sql/pgxpool"
	internalpostgresmodel "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres/model"
)

type (
	// Service is the Postgres service struct
	Service struct {
		godatabasespgxpool.Service
	}
)

// NewService creates a new Postgres service
func NewService(pool *pgxpool.Pool) (
	*Service,
	error,
) {
	// Create the default service
	service, err := godatabasespgxpool.NewDefaultService(pool)
	if err != nil {
		return nil, err
	}

	// Create the instance
	instance := &Service{
		Service: service,
	}

	// Migrate the database
	/**/
	err = instance.Migrate()
	if err != nil {
		return nil, err
	}
	/**/

	return instance, nil
}

// Migrate migrates the database
func (s *Service) Migrate() error {
	return s.Service.Migrate(
		internalpostgresmodel.CreateUsers,
		internalpostgresmodel.CreateUserUsernames,
		internalpostgresmodel.CreateUserPasswordHashes,
		internalpostgresmodel.CreateUserResetPasswords,
		internalpostgresmodel.CreateUserEmails,
		internalpostgresmodel.CreateUserEmailVerifications,
		internalpostgresmodel.CreateUserPhoneNumbers,
		internalpostgresmodel.CreateUserPhoneNumberVerifications,
		internalpostgresmodel.CreateUserFailedLogInAttempts,
		internalpostgresmodel.CreateUserRefreshTokens,
		internalpostgresmodel.CreateUserAccessTokens,
		internalpostgresmodel.CreateUserTOTPs,
		internalpostgresmodel.CreateUserTOTPRecoveryCodes,
		internalpostgresmodel.CreateNotes,
		internalpostgresmodel.CreateTags,
		internalpostgresmodel.CreateNoteTags,
		internalpostgresmodel.CreateNoteVersions,
		internalpostgresmodel.CreateGetUserRefreshTokenByIDFn,
		internalpostgresmodel.CreateListUserRefreshTokensFn,
		internalpostgresmodel.CreateListUserTokensFn,
		internalpostgresmodel.CreateGetUserEmailIDProc,
		internalpostgresmodel.CreateSendEmailVerificationTokenProc,
		internalpostgresmodel.CreateSignUpProc,
		internalpostgresmodel.CreateRevokeTOTPProc,
		internalpostgresmodel.CreateGenerateTokensProc,
		internalpostgresmodel.CreateRevokeTokensByIDProc,
		internalpostgresmodel.CreateRefreshTokenProc,
		internalpostgresmodel.CreateRevokeTokensProc,
		internalpostgresmodel.CreateGetAccessTokenIDByRefreshTokenIDProc,
		internalpostgresmodel.CreatePreLogInProc,
		internalpostgresmodel.CreateRegisterFailedLogInAttemptProc,
		internalpostgresmodel.CreateGetUserTOTPProc,
		internalpostgresmodel.CreateGetUserEmailProc,
		internalpostgresmodel.CreateGenerateTOTPUrlProc,
		internalpostgresmodel.CreateIsRefreshTokenValidProc,
		internalpostgresmodel.CreateIsAccessTokenValidProc,
		internalpostgresmodel.CreateRevokeTOTPRecoveryCodeProc,
		internalpostgresmodel.CreateVerifyTOTPProc,
		internalpostgresmodel.CreateVerifyEmailProc,
		internalpostgresmodel.CreateIsUserEmailVerifiedProc,
		internalpostgresmodel.CreateRevokeUserEmailProc,
		internalpostgresmodel.CreateChangeEmailProc,
		internalpostgresmodel.CreateForgotPasswordProc,
		internalpostgresmodel.CreateRevokeResetPasswordTokenProc,
		internalpostgresmodel.CreateResetPasswordProc,
		internalpostgresmodel.CreateRevokePasswordHashProc,
	)
}
