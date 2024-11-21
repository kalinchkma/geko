package routes

import (
	authController "ganja/controller/auth"
	"ganja/interfaces"

	"github.com/gin-gonic/gin"
)

// routes registry function
func RegisterAuthRoutes(actx *interfaces.AppContext, rootRouter *gin.Engine) {
	// create new route group
	router := rootRouter.Group("/auth")

	router.POST("/login", func(ctx *gin.Context) {
		authController.Login(actx, ctx)
	})

	router.POST("/register", func(ctx *gin.Context) {
		authController.Register(actx, ctx)
	})

	router.POST("/verify-otp", func(ctx *gin.Context) {
		authController.VerifyOtp(actx, ctx)
	})

	router.POST("/resend-otp", func(ctx *gin.Context) {
		authController.ResendOtp(actx, ctx)
	})

	router.POST("/forgot-password", func(ctx *gin.Context) {
		authController.ForgotPassword(actx, ctx)
	})

	router.POST("/reset-password", func(ctx *gin.Context) {
		authController.ResetPassword(actx, ctx)
	})

	router.POST("/refresh-token", func(ctx *gin.Context) {
		authController.RefreshToken(actx, ctx)
	})
}
