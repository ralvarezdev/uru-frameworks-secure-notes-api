package v1

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
	internalrouterapiv1auth "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/auth"
	internalrouterapiv1note "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note"
	internalrouterapiv1notes "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/notes"
	internalrouterapiv1tag "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/tag"
	internalrouterapiv1tags "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/tags"
	internalrouterapiv1user "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/user"
)

var (
	Service    = &service{}
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Pattern:    "/v1",
		Controller: Controller,
		Service:    Service,
		Submodules: gonethttp.NewSubmodules(
			internalrouterapiv1auth.Module,
			internalrouterapiv1note.Module,
			internalrouterapiv1notes.Module,
			internalrouterapiv1tag.Module,
			internalrouterapiv1tags.Module,
			internalrouterapiv1user.Module,
		),
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterExactRoute(
				"GET /ping",
				Controller.Ping,
			)
			m.RegisterExactRoute(
				"POST /sync",
				Controller.SyncByLastSyncedAt,
				internalmiddleware.AuthenticateAccessToken,
			)
		},
	}
)
