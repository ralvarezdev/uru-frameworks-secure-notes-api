package v1

import (
	gojwtissuer "github.com/ralvarezdev/go-jwt/token/issuer"
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/database/postgres"
	internaldto "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/dto"
	"net/http"
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

// Ping pings the service
func (s *Service) Ping() *gonethttpresponse.Response {
	return gonethttpresponse.NewResponseWithCode(
		&internaldto.BasicResponse{
			Message: "Pong",
		}, http.StatusOK,
	)
}
