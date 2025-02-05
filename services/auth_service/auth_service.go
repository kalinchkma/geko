package authservice

import (
	"geko/internal/server"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	serverContext *server.HttpServerContext
	route         *gin.RouterGroup
}

// Service constructor
func (s *AuthService) Mount(serverContext *server.HttpServerContext, route *gin.RouterGroup) {
	s.serverContext = serverContext
	s.route = route
}

// Service routes
func (s *AuthService) Routes() {
	s.route.GET("/register", s.registerHandler)
}
