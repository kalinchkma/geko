package auth

import (
	"fmt"
	config "ganja/app/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password int    `json:"password" binding:"required"`
}

func Register(cfg *config.Config, ctx *gin.Context) {
	var registerBody User

	if err := ctx.BindJSON(&registerBody); err != nil {
		fmt.Printf("%#v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(),
		})
		return
	}

	go (*cfg).Mailer.SendEmail("no-replay@gmail.com", []string{registerBody.Email}, "Welcome to Battech", "Hello, good to see you here")

	ctx.SecureJSON(http.StatusOK, gin.H{"message": "user register routes"})
}
