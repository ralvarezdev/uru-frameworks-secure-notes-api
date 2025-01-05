package validator

import (
	govalidatormapper "github.com/ralvarezdev/go-validator/structs/mapper"
	govalidatormapperparserjson "github.com/ralvarezdev/go-validator/structs/mapper/parser/json"
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
	govalidatormappervalidations "github.com/ralvarezdev/go-validator/structs/mapper/validations"
	govalidatormappervalidator "github.com/ralvarezdev/go-validator/structs/mapper/validator"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
)

var (
	// JSONGenerator is the mapper JSON generator
	JSONGenerator = govalidatormapper.NewJSONGenerator(internallogger.MapperGenerator)

	// ValidationsValidator is the mapper validations validator
	ValidationsValidator = govalidatormappervalidator.NewDefaultValidator(
		govalidatormappervalidations.NewDefaultValidations,
		internallogger.MapperValidator,
	)

	// ValidationsParser is the mapper validations parser
	ValidationsParser = govalidatormapperparserjson.NewParser(govalidatormapperparserjson.NewDefaultParsedValidations)

	// ValidationsService is the mapper validations service
	ValidationsService, _ = govalidatorservice.NewDefaultService(
		ValidationsParser,
		ValidationsValidator,
	)
)
