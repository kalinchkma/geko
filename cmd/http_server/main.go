package main

import (
	"fmt"
	"geko/internal/cache"
	"geko/internal/db"
	"geko/internal/env"
	"geko/internal/server"
	"log"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
		return
	}

	// Server config
	cfg := server.Config{
		Addr: fmt.Sprintf(":%v", env.GetString("PORT", "8080")),
		DbCfg: db.DatabaseConfig{
			Host:         env.GetString("DB_HOST", "127.0.0.1"),
			Port:         env.GetString("DB_PORT", "5432"),
			DBUserName:   env.GetString("DB_USERNAME", "admin"),
			DBDatabase:   env.GetString("DB_DATABASE", "geko"),
			DBPassword:   env.GetString("DB_PASSWORD", ""),
			DBSchema:     env.GetString("DB_SCHEMA", "public"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		RedisCfg: cache.RedisConfig{
			Host:    env.GetString("REDIS_HOST", "127.0.0.1"),
			Port:    env.GetString("REDIS_PORT", "6379"),
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

	logger.Fatal(srv.RunServer(router))

}
