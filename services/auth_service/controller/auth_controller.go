package authcontroller

import (
	"geko/internal/server"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	serverContext *server.HttpServerContext
}

func NewAuthController(serverContext *server.HttpServerContext) *AuthController {
	return &AuthController{
		serverContext: serverContext,
	}
}

func (a *AuthController) Login(ctx *gin.Context) {

}

func (a *AuthController) VerifyOtp(ctx *gin.Context) {

}

func (a *AuthController) ResendOtp(ctx *gin.Context) {

}

func (a *AuthController) ForgotPassword(ctx *gin.Context) {

}
func (a *AuthController) ResetPassword(ctx *gin.Context) {

}
func (a *AuthController) RefreshToken(ctx *gin.Context) {

}
