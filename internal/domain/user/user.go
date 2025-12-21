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

func FromApi(
	username string,
	password string,
	passwordResetCode string,
	email string,
	verificationCode string,
) User {
	expiration := time.Now().Add(2 * time.Hour)
	return User{
		Username:               username,
		Password:               password,
		PasswordResetCode:      passwordResetCode,
		Email:                  email,
		IsVerified:             false,
		VerificationCode:       verificationCode,
		VerificationExpiration: expiration,
	}
}

func FromDb(
	id uint32,
	username string,
	password string,
	passwordResetCode string,
	email string,
	isVerified bool,
	verificationCode string,
	verificationExpiration time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) User {
	return User{
		ID:                     id,
		Username:               username,
		Password:               password,
		PasswordResetCode:      passwordResetCode,
		Email:                  email,
		IsVerified:             isVerified,
		VerificationCode:       verificationCode,
		VerificationExpiration: verificationExpiration,
		CreatedAt:              createdAt,
		UpdatedAt:              updatedAt,
	}
}
