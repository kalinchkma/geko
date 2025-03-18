package authstore

import (
	"time"

	"github.com/kalinchkma/geko/internal/db"
	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expire_at"`
	RevokeAt  time.Time `json:"revoke_at"`
	UserID    uint      `json:"user_id" gorm:"not null;constraint:OnDelete:CASCADE;"`
}

type RefreshTokenStore struct {
	db *db.Database
}

// Constructor
func NewRefreshTokenStore(db *db.Database) *RefreshTokenStore {
	return &RefreshTokenStore{db}
}

// Create Refresh Token
