package api

import (
	"RenewCMS/adapters/secondary/gateways"
	"RenewCMS/adapters/secondary/gateways/models"
	domainGateways "RenewCMS/domain/gateways"
	"RenewCMS/useCases"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type UseCases struct {
	CreateArticleUseCase *useCases.CreateArticleUseCase
	GetArticleUseCase    *useCases.GetArticleUseCase
	ListArticlesUseCase  *useCases.ListArticlesUseCase
	UpdateArticleUseCase *useCases.UpdateArticleUseCase
	DeleteArticleUseCase *useCases.DeleteArticleUseCase
	GetUserUseCase       *useCases.GetUserUseCase
	CreateUserUseCase    *useCases.CreateUserUseCase
	UpdateUserUseCase    *useCases.UpdateUserUseCase
	DeleteUserUseCase    *useCases.DeleteUserUseCase
	ListUsersUseCase     *useCases.ListUsersUseCase
	GetPageUseCase       *useCases.GetPageUseCase
	SendMailUseCase      *useCases.SendMailUseCase
	CreateImageUseCase   *useCases.CreateImageUseCase
	DeleteImageUseCase   *useCases.DeleteImageUseCase
}

type Repositories struct {
	ArticleRepo domainGateways.IArticleRepository
	UserRepo    domainGateways.IUserRepository
	ImageRepo   domainGateways.IImageRepository
	MailRepo    domainGateways.IMailRepository
	PageRepo    domainGateways.IPageRepository
}

var Container *UseCases

func getDb() *gorm.DB {
	dbName := os.Getenv("DB_FILE")
	if err := os.MkdirAll(filepath.Dir(dbName), os.ModePerm); err != nil {
		panic("Unable to create necessary subdirectories: " + err.Error())
	}
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("Unable to open the database: " + err.Error())
	}

	if err := db.AutoMigrate(&models.Article{}, &models.User{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	return db
}

func initRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		ArticleRepo: gateways.NewArticleRepository(db),
		UserRepo:    gateways.NewUserRepository(db),
		ImageRepo:   gateways.NewImageRepository(db),
		MailRepo:    gateways.NewMailRepository(),
		PageRepo:    gateways.NewPageRepository(),
	}
}

func initUseCases(repos *Repositories) *UseCases {
	return &UseCases{
		CreateArticleUseCase: useCases.NewCreateArticleUseCase(repos.ArticleRepo),
		GetArticleUseCase:    useCases.NewGetArticleUseCase(repos.ArticleRepo),
		ListArticlesUseCase:  useCases.NewListArticlesUseCase(repos.ArticleRepo),
		UpdateArticleUseCase: useCases.NewUpdateArticleUseCase(repos.ArticleRepo),
		DeleteArticleUseCase: useCases.NewDeleteArticleUseCase(repos.ArticleRepo),
		GetUserUseCase:       useCases.NewGetUserUseCase(repos.UserRepo),
		CreateUserUseCase:    useCases.NewCreateUserUseCase(repos.UserRepo),
		UpdateUserUseCase:    useCases.NewUpdateUserUseCase(repos.UserRepo),
		DeleteUserUseCase:    useCases.NewDeleteUserUseCase(repos.UserRepo),
		ListUsersUseCase:     useCases.NewListUsersUseCase(repos.UserRepo),
		GetPageUseCase:       useCases.NewGetPageUseCase(repos.PageRepo),
		SendMailUseCase:      useCases.NewSendMailUseCase(repos.MailRepo),
		CreateImageUseCase:   useCases.NewCreateImageUseCase(repos.ImageRepo),
		DeleteImageUseCase:   useCases.NewDeleteImageUseCase(repos.ImageRepo),
	}
}

func InitContainer() {
	db := getDb()
	repos := initRepositories(db)
	Container = initUseCases(repos)
}
