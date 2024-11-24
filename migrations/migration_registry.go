package migrations

import "ganja/models"

// Register model for migration
var modelLists = map[string]interface{}{
	"users":  &models.User{},
	"guilds": &models.Guild{},
}
