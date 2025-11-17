package route

import (
	"RenewCMS/api/controllers/article"
	"RenewCMS/api/controllers/auth"
	"RenewCMS/api/controllers/pages"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/MadAppGang/httplog"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
)

const keyContentType = "Content-Type"

func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(keyContentType, "application/json")
		next.ServeHTTP(w, r)
	})
}

func HtmlContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(keyContentType, "text/html")
		next.ServeHTTP(w, r)
	})
}

func InitJwt() {
	auth.Token = jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)
}

func GetHelloWorld(w http.ResponseWriter, _ *http.Request) {
	msg, _ := json.Marshal(map[string]string{"message": "Hello World"})
	_, _ = w.Write(msg)
}

func InitBackendRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(httplog.LoggerWithName("backend"))
	r.Use(JsonContentTypeMiddleware)
	r.Get("/", GetHelloWorld)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(auth.Token))
		r.Use(jwtauth.Authenticator(auth.Token))
		r.Mount("/article", article.NewArticleRouter())
	})
	r.Mount("/auth", auth.NewAuthRouter())

	return r
}

func InitFrontendRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(httplog.LoggerWithName("frontend"))
	r.Use(HtmlContentTypeMiddleware)
	r.Mount("/", pages.NewPageRouter())

	return r
}

func InitRoutes() *chi.Mux {
	backend := InitBackendRoutes()
	frontend := InitFrontendRoutes()

	apiRouter := chi.NewRouter()
	apiRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";"),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", keyContentType, "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	apiRouter.Mount("/v1", backend)
	apiRouter.Mount("/", frontend)

	return apiRouter
}
