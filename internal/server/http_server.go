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
	Config        Config
	Store         store.Storage
	Mailer        mailers.Client
	CacheStore    cache.Storage
	Logger        *zap.SugaredLogger
	Authenticator auth.Authenticator
	RateLimiter   ratelimiter.Limiter
}

// Mount the server router
func (server *HttpServer) Mount() http.Handler {
	// Configure gin routing mode
	// Based on environtment
	if server.Config.Env == "development" {
		gin.SetMode(gin.DebugMode)
	} else if server.Config.Env == "testing" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Root router hander of server
	hander := gin.Default()

	return hander
}

// Run the HttpServer
func (server *HttpServer) RunServer(handler http.Handler) error {

	// Initialize http server base on configuration
	// You can customize as you want
	srv := &http.Server{
		Addr:         server.Config.Addr,
		Handler:      handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	// Error channel to catch shutdown server
	shutdown := make(chan error)

	// Shutdown gorutines
	go func() {

		// Server quite channel
		quit := make(chan os.Signal, 1)

		// Signal notification
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// Catch signal
		s := <-quit

		// gracefully shutdown context
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Shutdown log
		server.Logger.Infow("signal", s.String())

		// Catch shutdown log if any
		shutdown <- srv.Shutdown(ctx)
	}()

	// Server started log
	server.Logger.Infow("Server has started", "addr", server.Config.Addr, "env", server.Config.Env)

	// Serve server
	err := srv.ListenAndServe()

	// Check and return if any error
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	// Catch  shutdown error
	err = <-shutdown

	// Check and return error
	if err != nil {
		return err
	}

	// Server stopped error
	server.Logger.Infow("Server has stopped", "addr", server.Config.Addr, "env", server.Config.Env)

	return nil
}
