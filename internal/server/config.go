package server

import (
	"geko/internal/cache"
	"geko/internal/ratelimiter"
)

type Config struct {
	Addr           string
	Env            string
	AuthCfg        AuthConfig
	RedisCfg       cache.RedisConfig
	RateLimiterCfg ratelimiter.RateLimiterConfig
}

type AuthConfig struct {
}
