package bcrypt

import (
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
)

const (
	// EnvCost is the key for the cost parameter in the bcrypt hash
	EnvCost = "URU_FRAMEWORKS_SECURE_NOTES_BCRYPT_COST"
)

var (
	// Cost is the cost parameter for the bcrypt hash
	Cost int
)

// Load loads the bcrypt constants
func Load() {
	// Get the cost parameter for the bcrypt hash
	cost, err := internalloader.Loader.LoadIntVariable(EnvCost)
	if err != nil {
		panic(err)
	}
	internallogger.Environment.EnvironmentVariableLoaded(EnvCost)
	Cost = cost
}
