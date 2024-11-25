package interfaces

import (
	"ganja/initializers/database"
	"ganja/initializers/mailers"
)

type AppContext struct {
	DB     *database.Database
	Mailer *mailers.Mailer
}
