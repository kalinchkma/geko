package mailers

import (
	"log"
	"net/smtp"
	"os"
)

type Mailer interface {
	// send email
	SendEmail(from string, to []string, msg string) bool
}

type mailer struct {
	auth *smtp.Auth
}

var (
	smtpHost       = os.Getenv("SMTP_HOST")
	smtpPort       = os.Getenv("SMTP_PORT")
	username       = os.Getenv("SMTP_USERNAME")
	password       = os.Getenv("SMTP_PASSWORD")
	mailerInstance *mailer
)

func New() Mailer {
	// Reuse mailer connection
	if mailerInstance != nil {
		return mailerInstance
	}

	// Auth setup
	auth := smtp.PlainAuth("", username, password, smtpHost)

	mailerInstance = &mailer{
		auth: &auth,
	}
	return mailerInstance
}

func (m *mailer) SendEmail(from string, to []string, msg string) bool {
	// @TODO implement to send a email
	err := smtp.SendMail(smtpHost+":"+smtpPort, *m.auth, from, to, []byte(msg))

	if err != nil {
		log.Fatalln("Error sending email")
		return false
	}
	return true
}
