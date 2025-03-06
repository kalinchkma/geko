package authcontroller

import (
	"geko/internal/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResendOTPPayload struct {
	UserId uint `json:"user_id"`
}

func (a *AuthController) ResendOTP(ctx *gin.Context) {
	var payload ResendOTPPayload

	// Validate payload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		server.ErrorJSONResponse(ctx, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

}
