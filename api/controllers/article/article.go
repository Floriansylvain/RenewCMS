package article

import (
	"GoCMS/api"
	"GoCMS/api/controllers/auth"
	"GoCMS/domain/article"
	"GoCMS/useCases"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PostArticle struct {
	Title string `json:"title" validate:"required,min=3,max=50"`
	Body  string `json:"body" validate:"required,max=10000"`
}

const IdUint32ErrorMessage = "The server expects the ID to be in the format of an unsigned 32-bit integer (uint32)."

func getArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, IdUint32ErrorMessage, http.StatusBadRequest)
		return
	}

	localArticle, err := api.Container.GetArticleUseCase.GetArticle(uint32(id))
	if err != nil || !localArticle.IsOnline {
		http.Error(w, "The requested resource, identified by its unique ID, could not be found on the server.", http.StatusNotFound)
		return
	}

	articleJson, _ := json.Marshal(localArticle)
	_, _ = w.Write(articleJson)
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	var localArticle PostArticle
	err := json.NewDecoder(r.Body).Decode(&localArticle)
	if err != nil {
		http.Error(w, auth.BodyErrorMessage, http.StatusBadRequest)
		return
	}

	err = api.Validate.Struct(localArticle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdArticle, err := api.Container.CreateArticleUseCase.CreateArticle(useCases.CreateArticleCommand{
		Title: localArticle.Title,
		Body:  localArticle.Body,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	articleJson, _ := json.Marshal(createdArticle)

	_, _ = w.Write(articleJson)
}

func listArticles(w http.ResponseWriter, _ *http.Request) {
	articles := api.Container.ListArticlesUseCase.ListArticles()
	onlineArticles := make([]article.Article, 0)
	for _, localArticle := range articles {
		if localArticle.IsOnline {
			onlineArticles = append(onlineArticles, localArticle)
		}
	}
	articlesJson, _ := json.Marshal(onlineArticles)
	_, _ = w.Write(articlesJson)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, IdUint32ErrorMessage, http.StatusBadRequest)
	}

	err = api.Container.DeleteArticleUseCase.DeleteArticle(uint32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	_, _ = w.Write([]byte("article deleted"))
}

func NewArticleRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", getArticle)
	r.Post("/", postArticle)
	r.Get("/", listArticles)
	r.Delete("/{id}", deleteArticle)
	return r
}
