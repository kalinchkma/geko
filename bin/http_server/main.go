package main

import (
	"fmt"

	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kalinchkma/geko/auth"
	"github.com/kalinchkma/geko/internal/authenticator"
	"github.com/kalinchkma/geko/internal/cache"
	"github.com/kalinchkma/geko/internal/db"
	"github.com/kalinchkma/geko/internal/env"
	"github.com/kalinchkma/geko/internal/mailers"
	"github.com/kalinchkma/geko/internal/ratelimiter"
	"github.com/kalinchkma/geko/internal/server"
	"github.com/kalinchkma/geko/internal/store"
	orderservice "github.com/kalinchkma/geko/services/order_service"
	"go.uber.org/zap"
)

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
		return
	}

	// Mailers config
	mailerConfig := mailers.MailerConfig{
		Host:     env.GetString("SMTP_HOST", ""),
		Port:     env.GetInt("SMTP_PORT", 0),
		Username: env.GetString("SMTP_USER", ""),
		Password: env.GetString("SMTP_PASSWORD", ""),
		Domain:   env.GetString("EMAIL_DOMAIN", ""),
	}

	// Auth config
	authConfig := authenticator.AuthConfig{
		Token: authenticator.TokenConfig{
			Secret: env.GetString("TOKEN_SECRET", "super_secret"),
			Exp:    time.Duration(env.GetInt("TOKEN_VALIDITY", 24) * time.Now().Hour()),
			Iss:    env.GetString("TOKEN_ISS", "GEKO_"),
		},
	}

	// Server config
	cfg := server.Config{
		Addr: fmt.Sprintf(":%v", env.GetString("PORT", "8080")),
		RedisCfg: cache.RedisConfig{
			Host:    env.GetString("REDIS_HOST", "127.0.0.1"),
			Port:    env.GetString("REDIS_PORT", "6379"),
			PW:      env.GetString("REDIS_PW", ""),
			DB:      env.GetInt("REDIS_DB", 0),
			Enabled: env.GetBool("REDIS_ENABLED", false),
		},
		Env:                        env.GetString("ENV", "development"),
		AppName:                    env.GetString("APP_NAME", ""),
		OTPValidationTime:          env.GetInt("OTP_VALIDATION_TIME", 3),
		AccessTokenValidationTime:  env.GetInt("ACCESS_TOKEN_VALIDATION_TIME", 30),
		RefreshTokenValidationTime: env.GetInt("REFRESH_TOKEN_VALIDATION_TIME", 1440),
		MailerConfig:               mailerConfig,
		AuthCfg:                    authConfig,
		RateLimiterCfg: ratelimiter.RateLimiterConfig{
			RequestsPerTimeFrame: env.GetInt("RATELIMITER_REQUEST_COUNT", 20),
			TimeFrame:            time.Second * 5,
			Enabled:              env.GetBool("RATE_LIMITER_ENABLED", true),
		},
	}

	// Database configuration
	dbCfg := db.DatabaseConfig{
		Host:         env.GetString("DB_HOST", "127.0.0.1"),
		Port:         env.GetString("DB_PORT", "5432"),
		DBUserName:   env.GetString("DB_USERNAME", "admin"),
		DBName:       env.GetString("DB_NAME", "geko"),
		DBPassword:   env.GetString("DB_PASSWORD", ""),
		DBSchema:     env.GetString("DB_SCHEMA", "public"),
		MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
		MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
		MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
	}

	// mailer
	newMailers := mailers.NewMailer(mailerConfig)

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Server context
	ctx := &server.HttpServerContext{
		Config: cfg,
		Logger: logger,
		Store:  *store.NewStorage(dbCfg),
		Mailer: newMailers,
		Authenticator: authenticator.Authenticator{
			JWTAuth: *authenticator.NewJWTAuthenticator(authConfig.Token.Secret, authConfig.Token.Iss, authConfig.Token.Iss),
		},
	}

	// Server
	srv := server.NewHttpServer(ctx)

	srv.MountService("/auth", &auth.AuthService{})
	srv.MountService("/order", &orderservice.OrderService{})

	logger.Fatal(srv.Start())

}
