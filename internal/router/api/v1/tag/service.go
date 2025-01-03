package tag

import (
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/database/postgres"
)

type (
	// Service is the structure for the API V1 service for the tag route group
	Service struct {
		PostgresService *internalpostgres.Service
	}
)
