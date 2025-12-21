package article

import (
	"RenewCMS/api/pkg/apperr"
	validator "RenewCMS/api/pkg/validator"
	domainArticle "RenewCMS/internal/domain/article"
	"RenewCMS/internal/infrastructure/useCases"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	CreateUseCase *useCases.CreateArticleUseCase
	GetUseCase    *useCases.GetArticleUseCase
	ListUseCase   *useCases.ListArticlesUseCase
	DeleteUseCase *useCases.DeleteArticleUseCase
}

type PostArticle struct {
	Title string `json:"title" validate:"required,min=3,max=50"`
	Body  string `json:"body" validate:"required,max=10000"`
}

func NewHandler(create *useCases.CreateArticleUseCase, get *useCases.GetArticleUseCase, list *useCases.ListArticlesUseCase, delete *useCases.DeleteArticleUseCase) *Handler {
	return &Handler{
		CreateUseCase: create,
		GetUseCase:    get,
		ListUseCase:   list,
		DeleteUseCase: delete,
	}
}

func (h *Handler) GetArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, apperr.IdUint32ErrorMessage, http.StatusBadRequest)
		return
	}

	localArticle, err := h.GetUseCase.GetArticle(uint32(id))
	if err != nil || !localArticle.IsOnline {
		http.Error(w, "The requested resource, identified by its unique ID, could not be found on the server.", http.StatusNotFound)
		return
	}

	articleJson, _ := json.Marshal(localArticle)
	_, _ = w.Write(articleJson)
}

func (h *Handler) PostArticle(w http.ResponseWriter, r *http.Request) {
	var localArticle PostArticle
	err := json.NewDecoder(r.Body).Decode(&localArticle)
	if err != nil {
		http.Error(w, apperr.BodyErrorMessage, http.StatusBadRequest)
		return
	}

	err = validator.Validate.Struct(localArticle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdArticle, err := h.CreateUseCase.CreateArticle(useCases.CreateArticleCommand{
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

func (h *Handler) ListArticles(w http.ResponseWriter, _ *http.Request) {
	articles := h.ListUseCase.ListArticles()
	onlineArticles := make([]domainArticle.Article, 0)
	for _, localArticle := range articles {
		if localArticle.IsOnline {
			onlineArticles = append(onlineArticles, localArticle)
		}
	}
	articlesJson, _ := json.Marshal(onlineArticles)
	_, _ = w.Write(articlesJson)
}

func (h *Handler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, apperr.IdUint32ErrorMessage, http.StatusBadRequest)
		return
	}

	err = h.DeleteUseCase.DeleteArticle(uint32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, _ = w.Write([]byte("article deleted"))
}
