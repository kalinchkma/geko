package auth_mailer

import (
	"ganja/initializers/mailers"
	"ganja/models"
)

type AuthMailer struct{}

type Welcome struct {
	Name           string
	ActivationLink string
}

// Send welcome and active link
func (a *AuthMailer) Welcome(mailer *mailers.Mailer, user models.User) {

}

func (a *AuthMailer) SendOTP(mailer *mailers.Mailer, user models.User) {

}
