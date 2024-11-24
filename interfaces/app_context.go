package interfaces

import "ganja/initializers/database"

type AppContext struct {
	DB     *database.Database
	Mailer Mailer
}
