package auth

import (
	"fmt"
	govalidatormapper "github.com/ralvarezdev/go-validator/structs/mapper"
	govalidatormapperservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
	govalidatormappervalidations "github.com/ralvarezdev/go-validator/structs/mapper/validations"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
)

var (
	// Mappers
	LogInRequestMapper *govalidatormapper.Mapper
)

// LoadMappers loads the mappers
func LoadMappers() {
	LogInRequestMapper, _ = internalvalidator.JSONGenerator.NewMapper(&LogInRequest{})
}

type (
	// Validator is the structure for API V1 auth validator
	Validator struct {
		govalidatormapperservice.Service
	}
)

// ValidateLogInRequest validates the LogInRequest
func (v *Validator) ValidateLogInRequest(body interface{}) (
	interface{},
	error,
) {
	// Parse body
	parsedBody, ok := body.(*LogInRequest)
	if !ok {
		return nil, fmt.Errorf(
			govalidatormapperservice.ErrInvalidBodyType,
			LogInRequestMapper.Type(),
		)
	}

	return v.RunAndParseValidations(
		parsedBody,
		func(validations *govalidatormappervalidations.StructValidations) error {
			return v.ValidateRequiredFields(
				validations,
				LogInRequestMapper,
			)
		},
	)
}
