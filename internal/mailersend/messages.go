package mailersend

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mailersend/mailersend-go"
	"github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal"
	internallogger "github.com/ralvarezdev/uru-frameworks-secure-notes-api/internal/logger"
)

const (
	// WelcomeEmailSubject is the subject for the welcome email
	WelcomeEmailSubject = "Welcome to Secure Notes"

	// WelcomeEmailHTML is the HTML content for the welcome email
	WelcomeEmailHTML = `<p>Welcome to Secure Notes!</p>
<p>Thank you for signing up. We hope you enjoy using our service.</p>
<p>Best regards,</p>
<p>The Secure Notes Team</p>`

	// WelcomeEmailText is the text content for the welcome email
	WelcomeEmailText = `Welcome to Secure Notes!
Thank you for signing up. We hope you enjoy using our service.
Best regards,
The Secure Notes Team`

	// VerificationEmailSubject is the subject for the verification email
	VerificationEmailSubject = "Verify your email"

	// VerificationEmailHTML is the HTML content for the verification email
	VerificationEmailHTML = `<p>Click the link below to verify your email:</p>
<a href="%s">%s</a>`

	// VerificationEmailText is the text content for the verification email
	VerificationEmailText = `Opps! Your email client does not support HTML. Please copy and paste the link below in your browser to verify your email:
%s`

	// ResetPasswordEmailSubject is the subject for the reset password email
	ResetPasswordEmailSubject = "Reset your password"

	// ResetPasswordEmailHTML is the HTML content for the reset password email
	ResetPasswordEmailHTML = `<p>Click the link below to reset your password:</p>
<a href="%s">%s</a>`

	// ResetPasswordEmailText is the text content for the reset password email
	ResetPasswordEmailText = `Opps! Your email client does not support HTML. Please copy and paste the link below in your browser to reset your password:
%s`
)

// NewMessage creates a new mail message with the default from
func NewMessage() *mailersend.Message {
	mail := Client.Email.NewMessage()
	mail.SetFrom(From)
	return mail
}

// NewSingleRecipientMessage creates a new mail message with a single recipient
func NewSingleRecipientMessage(
	fullName string,
	email string,
) *mailersend.Message {
	mail := NewMessage()
	mail.SetRecipients(
		[]mailersend.Recipient{
			{
				Name:  fullName,
				Email: email,
			},
		},
	)
	return mail
}

// SendWelcomeEmail sends an email to welcome the user
func SendWelcomeEmail(
	fullName string,
	email string,
) {
	// Create a new mail message
	mail := NewSingleRecipientMessage(fullName, email)
	mail.SetSubject(WelcomeEmailSubject)
	mail.SetHTML(WelcomeEmailHTML)
	mail.SetText(WelcomeEmailText)

	// Send the email
	_, err := Client.Email.Send(context.Background(), mail)
	if err != nil {
		internallogger.Api.FailedToSendWelcomeEmail(email, err)
		return
	}
	internallogger.Api.SentWelcomeEmail(email)
}

// SendVerificationEmail sends an email to verify the email
func SendVerificationEmail(
	fullName string,
	email string,
	emailVerificationToken uuid.UUID,
) {
	// Format the verification URL
	verificationURL := fmt.Sprintf(
		"%s/%s",
		internal.VerifyEmailURL,
		emailVerificationToken,
	)

	// Create a new mail message
	mail := NewSingleRecipientMessage(fullName, email)
	mail.SetSubject(VerificationEmailSubject)
	mail.SetHTML(
		fmt.Sprintf(
			VerificationEmailHTML, verificationURL, verificationURL,
		),
	)
	mail.SetText(
		fmt.Sprintf(
			VerificationEmailText, verificationURL,
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

// SendResetPasswordEmail sends an email to reset the password
func SendResetPasswordEmail(
	fullName string,
	email string,
	resetPasswordToken uuid.UUID,
) {
	// Format the reset password URL
	resetPasswordURL := fmt.Sprintf(
		"%s/%s",
		internal.ResetPasswordURL,
		resetPasswordToken,
	)

	// Create a new mail message
	mail := NewSingleRecipientMessage(fullName, email)
	mail.SetSubject(ResetPasswordEmailSubject)
	mail.SetHTML(
		fmt.Sprintf(
			ResetPasswordEmailHTML, resetPasswordURL, resetPasswordURL,
		),
	)
	mail.SetText(
		fmt.Sprintf(
			ResetPasswordEmailText, resetPasswordURL,
		),
	)

	// Send the email
	_, err := Client.Email.Send(context.Background(), mail)
	if err != nil {
		internallogger.Api.FailedToSendResetPasswordEmail(email, err)
		return
	}
	internallogger.Api.SentResetPasswordEmail(email)
}
