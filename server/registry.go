package server

import (
	"ganja/models"
	"ganja/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register all your route group here
func (s *Server) routesRegistry() http.Handler {
	router := gin.Default()

	// register auth routes
	routes.RegisterAuthRoutes(s.actx, router)

	// register user routes
	routes.RegisterUserRoutes(s.actx, router)

	// register server check routes
	routes.RegisterCheckerRoutes(s.actx, router)

	return router
}

// load adition dependency services
func (s *Server) bootstrap() {
	// register models
	db := s.actx.DB.GetDB()

	// auto migrate user model
	db.AutoMigrate(&models.User{})
}
