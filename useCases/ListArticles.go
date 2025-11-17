package useCases

import (
	"RenewCMS/domain/article"
	"RenewCMS/domain/gateways"
)

type ListArticlesUseCase struct {
	articleRepository gateways.IArticleRepository
}

func NewListArticlesUseCase(articleRepository gateways.IArticleRepository) *ListArticlesUseCase {
	return &ListArticlesUseCase{articleRepository}
}

func (g *ListArticlesUseCase) ListArticles() []article.Article {
	return g.articleRepository.GetAll()
}
