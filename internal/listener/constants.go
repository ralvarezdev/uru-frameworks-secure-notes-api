package listener

import (
	goloaderlistener "github.com/ralvarezdev/go-loader/http/listener"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
)

const (
	// PortKey is the key of the default port for the server
	PortKey = "PORT"
)

var (
	// Port is the default port for the server
	Port *goloaderlistener.Port
)

// Load loads the listener constants
func Load() {
	// Get the default port for the server
	port, err := goloaderlistener.LoadPort(
		internalloader.Loader,
		"0.0.0.0",
		PortKey,
	)
	if err != nil {
		panic(err)
	}
	internallogger.Environment.EnvironmentVariableLoaded(PortKey)
	Port = port
}
