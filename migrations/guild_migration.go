package migrations

import (
	"ganja/models"

	"gorm.io/gorm"
)

func Up_guild(db *gorm.DB) error {
	return db.AutoMigrate(&models.Guild{})
}

func Down_guild(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Guild{})
}
