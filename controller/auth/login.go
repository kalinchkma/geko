package authController

import (
	"ganja/interfaces"
	"ganja/library"
	"ganja/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(actx *interfaces.AppContext, ctx *gin.Context) {
	var loginBody LoginBody

	// Validate user inputs
	if err := ctx.ShouldBindJSON(&loginBody); err != nil {
		// @TODO validate user inputs
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Find user on database
	var user models.User
	res := actx.DB.GetDB().Where("email = ?", loginBody.Email).First(&user)
	if res.RowsAffected == 0 {
		// @TODO handle error
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user",
		})
		return
	}

	// Check user active or not
	if !user.AcountStatus {
		// Return error is user not active
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "Inactive user",
		})
		return
	}

	// @TODO handle is email verifyed of not
	// -----

	// Check the password
	if err := library.ComparePassword(user.Password, loginBody.Password); !err {
		// Password not match return error
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user",
		})
		return
	}

	// Login the user with access token

	// token data

	if token, err := library.GenerateJWT(user.Name, time.Duration(time.Minute*10)); err != nil {
		// Respose error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	} else {
		// Respose with token
		ctx.JSON(http.StatusOK, gin.H{
			"success": "Login success",
			"token":   token,
		})
	}

}
