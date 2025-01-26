package app

import "time"

type Config struct {
	Addr           string
	Env            string
	DatabaseCfg    DatabaseConfig
	MailerCfg      MailerConfig
	AuthCfg        AuthConfig
	RedisCfg       RedisConfig
	RateLimiterCfg RateLimiterConfig
}

type DatabaseConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

type MailerConfig struct {
}

type AuthConfig struct {
}

type RedisConfig struct {
}

type RateLimiterConfig struct {
	RequestsPerTimeFrame int
	TimeFrame            time.Duration
	Enabled              bool
}
