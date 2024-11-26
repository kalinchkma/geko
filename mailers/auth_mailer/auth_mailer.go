package auth_mailer

import (
	"fmt"
	"ganja/initializers/mailers"
	"ganja/library"
	"ganja/models"
)

type WelcomeTemplate struct {
	Name           string
	ActivationLink string
}

// Send welcome and active link
func Welcome(mailer *mailers.Mailer, user models.User) {
	data := WelcomeTemplate{
		Name:           user.Name,
		ActivationLink: "http://localhost:9090",
	}

	// parse email body
	emailBody, err := library.LoadHtmlTemplateToString("./mailers/auth_mailer/templates/welcome.html", data)

	if err != nil {
		// @TODO implement error logging
		fmt.Println("error executing template: ", err)
	}

	// Send the email to the user
	if err := mailer.SendEmail("no-replay@demomailtrap.com", []string{user.Email}, "Welcome", emailBody); err != nil {
		// @TODO implement error logging
		fmt.Println(err)
	}

}

func SendOTP(mailer *mailers.Mailer, user models.User) {

}
