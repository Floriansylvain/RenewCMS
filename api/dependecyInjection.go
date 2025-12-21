package api

import (
	domainArticle "RenewCMS/internal/domain/article"
	domainImage "RenewCMS/internal/domain/image"
	domainMail "RenewCMS/internal/domain/mail"
	domainUser "RenewCMS/internal/domain/user"
	"RenewCMS/internal/infrastructure/mailer"
	"RenewCMS/internal/infrastructure/persistence"
	"RenewCMS/internal/infrastructure/persistence/models"
	"RenewCMS/internal/infrastructure/useCases"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/jwtauth/v5"
	"gorm.io/gorm"
)

type UseCases struct {
	Token                *jwtauth.JWTAuth
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
	SendMailUseCase      *useCases.SendMailUseCase
	CreateImageUseCase   *useCases.CreateImageUseCase
	DeleteImageUseCase   *useCases.DeleteImageUseCase
}

type Repositories struct {
	ArticleRepo domainArticle.Repository
	UserRepo    domainUser.Repository
	ImageRepo   domainImage.Repository
	MailRepo    domainMail.Repository
}

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

func getRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		ArticleRepo: persistence.NewArticleRepository(db),
		UserRepo:    persistence.NewUserRepository(db),
		ImageRepo:   persistence.NewImageRepository(db),
		MailRepo:    mailer.NewMailRepository(),
	}
}

func getUseCases(repos *Repositories) *UseCases {
	return &UseCases{
		Token:                jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil),
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
		SendMailUseCase:      useCases.NewSendMailUseCase(repos.MailRepo),
		CreateImageUseCase:   useCases.NewCreateImageUseCase(repos.ImageRepo),
		DeleteImageUseCase:   useCases.NewDeleteImageUseCase(repos.ImageRepo),
	}
}

func InitContainer() *UseCases {
	db := getDb()
	repos := getRepositories(db)
	return getUseCases(repos)
}
