package user

import (
	govalidatorfield "github.com/ralvarezdev/go-validator/field"
	govalidatorbirthdate "github.com/ralvarezdev/go-validator/field/birthdate"
	govalidatormail "github.com/ralvarezdev/go-validator/field/mail"
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
	govalidatorvalidations "github.com/ralvarezdev/go-validator/structs/mapper/validations"
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
		govalidatorservice.Service
	}
)

// ValidateEmail validates the email address field
func (v *Validator) ValidateEmail(
	emailField string,
	email string,
	validations govalidatorvalidations.Validations,
) {
	if _, err := govalidatormail.ValidMailAddress(email); err != nil {
		validations.AddFailedFieldValidationError(
			emailField,
			govalidatormail.ErrInvalidMailAddress,
		)
	}
}

// ValidateBirthdate validates the birthdate field
func (v *Validator) ValidateBirthdate(
	birthdateField string,
	birthdate *time.Time,
	validations govalidatorvalidations.Validations,
) {
	if birthdate == nil || birthdate.After(time.Now()) {
		validations.AddFailedFieldValidationError(
			birthdateField,
			govalidatorbirthdate.ErrInvalidBirthdate,
		)
	}
}

// ValidateName validates the name field
func (v *Validator) ValidateName(
	nameField string,
	name string,
	validations govalidatorvalidations.Validations,
) {
	if name == "" {
		validations.AddFailedFieldValidationError(
			nameField,
			govalidatorfield.ErrEmptyField,
		)
	}
}

// ValidateSignUpRequest validates the sign-up request
func (v *Validator) ValidateSignUpRequest(request *SignUpRequest) error {
	validations, _ := v.ValidateNilFields(
		request,
		SignUpRequestMapper,
	)

	// Check if the email is valid
	v.ValidateEmail("email", request.Email, validations)

	return v.CheckValidations(validations)
}

// ValidateUpdateProfileRequest validates the update profile request
func (v *Validator) ValidateUpdateProfileRequest(request *UpdateProfileRequest) error {
	validations := govalidatorvalidations.NewDefaultValidations()

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

	return v.CheckValidations(validations)
}
