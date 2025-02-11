package note

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
	internalrouterapiv1notetags "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note/tags"
	internalrouterapiv1noteversion "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note/version"
	internalrouterapiv1noteversions "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note/versions"
)

var (
	Service    = &service{}
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Pattern:    "/note",
		Service:    Service,
		Controller: Controller,
		BeforeLoadFn: func(m *gonethttp.Module) {
			m.Middlewares = gonethttp.NewMiddlewares(
				internalmiddleware.Authenticate,
			)
		},
		Submodules: gonethttp.NewSubmodules(
			internalrouterapiv1noteversion.Module,
			internalrouterapiv1noteversions.Module,
			internalrouterapiv1notetags.Module,
		),
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterExactRoute(
				"POST /",
				Controller.CreateUserNote,
				internalmiddleware.Validate(
					&CreateUserNoteRequest{},
				),
			)
			m.RegisterExactRoute(
				"PUT /",
				Controller.UpdateUserNote,
				internalmiddleware.Validate(
					&UpdateUserNoteRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /",
				Controller.DeleteUserNote,
				internalmiddleware.Validate(
					&DeleteUserNoteRequest{},
				),
			)
			m.RegisterExactRoute(
				"GET /",
				Controller.GetUserNoteByID,
				internalmiddleware.Validate(
					&GetUserNoteByIDRequest{},
				),
			)
			m.RegisterExactRoute(
				"PUT /pin",
				Controller.UpdateUserNotePin,
				internalmiddleware.Validate(
					&UpdateUserNotePinRequest{},
				),
			)
			m.RegisterExactRoute(
				"PUT /archive",
				Controller.UpdateUserNoteArchive,
				internalmiddleware.Validate(
					&UpdateUserNoteArchiveRequest{},
				),
			)
			m.RegisterExactRoute(
				"PUT /trash",
				Controller.UpdateUserNoteTrash,
				internalmiddleware.Validate(
					&UpdateUserNoteTrashRequest{},
				),
			)
			m.RegisterExactRoute(
				"PUT /star",
				Controller.UpdateUserNoteStar,
				internalmiddleware.Validate(
					&UpdateUserNoteStarRequest{},
				),
			)
		},
	}
)
