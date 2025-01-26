package notes

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
)

var (
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Path:       "/notes",
		Controller: Controller,
	}
)