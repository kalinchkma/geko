package authcontroller

import (
	authmailer "geko/auth/mailers"
	"geko/internal/server"
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
