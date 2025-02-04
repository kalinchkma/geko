package main

import (
	"fmt"
	"geko/internal/db"
	"geko/internal/env"
	"geko/internal/store"
	authstore "geko/internal/store/auth_store"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
		return
	}

	dbConfig := db.DatabaseConfig{
		Host:         env.GetString("DB_HOST", "127.0.0.1"),
		Port:         env.GetString("DB_PORT", "5432"),
		DBUserName:   env.GetString("DB_USERNAME", "admin"),
		DBDatabase:   env.GetString("DB_DATABASE", "geko"),
		DBPassword:   env.GetString("DB_PASSWORD", ""),
		DBSchema:     env.GetString("DB_SCHEMA", "public"),
		MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
		MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
		MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
	}

	database := db.New(dbConfig)
	storage := store.Storage{
		UserStore:         *authstore.NewUserStore(database),
		RoleStore:         *authstore.NewRoleStore(database),
		RefreshTokenStore: *authstore.NewRefreshTokenStore(database),
	}
	fmt.Println("Migration start....")
	storage.Migrate()
	fmt.Println("Migration done")
}
