package authcontroller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequestPayload struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

func (a *AuthController) Login(ctx *gin.Context) {
	var payload LoginRequestPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	user, err := a.serverContext.Store.UserStore.FindByEmail(payload.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	if ok := a.serverContext.Store.UserStore.ComparePassword(user.Password, payload.Password); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"errors": "Unauthorized",
		})
		return
	}

	// Generate access token
	claims := jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Duration(a.serverContext.Config.AccessTokenValidationTime)).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": a.serverContext.Config.AuthCfg.Token.Iss,
		"aud": a.serverContext.Config.AuthCfg.Token.Iss,
	}
	accessToken, err := a.serverContext.Authenticator.JWTAuth.GenerateToken(claims)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": "Internal server error",
		})
		return
	}

	// Generate refresh token

	claims = jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Duration(a.serverContext.Config.RefreshTokenValidationTime)).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": a.serverContext.Config.AuthCfg.Token.Iss,
		"aud": a.serverContext.Config.AuthCfg.Token.Iss,
	}

	refreshToken, err := a.serverContext.Authenticator.JWTAuth.GenerateToken(claims)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": "Internal server error",
		})
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message":       "Login success",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
