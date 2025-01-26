package app

import (
	"ganja/internal/auth"
	"ganja/internal/cache"
	"ganja/internal/mailers"
	"ganja/internal/ratelimiter"
	"ganja/internal/store"

	"go.uber.org/zap"
)

type Application struct {
	Config         Config
	Store          store.Storage
	Mailer         mailers.Client
	CacheStore     cache.Storage
	Logger         *zap.SugaredLogger
	Authentication auth.Authenticator
	RateLimiter    ratelimiter.Limiter
}
