package geko

import (
	"github.com/kalinchkma/geko/authenticator"
	"github.com/kalinchkma/geko/cache"
	"github.com/kalinchkma/geko/mailers"
	"github.com/kalinchkma/geko/ratelimiter"
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
