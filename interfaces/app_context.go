package interfaces

type AppContext struct {
	DB     Database
	Mailer Mailer
}
