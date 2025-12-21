package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID        uint32 `gorm:"primary_key;auto_increment;not_null"`
	Title     string
	Body      string
	Images    []*Image  `gorm:"many2many:article_images;"`
	IsOnline  bool      `gorm:"not_null;default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
