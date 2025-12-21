package useCases

import "RenewCMS/internal/domain/article"

type ListArticlesUseCase struct {
	articleRepository article.Repository
}

func NewListArticlesUseCase(articleRepository article.Repository) *ListArticlesUseCase {
	return &ListArticlesUseCase{articleRepository}
}

func (g *ListArticlesUseCase) ListArticles() []article.Article {
	return g.articleRepository.GetAll()
}
