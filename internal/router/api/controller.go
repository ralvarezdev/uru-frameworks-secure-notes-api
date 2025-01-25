package api

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
)

type (
	// controller is the structure for the API controller
	controller struct {
		gonethttpfactory.Controller
	}
)
