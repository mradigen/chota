package api

import (
	"net/http"
	"strconv"

	"github.com/mradigen/short/internal/logger"
	"github.com/mradigen/short/internal/shortener"
)

func Start(address string, port int, s *shortener.Shortener) {

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		slug := r.URL.Path[1:]
		longURL, err := s.Retrieve(slug)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		http.Redirect(w, r, longURL, http.StatusSeeOther)
	})

	http.HandleFunc("GET /shorten", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		longURL := r.URL.Query().Get("url")
		if longURL == "" {
			http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
			return
		}

		shortURL, err := s.Shorten(longURL)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"short_url":"` + shortURL + `"}`))
	})

	logger.Info("API server listening on " + address + ":" + strconv.Itoa(port))
	http.ListenAndServe(address+":"+strconv.Itoa(port), nil)
}
