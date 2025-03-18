package store

import (
	"geko/internal/db"
	authstore "geko/internal/store/auth_store"
)

// Application storage
type Storage struct {
	DB                *db.Database
	UserStore         authstore.UserStore
	OTPStore          authstore.OTPStore
	RoleStore         authstore.RoleStore
	RefreshTokenStore authstore.RefreshTokenStore
}

// Constructor
func NewStorage(dbCfg db.DatabaseConfig) *Storage {
	db := db.New(dbCfg)
	return &Storage{
		DB:                db,
		UserStore:         *authstore.NewUserStore(db),
		OTPStore:          *authstore.NewOTPStore(db),
		RoleStore:         *authstore.NewRoleStore(db),
		RefreshTokenStore: *authstore.NewRefreshTokenStore(db),
	}
}

// Model list
func (s *Storage) Models() map[string]interface{} {
	return map[string]any{
		"user":          &authstore.User{},
		"role":          &authstore.Role{},
		"refresh_token": &authstore.RefreshToken{},
		"otp":           &authstore.OTP{},
	}
}
