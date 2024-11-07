package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"ganja/app/database"
	"ganja/internal/interfaces"

	_ "github.com/joho/godotenv/autoload"

	"ganja/app/mailers"
)

type Server struct {
	port int
	cfg  *interfaces.Config
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		cfg: &interfaces.Config{
			Mailer: mailers.New(),
			DB:     database.New(),
		},
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
