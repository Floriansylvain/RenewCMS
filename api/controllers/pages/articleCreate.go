package pages

import (
	"GoCMS/api"
	"GoCMS/useCases"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

type ArticleCreatePageError struct {
	IsError bool
	Message string
}

func GetArticleCreatePageTemplate(postName string, errorMessage string) ([]byte, error) {
	navbarTmpl, _ := api.Container.GetPageUseCase.GetPage("componentNavbar", nil)
	return api.Container.GetPageUseCase.GetPage("articleCreate", map[string]any{
		"Navbar": template.HTML(navbarTmpl),
		"Head":   headTmpl,
		"PageError": ArticleCreatePageError{
			IsError: errorMessage != "",
			Message: errorMessage,
		},
		"Name": postName,
	})
}

func PostArticleCreatePage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	postName := r.FormValue("name")
	pattern := regexp.MustCompile("^[a-zA-Z0-9À-ÖØ-öø-ÿĀ-ſḀ-ỿ ]{3,50}$")

	if !pattern.MatchString(postName) {
		postsTmpl, _ := GetArticleCreatePageTemplate(postName, "Name should be alphanumeric, and between 3 and 50 characters.")
		_, _ = w.Write(postsTmpl)
		return
	}

	post, err := api.Container.CreateArticleUseCase.CreateArticle(useCases.CreateArticleCommand{
		Title: postName,
		Body:  "",
	})
	if err != nil {
		postsTmpl, _ := GetArticleCreatePageTemplate(postName, "Something went wrong when creating the article, please contact admin.")
		_, _ = w.Write(postsTmpl)
		return
	}

	http.Redirect(w, r, "/article/"+strconv.Itoa(int(post.ID))+"/edit", http.StatusSeeOther)
}

func GetArticleCreatePage(w http.ResponseWriter, _ *http.Request) {
	postsTmpl, _ := GetArticleCreatePageTemplate("", "")
	_, _ = w.Write(postsTmpl)
}
