package geko

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kalinchkma/geko/authenticator"
	"github.com/kalinchkma/geko/cache"
	"github.com/kalinchkma/geko/mailers"
	"github.com/kalinchkma/geko/ratelimiter"
	"github.com/kalinchkma/geko/store"
	"go.uber.org/zap"
)

type HttpServerContext struct {
	Config        Config
	Store         store.Storage
	Mailer        mailers.Mailer
	CacheStore    cache.Storage
	Logger        *zap.SugaredLogger
	Authenticator authenticator.Authenticator
	RateLimiter   ratelimiter.Limiter
}

// Http server
type HttpServer struct {
	context *HttpServerContext
	handler *gin.Engine
	server  *http.Server
}

// Constructor
func NewHttpServer(context *HttpServerContext) *HttpServer {

	// Configure gin routing mode
	// Based on environtment
	if context.Config.Env == "development" {
		gin.SetMode(gin.DebugMode)
	} else if context.Config.Env == "testing" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Root router hander of server
	handler := gin.Default()

	server := &http.Server{
		Addr:         context.Config.Addr,
		Handler:      handler,
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	return &HttpServer{
		context: context,
		handler: handler,
		server:  server,
	}
}

// Mount the service
func (server *HttpServer) MountService(mountPath string, service Service) {
	// base route group
	group := server.handler.Group(mountPath)

	// Mount the service to current server
	service.Mount(server.context, group)

	// Register the service
	service.RouteHandler()
}

// Run the HttpServer
func (server *HttpServer) Start() error {

	srv := server.server
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
		Ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Shutdown log
		log.Println("signal", s.String())

		// Catch shutdown log if any
		shutdown <- srv.Shutdown(Ctx)
	}()

	// Server started log
	log.Println("Server has started PORT ", server.context.Config.Addr, " env:", server.context.Config.Env)

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
	log.Println("Server has stopped", "addr", server.context.Config.Addr, "env", server.context.Config.Env)

	return nil
}
