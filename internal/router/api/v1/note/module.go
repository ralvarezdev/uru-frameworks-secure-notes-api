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
				Controller.GetUserNote,
				internalmiddleware.Validate(
					&GetUserNoteRequest{},
				),
			)
			m.RegisterExactRoute(
				"GET /tags",
				Controller.ListUserNoteTags,
				internalmiddleware.Validate(
					&ListUserNoteTagsRequest{},
				),
			)
			m.RegisterExactRoute(
				"PATCH /tags",
				Controller.AddUserNoteTags,
				internalmiddleware.Validate(
					&AddUserNoteTagsRequest{},
				),
			)
			m.RegisterExactRoute(
				"DELETE /tags",
				Controller.RemoveUserNoteTags,
				internalmiddleware.Validate(
					&RemoveUserNoteTagsRequest{},
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
