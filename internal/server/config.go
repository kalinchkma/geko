package server

import (
	"geko/internal/auth"
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
	AuthCfg                    auth.AuthConfig
	RedisCfg                   cache.RedisConfig
	RateLimiterCfg             ratelimiter.RateLimiterConfig
}
