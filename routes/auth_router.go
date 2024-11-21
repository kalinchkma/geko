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

	})

	router.POST("/register", func(ctx *gin.Context) {
		authController.Register(actx, ctx)
	})

	router.POST("/verify-otp", func(ctx *gin.Context) {

	})

	router.POST("/resend-otp", func(ctx *gin.Context) {

	})

	router.POST("/forgot-password", func(ctx *gin.Context) {

	})

	router.POST("/refresh-token", func(ctx *gin.Context) {

	})
}
