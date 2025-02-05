package note

import (
	gonethttp "github.com/ralvarezdev/go-net/http"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
	internalrouteapiv1noteversion "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note/version"
	internalrouteapiv1noteversions "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router/api/v1/note/versions"
	"net/http"
)

var (
	Service    = &service{}
	Controller = &controller{}
	Module     = &gonethttp.Module{
		Pattern:    "/note",
		Service:    Service,
		Controller: Controller,
		BeforeLoadFn: func(m *gonethttp.Module) {
			m.Middlewares = &[]func(http.Handler) http.Handler{
				internalmiddleware.AuthenticateAccessToken,
			}
		},
		Submodules: gonethttp.NewSubmodules(
			internalrouteapiv1noteversion.Module,
			internalrouteapiv1noteversions.Module,
		),
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterExactRoute(
				"POST /",
				Controller.CreateNote,
				internalmiddleware.Validate(
					&CreateNoteRequest{},
				),
			)
			m.RegisterExactRoute(
				"PUT /",
				Controller.UpdateNote,
				internalmiddleware.Validate(
					&UpdateNoteRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /",
				Controller.DeleteNote,
				internalmiddleware.Validate(
					&DeleteNoteRequest{},
				),
			)
			m.RegisterExactRoute(
				"GET /",
				Controller.GetNote,
				internalmiddleware.Validate(
					&GetNoteRequest{},
				),
			)
			m.RegisterExactRoute(
				"GET /tags",
				Controller.ListNoteTags,
				internalmiddleware.Validate(
					&ListNoteTagsRequest{},
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
