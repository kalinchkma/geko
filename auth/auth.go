package auth

import (
	"github.com/gin-gonic/gin"
	authcontroller "github.com/kalinchkma/geko/auth/controller"
	"github.com/kalinchkma/geko/internal/server"
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
