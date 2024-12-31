package v1

import (
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/database/postgres"
)

type (
	// Service is the structure for the API V1 service
	Service struct {
		jwtIssuer         gojwtissuer.Issuer
		postgresqlService *internalpostgres.Service
	}
)

// NewService creates a new API V1 service
func NewService(
	jwtIssuer gojwtissuer.Issuer, postgresqlService *internalpostgres.Service,
) (*Service, error) {
	return &Service{
		postgresqlService: postgresqlService,
		jwtIssuer:         jwtIssuer,
	}, nil
}
