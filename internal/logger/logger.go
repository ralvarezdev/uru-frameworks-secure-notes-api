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

// SignUp logs the sign-up event
func (l *Logger) SignUp(id int64) {
	l.logger.Info(
		"user signed up",
		fmt.Sprintf("user id: %d", id),
	)
}

// LogIn logs the log-in event
func (l *Logger) LogIn(id int64) {
	l.logger.Info(
		"user logged in",
		fmt.Sprintf("user id: %d", id),
	)
}
