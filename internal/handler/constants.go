package handler

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gonethttphandler "github.com/ralvarezdev/go-net/http/handler"
	internaljson "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/json"
)

var (
	// Handler is the default handler for the requests and responses
	Handler, _ = gonethttphandler.NewDefaultHandler(
		goflagsmode.ModeFlag,
		internaljson.Encoder,
		internaljson.Decoder,
	)
)
