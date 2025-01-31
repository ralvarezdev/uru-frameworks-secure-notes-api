package v1

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalrouterapiv1auth "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/auth"
	internalrouterapiv1note "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note"
	internalrouterapiv1notes "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/notes"
	internalrouterapiv1tag "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/tag"
	internalrouterapiv1tags "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/tags"
	internalrouterapiv1user "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/user"
)

var (
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Path:       "/v1",
		Controller: Controller,
		Submodules: gonethttp.NewSubmodules(
			internalrouterapiv1auth.Module,
			internalrouterapiv1note.Module,
			internalrouterapiv1notes.Module,
			internalrouterapiv1tag.Module,
			internalrouterapiv1tags.Module,
			internalrouterapiv1user.Module,
		),
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterRoute(
				"GET /ping",
				Controller.Ping,
			)
		},
	}
)
