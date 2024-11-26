package main

import (
	"ganja/initializers/database"
	"ganja/migrations"
	"ganja/server"
	"os"
)

const (
	SERVER       = "server"
	MIGRATEUP    = "migrate:up"
	MIGRATECLEAN = "migrate:clean"
	MIGRATEDOWN  = "migrate:down"
)

func main() {
	args := os.Args

	// check commandline arguments
	// if argument not pass panic
	if len(args) < 2 {
		panic("I don't know what to run")
	}

	a := args[1]
	if a == SERVER {
		server.RunServer()
	} else if a == MIGRATEUP {
		db := database.New()
		migrations.Migrate(db.GetDB())
	} else if a == MIGRATEDOWN {
		db := database.New()
		migrations.Rollback(db.GetDB())
	} else if a == MIGRATECLEAN {
		db := database.New()
		migrations.MigrateWithCleanUp(db.DB)
	}
}
