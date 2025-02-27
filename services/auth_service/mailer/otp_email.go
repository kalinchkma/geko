package authmailer

import (
	"bytes"
	"fmt"
	"html/template"
)

type OtpEmailTemplateData struct {
	Email string
	Otp   string
}

func (authMailer *AuthMailer) OnboardOTPEmail(templData OtpEmailTemplateData) {

	// emailBody, err := helper.LoadHtmlTemplateToString(FS, "templates/onboard.templ", templData)
	// if err != nil {
	// 	// @TODO implement error logging
	// 	fmt.Println("Template parsing error:", err)
	// 	return
	// }

	temp, err := template.ParseFS(FS, "templates/onboard.templ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var tempBuffer bytes.Buffer
	if err := temp.Execute(&tempBuffer, templData); err != nil {
		fmt.Printf("error parsing template: %s", err)
	}

	// Send the email to user
	if err := (*authMailer.mailer).SendHTML([]string{templData.Email}, "OTP", tempBuffer.String()); err != nil {
		// @TODO implement error logging
		fmt.Println(err)
	}
}
