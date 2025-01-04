package postgres

import (
	"database/sql"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	// DSN is the Postgres DSN
	DSN string

	// Database is the Postgres database connection
	Database *gorm.DB

	// Connection is the Postgres connection
	Connection *sql.DB
)

// init initializes the Postgres database connection
func init() {
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
	DSN = Uri + "/" + DatabaseName + "?sslmode=require"

	// Connect to Postgres with GORM
	database, err := gorm.Open(
		postgres.Open(DSN), &gorm.Config{TranslateError: true},
	)
	if err != nil {
		panic(err)
	}
	internallogger.Postgres.ConnectedToDatabase()
	Database = database

	// Retrieve the underlying SQL database connection
	connection, err := Database.DB()
	if err != nil {
		panic(err)
	}
	Connection = connection
}
