package user

import (
	govalidatormappervalidator "github.com/ralvarezdev/go-validator/struct/mapper/validator"
)

var (
// Mappers
)

// LoadMappers loads the mappers
func LoadMappers() {
}

type (
	// Validator is the structure for API V1 user validator
	Validator struct {
		govalidatormappervalidator.Service
	}
)

// ValidateUpdateProfileRequest validates the UpdateProfileRequest
/*
func (v *Validator) ValidateUpdateProfileRequest(request *UpdateProfileRequest) (
	interface{},
	error,
) {
	return v.RunAndParseValidations(
		func(validations *govalidatormappervalidation.StructValidations) error {
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
