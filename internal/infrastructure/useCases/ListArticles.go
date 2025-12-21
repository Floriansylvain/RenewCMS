package useCases

import "RenewCMS/internal/domain/article"

type ListArticlesUseCase struct {
	articleRepository article.Repository
}

func NewListArticlesUseCase(articleRepository article.Repository) *ListArticlesUseCase {
	return &ListArticlesUseCase{articleRepository}
}

func (u *ListArticlesUseCase) ListOnlineArticles() ([]article.Article, error) {
	online := true
	return u.articleRepository.FindByFilters(article.Filters{
		IsOnline: &online,
	})
}

func (u *ListArticlesUseCase) ListAllArticles() ([]article.Article, error) {
	return u.articleRepository.FindByFilters(article.Filters{})
}
