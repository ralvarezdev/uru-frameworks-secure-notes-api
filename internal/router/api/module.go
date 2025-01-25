package api

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
	internalrouterapiv1 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1"
)

var (
	Controller = &controller{}
	Module     = gonethttpfactory.NewModule(
		"/api", nil, nil, Controller, nil, internalrouterapiv1.Module,
	)
)
