package logger

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	gologgermode "github.com/ralvarezdev/go-logger/mode"
	gologgermodenamed "github.com/ralvarezdev/go-logger/mode/named"
)

// Logger is the logger for the API server
type Logger struct {
	logger gologgermodenamed.Logger
}

// NewLogger is the logger for the API server
func NewLogger(header string, modeLogger gologgermode.Logger) (*Logger, error) {
	// Initialize the mode named logger
	namedLogger, err := gologgermodenamed.NewDefaultLogger(header, modeLogger)
	if err != nil {
		return nil, err
	}

	return &Logger{logger: namedLogger}, nil
}

// ServerStarted logs a success message when the server starts
func (l *Logger) ServerStarted(port string) {
	l.logger.Info(
		"server started",
		"port: "+port,
	)
}

// PoolStat logs the pool stat
func (l *Logger) PoolStat(
	stat *pgxpool.Stat,
) {
	l.logger.Info(
		"pool stat",
		fmt.Sprintf("total connections: %d", stat.TotalConns()),
		fmt.Sprintf("new conns count: %d", stat.NewConnsCount()),
		fmt.Sprintf("acquire count: %d", stat.AcquireCount()),
		fmt.Sprintf("acquired connections: %d", stat.AcquiredConns()),
		fmt.Sprintf("max connections: %d", stat.MaxConns()),
		fmt.Sprintf("idle connections: %d", stat.IdleConns()),
		fmt.Sprintf("constructing conns: %d", stat.ConstructingConns()),
		fmt.Sprintf("max idle destroy count: %d", stat.MaxIdleDestroyCount()),
		fmt.Sprintf(
			"max lifetime destroy count: %d",
			stat.MaxLifetimeDestroyCount(),
		),
	)
}

// LogIn logs the log-in event
func (l *Logger) LogIn(id int64) {
	l.logger.Info(
		"user logged in",
		fmt.Sprintf("user id: %d", id),
	)
}

// RevokeRefreshToken logs the revoke refresh token event
func (l *Logger) RevokeRefreshToken(id int64) {
	l.logger.Info(
		"user revoked refresh token",
		fmt.Sprintf("user refresh token id: %d", id),
	)
}

// LogOut logs the log-out event
func (l *Logger) LogOut(id int64) {
	l.logger.Info(
		"user logged out",
		fmt.Sprintf("user refresh token id: %d", id),
	)
}

// RevokeRefreshTokens logs the revoke refresh tokens event
func (l *Logger) RevokeRefreshTokens(id int64) {
	l.logger.Info(
		"user revoked all refresh tokens",
		fmt.Sprintf("user id: %d", id),
	)
}

// RefreshToken logs the refresh token event
func (l *Logger) RefreshToken(id int64) {
	l.logger.Info(
		"user refreshed token",
		fmt.Sprintf("user id: %d", id),
	)
}

// SignUp logs the sign-up event
func (l *Logger) SignUp(id int64) {
	l.logger.Info(
		"user signed up",
		fmt.Sprintf("user id: %d", id),
	)
}

// GenerateTOTPUrl logs the generate TOTP URL event
func (l *Logger) GenerateTOTPUrl(id int64) {
	l.logger.Info(
		"user generated totp url",
		fmt.Sprintf("user id: %d", id),
	)
}

// VerifyTOTP logs the verify TOTP event
func (l *Logger) VerifyTOTP(id int64) {
	l.logger.Info(
		"user verified totp",
		fmt.Sprintf("user id: %d", id),
	)
}

// ListRefreshTokens logs the list refresh tokens event
func (l *Logger) ListRefreshTokens(id int64) {
	l.logger.Info(
		"user listed refresh tokens",
		fmt.Sprintf("user id: %d", id),
	)
}

// GetRefreshToken logs the get refresh token event
func (l *Logger) GetRefreshToken(id, refreshTokenID int64) {
	l.logger.Info(
		"user got refresh token",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("refresh token id: %d", refreshTokenID),
	)
}

// RevokeTOTP logs the revoke TOTP event
func (l *Logger) RevokeTOTP(id int64) {
	l.logger.Info(
		"user revoked totp",
		fmt.Sprintf("user id: %d", id),
	)
}

