package user

import (
	govalidatorfield "github.com/ralvarezdev/go-validator/field"
	govalidatorbirthdate "github.com/ralvarezdev/go-validator/field/birthdate"
	govalidatormail "github.com/ralvarezdev/go-validator/field/mail"
	govalidatormapperservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
	govalidatormappervalidations "github.com/ralvarezdev/go-validator/structs/mapper/validations"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
	"time"
)

var (
	// Mappers
	SignUpRequestMapper, _        = internalvalidator.JSONGenerator.NewMapper(&SignUpRequest{})
	UpdateProfileRequestMapper, _ = internalvalidator.JSONGenerator.NewMapper(&UpdateProfileRequest{})
)

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
	validations govalidatormappervalidations.Validations,
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
	validations govalidatormappervalidations.Validations,
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
	validations govalidatormappervalidations.Validations,
) {
	if name == "" {
		validations.AddFieldValidationError(
			nameField,
			govalidatorfield.ErrEmptyField,
		)
	}
}

// ValidateSignUpRequest validates the SignUpRequest
func (v *Validator) ValidateSignUpRequest(request *SignUpRequest) (
	interface{},
	error,
) {
	return v.RunAndParseValidations(
		func(validations govalidatormappervalidations.Validations) (err error) {
			err = v.ValidateNilFields(
				validations,
				request,
				SignUpRequestMapper,
			)
			if err != nil {
				return err
			}

			// Check if the email is valid
			v.ValidateEmail("email", request.Email, validations)
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
		func(validations govalidatormappervalidations.Validations) error {
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
