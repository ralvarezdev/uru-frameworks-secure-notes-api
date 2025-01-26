package v1

import (
	gonethttpresponse "github.com/ralvarezdev/go-net/http/response"
	internalhandler "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/handler"
	"net/http"
)

type (
	// controller is the structure for the API V1 controller
	controller struct{}
)

// Ping pings the service
// @Summary Ping the service
// @Description Returns a pong response to check if the service is running
// @Tags api v1
// @Accept json
// @Produce json
// @Success 200 {object} BasicResponse
// @Router /api/v1/ping [get]
func (c *controller) Ping(w http.ResponseWriter, r *http.Request) {
	// Handle the response
	internalhandler.Handler.HandleResponse(
		w, gonethttpresponse.NewSuccessResponse(
			nil, http.StatusOK,
		),
	)
}
