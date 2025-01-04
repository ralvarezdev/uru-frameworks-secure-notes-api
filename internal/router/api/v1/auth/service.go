package auth

import (
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
)

type (
	// Service is the structure for the API V1 service for the auth route group
	Service struct {
		JwtIssuer       gojwtissuer.Issuer
		PostgresService *internalpostgres.Service
	}
)
