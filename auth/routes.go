package auth

// Route handler
func (s *AuthService) RouteHandler() {

	// Service controller
	controller := s.controller

	// Service router
	router := s.router

	// User register routes
	router.POST("/register", controller.Register)

	// User login routes
	router.POST("/login", controller.Login)

	// Verify otp routes
	router.POST("/verify-opt", controller.VerifyOtp)

	// Resend otp routes
	router.POST("/resend-opt", controller.ResendOTP)

	// Forgot password routes
	router.POST("/forgot-password", controller.ForgotPassword)

	// Reset password routes
	router.POST("/reset-password", controller.ResetPassword)

	// Refresh auth token routes
	router.POST("/refresh", controller.RefreshToken)

	// Revoke auth token routes
	router.POST("/revoke", controller.RefreshToken)
}
