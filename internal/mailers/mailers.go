package mailers

type Client interface {
	Send(email string, data any)
}