// SendEmailVerificationToken logs the send email verification token event
func (l *Logger) SendEmailVerificationToken(id int64) {
	l.logger.Info(
		"user requested email verification token",
		fmt.Sprintf("user id: %d", id),
	)
}

// SentVerificationEmail logs that the verification email was sent successfully
func (l *Logger) SentVerificationEmail(email string) {
	l.logger.Info(
		"sent verification email to: " + email,
	)
}

// FailedToSendVerificationEmail logs the failed to send verification email event
func (l *Logger) FailedToSendVerificationEmail(email string, err error) {
	l.logger.Error(
		"failed to send verification email to: "+email,
		err,
	)
}

// VerifyEmail logs the verify email event
func (l *Logger) VerifyEmail(id int64) {
	l.logger.Info(
		"user verified email",
		fmt.Sprintf("user id: %d", id),
	)
}

// FailedToSendResetPasswordEmail logs the failed to send reset password email event
func (l *Logger) FailedToSendResetPasswordEmail(email string, err error) {
	l.logger.Error(
		"failed to send reset password email to: "+email,
		err,
	)
}

// SentResetPasswordEmail logs that the reset password email was sent successfully
func (l *Logger) SentResetPasswordEmail(email string) {
	l.logger.Info(
		"sent reset password email to: " + email,
	)
}

// FailedToSendWelcomeEmail logs the failed to send welcome email event
func (l *Logger) FailedToSendWelcomeEmail(email string, err error) {
	l.logger.Error(
		"failed to send welcome email to: "+email,
		err,
	)
}

// SentWelcomeEmail logs that the welcome email was sent successfully
func (l *Logger) SentWelcomeEmail(email string) {
	l.logger.Info(
		"sent welcome email to: " + email,
	)
}

// ForgotPassword logs the forgot password event
func (l *Logger) ForgotPassword(id int64) {
	l.logger.Info(
		"user forgot password",
		fmt.Sprintf("user id: %d", id),
	)
}

// ResetPassword logs the reset password event
func (l *Logger) ResetPassword(id int64) {
	l.logger.Info(
		"user reset password",
		fmt.Sprintf("user id: %d", id),
	)
}

// ChangePassword logs the change password event
func (l *Logger) ChangePassword(id int64) {
	l.logger.Info(
		"user changed password",
		fmt.Sprintf("user id: %d", id),
	)
}

// DeleteUser logs the delete user event
func (l *Logger) DeleteUser(id int64) {
	l.logger.Info(
		"user deleted",
		fmt.Sprintf("user id: %d", id),
	)
}

// ChangeUsername logs the change username event
func (l *Logger) ChangeUsername(id int64, newUsername string) {
	l.logger.Info(
		"user changed username",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("new username: %s", newUsername),
	)
}

// ChangeEmail logs the change email event
func (l *Logger) ChangeEmail(id int64, newEmail string) {
	l.logger.Info(
		"user changed email",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("new email: %s", newEmail),
	)
}

// UpdateProfile logs the update profile event
func (l *Logger) UpdateProfile(id int64) {
	l.logger.Info(
		"user updated profile",
		fmt.Sprintf("user id: %d", id),
	)
}

// GetMyProfile logs the get my profile event
func (l *Logger) GetMyProfile(id int64) {
	l.logger.Info(
		"user got his/her profile",
		fmt.Sprintf("user id: %d", id),
	)
}

// ListUserTags logs the list user tags event
func (l *Logger) ListUserTags(id int64) {
	l.logger.Info(
		"user listed tags",
		fmt.Sprintf("user id: %d", id),
	)
}

// CreateUserTag logs the user tag creation event
func (l *Logger) CreateUserTag(id, tagID int64) {
	l.logger.Info(
		"user created tag",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("tag id: %d", tagID),
	)
}

// UpdateUserTag logs the update user tag event
func (l *Logger) UpdateUserTag(id, tagID int64) {
	l.logger.Info(
		"user updated tag",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("tag id: %d", tagID),
	)
}

// DeleteUserTag logs the delete user tag event
func (l *Logger) DeleteUserTag(id, tagID int64) {
	l.logger.Info(
		"user deleted tag",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("tag id: %d", tagID),
	)
}

// GetUserTagByID logs the get user tag by ID event
func (l *Logger) GetUserTagByID(id, tagID int64) {
	l.logger.Info(
		"user got tag",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("tag id: %d", tagID),
	)
}

