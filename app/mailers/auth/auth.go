package mailers

import "ganja/app/mailers"

type AuthMailer struct{}

func (a *AuthMailer) Welcome(mailer *mailers.Mailer, email string) string {

	return "Welcome to the ganja"
}
