package router

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalrouterapi "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api"
)

var (
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Path:       "/",
		Controller: Controller,
		Submodules: gonethttp.NewSubmodules(
			internalrouterapi.Module,
		),
	}
)
