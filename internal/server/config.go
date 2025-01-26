package server

import "time"

type Config struct {
	Addr           string
	Env            string
	DbCfg          DatabaseConfig
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
	Addr    string
	PW      string // Password
	DB      int
	Enabled bool
}

type RateLimiterConfig struct {
	RequestsPerTimeFrame int
	TimeFrame            time.Duration
	Enabled              bool
}
