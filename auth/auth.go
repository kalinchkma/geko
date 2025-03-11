package auth

import (
	authcontroller "geko/auth/controller"
	"geko/internal/server"

	"github.com/gin-gonic/gin"
)

// Auth service struct
type AuthService struct {
	serverContext *server.HttpServerContext
	router        *gin.RouterGroup
	controller    *authcontroller.AuthController
}

// Service constructor
func (s *AuthService) Mount(serverContext *server.HttpServerContext, router *gin.RouterGroup) {
	// Configure service nessasary dependencies
	s.serverContext = serverContext
	s.router = router
	s.controller = authcontroller.NewAuthController(s.serverContext)
}
