package authstore

import (
	"geko/internal/db"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description"`
}

type RoleStore struct {
	db *db.Database
}

// Constructor
func NewRoleStore(db *db.Database) *RoleStore {
	return &RoleStore{db}
}

// Migrate role model
func (r *RoleStore) Migrate() error {
	// Migrate role model
	if err := r.db.ORM.Migrator().AutoMigrate(&Role{}); err != nil {
		// Return error if migration failed any
		return err
	}
	return nil
}

// Create role
func (r *RoleStore) Create(role Role) error {
	if res := r.db.ORM.Create(&role); res.Error != nil {
		return res.Error
	}
	return nil
}

//
