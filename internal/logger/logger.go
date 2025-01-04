package logger

import (
	gologger "github.com/ralvarezdev/go-logger"
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
)

// Logger is the logger for the API server
type Logger struct {
	logger gologger.Logger
}

// NewLogger is the logger for the API server
func NewLogger(logger gologger.Logger) (*Logger, error) {
	// Check if the logger is nil
	if logger == nil {
		return nil, gologger.ErrNilLogger
	}

	return &Logger{logger: logger}, nil
}

// SignUp logs the sign-up event
func (l *Logger) SignUp(id string) {
	l.logger.LogMessage(
		gologger.NewLogMessage(
			"User signed up",
			gologgerstatus.StatusInfo,
			nil,
			id,
		),
	)
}
