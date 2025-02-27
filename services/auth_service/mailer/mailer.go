package authmailer

import (
	"embed"
	"geko/internal/mailers"
)

//go:embed "templates"
var FS embed.FS

type AuthMailer struct {
	mailer *mailers.Mailer
}

func NewAuthMailer(mailer *mailers.Mailer) *AuthMailer {
	return &AuthMailer{
		mailer,
	}
}
