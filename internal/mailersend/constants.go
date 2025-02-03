package mailersend

import (
	"github.com/mailersend/mailersend-go"
	internalloader "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/loader"
)

const (
	// EnvAPIKey is the key for the API key of the mailer send service in the environment variables
	EnvAPIKey = "URU_FRAMEWORKS_SECURE_NOTES_MAILER_SEND_API_KEY"

	// EnvDomain is the key for the domain of the mailer send service in the environment variables
	EnvDomain = "URU_FRAMEWORKS_SECURE_NOTES_MAILER_SEND_DOMAIN"
)

var (
	// APIKey is the API key of the mailer send service
	APIKey string

	// Domain is the domain of the mailer send service
	Domain string

	// From is the email of the mailer send service
	From mailersend.From

	// Client is the client for the mailer send service
	Client *mailersend.Mailersend
)

// Load loads the Mailer
func Load() {
	// Load the environment variables
	for env, dest := range map[string]*string{
		EnvAPIKey: &APIKey,
		EnvDomain: &Domain,
	} {
		if err := internalloader.Loader.LoadVariable(
			env,
			dest,
		); err != nil {
			panic(err)
		}
	}

	// Create a new MailerSend client
	Client = mailersend.NewMailersend(APIKey)

	// Set the email for the mailer send service
	From = mailersend.From{
		Name:  "Secure Notes",
		Email: "noreply@" + Domain,
	}
}
