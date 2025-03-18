package main

import (
	"fmt"

	"log"

	"github.com/joho/godotenv"
	"github.com/kalinchkma/geko/internal/db"
	"github.com/kalinchkma/geko/internal/env"
	"github.com/kalinchkma/geko/internal/store"
	"gorm.io/gorm"
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
		DBUserName:   env.GetString("DB_USERNAME", "geko"),
		DBName:       env.GetString("DB_DATABASE", "geko"),
		DBPassword:   env.GetString("DB_PASSWORD", ""),
		DBSchema:     env.GetString("DB_SCHEMA", "public"),
		MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
		MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
		MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
	}

	storage := store.NewStorage(dbConfig)
	fmt.Println("Start migration process........")
	migrate(storage.Models(), storage.DB)
	fmt.Println("Migration done..........")
}

func migrate(models map[string]interface{}, db *db.Database) {
	for _, model := range models {
		modelName, err := getModelName(db.ORM, model)
		if err != nil {
			log.Fatalf("error migrating %T: %v", model, err)
		}
		fmt.Printf("Migrating %T: %v..........\n", model, modelName)

		db.ORM.AutoMigrate(model)

		fmt.Printf("\nMigration Complete %v\n", modelName)
	}

}

func getModelName(db *gorm.DB, model interface{}) (string, error) {

	stmt := &gorm.Statement{DB: db}

	if err := stmt.Parse(model); err != nil {
		return "", fmt.Errorf("error parsing model: %w", err)
	}

	return stmt.Schema.Table, nil

}
