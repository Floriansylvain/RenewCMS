package pages

import (
	"GoCMS/api"
	"html/template"
	"net/http"
)

func GetArticlesPage(w http.ResponseWriter, _ *http.Request) {
	articles := api.Container.ListArticlesUseCase.ListArticles()

	var formattedArticles []map[string]any
	for _, article := range articles {
		formattedArticles = append(formattedArticles, map[string]any{
			"ID":        article.ID,
			"Title":     article.Title,
			"IsOnline":  article.IsOnline,
			"CreatedAt": article.CreatedAt.Format("2006-01-02 15:04:05"),
			"UpdatedAt": article.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	navbarTmpl, _ := api.Container.GetPageUseCase.GetPage("componentNavbar", nil)
	articlesTmpl, _ := api.Container.GetPageUseCase.GetPage("articles", map[string]any{
		"Navbar":   template.HTML(navbarTmpl),
		"Head":     headTmpl,
		"Articles": formattedArticles,
	})

	_, _ = w.Write(articlesTmpl)
}
