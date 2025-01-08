package user

import (
	"fmt"
	govalidatorfield "github.com/ralvarezdev/go-validator/field"
	govalidatorbirthdate "github.com/ralvarezdev/go-validator/field/birthdate"
	govalidatormail "github.com/ralvarezdev/go-validator/field/mail"
	govalidatormapper "github.com/ralvarezdev/go-validator/structs/mapper"
	govalidatormapperservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
	govalidatormappervalidations "github.com/ralvarezdev/go-validator/structs/mapper/validations"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
	"time"
)

var (
	// Mappers
	SignUpRequestMapper *govalidatormapper.Mapper
)

// LoadMappers loads the mappers
func LoadMappers() {
	SignUpRequestMapper, _ = internalvalidator.JSONGenerator.NewMapper(&SignUpRequest{})
}

type (
	// Validator is the structure for API V1 user validator
	Validator struct {
		govalidatormapperservice.Service
	}
)

// ValidateEmail validates the email address field
func (v *Validator) ValidateEmail(
	emailField string,
	email string,
	validations *govalidatormappervalidations.StructValidations,
) {
	if _, err := govalidatormail.ValidMailAddress(email); err != nil {
		validations.AddFieldValidationError(
			emailField,
			govalidatormail.ErrInvalidMailAddress,
		)
	}
}

// ValidateBirthdate validates the birthdate field
func (v *Validator) ValidateBirthdate(
	birthdateField string,
	birthdate *time.Time,
	validations *govalidatormappervalidations.StructValidations,
) {
	if birthdate == nil || birthdate.After(time.Now()) {
		validations.AddFieldValidationError(
			birthdateField,
			govalidatorbirthdate.ErrInvalidBirthdate,
		)
	}
}

// ValidateName validates the name field
func (v *Validator) ValidateName(
	nameField string,
	name string,
	validations *govalidatormappervalidations.StructValidations,
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
			govalidatormapperservice.ErrInvalidBodyType,
			SignUpRequestMapper.Type(),
		)
	}

	return v.RunAndParseValidations(
		parsedBody,
		func(validations *govalidatormappervalidations.StructValidations) (err error) {
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

// ValidateUpdateProfileRequest validates the UpdateProfileRequest
/*
func (v *Validator) ValidateUpdateProfileRequest(request *UpdateProfileRequest) (
	interface{},
	error,
) {
	return v.RunAndParseValidations(
		func(validations *govalidatormappervalidations.StructValidations) error {
			// Check if the birthdate is valid
			if birthdate := request.Birthdate; birthdate != nil {
				v.ValidateBirthdate(
					"birthdate",
					birthdate,
					validations,
				)
			}

			// Check if the first name and last name are valid
			var names = map[string]*string{
				"first_name": request.FirstName,
				"last_name":  request.LastName,
			}
			for nameField, name := range names {
				v.ValidateName(nameField, *name, validations)
			}
			return nil
		},
	)
}
*/
