package frontend

import (
	"RenewCMS/internal/infrastructure/assets"
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

type spaHandler struct {
	staticFS fs.FS
	modTime  time.Time
}

func NewFrontendRouter() http.Handler {
	r := chi.NewRouter()

	distFS, _ := fs.Sub(assets.DistEmbed, "dist")
	handler := &spaHandler{
		staticFS: distFS,
		modTime:  time.Now(),
	}

	r.Get("/*", handler.ServeHTTP)
	return r
}

func (h *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.cleanPath(r.URL.Path)
	if h.fileExists(path) {
		h.serveAsset(w, r, path)
		return
	}
	h.serveIndex(w, r)
}

func (h *spaHandler) cleanPath(path string) string {
	return strings.TrimPrefix(filepath.Clean(path), "/")
}

func (h *spaHandler) fileExists(path string) bool {
	if path == "" || path == "." {
		return false
	}
	f, err := h.staticFS.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()
	return true
}

func (h *spaHandler) serveAsset(w http.ResponseWriter, r *http.Request, path string) {
	w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	http.FileServer(http.FS(h.staticFS)).ServeHTTP(w, r)
}

func (h *spaHandler) serveIndex(w http.ResponseWriter, r *http.Request) {
	f, err := h.staticFS.Open("index.html")
	if err != nil {
		http.Error(w, "Index not found", http.StatusNotFound)
		return
	}
	defer f.Close()

	content, ok := f.(io.ReadSeeker)
	if !ok {
		http.Error(w, "File not seekable", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	http.ServeContent(w, r, "index.html", h.modTime, content)
}
