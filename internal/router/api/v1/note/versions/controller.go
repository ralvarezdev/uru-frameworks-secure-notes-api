package versions

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
)

type (
	// controller is the structure for the API V1 versions controller
	controller struct {
		gonethttpfactory.Controller
	}
)

// RegisterRoutes registers the routes for the API V1 versions controller
func (c *controller) RegisterRoutes() {}
