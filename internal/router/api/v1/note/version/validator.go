package version

import (
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
)

type (
	// Validator is the structure for API V1 note version validator
	Validator struct {
		govalidatorservice.Service
	}
)
