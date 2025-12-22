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
	if len(createArticle.Title) < 3 || len(createArticle.Title) > 200 {
		return article.Article{}, article.ErrInvalidTitle
	}

	return g.articleRepository.Create(article.Article{
		Title:    createArticle.Title,
		Body:     createArticle.Body,
		IsOnline: false,
	})
}
