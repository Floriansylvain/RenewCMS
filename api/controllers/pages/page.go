package pages

import (
	"RenewCMS/api"
	"RenewCMS/api/controllers/auth"
	"RenewCMS/api/controllers/image"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
)

var headTmpl template.HTML

var contentTypes = map[string]string{
	".css":  "text/css",
	".js":   "application/javascript",
	".png":  "image/png",
	".jpg":  "image/jpeg",
	".webp": "image/webp",
	".svg":  "image/svg+xml",
	".ico":  "image/x-icon",
}

type PageError struct {
	Message string `json:"message"`
	IsError bool   `json:"isError"`
}

func NewPageError(message string) *PageError {
	return &PageError{
		Message: message,
		IsError: strings.Compare(message, "") != 0,
	}
}

func IsLoggedInMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !auth.IsLoggedIn(r) {
			http.Redirect(w, r, LoginRoute, http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func IsVerifiedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !auth.IsVerified(r) {
			http.Redirect(w, r, "/register/pending", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func IsNotVerifiedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth.IsVerified(r) {
			http.Redirect(w, r, "/register/pending", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, LoginRoute, http.StatusPermanentRedirect)
}

func GetLogout(w http.ResponseWriter, r *http.Request) {
	auth.RemoveJwtCookie(w)
	http.Redirect(w, r, LoginRoute, http.StatusSeeOther)
}

func StaticFileServerWithContentType(fsys http.FileSystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		filePath := filepath.Join("api/static", r.URL.Path)
		fileInfo, err := os.Stat(filePath)
		if err == nil && fileInfo.IsDir() {
			http.Error(w, "Folder direct access is disabled.", http.StatusForbidden)
			return
		}
		if ext := filepath.Ext(path); ext != "" {
			if ct, ok := contentTypes[ext]; ok {
				w.Header().Set("Cache-Control", "public, max-age=31536000")
				w.Header().Set("Content-Type", ct)
			}
		}
		http.FileServer(fsys).ServeHTTP(w, r)
	})
}

func InitHeadTmpl() {
	headTmplHtml, _ := api.Container.GetPageUseCase.GetPage("utilsHead", nil)
	headTmpl = template.HTML(headTmplHtml)
}

func NewPageRouter() http.Handler {
	r := chi.NewRouter()

	InitHeadTmpl()

	fileServer := StaticFileServerWithContentType(http.Dir("api/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	r.Get("/", GetLogin)

	r.Get(LoginRoute, GetLoginPageHandler(EmptyLoginPage))
	r.Post(LoginRoute, GetLoginPageHandler(EmptyLoginPage))

	r.Get("/register", GetRegisterPageHandler(EmptyRegisterPage))
	r.Post("/register", GetRegisterPageHandler(EmptyRegisterPage))

	r.Get("/logout", GetLogout)

	r.Get("/register/reset/request", GetPasswordResetRequest)
	r.Post("/register/reset/request", PostPasswordResetRequest)

	r.Get("/register/reset/validate", GetPasswordResetValidate)
	r.Post("/register/reset/validate", PostPasswordResetValidate)

	r.Group(func(r chi.Router) {
		r.Use(IsLoggedInMiddleware)
		r.Use(IsNotVerifiedMiddleware)

		r.Get("/register/pending", GetRegisterPendingPage)
		r.Post("/register/pending", PostRegisterPendingPage)

		r.Get("/register/validate", GetRegisterValidatePage)
	})

	r.Group(func(r chi.Router) {
		r.Use(IsLoggedInMiddleware)
		r.Use(IsVerifiedMiddleware)

		r.Get("/home", GetHomePage)

		r.Get("/article", GetArticlesPage)

		r.Get("/article/{id}/edit", GetArticleEditPage)
		r.Post("/article/{id}/edit", PostArticleEditPage)

		r.Get("/article/{id}/delete", GetArticleDeletePage)

		r.Get("/article/create", GetArticleCreatePage)
		r.Post("/article/create", PostArticleCreatePage)

		r.Post("/article/{id}/image/create", image.PostImage)

		r.Get("/article/{id}/publish", GetArticlePublishPage)
		r.Get("/article/{id}/unpublish", GetArticleUnpublishPage)

		r.Get("/integration", GetPageIntegration)
	})

	return r
}
