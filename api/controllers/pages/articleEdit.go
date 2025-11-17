package pages

import (
	"GoCMS/api"
	"GoCMS/domain/article"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ArticleEditPageAlert struct {
	IsError bool
	Message string
}

func getArticleEditPageTemplate(post article.Article, alert ArticleEditPageAlert) []byte {
	navbarTmpl, _ := api.Container.GetPageUseCase.GetPage("componentNavbar", nil)
	articleTmpl, _ := api.Container.GetPageUseCase.GetPage("articleEdit", map[string]any{
		"Navbar":  template.HTML(navbarTmpl),
		"Head":    headTmpl,
		"Article": post,
		"Alert":   alert,
		"Secured": os.Getenv("ENVIRONMENT") == "production",
	})
	return articleTmpl
}

func PostArticleEditPage(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "id")
	articleIDint, err := strconv.Atoi(articleID)
	if err != nil {
		_, _ = w.Write(getArticleEditPageTemplate(article.Article{}, ArticleEditPageAlert{
			IsError: true,
			Message: "Could not find the requested article.",
		}))
		return
	}

	_ = r.ParseForm()
	articleBody := r.FormValue("articleBody")

	getArticle, _ := api.Container.GetArticleUseCase.GetArticle(uint32(articleIDint))
	updatedArticle, err := api.Container.UpdateArticleUseCase.UpdateBody(getArticle.ID, articleBody)
	if err != nil {
		_, _ = w.Write(getArticleEditPageTemplate(article.Article{Body: articleBody}, ArticleEditPageAlert{
			IsError: true,
			Message: "Could not save the article: " + err.Error(),
		}))
		return
	}

	_, _ = w.Write(getArticleEditPageTemplate(updatedArticle, ArticleEditPageAlert{
		IsError: false,
		Message: "Article successfully edited!",
	}))
}

func GetArticleEditPage(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "id")
	articleIDint, err := strconv.Atoi(articleID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	getArticle, _ := api.Container.GetArticleUseCase.GetArticle(uint32(articleIDint))

	_, _ = w.Write(getArticleEditPageTemplate(getArticle, ArticleEditPageAlert{
		IsError: false,
		Message: "",
	}))
}
