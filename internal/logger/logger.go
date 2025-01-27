package logger

import (
	"fmt"
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
