package authcontroller

import (
	authmailer "geko/auth/mailers"
	"geko/internal/server"
	authstore "geko/internal/store/auth_store"
	"geko/internal/validators"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterPayload struct {
	Name     string `json:"name" binding:"required,min=5,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

var validationMessages = map[string]string{
	"Name.required":     "Name is required",
	"Name.min":          "Name is too short at leat 5 alphabate",
	"Name.max":          "Name is too long it must less then 20 alphabate",
	"Email.required":    "Email is required",
	"Email.email":       "Invalid Email",
	"Password.required": "Password is required",
	"Password.min":      "Password must be at least 8 character long",
}

func (s *AuthController) Register(ctx *gin.Context) {
	var registerBody RegisterPayload

	if err := ctx.ShouldBindJSON(&registerBody); err != nil {
		server.ErrorJSONResponse(ctx, http.StatusBadRequest, "Bad Request", validators.NormalizeJsonValidationErrorWithType(err, validationMessages))
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
