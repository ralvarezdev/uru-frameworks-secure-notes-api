package auth

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
)

var (
	Service    = &service{}
	Validator  = &validator{}
	Controller = &controller{}
	Module     = gonethttpfactory.NewModule(
		"/auth", Service, Validator, Controller,
		func() {
			// Load the mappers
			LogInRequestMapper = internalvalidator.JSONGenerator.NewMapperWithNoError(&LogInRequest{})
			VerifyTOTPRequestMapper = internalvalidator.JSONGenerator.NewMapperWithNoError(&VerifyTOTPRequest{})
			SignUpRequestMapper = internalvalidator.JSONGenerator.NewMapperWithNoError(&SignUpRequest{})
		},
	)
)
