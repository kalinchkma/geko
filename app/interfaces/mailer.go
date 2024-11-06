package config

type Mailer interface {
	// send email
	SendEmail(from string, to []string, subject string, msg string)
}
