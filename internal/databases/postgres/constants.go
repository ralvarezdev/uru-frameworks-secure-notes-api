package postgres

import (
	godatabasessql "github.com/ralvarezdev/go-databases/sql"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"time"
)

const (
	// UriKey is the key of the default URI for the Postgres database
	UriKey = "URU_FRAMEWORKS_SECURE_NOTES_POSTGRESQL_HOST"

	// DatabaseNameKey is the key of the default database name for the Postgres database
	DatabaseNameKey = "URU_FRAMEWORKS_SECURE_NOTES_POSTGRESQL_NAME"
)

var (
	// Uri is the default URI for the Postgres database
	Uri string

	// DatabaseName is the default database name for the Postgres database
	DatabaseName string

	// DataSourceName is the Postgres DSN
	DataSourceName string

	// Config is the Postgres configuration
	Config = godatabasessql.NewConfig(2, 10, time.Hour)
)

// Load loads the Postgres constants
func Load() {
	// Get the default URI for the Postgres database
	uri, err := internalloader.Loader.LoadVariable(UriKey)
	if err != nil {
		panic(err)
	}
	internallogger.Environment.EnvironmentVariableLoaded(UriKey)
	Uri = uri

	// Get the default database name for the Postgres database
	databaseName, err := internalloader.Loader.LoadVariable(DatabaseNameKey)
	if err != nil {
		panic(err)
	}
	internallogger.Environment.EnvironmentVariableLoaded(DatabaseNameKey)
	DatabaseName = databaseName

	// Create the Postgres DSN
	DataSourceName = Uri + "/" + DatabaseName + "?sslmode=require"
}
