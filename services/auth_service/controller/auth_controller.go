package authcontroller

import (
	"geko/internal/server"
)

type AuthController struct {
	serverContext *server.HttpServerContext
}

func NewAuthController(serverContext *server.HttpServerContext) *AuthController {
	return &AuthController{
		serverContext: serverContext,
	}
}
