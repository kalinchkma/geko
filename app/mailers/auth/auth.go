package mailers

import "ganja/internal/interfaces"

type AuthMailer struct{}

func (a *AuthMailer) Welcome(mailer *interfaces.Mailer, email string) string {

	return "Welcome to the ganja"
}
