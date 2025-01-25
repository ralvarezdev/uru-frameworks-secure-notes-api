package notes

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
)

type (
	// controller is the structure for the API V1 notes controller
	controller struct {
		gonethttpfactory.Controller
	}
)

// RegisterRoutes registers the routes for the API V1 notes controller
func (c *controller) RegisterRoutes() {}
