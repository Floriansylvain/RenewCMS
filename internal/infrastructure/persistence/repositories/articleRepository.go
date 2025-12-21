package persistence

import (
	domainArticle "RenewCMS/internal/domain/article"
	"RenewCMS/internal/infrastructure/persistence/mappers"
	entity "RenewCMS/internal/infrastructure/persistence/models"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db}
}

func (a *ArticleRepository) Get(id uint32) (domainArticle.Article, error) {
	var article entity.Article
	err := a.db.Model(&entity.Article{}).Preload("Images").First(&article, id).Error
	if err != nil {
		return domainArticle.Article{}, err
	}

	return mappers.ArticleToDomain(article), nil
}

func (a *ArticleRepository) GetByName(name string) (domainArticle.Article, error) {
	var article entity.Article
	err := a.db.Model(&entity.Article{}).Where("title = ?", name).First(&article).Error
	if err != nil {
		return domainArticle.Article{}, err
	}

	return mappers.ArticleToDomain(article), nil
}

func (a *ArticleRepository) Create(article domainArticle.Article) (domainArticle.Article, error) {
	creationResult := a.db.Create(&entity.Article{
		Title: article.Title,
		Body:  article.Body,
	})
	if creationResult.Error != nil {
		return domainArticle.Article{}, creationResult.Error
	}

	var createdArticle entity.Article
	creationResult.Scan(&createdArticle)

	return mappers.ArticleToDomain(createdArticle), nil
}

func (a *ArticleRepository) GetAll() []domainArticle.Article {
	var articles []entity.Article
	err := a.db.Model(&entity.Article{}).Find(&articles).Error
	if err != nil {
		return []domainArticle.Article{}
	}

	var domainArticles = make([]domainArticle.Article, 0)
	for _, article := range articles {
		domainArticles = append(domainArticles, mappers.ArticleToDomain(article))
	}

	return domainArticles
}

func (a *ArticleRepository) UpdateBody(id uint32, body string) (domainArticle.Article, error) {
	var localArticle entity.Article
	err := a.db.Model(&entity.Article{}).First(&localArticle, id).Error
	if err != nil {
		return domainArticle.Article{}, err
	}

	localArticle.Body = body
	err = a.db.Save(&localArticle).Error
	if err != nil {
		return domainArticle.Article{}, err
	}

	newArticle := mappers.ArticleToDomain(localArticle)

	return newArticle, nil
}

func (a *ArticleRepository) UpdateIsOnline(id uint32, isOnline bool) (domainArticle.Article, error) {
	var localArticle entity.Article
	err := a.db.Model(&entity.Article{}).First(&localArticle, id).Error
	if err != nil {
		return domainArticle.Article{}, err
	}

	localArticle.IsOnline = isOnline
	err = a.db.Save(&localArticle).Error
	if err != nil {
		return domainArticle.Article{}, err
	}

	return mappers.ArticleToDomain(localArticle), nil
}

func (a *ArticleRepository) Delete(id uint32) error {
	return a.db.Delete(&entity.Article{}, id).Error
}

func (a *ArticleRepository) AddImage(articleId uint32, imageId uint32) error {
	var localArticle entity.Article
	err := a.db.Model(&entity.Article{}).First(&localArticle, articleId).Error
	if err != nil {
		return err
	}

	var localImage entity.Image
	err = a.db.Model(&entity.Image{}).First(&localImage, imageId).Error
	if err != nil {
		return err
	}

	err = a.db.Model(&localArticle).Association("Images").Append(&localImage)
	if err != nil {
		return err
	}

	return nil
}

var _ domainArticle.Repository = &ArticleRepository{}
