package note

import (
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
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
		Path:       "/note",
		Service:    Service,
		Controller: Controller,
		Middlewares: &[]func(http.Handler) http.Handler{
			internalmiddleware.Authenticate(gojwtinterception.AccessToken),
		},
		Submodules: gonethttp.NewSubmodules(
			internalrouteapiv1noteversion.Module,
			internalrouteapiv1noteversions.Module,
		),
		RegisterRoutesFn: func(m *gonethttp.Module) {
			m.RegisterRoute(
				"POST /",
				Controller.CreateNote,
				internalmiddleware.Validate(
					&CreateNoteRequest{},
				),
			)
			m.RegisterRoute(
				"PUT /",
				Controller.UpdateNote,
				internalmiddleware.Validate(
					&UpdateNoteRequest{},
				),
			)
			m.RegisterRoute(
				"DELETE /",
				Controller.DeleteNote,
				internalmiddleware.Validate(
					&DeleteNoteRequest{},
				),
			)
			m.RegisterRoute(
				"GET /",
				Controller.GetNote,
				internalmiddleware.Validate(
					&GetNoteRequest{},
				),
			)
			m.RegisterRoute(
				"GET /tags",
				Controller.ListNoteTags,
				internalmiddleware.Validate(
					&ListNoteTagsRequest{},
				),
			)
		},
	}
)
