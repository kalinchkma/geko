package authmailer

import (
	"embed"

	"github.com/kalinchkma/geko/internal/mailers"
)

//go:embed "templates"
var FS embed.FS

type AuthMailer struct {
	mailer       *mailers.Mailer
	managerEmail string
}

func NewAuthMailer(mailer *mailers.Mailer, managerEmail string) *AuthMailer {

	return &AuthMailer{
		mailer:       mailer,
		managerEmail: managerEmail,
	}
}
