package note

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
	internalrouteapiv1noteversion "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note/version"
	internalrouteapiv1noteversions "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note/versions"
)

var (
	Controller = &controller{}
	Module     = gonethttpfactory.NewModule(
		"/note", nil, nil, Controller,
		internalrouteapiv1noteversion.Module,
		internalrouteapiv1noteversions.Module,
	)
)
