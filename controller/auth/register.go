package authController

import (
	"fmt"
	"ganja/interfaces"
	"ganja/library"
	"ganja/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(actx *interfaces.AppContext, ctx *gin.Context) {
	var requestBody RequestBody

	// var errorMessage map[string]string
	// validate requesting inputs
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		// @TODO handle user input error
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// create new user
	user := models.User{
		Name:     requestBody.Name,
		Email:    requestBody.Email,
		Password: requestBody.Password,
	}

	// save the new user
	res := actx.DB.DB.Create(&user)

	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	go func() {
		otp := library.GenerateOTP(4)
		fmt.Println("otp", otp)
		(*actx).Mailer.SendEmail("no-replay@demomailtrap.com", []string{user.Email}, "Welcome to Battech", fmt.Sprintf("Your OTP: <p> %v </p>", otp))
	}()

	ctx.SecureJSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}
