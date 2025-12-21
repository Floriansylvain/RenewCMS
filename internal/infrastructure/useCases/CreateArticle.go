package useCases

import (
	"RenewCMS/internal/domain/article"
)

type CreateArticleUseCase struct {
	articleRepository article.Repository
}

type CreateArticleCommand struct {
	Title string
	Body  string
}

func NewCreateArticleUseCase(articleRepository article.Repository) *CreateArticleUseCase {
	return &CreateArticleUseCase{articleRepository}
}

func (g *CreateArticleUseCase) CreateArticle(createArticle CreateArticleCommand) (article.Article, error) {
	return g.articleRepository.Create(article.Article{
		Title: createArticle.Title,
		Body:  createArticle.Body,
	})
}
