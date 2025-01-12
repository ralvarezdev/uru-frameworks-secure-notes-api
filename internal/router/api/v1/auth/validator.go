package auth

import (
	"fmt"
	gonethttp "github.com/ralvarezdev/go-net/http"
	govalidatormapper "github.com/ralvarezdev/go-validator/struct/mapper"
	govalidatormappervalidation "github.com/ralvarezdev/go-validator/struct/mapper/validation"
	govalidatormappervalidator "github.com/ralvarezdev/go-validator/struct/mapper/validator"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
)

var (
	// Mappers
	LogInRequestMapper              *govalidatormapper.Mapper
	RevokeRefreshTokenRequestMapper *govalidatormapper.Mapper
	VerifyTOTPRequestMapper         *govalidatormapper.Mapper
)

// LoadMappers loads the mappers
func LoadMappers() {
	LogInRequestMapper, _ = internalvalidator.JSONGenerator.NewMapper(&LogInRequest{})
	RevokeRefreshTokenRequestMapper, _ = internalvalidator.JSONGenerator.NewMapper(&RevokeRefreshTokenRequest{})
	VerifyTOTPRequestMapper, _ = internalvalidator.JSONGenerator.NewMapper(&VerifyTOTPRequest{})
}

type (
	// Validator is the structure for API V1 auth validator
	Validator struct {
		govalidatormappervalidator.Service
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
			gonethttp.ErrInvalidRequestBody,
			LogInRequestMapper.Type(),
		)
	}

	return v.RunAndParseValidations(
		parsedBody,
		func(validations *govalidatormappervalidation.StructValidations) error {
			return v.ValidateRequiredFields(
				validations,
				LogInRequestMapper,
			)
		},
	)
}

// ValidateRevokeRefreshTokenRequest validates the RevokeRefreshTokenRequest
func (v *Validator) ValidateRevokeRefreshTokenRequest(body interface{}) (
	interface{},
	error,
) {
	// Parse body
	parsedBody, ok := body.(*RevokeRefreshTokenRequest)
	if !ok {
		return nil, fmt.Errorf(
			gonethttp.ErrInvalidRequestBody,
			RevokeRefreshTokenRequestMapper.Type(),
		)
	}

	return v.RunAndParseValidations(
		parsedBody,
		func(validations *govalidatormappervalidation.StructValidations) error {
			return v.ValidateRequiredFields(
				validations,
				RevokeRefreshTokenRequestMapper,
			)
		},
	)
}

// ValidateVerifyTOTPRequest validates the VerifyTOTPRequest
func (v *Validator) ValidateVerifyTOTPRequest(body interface{}) (
	interface{},
	error,
) {
	// Parse body
	parsedBody, ok := body.(*VerifyTOTPRequest)
	if !ok {
		return nil, fmt.Errorf(
			gonethttp.ErrInvalidRequestBody,
			VerifyTOTPRequestMapper.Type(),
		)
	}

	return v.RunAndParseValidations(
		parsedBody,
		func(validations *govalidatormappervalidation.StructValidations) error {
			return v.ValidateRequiredFields(
				validations,
				VerifyTOTPRequestMapper,
			)
		},
	)
}
