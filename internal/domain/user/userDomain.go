package user

import (
	"time"
)

type User struct {
	ID                     uint32    `json:"id"`
	Username               string    `json:"username"`
	Password               string    `json:"password"`
	PasswordResetCode      string    `json:"password_reset_code"`
	Email                  string    `json:"email"`
	IsVerified             bool      `json:"is_verified"`
	VerificationCode       string    `json:"verification_code"`
	VerificationExpiration time.Time `json:"verification_expiration"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}
