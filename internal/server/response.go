package server

import "github.com/gin-gonic/gin"

func SuccessJSONResponse(ctx *gin.Context, code int, message string, data any) {
	response := gin.H{
		"status":  true,
		"message": message,
		"data":    data,
	}
	ctx.JSON(code, response)
}

func ErrorJSONResponse(ctx *gin.Context, code int, message string, data any) {
	response := gin.H{
		"status":  false,
		"message": message,
		"errors":  data,
	}
	ctx.JSON(code, response)
}
