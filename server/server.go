package server

import (
	"fmt"

	"ganja/initializers/database"
	"ganja/interfaces"
	"ganja/mailers"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	actx *interfaces.AppContext
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		actx: &interfaces.AppContext{
			Mailer: mailers.New(),
			DB:     database.New(),
		},
	}

	// load dependency services
	NewServer.bootstrap()

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.routesRegistry(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
