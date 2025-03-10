package server

import (
	"geko/internal/authenticator"
	"geko/internal/cache"
	"geko/internal/mailers"
	"geko/internal/ratelimiter"
)

type Config struct {
	Addr                       string
	Env                        string
	AppName                    string
	OTPValidationTime          int
	AccessTokenValidationTime  int
	RefreshTokenValidationTime int
	MailerConfig               mailers.MailerConfig
	AuthCfg                    authenticator.AuthConfig
	RedisCfg                   cache.RedisConfig
	RateLimiterCfg             ratelimiter.RateLimiterConfig
}
