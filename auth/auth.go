package auth

import (
	authcontroller "geko/auth/controller"
	"geko/internal/server"

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

// Route handler
func (s *AuthService) RouteHandler(c *authcontroller.AuthController) {

	s.route.POST("/register", (*c).Register)
	s.route.POST("/login", (*c).Login)
	s.route.POST("/verify-opt", (*c).VerifyOtp)
	s.route.POST("/resend-opt", (*c).ResendOTP)
	s.route.POST("/forgot-password", (*c).ForgotPassword)
	s.route.POST("/reset-password", (*c).ResetPassword)
	s.route.POST("/refresh", (*c).RefreshToken)
	s.route.POST("/revoke", (*c).RefreshToken)
}
