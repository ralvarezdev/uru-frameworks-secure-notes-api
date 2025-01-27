package auth

import (
	govalidatorfieldbirthdate "github.com/ralvarezdev/go-validator/struct/field/birthdate"
	govalidatorfieldmail "github.com/ralvarezdev/go-validator/struct/field/mail"
	govalidatormapper "github.com/ralvarezdev/go-validator/struct/mapper"
	govalidatormappervalidation "github.com/ralvarezdev/go-validator/struct/mapper/validation"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
	"time"
)

type (
	// validator is the structure for API V1 auth validator
	validator struct{}
)

// Email validates the email address field
func (v *validator) Email(
	emailField string,
	email string,
	validations *govalidatormappervalidation.StructValidations,
) error {
	if _, err := govalidatorfieldmail.ValidMailAddress(email); err != nil {
		validations.AddFieldValidationError(
			emailField,
			govalidatorfieldmail.ErrInvalidMailAddress,
		)
	}
	return nil
}

// Birthdate validates the birthdate field
func (v *validator) Birthdate(
	birthdateField string,
	birthdate *time.Time,
	validations *govalidatormappervalidation.StructValidations,
) error {
	if birthdate == nil || birthdate.After(time.Now()) {
		validations.AddFieldValidationError(
			birthdateField,
			govalidatorfieldbirthdate.ErrInvalidBirthdate,
		)
	}
	return nil
}

// SignUp validates the SignUpRequest
func (v *validator) SignUp(
	body *SignUpRequest,
	mapper *govalidatormapper.Mapper,
) func() (
	interface{},
	error,
) {
	return internalvalidator.Validate(
		body,
		mapper,
		func(validations *govalidatormappervalidation.StructValidations) (err error) {
			return v.Email("email", body.Email, validations)
		},
	)
}

// LogIn validates the LogInRequest
func (v *validator) LogIn(
	body *LogInRequest,
	mapper *govalidatormapper.Mapper,
) func() (
	interface{},
	error,
) {
	return internalvalidator.Validate(
		body,
		mapper,
	)
}

// VerifyTOTP validates the VerifyTOTPRequest
func (v *validator) VerifyTOTP(
	body *VerifyTOTPRequest,
	mapper *govalidatormapper.Mapper,
) func() (
	interface{},
	error,
) {
	return internalvalidator.Validate(
		body,
		mapper,
	)
}
