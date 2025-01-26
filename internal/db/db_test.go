package db

import (
	"context"
	"geko/internal/env"
	"log"
	"os"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func mustStartPostgresContainer() (func(context.Context) error, error) {
	var (
		dbName = "database"
		dbPwd  = "password"
		dbUser = "user"
	)

	dbContainer, err := postgres.Run(
		context.Background(),
		"postgres:latest",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPwd),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}

	database = dbName
	password = dbPwd
	username = dbUser

	dbHost, err := dbContainer.Host(context.Background())
	if err != nil {
		return dbContainer.Terminate, err
	}

	dbPort, err := dbContainer.MappedPort(context.Background(), "5432/tcp")
	if err != nil {
		return dbContainer.Terminate, err
	}

	host = dbHost
	port = dbPort.Port()

	return dbContainer.Terminate, err
}

func TestMain(m *testing.M) {
	teardown, err := mustStartPostgresContainer()
	if err != nil {
		log.Fatalf("could not start postgres container: %v", err)
	}

	m.Run()

	if teardown != nil && teardown(context.Background()) != nil {
		log.Fatalf("could not teardown postgres container: %v", err)
	}
}

func TestNew(t *testing.T) {
	srv := New(DatabaseConfig{
		Host:         env.GetString("DB_HOST", "127.0.0.1"),
		Port:         env.GetString("DB_PORT", "5432"),
		DBUserName:   env.GetString("DB_USERNAME", "admin"),
		DBDatabase:   env.GetString("DB_DATABASE", "geko"),
		DBPassword:   env.GetString("DB_PASSWORD", ""),
		DBSchema:     env.GetString("DB_SCHEMA", "public"),
		MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
		MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
		MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
	})
	if srv == nil {
		t.Fatal("New() returned nil")
	}
}

func TestHealth(t *testing.T) {
	srv := New(DatabaseConfig{
		Host:         env.GetString("DB_HOST", "127.0.0.1"),
		Port:         env.GetString("DB_PORT", "5432"),
		DBUserName:   env.GetString("DB_USERNAME", "admin"),
		DBDatabase:   env.GetString("DB_DATABASE", "geko"),
		DBPassword:   env.GetString("DB_PASSWORD", ""),
		DBSchema:     env.GetString("DB_SCHEMA", "public"),
		MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
		MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
		MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
	})

	stats := srv.Health()

	if stats["status"] != "up" {
		t.Fatalf("expected status to be up, got %s", stats["status"])
	}

	if _, ok := stats["error"]; ok {
		t.Fatalf("expected error not to be present")
	}

	if stats["message"] != "It's healthy" {
		t.Fatalf("expected message to be 'It's healthy', got %s", stats["message"])
	}
}

func TestClose(t *testing.T) {
	srv := New(DatabaseConfig{
		Host:         env.GetString("DB_HOST", "127.0.0.1"),
		Port:         env.GetString("DB_PORT", "5432"),
		DBUserName:   env.GetString("DB_USERNAME", "admin"),
		DBDatabase:   env.GetString("DB_DATABASE", "geko"),
		DBPassword:   env.GetString("DB_PASSWORD", ""),
		DBSchema:     env.GetString("DB_SCHEMA", "public"),
		MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
		MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
		MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
	})

	if srv.Close() != nil {
		t.Fatalf("expected Close() to return nil")
	}
}
