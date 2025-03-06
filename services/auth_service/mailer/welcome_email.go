package authmailer

import (
	"fmt"
	"geko/internal/utils"
)

type WelcomeEmailTemplateData struct {
	Name    string
	AppName string
}

func (authMailer *AuthMailer) SendWelcomeEmail(templData WelcomeEmailTemplateData) {
	emailBody, err := utils.LoadHtmlTemplateToString(FS, "templates/welcome.templ", templData)
	if err != nil {
		// @TODO implement error logging
		fmt.Println("Template parsing error:", err)
		return
	}

	// Send the email to user
	if err := (*authMailer.mailer).SendHTML(authMailer.managerEmail, []string{templData.Name}, "Welcome", emailBody); err != nil {
		fmt.Println(err)
		return
	}
}
