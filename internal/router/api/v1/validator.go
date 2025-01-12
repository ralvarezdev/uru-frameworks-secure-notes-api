package v1

import (
	govalidatormappervalidator "github.com/ralvarezdev/go-validator/struct/mapper/validator"
)

type (
	// Validator is the structure for API V1 validator
	Validator struct {
		govalidatormappervalidator.Service
	}
)
