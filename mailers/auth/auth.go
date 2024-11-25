package auth_mailer

import (
	"ganja/initializers/mailers"
)

type AuthMailer struct{}

func (a *AuthMailer) Welcome(mailer *mailers.Mailer, email string) string {

	return "Welcome to the ganja"
}
