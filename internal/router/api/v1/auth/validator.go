package auth

import (
	govalidatormapper "github.com/ralvarezdev/go-validator/structs/mapper"
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
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
		govalidatorservice.Service
	}
)

// ValidateLogInRequest validates the LogInRequest
/*
func (v *Validator) ValidateLogInRequest(request *LogInRequest) error {
	validations, _ := v.ValidateNilFields(
		request,
		LogInRequestMapper,
	)

	// Check if the email is valid
	v.ValidateEmail("email", request.Email, validations)

	return v.CheckValidations(validations)
}
*/
