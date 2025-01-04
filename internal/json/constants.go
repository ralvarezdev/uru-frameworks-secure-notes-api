package json

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gonethttpjson "github.com/ralvarezdev/go-net/http/json"
)

var (
	// Encoder is the JSON encoder
	Encoder = gonethttpjson.NewDefaultEncoder(goflagsmode.Mode)

	// Decoder is the JSON decoder
	Decoder = gonethttpjson.NewDefaultDecoder(goflagsmode.Mode, Encoder)
)
