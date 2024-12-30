package v1

import (
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	internalpostgresql "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/database/postgresql"
)

type (
	// Service is the structure for the API V1 service
	Service struct {
		jwtIssuer         gojwtissuer.Issuer
		postgresqlService internalpostgresql.Service
		gonethttproute.Service
	}
)

// NewService creates a new API V1 service
func NewService(
	jwtIssuer gojwtissuer.Issuer, postgresqlService internalpostgresql.Service,
) (*Service, error) {
	return &Service{
		postgresqlService: postgresqlService,
		jwtIssuer:         jwtIssuer,
	}, nil
}
