package validator

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	govalidatormapper "github.com/ralvarezdev/go-validator/struct/mapper"
	govalidatormapperparser "github.com/ralvarezdev/go-validator/struct/mapper/parser"
	govalidatormapperparserjson "github.com/ralvarezdev/go-validator/struct/mapper/parser/json"
	govalidatormappervalidator "github.com/ralvarezdev/go-validator/struct/mapper/validator"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
)

var (
	// JSONGenerator is the mapper JSON generator
	JSONGenerator govalidatormapper.Generator

	// Validator is the mapper validations validator
	Validator govalidatormappervalidator.Validator

	// Parser is the mapper validations parser
	Parser govalidatormapperparser.Parser

	// Service is the mapper validations service
	Service govalidatormappervalidator.Service
)

// Load initializes the validator constants
func Load(mode *goflagsmode.Flag) {
	// Added the logger to the constants if the debug mode is enabled
	if mode != nil && mode.IsDebug() {
		JSONGenerator = govalidatormapper.NewJSONGenerator(internallogger.MapperGenerator)
		Validator = govalidatormappervalidator.NewDefaultValidator(
			internallogger.MapperValidator,
		)
		Parser = govalidatormapperparserjson.NewParser(internallogger.MapperParser)
	} else {
		JSONGenerator = govalidatormapper.NewJSONGenerator(nil)
		Validator = govalidatormappervalidator.NewDefaultValidator(nil)
		Parser = govalidatormapperparserjson.NewParser(nil)
	}

	// Initialize the service
	service, _ := govalidatormappervalidator.NewDefaultService(
		Parser,
		Validator,
	)
	Service = service
}
