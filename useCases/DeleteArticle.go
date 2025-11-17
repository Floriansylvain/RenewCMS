package useCases

import (
	"RenewCMS/domain/gateways"
)

type DeleteArticleUseCase struct {
	articleRepository gateways.IArticleRepository
}

func NewDeleteArticleUseCase(articleRepository gateways.IArticleRepository) *DeleteArticleUseCase {
	return &DeleteArticleUseCase{articleRepository}
}

func (g *DeleteArticleUseCase) DeleteArticle(userId uint32) error {
	return g.articleRepository.Delete(userId)
}
