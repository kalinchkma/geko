package authcontroller

import (
	authmailer "geko/auth/mailers"
	"geko/internal/server"
	authstore "geko/internal/store/auth_store"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *AuthController) ResendOTP(ctx *gin.Context) {
	var payload ResendOTPPayload

	// Validate payload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		server.ErrorJSONResponseWithFormatter(ctx, http.StatusBadRequest, "Bad request", err, ResendOTPValidationMessages)
		return
	}

	// Finde user by email
	user, err := a.serverContext.Store.UserStore.FindByEmail(payload.Email)
	if err != nil {
		server.ErrorJSONResponse(ctx, http.StatusForbidden, "Invalid email", nil)
		return
	}

	if user.AcountStatus && user.EmailVerified {
		server.ErrorJSONResponse(ctx, http.StatusBadRequest, "Account already activated", nil)
		return
	}

	// generate new otp
	newOtpCode, err := a.serverContext.Store.OTPStore.RegenerateOTP(payload.Email, 6)
	if err != nil {
		server.ErrorJSONResponse(ctx, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	// New otp
	otp := authstore.OTP{
		Code:      newOtpCode,
		Email:     user.Email,
		UserId:    user.ID,
		ExpiresAt: time.Now().Add(time.Duration(a.serverContext.Config.OTPValidationTime) * time.Minute),
	}

	// Store to database
	err = a.serverContext.Store.OTPStore.Create(otp)

	// Send otp it otp is created
	if err == nil {
		templData := authmailer.OtpEmailTemplateData{
			Email:      user.Email,
			Otp:        otp.Code,
			AppName:    a.serverContext.Config.AppName,
			Expiration: a.serverContext.Config.OTPValidationTime,
		}
		// Send otp to use email
		go a.mailer.SendOTPEmail(templData)
	}

	// send success response
	server.SuccessJSONResponse(ctx, http.StatusCreated, "Otp send successfully, check your email", nil)
}
