package authstore

import (
	"geko/internal/db"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expire_at"`
	RevokeAt  time.Time `json:"revoke_at"`
	UserID    uuid.UUID `json:"user_id"`
}

type RefreshTokenStore struct {
	db *db.Database
}

// Constructor
func NewRefreshTokenStore(db *db.Database) *RefreshTokenStore {
	return &RefreshTokenStore{db}
}
