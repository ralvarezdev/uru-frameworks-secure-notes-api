package note

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalrouteapiv1noteversion "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note/version"
	internalrouteapiv1noteversions "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note/versions"
)

var (
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Path:       "/note",
		Controller: Controller,
		Submodules: gonethttp.NewSubmodules(
			internalrouteapiv1noteversion.Module,
			internalrouteapiv1noteversions.Module,
		),
	}
)
