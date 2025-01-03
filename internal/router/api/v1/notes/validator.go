package notes

import (
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
)

type (
	// Validator is the structure for API V1 notes validator
	Validator struct {
		govalidatorservice.Service
	}
)
