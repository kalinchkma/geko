package authcontroller

import (
	"geko/internal/auth"
	authstore "geko/internal/store/auth_store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterBody struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *AuthController) Register(ctx *gin.Context) {
	var registerBody RegisterBody

	if err := ctx.ShouldBindJSON(&registerBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
	}

	userStore := s.serverContext.Store.UserStore
	// Find the user by email, if already exist
	// Return error if user already  exist
	if _, err := userStore.FindByEmail(registerBody.Email); err == nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"errors": "User already exists",
		})
		return
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(registerBody.Password)
	// Check error
	if err != nil {
		// Return error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
	}
	ctx.SecureJSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

}
