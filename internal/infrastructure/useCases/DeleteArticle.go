package useCases

import "RenewCMS/internal/domain/article"

type DeleteArticleUseCase struct {
	articleRepository article.Repository
}

func NewDeleteArticleUseCase(articleRepository article.Repository) *DeleteArticleUseCase {
	return &DeleteArticleUseCase{articleRepository}
}

func (g *DeleteArticleUseCase) DeleteArticle(userId uint32) error {
	return g.articleRepository.Delete(userId)
}
