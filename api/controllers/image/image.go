package image

import (
	"GoCMS/api"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func PostImage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newImage, err := api.Container.CreateImageUseCase.CreateImage(file, *fileHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = api.Container.UpdateArticleUseCase.AddImage(uint32(idInt), newImage.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newJson := map[string]any{"location": newImage.Path}
	newJsonBytes, _ := json.Marshal(newJson)

	_, _ = w.Write(newJsonBytes)
}
