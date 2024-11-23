package migrations

import (
	"ganja/models"

	"gorm.io/gorm"
)

func Up_user(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
}

func Down_user(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.User{})
}
