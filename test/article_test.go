package test

import (
	"RenewCMS/internal/domain/article"
	"RenewCMS/internal/infrastructure/useCases"
	"errors"
	"testing"
)

type mockArticleRepo struct {
	articles map[uint32]article.Article
}

func (m *mockArticleRepo) FindByFilters(f article.Filters) ([]article.Article, error) {
	var results []article.Article
	for _, a := range m.articles {
		if f.IsOnline != nil && a.IsOnline == *f.IsOnline {
			results = append(results, a)
		}
	}
	return results, nil
}

func (m *mockArticleRepo) Get(id uint32) (article.Article, error) {
	a, ok := m.articles[id]
	if !ok {
		return article.Article{}, article.ErrArticleNotFound
	}
	return a, nil
}

func (m *mockArticleRepo) Create(a article.Article) (article.Article, error) {
	a.ID = uint32(len(m.articles) + 1)
	m.articles[a.ID] = a
	return a, nil
}

func (m *mockArticleRepo) Delete(id uint32) error {
	delete(m.articles, id)
	return nil
}

func (m *mockArticleRepo) GetByName(n string) (article.Article, error) { return article.Article{}, nil }
func (m *mockArticleRepo) GetAll() []article.Article                   { return nil }
func (m *mockArticleRepo) UpdateBody(id uint32, b string) (article.Article, error) {
	return article.Article{}, nil
}
func (m *mockArticleRepo) UpdateIsOnline(id uint32, o bool) (article.Article, error) {
	return article.Article{}, nil
}
func (m *mockArticleRepo) AddImage(p, i uint32) error { return nil }

var _ article.Repository = &mockArticleRepo{}

func TestCreateArticle(t *testing.T) {
	repo := &mockArticleRepo{articles: make(map[uint32]article.Article)}
	uc := useCases.NewCreateArticleUseCase(repo)

	t.Run("Valid Title", func(t *testing.T) {
		cmd := useCases.CreateArticleCommand{Title: "Valid Title", Body: "Content"}
		res, err := uc.CreateArticle(cmd)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if res.Title != cmd.Title {
			t.Errorf("Expected %s, got %s", cmd.Title, res.Title)
		}
	})

	t.Run("Short Title", func(t *testing.T) {
		cmd := useCases.CreateArticleCommand{Title: "Ab", Body: "Content"}
		_, err := uc.CreateArticle(cmd)
		if !errors.Is(err, article.ErrInvalidTitle) {
			t.Errorf("Expected ErrInvalidTitle")
		}
	})
}

func TestGetOnlineArticle(t *testing.T) {
	repo := &mockArticleRepo{articles: map[uint32]article.Article{
		1: {ID: 1, IsOnline: true},
		2: {ID: 2, IsOnline: false},
	}}
	uc := useCases.NewGetArticleUseCase(repo)

	t.Run("Article Online", func(t *testing.T) {
		_, err := uc.GetOnlineArticle(1)
		if err != nil {
			t.Errorf("Expected article to be found")
		}
	})

	t.Run("Article Offline", func(t *testing.T) {
		_, err := uc.GetOnlineArticle(2)
		if !errors.Is(err, article.ErrArticleNotFound) {
			t.Errorf("Offline article should return not found")
		}
	})
}
