package persistence

import (
	domain "RenewCMS/internal/domain/article"
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

func (a *ArticleRepository) FindByFilters(filters domain.Filters) ([]domain.Article, error) {
	query := a.db.Model(&entity.Article{})

	if filters.IsOnline != nil {
		query = query.Where("is_online = ?", *filters.IsOnline)
	}

	var entityArticles []entity.Article
	err := query.Find(&entityArticles).Error

	domainArticles := make([]domain.Article, 0, len(entityArticles))
	for _, entityArticle := range entityArticles {
		domainArticles = append(domainArticles, mappers.ArticleToDomain(entityArticle))
	}

	return domainArticles, err
}

func (a *ArticleRepository) Get(id uint32) (domain.Article, error) {
	var article entity.Article
	err := a.db.Model(&entity.Article{}).Preload("Images").First(&article, id).Error
	if err != nil {
		return domain.Article{}, err
	}

	return mappers.ArticleToDomain(article), nil
}

func (a *ArticleRepository) GetByName(name string) (domain.Article, error) {
	var article entity.Article
	err := a.db.Model(&entity.Article{}).Where("title = ?", name).First(&article).Error
	if err != nil {
		return domain.Article{}, err
	}

	return mappers.ArticleToDomain(article), nil
}

func (a *ArticleRepository) Create(article domain.Article) (domain.Article, error) {
	creationResult := a.db.Create(&entity.Article{
		Title:    article.Title,
		Body:     article.Body,
		IsOnline: article.IsOnline,
	})
	if creationResult.Error != nil {
		return domain.Article{}, creationResult.Error
	}

	var createdArticle entity.Article
	creationResult.Scan(&createdArticle)

	return mappers.ArticleToDomain(createdArticle), nil
}

func (a *ArticleRepository) GetAll() []domain.Article {
	var articles []entity.Article
	err := a.db.Model(&entity.Article{}).Find(&articles).Error
	if err != nil {
		return []domain.Article{}
	}

	var domainArticles = make([]domain.Article, 0)
	for _, article := range articles {
		domainArticles = append(domainArticles, mappers.ArticleToDomain(article))
	}

	return domainArticles
}

func (a *ArticleRepository) UpdateBody(id uint32, body string) (domain.Article, error) {
	var localArticle entity.Article
	err := a.db.Model(&entity.Article{}).First(&localArticle, id).Error
	if err != nil {
		return domain.Article{}, err
	}

	localArticle.Body = body
	err = a.db.Save(&localArticle).Error
	if err != nil {
		return domain.Article{}, err
	}

	newArticle := mappers.ArticleToDomain(localArticle)

	return newArticle, nil
}

func (a *ArticleRepository) UpdateIsOnline(id uint32, isOnline bool) (domain.Article, error) {
	var localArticle entity.Article
	err := a.db.Model(&entity.Article{}).First(&localArticle, id).Error
	if err != nil {
		return domain.Article{}, err
	}

	localArticle.IsOnline = isOnline
	err = a.db.Save(&localArticle).Error
	if err != nil {
		return domain.Article{}, err
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

var _ domain.Repository = &ArticleRepository{}
