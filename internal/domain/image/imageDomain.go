package image

import (
	"io"
	"time"
)

type ImageInput struct {
	Content     io.Reader
	Filename    string
	Size        int64
	ContentType string
}

type Image struct {
	ID        uint32    `json:"id"`
	Path      string    `json:"path"`
	ArticleID uint32    `json:"article_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
