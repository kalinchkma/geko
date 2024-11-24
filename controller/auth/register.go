package authController

import (
	"ganja/interfaces"

	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password int    `json:"password" binding:"required"`
}

func Register(actx *interfaces.AppContext, ctx *gin.Context) {
	var registerBody RequestBody

	// var errorMessage map[string]string
	// validate requesting inputs
	if err := ctx.ShouldBindJSON(&registerBody.Name); err != nil {
		// @TODO handle error
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	go (*actx).Mailer.SendEmail("no-replay@demomailtrap.com", []string{registerBody.Email}, "Welcome to Battech", "Hello, good to see you here")

	ctx.SecureJSON(http.StatusOK, gin.H{"message": "user register routes"})
}
