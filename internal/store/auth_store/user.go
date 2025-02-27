package authstore

import (
	"fmt"
	"geko/internal/db"

	"golang.org/x/crypto/bcrypt"
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
	OTP           OTP
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

	if res.RowsAffected == 0 {
		return User{}, fmt.Errorf("user not found")
	}

	return user, nil

}

// Hash user password
func (u *UserStore) HashPassword(passwordString string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// Compate hashed password
func (u *UserStore) ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
