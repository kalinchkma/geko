package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// This function is responsible for registering routes
func (s *Server) RegisterRoutes() http.Handler {
	router := gin.Default()

	return router
}
