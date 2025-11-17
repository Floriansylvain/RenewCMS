package pages

import (
	"GoCMS/api"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func updateIsOnline(articleId string, isOnline bool) (error, int) {
	articleIdInt, err := strconv.Atoi(articleId)
	if err != nil {
		return errors.New("the server expects the ID to be in the format of an unsigned 32-bit integer (uint32)"), http.StatusBadRequest
	}

	_, err = api.Container.UpdateArticleUseCase.UpdateIsOnline(uint32(articleIdInt), isOnline)
	if err != nil {
		return errors.New("the requested resource, identified by its unique ID, could not be found on the server"), http.StatusNotFound
	}

	return nil, http.StatusOK
}

func GetArticleUnpublishPage(w http.ResponseWriter, r *http.Request) {
	articleId := chi.URLParam(r, "id")
	err, statusCode := updateIsOnline(articleId, false)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	http.Redirect(w, r, "/article", http.StatusSeeOther)
}

func GetArticlePublishPage(w http.ResponseWriter, r *http.Request) {
	articleId := chi.URLParam(r, "id")
	err, statusCode := updateIsOnline(articleId, true)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	http.Redirect(w, r, "/article", http.StatusSeeOther)
}
