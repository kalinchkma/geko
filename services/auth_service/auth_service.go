package authservice

import (
	"geko/internal/server"
	authcontroller "geko/services/auth_service/controller"

	"github.com/gin-gonic/gin"
)

// Auth service struct
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
	authController := authcontroller.NewAuthController(s.serverContext)
	s.RouteHandler(authController)
}

// Controller fallback
func (s *AuthService) RouteHandler(c *authcontroller.AuthController) {
	s.route.GET("/register", (*c).Register)
	s.route.GET("/login", (*c).Login)
	s.route.GET("/verify-opt", (*c).VerifyOtp)
	s.route.GET("/resend-opt", (*c).ResendOtp)
	s.route.GET("/forgot-password", (*c).ForgotPassword)
	s.route.GET("/reset-password", (*c).ResetPassword)
	s.route.GET("/refresh-token", (*c).RefreshToken)
}
