package useCases

import "RenewCMS/internal/domain/article"

type UpdateArticleUseCase struct {
	articleRepository article.Repository
}

func NewUpdateArticleUseCase(articleRepository article.Repository) *UpdateArticleUseCase {
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
