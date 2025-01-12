package listener

import (
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
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
	if err := internalloader.Loader.LoadVariable(
		EnvPort,
		&Port,
	); err != nil {
		panic(err)
	}
}
