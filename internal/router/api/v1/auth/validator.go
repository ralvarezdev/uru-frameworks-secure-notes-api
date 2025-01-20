package auth

import (
	"fmt"
	gonethttp "github.com/ralvarezdev/go-net/http"
	govalidatorfield "github.com/ralvarezdev/go-validator/struct/field"
	govalidatorfieldbirthdate "github.com/ralvarezdev/go-validator/struct/field/birthdate"
	govalidatorfieldmail "github.com/ralvarezdev/go-validator/struct/field/mail"
	govalidatormapper "github.com/ralvarezdev/go-validator/struct/mapper"
	govalidatormappervalidation "github.com/ralvarezdev/go-validator/struct/mapper/validation"
	govalidatormappervalidator "github.com/ralvarezdev/go-validator/struct/mapper/validator"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
	"time"
)

var (
	// Mappers
	SignUpRequestMapper     *govalidatormapper.Mapper
	LogInRequestMapper      *govalidatormapper.Mapper
	VerifyTOTPRequestMapper *govalidatormapper.Mapper
)

// LoadMappers loads the mappers
func LoadMappers() {
	LogInRequestMapper, _ = internalvalidator.JSONGenerator.NewMapper(&LogInRequest{})
	VerifyTOTPRequestMapper, _ = internalvalidator.JSONGenerator.NewMapper(&VerifyTOTPRequest{})
	SignUpRequestMapper, _ = internalvalidator.JSONGenerator.NewMapper(&SignUpRequest{})
}

type (
	// Validator is the structure for API V1 auth validator
	Validator struct {
		govalidatormappervalidator.Service
	}
)

// ValidateEmail validates the email address field
func (v *Validator) ValidateEmail(
	emailField string,
	email string,
	validations *govalidatormappervalidation.StructValidations,
) {
	if _, err := govalidatorfieldmail.ValidMailAddress(email); err != nil {
		validations.AddFieldValidationError(
			emailField,
			govalidatorfieldmail.ErrInvalidMailAddress,
		)
	}
}

// ValidateBirthdate validates the birthdate field
func (v *Validator) ValidateBirthdate(
	birthdateField string,
	birthdate *time.Time,
	validations *govalidatormappervalidation.StructValidations,
) {
	if birthdate == nil || birthdate.After(time.Now()) {
		validations.AddFieldValidationError(
			birthdateField,
			govalidatorfieldbirthdate.ErrInvalidBirthdate,
		)
	}
}

// ValidateName validates the name field
func (v *Validator) ValidateName(
	nameField string,
	name string,
	validations *govalidatormappervalidation.StructValidations,
) {
	if name == "" {
		validations.AddFieldValidationError(
			nameField,
			govalidatorfield.ErrEmptyField,
		)
	}
}

// ValidateSignUpRequest validates the SignUpRequest
func (v *Validator) ValidateSignUpRequest(body interface{}) (
	interface{},
	error,
) {
	// Parse body
	parsedBody, ok := body.(*SignUpRequest)
	if !ok {
		return nil, fmt.Errorf(
			gonethttp.ErrInvalidRequestBody,
			SignUpRequestMapper.Type(),
		)
	}

	return v.RunAndParseValidations(
		parsedBody,
		func(validations *govalidatormappervalidation.StructValidations) (err error) {
			err = v.ValidateRequiredFields(
				validations,
				SignUpRequestMapper,
			)
			if err != nil {
				return err
			}

			// Check if the email is valid
			v.ValidateEmail("email", parsedBody.Email, validations)
			return nil
		},
	)
}

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
