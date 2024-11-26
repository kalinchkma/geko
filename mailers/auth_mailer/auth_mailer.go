package auth_mailer

import (
	"bytes"
	"fmt"
	"ganja/initializers/mailers"
	"ganja/models"
	"html/template"
)

type WelcomeTemplate struct {
	Name           string
	ActivationLink string
}

var welcomeTemplate = template.Must(template.ParseFiles("./mailers/auth_mailer/templates/welcome.html"))

// Send welcome and active link
func Welcome(mailer *mailers.Mailer, user models.User) {
	data := WelcomeTemplate{
		Name:           user.Name,
		ActivationLink: "http://localhost:9090",
	}

	var tempBuffer bytes.Buffer

	if err := welcomeTemplate.Execute(&tempBuffer, data); err != nil {
		// @TODO implement error logging
		fmt.Println("error executing template: ", err)
	}
	// Use the buffer's content as the email body
	emailBody := tempBuffer.String()

	// Send the email to the user
	if err := mailer.SendEmail("no-replay@demomailtrap.com", []string{user.Email}, "Welcome", emailBody); err != nil {
		// @TODO implement error logging
		fmt.Println(err)
	}

}

func SendOTP(mailer *mailers.Mailer, user models.User) {

}
