package version

import (
	govalidatormappervalidations "github.com/ralvarezdev/go-validator/struct/mapper/validator"
)

type (
	// Validator is the structure for API V1 note version validator
	Validator struct {
		govalidatormappervalidations.Service
	}
)
