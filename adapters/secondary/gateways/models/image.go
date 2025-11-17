package models

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	ID        uint32 `gorm:"primary_key;auto_increment;not_null"`
	Path      string
	ArticleID uint32
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
