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
func (otpStore *OTPStore) FindOTPByUserID(userID uint) (OTP, error) {
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

// Delete otp by user id
// func (otpStore *OTPStore) DeleteOTPByUserID(userID uint) (OTP, error) {
// 	var otp OTP
// 	res := otpStore.db.ORM.Where("user_id = ?", userID).Delete(&otp)
// 	if res.Error != nil {
// 		return OTP{}, res.Error
// 	}
// 	return otp, nil
// }

// Generate otp
func (otpStore *OTPStore) GenerateOTP(length int) string {
	otp := ""
	for range length {
		otp += fmt.Sprintf("%d", rand.Intn(10))
	}
	return otp
}

// Verify otp
func (otpStore *OTPStore) VerifyOTP(userID uint, inputOTP string) error {
	var otp OTP

	// Fetch the latest OTP for the user
	res := otpStore.db.ORM.Where("user_id = ?", userID).Find(&otp)
	if res.Error != nil {
		fmt.Println("Otp error", res.Error.Error(), userID, inputOTP)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			// Invalid or expired OTP
			return errors.New("invalid OTP")
		}
		return errors.New("invalid OTP")
	}
	if res.RowsAffected == 0 {
		fmt.Println("otp not found", userID, inputOTP)
		return errors.New("invalid OTP")
	}

	// Check if OTP is expired
	if time.Now().After(otp.ExpiresAt) {
		fmt.Println("OTP expires")
		return errors.New("OTP expires")
	}

	// Validate OTP
	if otp.Code != inputOTP {
		fmt.Println("Invalid OTP")
		return errors.New("invalid OTP")
	}

	// If OTP is valid, delete it after successful verification
	otpStore.db.ORM.Unscoped().Delete(&otp)
	return nil
}
