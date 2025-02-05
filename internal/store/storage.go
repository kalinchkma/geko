package store

import (
	authstore "geko/internal/store/auth_store"
)

// Application storage
type Storage struct {
	UserStore         authstore.UserStore
	RoleStore         authstore.RoleStore
	RefreshTokenStore authstore.RefreshTokenStore
}

func (s *Storage) Migrate() {
	// Migrate user store
	s.UserStore.Migrate()

	// Migrate role store
	s.RoleStore.Migrate()

	// Migrate refresh token store
	s.RefreshTokenStore.Migrate()
}
