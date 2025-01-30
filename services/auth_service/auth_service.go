package authservice

import (
	"geko/internal/server"
	"net/http"

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

// Service route registry
func (s *AuthService) Registry() {
	s.route.GET("/login", s.login)
	s.route.GET("/register", s.register)
}

func (s *AuthService) login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Test": s.serverContext.Config.Env})
}

func (s *AuthService) register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Register Test": s.serverContext.Config.Addr})
}
