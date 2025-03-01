package authstore

import (
	"errors"
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
	UserId    uint      `json:"user_id" gorm:"not null;constraint:OnDelete:CASCADE;"`
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

// Verify otp
func (otpStore *OTPStore) VerifyOTP(userID uint, inputOTP string) bool {
	var otp OTP

	// Fetch the latest OTP for the user
	err := otpStore.db.ORM.Where("user_id = ?", userID).Order("created_at desc").First(&otp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Invalid or expired OTP
			return false
		}
		return false
	}

	// Check if OTP is expired
	if time.Now().After(otp.ExpiresAt) {
		return false
	}

	// Validate OTP
	if otp.Code != inputOTP {
		return false
	}

	// If OTP is valid, delete it after successful verification
	otpStore.db.ORM.Delete(&otp)

	return true
}
