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
		internalpostgresmodel.CreateUserNotes,
		internalpostgresmodel.CreateUserTags,
		internalpostgresmodel.CreateUserNoteTags,
		internalpostgresmodel.CreateUserNoteVersions,
		internalpostgresmodel.CreateGetUserRefreshTokenByIDFn,
		internalpostgresmodel.CreateListUserRefreshTokensFn,
		internalpostgresmodel.CreateListUserTokensFn,
		internalpostgresmodel.CreateListUserTagsFn,
		internalpostgresmodel.CreateListUserNoteVersionsFn,
		internalpostgresmodel.CreateListUserNoteTagsFn,
		internalpostgresmodel.CreateSyncUserTagsByLastSyncedAtFn,
		internalpostgresmodel.CreateSyncUserNoteVersionsByLastSyncedAtFn,
		internalpostgresmodel.CreateSyncUserNotesByLastSyncedAtFn,
		internalpostgresmodel.CreateSyncUserNoteTagsByLastSyncedAtFn,
		internalpostgresmodel.CreateGetUserTagByIDFn,
		internalpostgresmodel.CreateGetUserNoteVersionByIDFn,
		internalpostgresmodel.CreateGetUserNoteByIDFn,
		internalpostgresmodel.CreateGetLogInInformationFn,
		internalpostgresmodel.CreateGetUserEmailIDProc,
		internalpostgresmodel.CreateRevokeUserEmailVerificationTokenProc,
		internalpostgresmodel.CreateSendEmailVerificationTokenProc,
		internalpostgresmodel.CreateSignUpProc,
		internalpostgresmodel.CreateRevokeUserTOTPProc,
		internalpostgresmodel.CreateGenerateUserTokensProc,
		internalpostgresmodel.CreateRevokeUserTokensByIDProc,
		internalpostgresmodel.CreateRefreshTokenProc,
		internalpostgresmodel.CreateRevokeUserTokensProc,
		internalpostgresmodel.CreateGetUserAccessTokenIDByUserRefreshTokenIDProc,
		internalpostgresmodel.CreateRegisterFailedLogInAttemptProc,
		internalpostgresmodel.CreateGetUserTOTPProc,
		internalpostgresmodel.CreateGetUserEmailProc,
		internalpostgresmodel.CreateGenerateTOTPUrlProc,
		internalpostgresmodel.CreateIsRefreshTokenValidProc,
		internalpostgresmodel.CreateIsAccessTokenValidProc,
		internalpostgresmodel.CreateRevokeUserTOTPRecoveryCodeProc,
		internalpostgresmodel.CreateCreateUserTOTPRecoveryCodesProc,
		internalpostgresmodel.CreateVerifyTOTPProc,
		internalpostgresmodel.CreateVerifyEmailProc,
		internalpostgresmodel.CreateIsUserEmailVerifiedProc,
		internalpostgresmodel.CreateRevokeUserEmailProc,
		internalpostgresmodel.CreateChangeEmailProc,
		internalpostgresmodel.CreateForgotPasswordProc,
		internalpostgresmodel.CreateRevokeUserResetPasswordTokenProc,
		internalpostgresmodel.CreateResetPasswordProc,
		internalpostgresmodel.CreateRevokeUserPasswordHashProc,
		internalpostgresmodel.CreateChangePasswordProc,
		internalpostgresmodel.CreateChangePasswordProc,
		internalpostgresmodel.CreateGetUserPasswordHashProc,
		internalpostgresmodel.CreateRevokeUserUsernameProc,
		internalpostgresmodel.CreateRevokeUserPhoneNumberProc,
		internalpostgresmodel.CreateDeleteUserProc,
		internalpostgresmodel.CreateChangeUsernameProc,
		internalpostgresmodel.CreateGetUserBasicInfoProc,
		internalpostgresmodel.CreateUpdateProfileProc,
		internalpostgresmodel.CreatePreSendEmailVerificationTokenProc,
		internalpostgresmodel.CreateGetUserPhoneNumberProc,
		internalpostgresmodel.CreateGetUserUsernameProc,
		internalpostgresmodel.CreateHasUserTOTPEnabledProc,
		internalpostgresmodel.CreateIsUserPhoneNumberVerifiedProc,
		internalpostgresmodel.CreateGetMyProfileProc,
		internalpostgresmodel.CreateCreateUserTagProc,
		internalpostgresmodel.CreateUpdateUserTagProc,
		internalpostgresmodel.CreateDeleteUserTagProc,
		internalpostgresmodel.CreateUpdateUserNoteArchiveProc,
		internalpostgresmodel.CreateUpdateUserNoteTrashProc,
		internalpostgresmodel.CreateUpdateUserNoteStarProc,
		internalpostgresmodel.CreateUpdateUserNotePinProc,
		internalpostgresmodel.CreateCreateUserNoteVersionProc,
		internalpostgresmodel.CreateDeleteUserNoteVersionProc,
		internalpostgresmodel.CreateValidateUserTagsIDProc,
		internalpostgresmodel.CreateAddUserNoteTagsProc,
		internalpostgresmodel.CreateCreateUserNoteProc,
		internalpostgresmodel.CreateDeleteUserNoteProc,
		internalpostgresmodel.CreateRemoveUserNoteTagsProc,
		internalpostgresmodel.CreateUpdateUserNoteProc,
		internalpostgresmodel.CreateListUserNotesProc,
	)
}
