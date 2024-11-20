package authController

import (
	"fmt"
	"ganja/interfaces"
	"strings"

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
		// fmt.Printf("%#v \n%#v", err.Error(), strings.Split(err.Error(), "\n"))
		fmt.Println(err.Error())
		for _, filedError := range strings.Split(err.Error(), "\n") {
			fmt.Println(filedError)
		}
	}

	// go (*actx).Mailer.SendEmail("no-replay@demomailtrap.com", []string{registerBody.Email}, "Welcome to Battech", "Hello, good to see you here")

	ctx.SecureJSON(http.StatusOK, gin.H{"message": "user register routes"})
}
