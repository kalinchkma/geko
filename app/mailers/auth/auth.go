package mailers

import (
	config "ganja/app/interfaces"
)

type AuthMailer struct{}

func (a *AuthMailer) Welcome(mailer *config.Mailer, email string) string {

	return "Welcome to the ganja"
}
