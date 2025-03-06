package authcontroller

import (
	"geko/internal/server"
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
		server.ErrorJSONResponse(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	// Find user by email
	user, err := a.serverContext.Store.UserStore.FindByEmail(payload.Email)
	if err != nil {
		server.ErrorJSONResponse(ctx, http.StatusBadRequest, "Invalid user account", nil)
		return
	}

	// Check Password
	if ok := a.serverContext.Store.UserStore.ComparePassword(user.Password, payload.Password); !ok {
		server.ErrorJSONResponse(ctx, http.StatusForbidden, "Invalid user credentials", nil)
		return
	}

	// Check user account status
	if !user.AcountStatus || !user.EmailVerified {
		server.ErrorJSONResponse(ctx, http.StatusForbidden, "Inactive user account", nil)
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
		server.ErrorJSONResponse(ctx, http.StatusInternalServerError, "Internal server error", nil)
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
		server.ErrorJSONResponse(ctx, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	server.SuccessJSONResponse(ctx, http.StatusOK, "Login success", gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
