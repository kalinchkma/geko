package authmailer

import (
	"fmt"
	"geko/internal/helper"
)

type OtpEmailTemplateData struct {
	Email   string
	Otp     string
	AppName string
}

func (authMailer *AuthMailer) SendOTPEmail(templData OtpEmailTemplateData) {

	emailBody, err := helper.LoadHtmlTemplateToString(FS, "templates/onboard.templ", templData)
	if err != nil {
		// @TODO implement error logging
		fmt.Println("Template parsing error:", err)
		return
	}

	// Send the email to user
	if err := (*authMailer.mailer).SendHTML([]string{templData.Email}, "OTP", emailBody); err != nil {
		// @TODO implement error logging
		fmt.Println(err)
	}
}
