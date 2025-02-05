package authstore

import (
	"geko/internal/db"

	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Name          string `json:"name"`
	Email         string `gorm:"unique" json:"email"`
	Password      string
	EmailVerified bool `json:"email_verified"`
	AcountStatus  bool `json:"account_status"`
}

type UserStore struct {
	db *db.Database
}

// Constructor
func NewUserStore(db *db.Database) *UserStore {
	return &UserStore{db}
}

// Create user
func (u *UserStore) Create(user User) error {
	// Store user to database
	res := u.db.ORM.Create(&user)

	// Check if any error
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// Find User by email
func (u *UserStore) FindByEmail(email string) (User, error) {
	var user User

	res := u.db.ORM.Where("email = ?", email).Find(&user)
	if res.Error != nil {
		return User{}, res.Error
	}

	return user, nil

}
