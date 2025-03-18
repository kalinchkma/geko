package authstore

import (
	"errors"
	"fmt"
	"time"

	"github.com/kalinchkma/geko/internal/db"
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
	OTP           []OTP
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

// Update user account status
func (u *UserStore) UpdateAccountStatus(email string, status bool) (User, error) {
	var user User
	// find user by id
	res := u.db.ORM.Where("email = ?", email).Find(&user)
	if res.Error != nil {
		fmt.Println("User query error", res.Error, email)
		return User{}, res.Error
	}

	// Check if user found
	if res.RowsAffected == 0 {
		fmt.Println("User not affected")
		return User{}, errors.New("user not found")
	}

	// Update the user status
	user.AcountStatus = status

	// set email verifyed
	user.EmailVerified = true

	// Save updated user
	u.db.ORM.Save(&user)

	return user, nil
}

// Delete user
func (u *UserStore) DeleteByEmail(email string) error {
	res := u.db.ORM.Where("email = ?", email).Delete(&email)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("no user has been deleted")
	}
	return nil
}

// Normalize user
func (u *UserStore) Normalize(user User) any {
	return struct {
		ID            uint
		CreatedAt     time.Time
		UpdatedAt     time.Time
		DeletedAt     time.Time
		Name          string
		Email         string
		EmailVerified bool
		AcountStatus  bool
	}{
		ID:            user.ID,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
		DeletedAt:     user.DeletedAt.Time,
		Name:          user.Name,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		AcountStatus:  user.AcountStatus,
	}
}
