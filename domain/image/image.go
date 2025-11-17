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

func FromDB(id uint32, path string, articleId uint32, createdAt time.Time, updatedAt time.Time) Image {
	return Image{
		ID:        id,
		Path:      path,
		ArticleID: articleId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
