package article

import (
	domain "RenewCMS/internal/domain/image"
	"errors"
	"time"
)

var (
	ErrArticleNotFound = errors.New("article not found")
	ErrArticleOffline  = errors.New("article is offline")
)

type Article struct {
	ID        uint32          `json:"id"`
	Title     string          `json:"title"`
	Body      string          `json:"body"`
	Images    []*domain.Image `json:"images"`
	IsOnline  bool            `json:"is_online"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
