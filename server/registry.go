package server

import (
	"geko/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register all your route group here
func (s *Server) routesRegistry() http.Handler {
	router := gin.New()

	// register auth routes
	routes.RegisterAuthRoutes(s.actx, router)

	// register user routes
	routes.RegisterUserRoutes(s.actx, router)

	// register server check routes
	routes.RegisterCheckerRoutes(s.actx, router)

	return router
}
