package tag

import (
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
)

type (
	// Validator is the structure for API V1 tag validator
	Validator struct {
		govalidatorservice.Service
	}
)
