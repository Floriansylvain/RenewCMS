package gateways

import (
	"GoCMS/domain/article"
)

type IArticleRepository interface {
	Get(id uint32) (article.Article, error)
	GetByName(name string) (article.Article, error)
	GetAll() []article.Article
	Create(post article.Article) (article.Article, error)
	UpdateBody(id uint32, body string) (article.Article, error)
	UpdateIsOnline(id uint32, isOnline bool) (article.Article, error)
	Delete(id uint32) error
	AddImage(postId uint32, imageId uint32) error
}
