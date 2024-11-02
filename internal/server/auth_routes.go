package server

import (
	"ganja/internal/handlers/auth"

	"github.com/gin-gonic/gin"
)

func (s *Server) Register(ctx *gin.Context) {
	go s.mailer.SendEmail("dev@test.com", []string{"chakma.kalin10211@gmail.com"}, "Hello, good to see you here")
	auth.Register(ctx)
}
