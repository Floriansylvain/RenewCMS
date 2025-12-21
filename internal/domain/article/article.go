package article

import (
	domain "RenewCMS/internal/domain/image"
	"time"
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

func FromApi(
	title string,
	body string,
) Article {
	return Article{
		Title: title,
		Body:  body,
	}
}

func FromDb(
	id uint32,
	title string,
	body string,
	images []*domain.Image,
	isOnline bool,
	createdAt time.Time,
	updatedAt time.Time,
) Article {
	return Article{
		ID:        id,
		Title:     title,
		Body:      body,
		Images:    images,
		IsOnline:  isOnline,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
