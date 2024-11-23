package models

import "gorm.io/gorm"

type Guild struct {
	gorm.Model
	Name string
}

func (g *Guild) String() string {
	return "guilds"
}
