package validator

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	govalidatormapper "github.com/ralvarezdev/go-validator/structs/mapper"
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
	govalidatorvalidations "github.com/ralvarezdev/go-validator/structs/mapper/validations"
)

var (
	// JSONGenerator is the mapper JSON generator
	JSONGenerator = govalidatormapper.NewJSONGenerator(goflagsmode.Mode)

	// ValidationsValidator is the mapper validations validator
	ValidationsValidator = govalidatorvalidations.NewDefaultValidator(
		goflagsmode.Mode,
	)

	// ValidationsGenerator is the mapper validations generator
	ValidationsGenerator = govalidatorvalidations.NewDefaultGenerator(nil)

	// ValidationsService is the mapper validations service
	ValidationsService, _ = govalidatorservice.NewDefaultService(
		ValidationsGenerator,
		ValidationsValidator,
		goflagsmode.Mode,
	)
)
