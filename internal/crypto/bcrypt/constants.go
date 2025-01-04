package bcrypt

import (
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
)

const (
	// CostKey is the key for the cost parameter in the bcrypt hash
	CostKey = "URU_FRAMEWORKS_SECURE_NOTES_BCRYPT_COST"
)

var (
	// Cost is the cost parameter for the bcrypt hash
	Cost int
)

func init() {
	// Get the cost parameter for the bcrypt hash
	cost, err := internalloader.Loader.LoadIntVariable(CostKey)
	if err != nil {
		panic(err)
	}
	internallogger.Environment.EnvironmentVariableLoaded(CostKey)
	Cost = cost
}
