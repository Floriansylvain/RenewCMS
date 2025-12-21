package useCases

import "RenewCMS/internal/domain/article"

type GetArticleUseCase struct {
	articleRepository article.Repository
}

func NewGetArticleUseCase(articleRepository article.Repository) *GetArticleUseCase {
	return &GetArticleUseCase{articleRepository}
}

func (g *GetArticleUseCase) GetArticle(id uint32) (article.Article, error) {
	return g.articleRepository.Get(id)
}

func (g *GetArticleUseCase) GetArticleByName(name string) (article.Article, error) {
	return g.articleRepository.GetByName(name)
}
