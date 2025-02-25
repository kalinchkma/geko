package mailers

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

type MailerConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type Mailer interface {
	Send(to []string, subject, body string) error
	SendWithAttachments(to []string, subject, body string, attachements []string) error
	SendHTML(to []string, subject, htmlBody string) error
}

type mailer struct {
	dialer *gomail.Dialer
}

func NewMailer(mailerConfig MailerConfig) *mailer {
	// Create new dialer
	dialer := gomail.NewDialer(mailerConfig.Host, mailerConfig.Port, mailerConfig.Username, mailerConfig.Password)

	// For development skiping SSL checking
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// for production
	// dialer.TLSConfig = &tls.Config{
	// 	ServerName: mailerConfig.Host, // SMTP server hostname
	// }

	return &mailer{
		dialer: dialer,
	}
}

// Send email plain text body
func (m *mailer) Send(to []string, subject, body string) error {
	fmt.Println("test from email", to)
	return m.sendEmail(to, subject, body, false, nil)
}

// Send email with attachments
func (m *mailer) SendWithAttachments(to []string, subject, body string, attachments []string) error {
	return m.sendEmail(to, subject, body, false, attachments)
}

// Send email with HTML body
func (m *mailer) SendHTML(to []string, subject, htmlBody string) error {
	return m.sendEmail(to, subject, htmlBody, true, nil)
}

// Send email helper function
func (m *mailer) sendEmail(to []string, subject, body string, isHTML bool, attachments []string) error {
	msg := gomail.NewMessage()

	msg.SetHeader("From", "no-replay@demomailtrap.com") // @TODO make dynamic of sender email
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", subject)

	// Set body based on type
	if isHTML {
		msg.SetBody("text/html", body)
	} else {
		msg.SetBody("text/plain", body)
	}

	// Attach files if any
	for _, file := range attachments {
		msg.Attach(file)
	}

	return m.dialer.DialAndSend(msg)
}
