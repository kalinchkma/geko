package store

import "gorm.io/gorm"

const (
	GuildMaster = iota
	Adventurer
)

type Role int

func (r Role) String() string {
	return [...]string{"Guild Master", "Adventurer"}[r]
}

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

func (u *User) Create() {

}
