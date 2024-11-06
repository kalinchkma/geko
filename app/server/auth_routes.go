package server

import (
	"ganja/app/handlers/auth"

	"github.com/gin-gonic/gin"
)

func (s *Server) Register(ctx *gin.Context) {
	auth.Register(s.cfg, ctx)
}
