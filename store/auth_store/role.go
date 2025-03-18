package authstore

import (
	"github.com/kalinchkma/geko/db"
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

// Create role
func (r *RoleStore) Create(role Role) error {
	if res := r.db.ORM.Create(&role); res.Error != nil {
		return res.Error
	}
	return nil
}
