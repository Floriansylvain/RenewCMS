package api

import (
	"RenewCMS/api/handlers/article"
	"RenewCMS/api/handlers/auth"
	"RenewCMS/api/handlers/frontend"
	"RenewCMS/api/middleware"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/MadAppGang/httplog"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
)

func enableLoggingIfDev(r *chi.Mux, routerName string) {
	if os.Getenv("ENVIRONMENT") == "development" {
		r.Use(httplog.LoggerWithName(routerName))
	}
}

func GetHelloWorld(w http.ResponseWriter, _ *http.Request) {
	msg, _ := json.Marshal(map[string]string{"message": "Hello World"})
	_, _ = w.Write(msg)
}

func NewArticleRouter(container *UseCases) http.Handler {
	r := chi.NewRouter()

	articleHandler := article.NewHandler(container.CreateArticleUseCase, container.GetArticleUseCase, container.ListArticlesUseCase, container.DeleteArticleUseCase)

	r.Get("/{id}", articleHandler.GetArticle)
	r.Post("/", articleHandler.PostArticle)
	r.Get("/", articleHandler.ListArticles)
	r.Delete("/{id}", articleHandler.DeleteArticle)
	return r
}

func NewAuthRouter(container *UseCases) http.Handler {
	r := chi.NewRouter()

	authHandler := auth.NewHandler(container.Token, container.GetUserUseCase, container.ListUsersUseCase, container.CreateUserUseCase)

	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)
	r.Post("/logout", authHandler.Logout)
	return r
}

func InitBackendRoutes(container *UseCases) *chi.Mux {
	r := chi.NewRouter()

	enableLoggingIfDev(r, "backend")

	r.Use(middleware.JsonContentTypeMiddleware)
	r.Get("/", GetHelloWorld)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(container.Token))
		r.Use(jwtauth.Authenticator(container.Token))
		r.Mount("/article", NewArticleRouter(container))
	})
	r.Mount("/auth", NewAuthRouter(container))

	return r
}

func InitFrontendRoutes() *chi.Mux {
	r := chi.NewRouter()

	enableLoggingIfDev(r, "frontend")

	r.Mount("/", frontend.NewFrontendRouter())

	return r
}

func InitRoutes(container *UseCases) *chi.Mux {
	backend := InitBackendRoutes(container)
	frontend := InitFrontendRoutes()

	apiRouter := chi.NewRouter()
	apiRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";"),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	apiRouter.Mount("/v1", backend)
	apiRouter.Mount("/", frontend)

	return apiRouter
}
