package server

import (
	"ganja/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register all your route group here
func (s *Server) routesRegistry() http.Handler {
	router := gin.New()

	// pointing to the frontend
	router.StaticFile("/", "frontend/dist/index.html")
	router.StaticFS("/static", http.Dir("frontend/dist/static"))

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

}
