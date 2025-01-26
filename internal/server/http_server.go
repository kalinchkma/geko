package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"geko/internal/auth"
	"geko/internal/cache"
	"geko/internal/mailers"
	"geko/internal/ratelimiter"
	"geko/internal/store"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Http server
type HttpServer struct {
	Config         Config
	Store          store.Storage
	Mailer         mailers.Client
	CacheStore     cache.Storage
	Logger         *zap.SugaredLogger
	Authentication auth.Authenticator
	RateLimiter    ratelimiter.Limiter
}

// Mount the server router
func (server *HttpServer) Mount() http.Handler {
	// Setting gin routing mode
	if server.Config.Env == "development" {
		gin.SetMode(gin.DebugMode)
	} else if server.Config.Env == "testing" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Base route hander
	hander := gin.Default()

	// Implement the middleware

	return hander
}

// Run the HttpServer
func (server *HttpServer) RunServer(handler http.Handler) error {

	srv := &http.Server{
		Addr:         server.Config.Addr,
		Handler:      handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		server.Logger.Infow("signal caught", "signal", s.String())

		shutdown <- srv.Shutdown(ctx)
	}()

	server.Logger.Infow("Server has started", "addr", server.Config.Addr, "env", server.Config.Env)

	err := srv.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdown
	if err != nil {
		return err
	}

	server.Logger.Infow("Server has stopped", "addr", server.Config.Addr, "env", server.Config.Env)

	return nil
}
