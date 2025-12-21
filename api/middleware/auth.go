package middleware

import "net/http"

const KeyContentType = "Content-Type"

func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(KeyContentType, "application/json")
		next.ServeHTTP(w, r)
	})
}

func HtmlContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(KeyContentType, "text/html")
		next.ServeHTTP(w, r)
	})
}
