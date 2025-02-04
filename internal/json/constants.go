package json

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gonethttpjson "github.com/ralvarezdev/go-net/http/json"
)

var (
	// Encoder is the JSON encoder
	Encoder = gonethttpjson.NewDefaultStreamEncoder(goflagsmode.ModeFlag)

	// Decoder is the JSON decoder
	Decoder, _ = gonethttpjson.NewDefaultStreamDecoder(
		goflagsmode.ModeFlag,
		Encoder,
	)
)
