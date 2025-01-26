package server

import (
	"geko/internal/cache"
	"geko/internal/db"
	"time"
)

type Config struct {
	Addr           string
	Env            string
	DbCfg          db.DatabaseConfig
	MailerCfg      MailerConfig
	AuthCfg        AuthConfig
	RedisCfg       cache.RedisConfig
	RateLimiterCfg RateLimiterConfig
}

type MailerConfig struct {
}

type AuthConfig struct {
}

type RateLimiterConfig struct {
	RequestsPerTimeFrame int
	TimeFrame            time.Duration
	Enabled              bool
}
