package server

import (
	"github.com/kalinchkma/geko/internal/authenticator"
	"github.com/kalinchkma/geko/internal/cache"
	"github.com/kalinchkma/geko/internal/mailers"
	"github.com/kalinchkma/geko/internal/ratelimiter"
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
