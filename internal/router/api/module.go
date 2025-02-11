package api

import (
	"fmt"
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalrouterapiv1 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1"
)

var (
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Pattern:    "/api",
		Controller: Controller,
		Submodules: gonethttp.NewSubmodules(internalrouterapiv1.Module),
		AfterLoadFn: func(m *gonethttp.Module) {
			fmt.Println("2b", m.GetRouter().GetMiddlewares())
		},
	}
)
