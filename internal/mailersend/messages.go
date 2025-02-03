package mailersend

import (
	"context"
	"fmt"
	"github.com/mailersend/mailersend-go"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
)

// SendVerificationEmail sends an email to verify the email
func SendVerificationEmail(
	fullName string,
	email string,
	verificationURL string,
) {
	// Create a new mail message
	mail := Client.Email.NewMessage()
	mail.SetFrom(From)
	mail.SetRecipients(
		[]mailersend.Recipient{
			{
				Name:  fullName,
				Email: email,
			},
		},
	)
	mail.SetSubject("Verify your email")
	mail.SetHTML(
		fmt.Sprintf(
			`
		<p>Click the link below to verify your email:</p>
		<a href="%s">%s</a>
	`, verificationURL, verificationURL,
		),
	)
	mail.SetText(
		fmt.Sprintf(
			`
		Opps! Your email client does not support HTML. Please copy and paste the link below in your browser to verify your email:
		%s
	`, verificationURL,
		),
	)

	// Send the email
	_, err := Client.Email.Send(context.Background(), mail)
	if err != nil {
		internallogger.Api.FailedToSendVerificationEmail(email, err)
		return
	}
	internallogger.Api.SentVerificationEmail(email)
}
