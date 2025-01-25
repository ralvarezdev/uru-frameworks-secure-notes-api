package versions

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
)

var (
	Controller = &controller{}
	Module     = gonethttpfactory.NewModule(
		"/versions", nil, nil, Controller, nil,
	)
)
