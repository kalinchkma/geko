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

	router.POST("/register", func(ctx *gin.Context) {
		authController.Register(actx, ctx)
	})

}
