package server

import "github.com/gin-gonic/gin"

func SuccessJSONResponse(ctx *gin.Context, code int, message string, data any) {
	response := gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
		"error":   nil,
	}
	ctx.JSON(code, response)
}

func ErrorJSONResponse(ctx *gin.Context, code int, message string, data any) {
	response := gin.H{
		"status":  "error",
		"message": message,
		"data":    nil,
		"error":   data,
	}
	ctx.JSON(code, response)
}
