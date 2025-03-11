package authcontroller

import (
	"geko/internal/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *AuthController) ResendOTP(ctx *gin.Context) {
	var payload ResendOTPPayload

	// Validate payload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		server.ErrorJSONResponseWithFormatter(ctx, http.StatusBadRequest, "Bad request", err, ResendOTPValidationMessages)
		return
	}

}
