package store

import (
	"geko/internal/db"

	"gorm.io/gorm"
)

const (
	GuildMaster = iota
	Adventurer
)

type Role int

func (r Role) String() string {
	return [...]string{"Guild Master", "Adventurer"}[r]
}

// User model
type User struct {
	gorm.Model
	Name          string
	Email         string `gorm:"unique"`
	Password      string
	Role          Role
	EmailVerified bool
	AcountStatus  bool
	AccessToken   string
	RefreshToken  string
}

type UserStore struct {
	db db.Database
}

func (u *UserStore) Create(user User) error {
	// Store user to database
	res := u.db.DB.Create(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *UserStore) FindByEmail(email string) (User, error) {
	var user User

	res := u.db.DB.Where("email = ?", email).Find(&user)
	if res.Error != nil {
		return User{}, res.Error
	}

	return user, nil

}
