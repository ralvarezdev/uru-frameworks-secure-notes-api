package v1

import (
	gonethttpfactory "github.com/ralvarezdev/go-net/http/factory"
	internalrouterapiv1auth "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/auth"
	internalrouterapiv1note "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note"
	internalrouterapiv1notes "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/notes"
	internalrouterapiv1tag "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/tag"
	internalrouterapiv1user "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/user"
)

var (
	Controller = &controller{}
	Module     = gonethttpfactory.NewModule(
		"/v1", nil, nil, Controller,
		internalrouterapiv1auth.Module,
		internalrouterapiv1note.Module,
		internalrouterapiv1notes.Module,
		internalrouterapiv1tag.Module,
		internalrouterapiv1user.Module,
	)
)
