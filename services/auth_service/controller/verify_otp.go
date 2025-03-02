package authcontroller

import (
	authmailer "geko/services/auth_service/mailer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VerifyOTPRequestBody struct {
	Code   string `json:"code"`
	UserID uint   `json:"user_id"`
}

func (a *AuthController) VerifyOtp(ctx *gin.Context) {
	var verifyOTPBody VerifyOTPRequestBody

	// Verify request body
	if err := ctx.ShouldBindJSON(&verifyOTPBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	// Verify otp
	if err := a.serverContext.Store.OTPStore.VerifyOTP(verifyOTPBody.UserID, verifyOTPBody.Code); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	} else {
		// OTP is valid
		// Active the user
		user, err := a.serverContext.Store.UserStore.UpdateAccountStatus(verifyOTPBody.UserID, true)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errors": "Internal server errors",
			})
			return
		}

		// Send account activated welcome email
		emailData := authmailer.WelcomeEmailTemplateData{
			Name:    user.Name,
			AppName: a.serverContext.Config.AppName,
		}
		a.mailer.SendWelcomeEmail(emailData)
		ctx.JSON(http.StatusOK, gin.H{
			"success": "Your account activated succussfully",
		})
	}

}
