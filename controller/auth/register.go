package authController

import (
	"ganja/interfaces"
	"ganja/mailers/auth_mailer"
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
	// Create user instant
	var user models.User

	// Find the user by email, if it's already exist
	// Return error if user already exist
	if res := actx.DB.GetDB().Where("email = ?", requestBody.Email).Find(&user); res.RowsAffected != 0 {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "User already exist",
		})
		return
	}

	// Create new user
	user = models.User{
		Name:     requestBody.Name,
		Email:    requestBody.Email,
		Password: requestBody.Password,
	}

	// Save the new user
	res := actx.DB.DB.Create(&user)
	// Handle the error of user
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	// Send the active account otp
	go auth_mailer.Welcome(actx.Mailer, user)
	// Return success
	ctx.SecureJSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}
