package routes

import (
	"ganja/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterCheckerRoutes(actx *interfaces.AppContext, rootRouter *gin.Engine) {
	router := rootRouter.Group("/")
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, actx.DB.Health())
	})
}
