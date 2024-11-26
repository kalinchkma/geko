package authController

import (
	"ganja/interfaces"

	"github.com/gin-gonic/gin"
)

type LoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(actx *interfaces.AppContext, ctx *gin.Context) {
	// @TODO complete login

}
