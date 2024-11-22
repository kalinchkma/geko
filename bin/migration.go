package main

import "gorm.io/gorm"

type Migration struct {
	ID      string
	Up      func(*gorm.DB) error
	Down    func(*gorm.DB) error
	Applied bool
}
