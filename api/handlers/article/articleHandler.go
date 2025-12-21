package article

import (
	"RenewCMS/api/pkg/apperr"
	validator "RenewCMS/api/pkg/validator"
	domain "RenewCMS/internal/domain/article"
	"RenewCMS/internal/infrastructure/useCases"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

	article, err := h.GetUseCase.GetOnlineArticle(uint32(id))
	if err != nil {
		if errors.Is(err, domain.ErrArticleNotFound) {
			http.Error(w, "Article not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	articleJson, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(articleJson); err != nil {
		fmt.Printf("Failed to write response: %v", err)
	}
}

func (h *Handler) PostArticle(w http.ResponseWriter, r *http.Request) {
	var localArticle PostArticle
	if err := json.NewDecoder(r.Body).Decode(&localArticle); err != nil {
		http.Error(w, apperr.BodyErrorMessage, http.StatusBadRequest)
		return
	}

	if err := validator.Validate.Struct(localArticle); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdArticle, err := h.CreateUseCase.CreateArticle(useCases.CreateArticleCommand{
		Title: localArticle.Title,
		Body:  localArticle.Body,
	})
	if err != nil {
		http.Error(w, "Failed to create article", http.StatusInternalServerError)
		return
	}

	articleJson, err := json.Marshal(createdArticle)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(articleJson); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (h *Handler) ListArticles(w http.ResponseWriter, _ *http.Request) {
	articles, err := h.ListUseCase.ListOnlineArticles()
	if err != nil {
		http.Error(w, "Failed to retrieve articles", http.StatusInternalServerError)
		return
	}

	articlesJson, err := json.Marshal(articles)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(articlesJson); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (h *Handler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, apperr.IdUint32ErrorMessage, http.StatusBadRequest)
		return
	}

	if err := h.DeleteUseCase.DeleteArticle(uint32(id)); err != nil {
		http.Error(w, "Failed to delete article", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
