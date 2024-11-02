package server

import (
	"ganja/internal/handlers/auth"

	"github.com/gin-gonic/gin"
)

func (s *Server) Register(ctx *gin.Context) {

	auth.Register(&s.mailer, ctx)
}
