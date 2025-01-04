package validator

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	govalidatormapper "github.com/ralvarezdev/go-validator/structs/mapper"
	govalidatormapperparser "github.com/ralvarezdev/go-validator/structs/mapper/parser"
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
	govalidatormappervalidator "github.com/ralvarezdev/go-validator/structs/mapper/validator"
)

var (
	// JSONGenerator is the mapper JSON generator
	JSONGenerator = govalidatormapper.NewJSONGenerator(goflagsmode.Mode)

	// ValidationsValidator is the mapper validations validator
	ValidationsValidator = govalidatormappervalidator.NewDefaultValidator(goflagsmode.Mode)

	// ValidationsParser is the mapper validations parser
	ValidationsParser = govalidatormapperparser.NewJSONParser()

	// ValidationsService is the mapper validations service
	ValidationsService, _ = govalidatorservice.NewDefaultService(
		ValidationsParser,
		ValidationsValidator,
		goflagsmode.Mode,
	)
)
