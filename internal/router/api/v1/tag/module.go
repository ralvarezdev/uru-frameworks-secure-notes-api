package tag

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
)

var (
	Controller = &controller{}
	Module     = gonethttpfactory.NewModule(
		"/tag", nil, nil, Controller, nil,
	)
)
