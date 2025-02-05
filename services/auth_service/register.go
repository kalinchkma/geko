package authservice

import "github.com/gin-gonic/gin"

type RegisterBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *AuthService) registerHandler(ctx *gin.Context) {
	reqisterBody := new(RegisterBody)
	println(reqisterBody)
}
