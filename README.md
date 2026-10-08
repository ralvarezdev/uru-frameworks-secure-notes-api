# uru-frameworks-secure-notes-api

**Note:** This repository is archived and read-only.

Secure Notes REST API from the Frameworks college course (URU): a Go service backed by PostgreSQL. Its front end is [uru-frameworks-secure-notes-app](https://github.com/ralvarezdev/uru-frameworks-secure-notes-app).

## Features

From the generated Swagger spec (`docs/swagger.json`):

- Signup, login/logout, refresh tokens, password change/forgot/reset
- Email and phone-number verification and update
- Two-factor authentication: TOTP, email codes and recovery codes
- Notes with archive, pin, star and trash states, tags and versions
- Sync endpoints for notes, tags and a combined sync, for offline-capable clients
- User profile and username management

## Tech stack

Go 1.23.4, PostgreSQL via `pgx/v5` (tables and stored procedures defined under `internal/databases/postgres/model`), JWT, bcrypt/PBKDF2/AES/TOTP helpers, MailerSend for email, Swagger docs generated with `swag` and served from `./docs` at `/docs/`, plus several author-owned `go-*` libraries.

## Project structure

- **`cmd/server/main.go`** — entry point (mode flag, config loading, API and Swagger)
- **`internal/router/api/v1/`** — route modules `auth`, `note`, `notes`, `tag`, `tags`, `user`
- **`internal/`** — also `databases/postgres`, `crypto`, `jwt`, `mailersend`, `middleware`
- **`docs/`** — Swagger UI, `swagger.json`, `swagger.yaml` (full list of operations under base path `/api/v1`)
- **`Dockerfile`** — multi-stage Alpine build, exposes 8080, runs with `-m=prod`

## Configuration

Environment variables (loaded from `.env`), prefixed `URU_FRAMEWORKS_SECURE_NOTES_` except `PORT`. Names come from the `constants.go` files; check the code for the exact required set.

- **Server and database** — `HOST`, `PORT`, `BODY_LIMIT`, `POSTGRES_DSN`, `POSTGRES_MAX_IDLE_CONNECTIONS`, `POSTGRES_MAX_OPEN_CONNECTIONS`
- **JWT and tokens** — `JWT_PRIVATE_KEY`, `JWT_PUBLIC_KEY`, `ACCESS_TOKEN_DURATION`, `REFRESH_TOKEN_DURATION`, `EMAIL_VERIFICATION_TOKEN_DURATION`, `RESET_PASSWORD_TOKEN_DURATION`, `RESET_PASSWORD_URL`, `VERIFY_EMAIL_URL`
- **Crypto** — `BCRYPT_COST`, `AES_KEY_SIZE`, `PBKDF2_ITERATIONS`, `PBKDF2_KEY_LENGTH`, `PBKDF2_SALT_LENGTH`
- **2FA** — `2FA_EMAIL_CODE_DURATION`, `2FA_EMAIL_CODE_LENGTH`, `TOTP_DIGITS`, `TOTP_PERIOD`, `TOTP_SECRET_LENGTH`, `TOTP_RECOVERY_CODES_COUNT`, `TOTP_RECOVERY_CODES_LENGTH`
- **Policy** — `MINIMUM_PASSWORD_LENGTH`, `MINIMUM_PASSWORD_CAPS_COUNT`, `MINIMUM_PASSWORD_NUMBERS_COUNT`, `MINIMUM_PASSWORD_SPECIAL_COUNT`, `MINIMUM_AGE`, `MAXIMUM_AGE`, `MAXIMUM_FAILED_ATTEMPTS_COUNT`, `MAXIMUM_FAILED_ATTEMPTS_PERIOD`
- **Email** — `MAILER_SEND_API_KEY`, `MAILER_SEND_DOMAIN`

## Build and run

```bash
cd cmd/server && go build -o ../../bin/server
./bin/server -m=dev                                              # from the repo root so ./docs resolves
swag init -g cmd/server/main.go --parseDependency --parseInternal  # regenerate Swagger docs
```

**Note:** the Dockerfile runs `go build` at the repository root while `main.go` lives in `cmd/server`, so check the build path before relying on the image. No tests are included.

## License

GNU General Public License v3.0 (see `LICENSE`).
