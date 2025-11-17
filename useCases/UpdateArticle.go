package useCases

import (
	"RenewCMS/domain/article"
	"RenewCMS/domain/gateways"
)

type UpdateArticleUseCase struct {
	articleRepository gateways.IArticleRepository
}

func NewUpdateArticleUseCase(articleRepository gateways.IArticleRepository) *UpdateArticleUseCase {
	return &UpdateArticleUseCase{articleRepository}
}

func (g *UpdateArticleUseCase) UpdateBody(id uint32, body string) (article.Article, error) {
	return g.articleRepository.UpdateBody(id, body)
}

func (g *UpdateArticleUseCase) AddImage(articleId uint32, imageId uint32) error {
	return g.articleRepository.AddImage(articleId, imageId)
}

func (g *UpdateArticleUseCase) UpdateIsOnline(id uint32, isOnline bool) (article.Article, error) {
	return g.articleRepository.UpdateIsOnline(id, isOnline)
}
