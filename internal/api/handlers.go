package api

import (
	"net/http"
	// "fmt"
	"github.com/mradigen/short/internal/shortener"
)

func Shorten(s* shortener.Shortener) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract the "url" query parameter
		longURL := r.URL.Query().Get("url")
		if longURL == "" {
			http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
			return
		}

		// Shorten the URL
		shortURL, err := s.Shorten(longURL)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Respond with the short URL
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"short_url":"` + shortURL + `"}`))
	}
}

func Resolve(s* shortener.Shortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slug := r.URL.Path[1:];
		longURL, err := s.Retrieve(slug);
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		http.Redirect(w, r, longURL, http.StatusSeeOther)
	}
}
