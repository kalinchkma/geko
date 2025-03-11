package authcontroller

import (
	authmailer "geko/auth/mailers"
	"geko/internal/server"
	authstore "geko/internal/store/auth_store"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *AuthController) Register(ctx *gin.Context) {
	var registerBody RegisterPayload

	if err := ctx.ShouldBindJSON(&registerBody); err != nil {
		// server.ErrorJSONResponse(ctx, http.StatusBadRequest, "Bad Request", validators.NormalizeJsonValidationError(err, RegisterValidationMessages))
		server.ErrorJSONResponseWithFormatter(ctx, http.StatusBadRequest, "Bad request", err, RegisterValidationMessages)
		return
	}

	userStore := s.serverContext.Store.UserStore
	// Find the user by email, if already exist
	// Return error if user already  exist
	if _, err := userStore.FindByEmail(registerBody.Email); err == nil {
		server.ErrorJSONResponse(ctx, http.StatusConflict, "User already exists", nil)
		return
	}

	// Hash password
	hashedPassword, err := userStore.HashPassword(registerBody.Password)
	// Check error
	if err != nil {
		// Return error
		server.ErrorJSONResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Create new user
	user := authstore.User{
		Name:     registerBody.Name,
		Email:    registerBody.Email,
		Password: hashedPassword,
	}

	// Store user to database
	err = userStore.Create(user)
	// Check error
	if err != nil {
		// Return error
		server.ErrorJSONResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Get Update user
	user, err = userStore.FindByEmail(user.Email)
	if err != nil {
		// If created user not found return error
		server.ErrorJSONResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// generate and store otp
	otpStore := s.serverContext.Store.OTPStore

	// New otp code
	newOTPCode := otpStore.GenerateOTP(6)

	// New otp object
	otp := authstore.OTP{
		Code:      newOTPCode,
		UserId:    user.ID,
		ExpiresAt: time.Now().Add(time.Duration(s.serverContext.Config.OTPValidationTime) * time.Minute),
	}

	// Store to database
	err = otpStore.Create(otp)

	// Only send the email if otp has been created on database
	if err == nil {
		templData := authmailer.OtpEmailTemplateData{
			Email:      user.Email,
			Otp:        otp.Code,
			AppName:    s.serverContext.Config.AppName,
			Expiration: s.serverContext.Config.OTPValidationTime,
		}
		// Send otp to user email
		go s.mailer.SendOTPEmail(templData)
	}

	// Send
	server.SuccessJSONResponse(ctx, http.StatusCreated, "User Register successfully", s.serverContext.Store.UserStore.Normalize(user))

}
