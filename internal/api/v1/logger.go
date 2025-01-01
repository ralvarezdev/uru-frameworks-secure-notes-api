package v1

import (
	gologger "github.com/ralvarezdev/go-logger"
)

// Logger is the logger for the API V1
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
