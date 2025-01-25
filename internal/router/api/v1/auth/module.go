package auth

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
)

var (
	Service   = &service{}
	Validator = &validator{
		Service:                 internalvalidator.ValidationsService,
		LogInRequestMapper:      internalvalidator.JSONGenerator.NewMapperWithNoError(&LogInRequest{}),
		VerifyTOTPRequestMapper: internalvalidator.JSONGenerator.NewMapperWithNoError(&VerifyTOTPRequest{}),
		SignUpRequestMapper:     internalvalidator.JSONGenerator.NewMapperWithNoError(&SignUpRequest{}),
	}
	Controller = &controller{
		Service:   Service,
		Validator: Validator,
	}
	Module = gonethttpfactory.NewModule("/auth", Service, Validator, Controller)
)
