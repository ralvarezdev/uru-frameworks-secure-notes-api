package router

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
)

type (
	// controller is the structure for the API V1 controller
	controller struct {
		gonethttpfactory.Controller
	}
)
