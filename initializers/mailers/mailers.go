package mailers

import (
	"crypto/tls"
	"fmt"

	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type Mailer struct {
	dialer *gomail.Dialer
}

var (
	smtpHost = os.Getenv("SMTP_HOST")
	smtpPort = os.Getenv("SMTP_PORT")
	user     = os.Getenv("SMTP_USER")
	password = os.Getenv("SMTP_PASSWORD")
	mailer   *Mailer
)

func New() *Mailer {
	// Reuse mailer connection
	if mailer != nil {
		return mailer
	}

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Fatal("PORT must be an interger")
	}
	// configure mailer connection
	dialer := gomail.NewDialer(smtpHost, port, user, password)

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	mailer = &Mailer{
		dialer,
	}
	return mailer
}

func (m *Mailer) SendEmail(from string, to []string, subject string, body string) {
	fmt.Println("Email body", body)
	newMessage := gomail.NewMessage()
	newMessage.SetHeader("From", from)
	newMessage.SetHeader("To", to...)
	newMessage.SetHeader("Subject", subject)
	newMessage.SetBody("text/html", body)

	if err := (*m).dialer.DialAndSend(newMessage); err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}
}
