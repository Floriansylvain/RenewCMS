package image

import (
	"time"
)

type Image struct {
	ID        uint32    `json:"id"`
	Path      string    `json:"path"`
	ArticleID uint32    `json:"article_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
