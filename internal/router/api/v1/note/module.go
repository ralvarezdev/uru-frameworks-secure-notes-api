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
				"POST /pin",
				Controller.PinNote,
				internalmiddleware.Validate(
					&PinNoteRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /pin",
				Controller.UnpinNote,
				internalmiddleware.Validate(
					&UnpinNoteRequest{},
				),
			)
			m.RegisterExactRoute(
				"POST /archive",
				Controller.ArchiveNote,
				internalmiddleware.Validate(
					&ArchiveNoteRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /archive",
				Controller.UnarchiveNote,
				internalmiddleware.Validate(
					&UnarchiveNoteRequest{},
				),
			)
			m.RegisterExactRoute(
				"POST /trash",
				Controller.TrashNote,
				internalmiddleware.Validate(
					&TrashNoteRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /trash",
				Controller.UntrashNote,
				internalmiddleware.Validate(
					&UntrashNoteRequest{},
				),
			)
			m.RegisterExactRoute(
				"POST /star",
				Controller.StarNote,
				internalmiddleware.Validate(
					&StarNoteRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /star",
				Controller.UnstarNote,
				internalmiddleware.Validate(
					&UnstarNoteRequest{},
				),
			)
		},
	}
)
