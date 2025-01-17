package postgres

import (
	godatabasessql "github.com/ralvarezdev/go-databases/sql"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	"time"
)

const (
	// EnvUri is the key of the default URI for the Postgres database
	EnvUri = "URU_FRAMEWORKS_SECURE_NOTES_POSTGRES_HOST"

	// EnvDatabaseName is the key of the default database name for the Postgres database
	EnvDatabaseName = "URU_FRAMEWORKS_SECURE_NOTES_POSTGRES_NAME"

	// EnvMaxOpenConnections is the key of the maximum number of open connections for the Postgres database
	EnvMaxOpenConnections = "URU_FRAMEWORKS_SECURE_NOTES_POSTGRES_MAX_OPEN_CONNECTIONS"

	// EnvMaxIdleConnections is the key of the maximum number of idle connections for the Postgres database
	EnvMaxIdleConnections = "URU_FRAMEWORKS_SECURE_NOTES_POSTGRES_MAX_IDLE_CONNECTIONS"
)

var (
	// Uri is the default URI for the Postgres database
	Uri string

	// DatabaseName is the default database name for the Postgres database
	DatabaseName string

	// DataSourceName is the Postgres DSN
	DataSourceName string

	// MaxOpenConnections is the maximum number of open connections for the Postgres database
	MaxOpenConnections int

	// MaxIdleConnections is the maximum number of idle connections for the Postgres database
	MaxIdleConnections int

	// Config is the Postgres configuration
	Config *godatabasessql.Config
)

// Load loads the Postgres constants
func Load() {
	// Load the default URI and database name for the Postgres database
	for key, variable := range map[string]*string{
		EnvUri:          &Uri,
		EnvDatabaseName: &DatabaseName,
	} {
		if err := internalloader.Loader.LoadVariable(
			key,
			variable,
		); err != nil {
			panic(err)
		}
	}

	// Load the maximum number of open and idle connections for the Postgres database
	for key, variable := range map[string]*int{
		EnvMaxOpenConnections: &MaxOpenConnections,
		EnvMaxIdleConnections: &MaxIdleConnections,
	} {
		if err := internalloader.Loader.LoadIntVariable(
			key,
			variable,
		); err != nil {
			panic(err)
		}
	}

	// Create the Postgres DSN
	DataSourceName = Uri + "/" + DatabaseName + "?sslmode=require"

	Config = godatabasessql.NewConfig(
		MaxOpenConnections,
		MaxIdleConnections,
		time.Hour,
	)
}
