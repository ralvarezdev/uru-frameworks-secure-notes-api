package notes

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
)

var (
	Controller = &controller{}
	Module     = gonethttpfactory.NewModule(
		"/notes", nil, nil, Controller, nil,
	)
)
