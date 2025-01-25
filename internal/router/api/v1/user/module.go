package user

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
)

var (
	Controller = &controller{}
	Module     = gonethttpfactory.NewModule(
		"/user", nil, nil, Controller,
	)
)
