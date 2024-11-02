package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := gin.Default()

	// API router group
	api := router.Group("/api")

	// auth api
	auth := api.Group("/auth")
	{
		auth.POST("/register", s.Register)
	}

	// view route
	router.GET("/", func(ctx *gin.Context) {
		ctx.SecureJSON(http.StatusOK, map[string]string{"message": "Default view routes"})
	})

	return router
}
