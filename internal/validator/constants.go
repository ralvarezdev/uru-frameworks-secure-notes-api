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

	// ValidationsValidator is the mapper validations validator
	ValidationsValidator govalidatormappervalidator.Validator

	// ValidationsParser is the mapper validations parser
	ValidationsParser govalidatormapperparser.Parser

	// ValidationsService is the mapper validations service
	ValidationsService govalidatormappervalidator.Service
)

// Load initializes the validator constants
func Load() {
	// Added the logger to the constants if the debug mode is enabled
	if goflagsmode.ModeFlag.IsDebug() {
		JSONGenerator = govalidatormapper.NewJSONGenerator(internallogger.MapperGenerator)
		ValidationsValidator = govalidatormappervalidator.NewDefaultValidator(
			internallogger.MapperValidator,
		)
		ValidationsParser = govalidatormapperparserjson.NewParser(internallogger.MapperParser)
	} else {
		JSONGenerator = govalidatormapper.NewJSONGenerator(nil)
		ValidationsValidator = govalidatormappervalidator.NewDefaultValidator(nil)
		ValidationsParser = govalidatormapperparserjson.NewParser(nil)
	}

	ValidationsService, _ = govalidatormappervalidator.NewDefaultService(
		ValidationsParser,
		ValidationsValidator,
	)
}
