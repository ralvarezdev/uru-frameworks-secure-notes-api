package listener

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
)

const (
	// EnvPort is the key of the default port for the server
	EnvPort = "PORT"

	// EnvHost is the key of the default host for the server
	EnvHost = "HOST"
)

var (
	// Port is the default port for the server
	Port string

	// Host is the default host for the server
	Host string
)

// Load loads the listener constants
func Load(mode *goflagsmode.Flag) {
	// Load the port and host variables
	for env, dest := range map[string]*string{
		EnvHost: &Host,
		EnvPort: &Port,
	} {
		if err := internalloader.Loader.LoadVariable(
			env,
			dest,
		); err != nil {
			panic(err)
		}
	}

	// Check if it's not on production mode
	if mode != nil && !mode.IsProd() {
		Host += ":" + Port
	}
}
