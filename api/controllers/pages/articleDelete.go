package pages

import (
	"RenewCMS/api"
	"RenewCMS/api/controllers/article"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetArticleDeletePage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, article.IdUint32ErrorMessage, http.StatusBadRequest)
		return
	}

	localArticle, err := api.Container.GetArticleUseCase.GetArticle(uint32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Println(localArticle.Images)
	for _, image := range localArticle.Images {
		err = api.Container.DeleteImageUseCase.DeleteImage(image.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	err = api.Container.DeleteArticleUseCase.DeleteArticle(uint32(id))
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/article", http.StatusSeeOther)
}
