package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResendOTPPayload struct {
	UserId uint `json:"user_id"`
}

func (a *AuthController) ResendOTP(ctx *gin.Context) {
	var payload ResendOTPPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

}
