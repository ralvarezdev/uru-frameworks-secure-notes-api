package main

import (
	"database/sql"
	"flag"
	_ "github.com/jackc/pgx/v5/stdlib"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	gosecurityheadersnethttp "github.com/ralvarezdev/go-security-headers/net/http"
	internalbcrypt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/bcrypt"
	internaltotp "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/otp/totp"
	internalpbkdf2 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/pbkdf2"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internaljwt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt"
	internaljwtcache "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/cache"
	internallistener "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/listener"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
	internalrouter "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
	"log"
	"net/http"
)

// init initializes the flags and calls the load functions
func init() {
	// Parse the flags
	flag.Parse()

	// Log the mode flag
	mode := goflagsmode.ModeFlag
	log.Printf("Running in %s mode...\n", mode.Value())

	// Call the load functions
	internalloader.Load()
	internalbcrypt.Load()
	internaltotp.Load()
	internalpbkdf2.Load()
	internalpostgres.Load()
	internaljwtcache.Load(mode)
	internaljwt.Load()
	internallistener.Load()
	internalvalidator.Load(mode)
	internalmiddleware.Load()
}

func main() {
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
		internallogger.Postgres.DisconnectedFromDatabase()
	}(internalpostgres.DB)

	// Create the main router module
	if err := internalrouter.Module.Create(
		gonethttproute.NewRouter(
			"",
			goflagsmode.ModeFlag,
			internallogger.Router,
			gosecurityheadersnethttp.Handler,
		),
	); err != nil {
		panic(err)
	}

	// Serve the API server
	internallogger.Api.ServerStarted(internallistener.Port)
	if err := http.ListenAndServe(
		":"+internallistener.Port,
		internalrouter.Module.Handler(),
	); err != nil {
		panic(err)
	}
}
