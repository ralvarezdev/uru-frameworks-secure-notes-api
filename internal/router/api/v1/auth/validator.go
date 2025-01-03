package auth

import (
	govalidatorservice "github.com/ralvarezdev/go-validator/structs/mapper/service"
)

type (
	// Validator is the structure for API V1 auth validator
	Validator struct {
		govalidatorservice.Service
	}
)