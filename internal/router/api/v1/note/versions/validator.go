package versions

import (
	govalidatormappervalidations "github.com/ralvarezdev/go-validator/struct/mapper/validator"
)

type (
	// Validator is the structure for API V1 note versions validator
	Validator struct {
		govalidatormappervalidations.Service
	}
)
