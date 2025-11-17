package useCases

import (
	"GoCMS/domain/article"
	"GoCMS/domain/gateways"
)

type CreateArticleUseCase struct {
	articleRepository gateways.IArticleRepository
}

type CreateArticleCommand struct {
	Title string
	Body  string
}

func NewCreateArticleUseCase(articleRepository gateways.IArticleRepository) *CreateArticleUseCase {
	return &CreateArticleUseCase{articleRepository}
}

func (g *CreateArticleUseCase) CreateArticle(createArticle CreateArticleCommand) (article.Article, error) {
	return g.articleRepository.Create(article.FromApi(
		createArticle.Title,
		createArticle.Body,
	))
}
