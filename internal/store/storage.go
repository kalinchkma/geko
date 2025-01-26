package store

// Application storage
type Storage struct {
	Users users
}

// User Storage interface
type users interface {
	GetByID(string) (any, error)
	GetByEmail(string) (any, error)
	Create(any) error
}
