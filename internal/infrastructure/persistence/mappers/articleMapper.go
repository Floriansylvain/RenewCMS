package mappers

import (
	"RenewCMS/internal/domain/article"
	"RenewCMS/internal/domain/image"
	"RenewCMS/internal/infrastructure/persistence/models"
)

func ArticleToDomain(m models.Article) article.Article {
	domainImages := make([]*image.Image, len(m.Images))
	for i, img := range m.Images {
		domainImg := ImageToDomain(*img)
		domainImages[i] = &domainImg
	}

	return article.Article{
		ID:        m.ID,
		Title:     m.Title,
		Body:      m.Body,
		Images:    domainImages,
		IsOnline:  m.IsOnline,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ArticleToModel(a article.Article) models.Article {
	modelImages := make([]*models.Image, len(a.Images))
	for i, img := range a.Images {
		modelImage := ImageToModel(*img)
		modelImages[i] = &modelImage
	}

	return models.Article{
		Title:     a.Title,
		Body:      a.Body,
		Images:    modelImages,
		IsOnline:  a.IsOnline,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}
