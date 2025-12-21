package useCases

import "RenewCMS/internal/domain/article"

type GetArticleUseCase struct {
	articleRepository article.Repository
}

func NewGetArticleUseCase(articleRepository article.Repository) *GetArticleUseCase {
	return &GetArticleUseCase{articleRepository}
}

func (u *GetArticleUseCase) GetOnlineArticle(id uint32) (article.Article, error) {
	art, err := u.articleRepository.Get(id)
	if err != nil {
		return article.Article{}, err
	}

	if !art.IsOnline {
		return article.Article{}, article.ErrArticleNotFound
	}

	return art, nil
}

func (u *GetArticleUseCase) GetArticle(id uint32) (article.Article, error) {
	return u.articleRepository.Get(id)
}

func (g *GetArticleUseCase) GetArticleByName(name string) (article.Article, error) {
	return g.articleRepository.GetByName(name)
}
