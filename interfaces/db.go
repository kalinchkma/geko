package interfaces

import "gorm.io/gorm"

type Database interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	// Get db instance
	GetDB() *gorm.DB
}
