package v1

import (
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
)

type (
	// Validator is the structure for API V1 validator
	Validator struct {
		service          Service
		validatorService govalidatorservice.Service
	}
)

// NewValidator creates a new API V1 validator
func NewValidator(
	service Service,
	validatorService govalidatorservice.Service,
) *Validator {
	return &Validator{
		service:          service,
		validatorService: validatorService,
	}
}
