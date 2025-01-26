package interfaces

import (
	"geko/initializers/database"
	"geko/initializers/mailers"
)

type AppContext struct {
	DB     *database.Database
	Mailer *mailers.Mailer
}
