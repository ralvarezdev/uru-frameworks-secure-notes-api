package notes

import (
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
)

type (
	// Service is the structure for the API V1 service for the notes route group
	Service struct {
		PostgresService *internalpostgres.Service
	}
)
