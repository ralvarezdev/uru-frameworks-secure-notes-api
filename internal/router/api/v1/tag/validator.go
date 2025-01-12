package tag

import (
	govalidatormappervalidations "github.com/ralvarezdev/go-validator/struct/mapper/validator"
)

type (
	// Validator is the structure for API V1 tag validator
	Validator struct {
		govalidatormappervalidations.Service
	}
)
