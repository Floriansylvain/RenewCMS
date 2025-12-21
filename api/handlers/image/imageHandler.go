package image

import (
	"RenewCMS/internal/infrastructure/useCases"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	CreateUseCase *useCases.CreateImageUseCase
	UpdateUseCase *useCases.UpdateArticleUseCase
}

func NewHandler(create *useCases.CreateImageUseCase, update *useCases.UpdateArticleUseCase) *Handler {
	return &Handler{
		CreateUseCase: create,
		UpdateUseCase: update,
	}
}

func (h *Handler) PostImage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newImage, err := h.CreateUseCase.CreateImage(file, *fileHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.UpdateUseCase.AddImage(uint32(idInt), newImage.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newJson := map[string]any{"location": newImage.Path}
	newJsonBytes, _ := json.Marshal(newJson)

	_, _ = w.Write(newJsonBytes)
}
