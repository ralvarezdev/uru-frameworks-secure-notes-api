package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/pgxpool"
	godatabasespgxpool "github.com/ralvarezdev/go-databases/sql/pgxpool"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	"time"
)

const (
	// EnvDSN is the key of the DSN for the Postgres database
	EnvDSN = "URU_FRAMEWORKS_SECURE_NOTES_POSTGRES_DSN"

	// EnvMaxOpenConnections is the key of the maximum number of open connections for the Postgres database
	EnvMaxOpenConnections = "URU_FRAMEWORKS_SECURE_NOTES_POSTGRES_MAX_OPEN_CONNECTIONS"

	// EnvMaxIdleConnections is the key of the maximum number of idle connections for the Postgres database
	EnvMaxIdleConnections = "URU_FRAMEWORKS_SECURE_NOTES_POSTGRES_MAX_IDLE_CONNECTIONS"
)

var (
	// DSN is the DSN for the Postgres database
	DSN string

	// MaxOpenConnections is the maximum number of open connections for the Postgres database
	MaxOpenConnections int

	// MaxIdleConnections is the maximum number of idle connections for the Postgres database
	MaxIdleConnections int

	// PoolHandler is the Postgres pool handler
	PoolHandler godatabasespgxpool.PoolHandler

	// PoolService is the Postgres pool service
	PoolService *Service
)

// Load loads the Postgres constants
func Load(mode *goflagsmode.Flag) {
	// Load the DSN for the Postgres database
	if err := internalloader.Loader.LoadVariable(
		EnvDSN,
		&DSN,
	); err != nil {
		panic(err)
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

	// Create the Postgres pool configuration
	config, err := godatabasespgxpool.NewPoolConfig(
		DSN,
		MaxOpenConnections,
		MaxIdleConnections,
		time.Hour,
		time.Hour,
		5*time.Minute,
		5*time.Minute,
	)

	// Create the Postgres database pool handler
	poolHandler, err := godatabasespgxpool.NewDefaultPoolHandler(config)
	if err != nil {
		panic(err)
	}
	PoolHandler = poolHandler

	// Connect to the Postgres database
	pool, err := poolHandler.Connect()
	if err != nil {
		panic(err)
	}
	internallogger.Postgres.ConnectedToDatabase()

	// Create the Postgres database service
	service, err := NewService(
		pool,
	)
	if err != nil {
		panic(err)
	}
	PoolService = service

	// Check if the mode is debug
	if !mode.IsDebug() {
		return
	}

	// Set ticker for the Postgres database
	service.SetStatTicker(
		10*time.Second, func(stat *pgxpool.Stat) {
			internallogger.Api.PoolStat(stat)
		},
	)
}
