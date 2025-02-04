package store

import (
	authstore "geko/internal/store/auth_store"
	"log"
)

// Application storage
type Storage struct {
	UserStore         authstore.UserStore
	RoleStore         authstore.RoleStore
	RefreshTokenStore authstore.RefreshTokenStore
}

func (s *Storage) Migrate() {
	// Migrate user store
	log.Fatal(s.UserStore.Migrate())

	// Migrate role store
	log.Fatal(s.RoleStore.Migrate())

	// Migrate refresh token store
	log.Fatal(s.RefreshTokenStore.Migrate())
}
