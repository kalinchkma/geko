package main

import (
	"ganja/internal/env"
	"ganja/internal/server"
	"time"

	"go.uber.org/zap"
)

func main() {

	// Server config
	cfg := server.Config{
		Addr: env.GetString("Addr", ":8080"),
		DbCfg: server.DatabaseConfig{
			Addr:         env.GetString("DB_ADDR", "postgres://admin:admin@localhost:5432/ganja"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		RedisCfg: server.RedisConfig{
			Addr:    env.GetString("REDIS_ADDR", "127.0.0.1:6379"),
			PW:      env.GetString("REDIS_PW", ""),
			DB:      env.GetInt("REDIS_DB", 0),
			Enabled: env.GetBool("REDIS_ENABLED", false),
		},
		Env:       env.GetString("ENV", "development"),
		MailerCfg: server.MailerConfig{
			// @TODO implement mailer config
		},
		AuthCfg: server.AuthConfig{
			// @TODO implement auth config
		},
		RateLimiterCfg: server.RateLimiterConfig{
			RequestsPerTimeFrame: env.GetInt("RATELIMITER_REQUEST_COUNT", 20),
			TimeFrame:            time.Second * 5,
			Enabled:              env.GetBool("RATE_LIMITER_ENABLED", true),
		},
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Server
	srv := &server.HttpServer{
		Config: cfg,
		Logger: logger,
	}

	router := srv.Mount()

	logger.Fatal(srv.Run(router))

}
