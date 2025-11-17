package useCases

import (
	"GoCMS/domain/article"
	"GoCMS/domain/gateways"
)

type GetArticleUseCase struct {
	articleRepository gateways.IArticleRepository
}

func NewGetArticleUseCase(articleRepository gateways.IArticleRepository) *GetArticleUseCase {
	return &GetArticleUseCase{articleRepository}
}

func (g *GetArticleUseCase) GetArticle(id uint32) (article.Article, error) {
	return g.articleRepository.Get(id)
}

func (g *GetArticleUseCase) GetArticleByName(name string) (article.Article, error) {
	return g.articleRepository.GetByName(name)
}
