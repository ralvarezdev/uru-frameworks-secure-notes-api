package router

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
	internalrouterapi "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api"
)

var (
	Controller = &controller{}
	Module     = gonethttpfactory.NewModule(
		"/", nil, nil, Controller, internalrouterapi.Module,
	)
)
