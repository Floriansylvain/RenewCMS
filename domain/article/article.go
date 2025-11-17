package article

import (
	entity "RenewCMS/adapters/secondary/gateways/models"
	domain "RenewCMS/domain/image"
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
	images []*entity.Image,
	isOnline bool,
	createdAt time.Time,
	updatedAt time.Time,
) Article {
	domainImages := make([]*domain.Image, len(images))
	for i, img := range images {
		domainImage := domain.FromDB(
			img.ID,
			img.Path,
			img.ArticleID,
			img.CreatedAt,
			img.UpdatedAt,
		)
		domainImages[i] = &domainImage
	}
	return Article{
		ID:        id,
		Title:     title,
		Body:      body,
		Images:    domainImages,
		IsOnline:  isOnline,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
