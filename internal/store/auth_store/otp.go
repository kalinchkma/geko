package authstore

import (
	"fmt"
	"geko/internal/db"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

// Otp model
type OTP struct {
	gorm.Model
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
	UserId    uint      `json:"user_id"`
}

// Otp store
type OTPStore struct {
	db *db.Database
}

func NewOTPStore(db *db.Database) *OTPStore {
	return &OTPStore{
		db,
	}
}

// Create
func (otpStore *OTPStore) Create(otp OTP) error {
	// Store otp to database
	res := otpStore.db.ORM.Create(&otp)

	// Check if any error
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// Find OTP by user id
func (otpStore OTPStore) FindOTPByUserID(userID uint) (OTP, error) {
	var otp OTP

	res := otpStore.db.ORM.Where("user_id = ?", userID).Find(&otp)
	if res.Error != nil {
		return OTP{}, res.Error
	}

	if res.RowsAffected == 0 {
		return OTP{}, fmt.Errorf("otp not found")
	}

	return otp, nil
}

// Generate otp
func (otpStore *OTPStore) GenerateOTP(length int) string {
	otp := ""
	for range length {
		otp += fmt.Sprintf("%d", rand.Intn(10))
	}
	return otp
}
