package version

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
)

var (
	Controller = &controller{}
	Module     = gonethttpfactory.NewModule(
		"/version", nil, nil, Controller, nil,
	)
)
