package tag

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
)

type (
	// controller is the structure for the API V1 tag controller
	controller struct {
		gonethttpfactory.Controller
	}
)

// RegisterRoutes registers the routes for the API V1 tag controller
func (c *controller) RegisterRoutes() {}
