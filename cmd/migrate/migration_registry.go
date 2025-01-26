package migrations

import "geko/models"

// Register model for migration
var modelLists = map[string]interface{}{
	"users":  &models.User{},
	"guilds": &models.Guild{},
}
