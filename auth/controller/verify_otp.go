package authcontroller

import (
	authmailer "geko/auth/mailers"
	"geko/internal/server"
	"geko/internal/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *AuthController) VerifyOtp(ctx *gin.Context) {
	var verifyOTPBody VerifyOTPPayload

	// Verify request body
	if err := ctx.ShouldBindJSON(&verifyOTPBody); err != nil {
		server.ErrorJSONResponse(ctx, http.StatusBadRequest, "Bad Request", validators.NormalizeJsonValidationError(err, VerifyOTPValidationMessages))
		return
	}

	// Verify otp
	if err := a.serverContext.Store.OTPStore.VerifyOTP(verifyOTPBody.Email, verifyOTPBody.Code); err != nil {
		server.ErrorJSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	} else {
		// OTP is valid
		// Active the user
		user, err := a.serverContext.Store.UserStore.UpdateAccountStatus(verifyOTPBody.Email, true)
		if err != nil {
			server.ErrorJSONResponse(ctx, http.StatusInternalServerError, "Internal server error", nil)
			return
		}

		// Send account activated welcome email
		emailData := authmailer.WelcomeEmailTemplateData{
			Name:    user.Name,
			AppName: a.serverContext.Config.AppName,
		}
		// Send welcome email
		go a.mailer.SendWelcomeEmail(emailData)

		server.SuccessJSONResponse(ctx, http.StatusOK, "Your account has been activated successfully", nil)
	}
}
