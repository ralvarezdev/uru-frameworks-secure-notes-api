package main

import (
	"flag"
	godatabasespgxpool "github.com/ralvarezdev/go-databases/sql/pgxpool"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gonethttproute "github.com/ralvarezdev/go-net/http/route"
	"github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal"
	internalaes "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/aes"
	internalbcrypt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/bcrypt"
	internaltotp "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/otp/totp"
	internalpbkdf2 "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/pbkdf2"
	internaltoken "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/crypto/token"
	internalpostgres "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/databases/postgres"
	internaljwt "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt"
	internaljwtcache "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/jwt/cache"
	internallistener "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/listener"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
	internalmailersend "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/mailersend"
	internalmiddleware "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/middleware"
	internalrouter "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/router"
	internalvalidator "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/validator"
	"log"
	"net/http"
	"path/filepath"
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
	internal.Load()
	internalaes.Load()
	internalbcrypt.Load()
	internaltotp.Load()
	internalpbkdf2.Load()
	internaltoken.Load()
	internalpostgres.Load(mode)
	internaljwtcache.Load(mode)
	internaljwt.Load()
	internallistener.Load(mode)
	internalvalidator.Load(mode)
	internalmiddleware.Load()
	internalmailersend.Load()
}

//	@Title			Secure Notes REST API
//	@Version		1.0
//	@Description	This is the REST API for the Secure Notes application.

//	@License.name	GPL-3.0
//	@License.url	http://www.gnu.org/licenses/gpl-3.0.html

//	@BasePath	/

// @securityDefinitions.apikey	CookieAuth
// @in							cookie
// @name						access_token
func main() {
	defer func(handler godatabasespgxpool.PoolHandler) {
		handler.Disconnect()
		internallogger.Postgres.DisconnectedFromDatabase()
	}(internalpostgres.PoolHandler)

	// Check if the mode is migrate
	if goflagsmode.ModeFlag.IsMigrate() {
		// Migrate the database
		err := internalpostgres.PoolService.Migrate()
		if err != nil {
			panic(err)
		}
		return
	}

	// Create the main router
	router, err := gonethttproute.NewBaseRouter(
		goflagsmode.ModeFlag,
		internallogger.Router,
	)
	if err != nil {
		panic(err)
	}

	// Log the serving of the Swagger UI
	absPath, err := filepath.Abs("./docs")
	if err != nil {
		panic(err)
	}
	internallogger.Api.ServingSwaggerUI(absPath)

	// Serve the Swaggers docs
	router.ServeStaticFiles(
		"/docs/", absPath,
	)

	// Create the main router module
	if err = internalrouter.Module.Create(router); err != nil {
		panic(err)
	}

	// Serve the API server
	internallogger.Api.ServerStarted(internallistener.Port)
	if err = http.ListenAndServe(
		":"+internallistener.Port,
		router.Handler(),
	); err != nil {
		panic(err)
	}
}
