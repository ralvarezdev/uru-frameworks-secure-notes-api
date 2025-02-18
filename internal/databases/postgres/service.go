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
		internalpostgresmodel.CreateUser2FATOTP,
		internalpostgresmodel.CreateUser2FARecoveryCodes,
		internalpostgresmodel.CreateUser2FAEmailCodes,
		internalpostgresmodel.CreateUserNotes,
		internalpostgresmodel.CreateUserTags,
		internalpostgresmodel.CreateUserNoteTags,
		internalpostgresmodel.CreateUserNoteVersions,
		internalpostgresmodel.CreateUser2FA,
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
		internalpostgresmodel.CreateRevokeUser2FATOTPProc,
		internalpostgresmodel.CreateGenerateUserTokensProc,
		internalpostgresmodel.CreateRevokeUserTokensByIDProc,
		internalpostgresmodel.CreateRefreshTokenProc,
		internalpostgresmodel.CreateRevokeUserTokensProc,
		internalpostgresmodel.CreateGetUserAccessTokenIDByUserRefreshTokenIDProc,
		internalpostgresmodel.CreateRegisterFailedLogInAttemptProc,
		internalpostgresmodel.CreateGetUser2FATOTPProc,
		internalpostgresmodel.CreateGetUserEmailProc,
		internalpostgresmodel.CreateGenerate2FATOTPUrlProc,
		internalpostgresmodel.CreateIsRefreshTokenValidProc,
		internalpostgresmodel.CreateIsAccessTokenValidProc,
		internalpostgresmodel.CreateCreateUser2FARecoveryCodesProc,
		internalpostgresmodel.CreateUseUser2FARecoveryCodeProc,
		internalpostgresmodel.CreateRevokeUser2FARecoveryCodesProc,
		internalpostgresmodel.CreateVerify2FATOTPProc,
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
		internalpostgresmodel.CreateHasUser2FAEnabledProc,
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
		internalpostgresmodel.CreateRevokeUser2FAEmailCodeProc,
		internalpostgresmodel.CreateCreateUser2FAEmailCodeProc,
		internalpostgresmodel.CreateUseUser2FAEmailCodeProc,
		internalpostgresmodel.CreateEnableUser2FAProc,
		internalpostgresmodel.CreateDisableUser2FAProc,
		internalpostgresmodel.CreateSendUser2FAEmailCodeProc,
		internalpostgresmodel.CreateHasUser2FATOTPEnabledProc,
		internalpostgresmodel.CreateGetUser2FAMethodsProc,
	)
}
