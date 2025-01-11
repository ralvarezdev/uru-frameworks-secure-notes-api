package listener

import (
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
)

const (
	// EnvPort is the key of the default port for the server
	EnvPort = "PORT"
)

var (
	// Port is the default port for the server
	Port string
)

// Load loads the listener constants
func Load() {
	// Get the default port for the server
	port, err := internalloader.Loader.LoadVariable(
		EnvPort,
	)
	if err != nil {
		panic(err)
	}
	internallogger.Environment.EnvironmentVariableLoaded(EnvPort)
	Port = port
}
