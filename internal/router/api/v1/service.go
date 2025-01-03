package v1

import (
	gonethttphandler "github.com/ralvarezdev/go-net/http/handler"
	"github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/common"
	"net/http"
)

type (
	// Service is the structure for the API V1 service
	Service struct{}
)

// Ping pings the service
func (s *Service) Ping() *gonethttphandler.Response {
	return gonethttphandler.NewResponseWithCode(
		&common.BasicResponse{
			Message: "pong",
		}, http.StatusOK,
	)
}
