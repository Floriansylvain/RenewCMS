package article

import (
	domain "RenewCMS/internal/domain/image"
	"errors"
	"time"
)

var (
	ErrInvalidTitle    = errors.New("article's title length should be between 3 and 200 caracters")
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
