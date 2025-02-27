package authcontroller

import (
	"geko/internal/server"
	authmailer "geko/services/auth_service/mailer"
)

type AuthController struct {
	serverContext *server.HttpServerContext
	mailer        *authmailer.AuthMailer
}

func NewAuthController(serverContext *server.HttpServerContext) *AuthController {
	return &AuthController{
		serverContext: serverContext,
		mailer:        authmailer.NewAuthMailer(&serverContext.Mailer),
	}
}
