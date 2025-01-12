package note

import (
	govalidatormappervalidations "github.com/ralvarezdev/go-validator/struct/mapper/validator"
)

type (
	// Validator is the structure for API V1 note validator
	Validator struct {
		govalidatormappervalidations.Service
	}
)
