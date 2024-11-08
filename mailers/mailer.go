package mailers

import (
	"crypto/tls"
	"fmt"
	"ganja/interfaces"

	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type mailer struct {
	dialer *gomail.Dialer
}

var (
	smtpHost       = os.Getenv("SMTP_HOST")
	smtpPort       = os.Getenv("SMTP_PORT")
	user           = os.Getenv("SMTP_USER")
	password       = os.Getenv("SMTP_PASSWORD")
	mailerInstance *mailer
)

func New() interfaces.Mailer {
	// Reuse mailer connection
	if mailerInstance != nil {
		return mailerInstance
	}

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Fatal("PORT must be an interger")
	}
	// configure mailer connection
	dialer := gomail.NewDialer(smtpHost, port, user, password)

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	mailerInstance = &mailer{
		dialer,
	}
	return mailerInstance
}

func (m *mailer) SendEmail(from string, to []string, subject string, body string) {
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