// UpdateUserNoteStar logs the update user note star event
func (l *Logger) UpdateUserNoteStar(id, noteID int64, star bool) {
	l.logger.Info(
		"user updated note star",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note id: %d", noteID),
		fmt.Sprintf("star: %t", star),
	)
}

// UpdateUserNoteTrash logs the update user note trash event
func (l *Logger) UpdateUserNoteTrash(id, noteID int64, trash bool) {
	l.logger.Info(
		"user updated note trash",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note id: %d", noteID),
		fmt.Sprintf("trash: %t", trash),
	)
}

// UpdateUserNoteArchive logs the update user note archive event
func (l *Logger) UpdateUserNoteArchive(id, noteID int64, archive bool) {
	l.logger.Info(
		"user updated note archive",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note id: %d", noteID),
		fmt.Sprintf("archive: %t", archive),
	)
}

// UpdateUserNotePin logs the update user note pin event
func (l *Logger) UpdateUserNotePin(id, noteID int64, pin bool) {
	l.logger.Info(
		"user updated note pin",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note id: %d", noteID),
		fmt.Sprintf("pin: %t", pin),
	)
}

// CreateUserNoteVersion logs the user note version creation event
func (l *Logger) CreateUserNoteVersion(id, noteID, noteVersionID int64) {
	l.logger.Info(
		"user created note version",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note id: %d", noteID),
		fmt.Sprintf("note version id: %d", noteVersionID),
	)
}

// DeleteUserNoteVersion logs the delete user note version event
func (l *Logger) DeleteUserNoteVersion(id, noteVersionID int64) {
	l.logger.Info(
		"user deleted note version",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note version id: %d", noteVersionID),
	)
}

// GetUserNoteVersionByID logs the get user note version by ID event
func (l *Logger) GetUserNoteVersionByID(id, noteVersionID int64) {
	l.logger.Info(
		"user got note version",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note version id: %d", noteVersionID),
	)
}

// ListUserNoteVersions logs the list user note versions event
func (l *Logger) ListUserNoteVersions(id, noteID int64) {
	l.logger.Info(
		"user listed note versions",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note id: %d", noteID),
	)
}

// SyncUserNoteVersions logs the sync user note versions event
func (l *Logger) SyncUserNoteVersions(id, noteID, latestNoteVersionID int64) {
	l.logger.Info(
		"user synced note versions",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note id: %d", noteID),
		fmt.Sprintf("latest note version id: %d", latestNoteVersionID),
	)
}

// ListUserNoteTags logs the list user note tags event
func (l *Logger) ListUserNoteTags(id, noteID int64) {
	l.logger.Info(
		"user listed note tags",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note id: %d", noteID),
	)
}

// AddUserNoteTags logs the add user note tags event
func (l *Logger) AddUserNoteTags(id, noteID int64) {
	l.logger.Info(
		"user added note tag",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note id: %d", noteID),
	)
}

// RemoveUserNoteTags logs the remove user note tags event
func (l *Logger) RemoveUserNoteTags(id, noteID int64) {
	l.logger.Info(
		"user removed note tag",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note id: %d", noteID),
	)
}

// GetUserNoteByID logs the get user note by ID event
func (l *Logger) GetUserNoteByID(id, noteID int64) {
	l.logger.Info(
		"user got note",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note id: %d", noteID),
	)
}

// CreateUserNote logs the user note creation event
func (l *Logger) CreateUserNote(id, noteID int64) {
	l.logger.Info(
		"user created note",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note id: %d", noteID),
	)
}

// UpdateUserNote logs the update user note event
func (l *Logger) UpdateUserNote(id, noteID int64) {
	l.logger.Info(
		"user updated note",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note id: %d", noteID),
	)
}

// DeleteUserNote logs the delete user note event
func (l *Logger) DeleteUserNote(id, noteID int64) {
	l.logger.Info(
		"user deleted note",
		fmt.Sprintf("user id: %d", id),
		fmt.Sprintf("note id: %d", noteID),
	)
}

// ListUserNotes logs the list user notes event
func (l *Logger) ListUserNotes(id int64) {
	l.logger.Info(
		"user listed notes",
		fmt.Sprintf("user id: %d", id),
	)
}
