package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID                     uint32 `gorm:"primaryKey;autoIncrement"`
	Username               string `gorm:"unique;not null"`
	Password               string `gorm:"not null"`
	PasswordResetCode      string `gorm:"unique"`
	Email                  string `gorm:"unique;not null"`
	IsVerified             bool   `gorm:"default=false;not null"`
	VerificationCode       string `gorm:"unique"`
	VerificationExpiration time.Time
	CreatedAt              time.Time `gorm:"autoCreateTime"`
	UpdatedAt              time.Time `gorm:"autoUpdateTime"`
}
