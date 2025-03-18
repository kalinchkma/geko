package authcontroller

import (
	authmailer "github.com/kalinchkma/geko/auth/mailers"
	"github.com/kalinchkma/geko/internal/server"
)

type AuthController struct {
	serverContext *server.HttpServerContext
	mailer        *authmailer.AuthMailer
}

func NewAuthController(serverContext *server.HttpServerContext) *AuthController {
	managerEmail := "no-replay" + "@" + serverContext.Config.MailerConfig.Domain
	return &AuthController{
		serverContext: serverContext,
		mailer:        authmailer.NewAuthMailer(&serverContext.Mailer, managerEmail),
	}
}
