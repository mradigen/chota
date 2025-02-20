package api

import (
	"net/http"
	"strconv"

	"github.com/mradigen/chota/internal/log"
	"github.com/mradigen/chota/internal/shortener"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "https://c.phy0.in" || origin == "https://b.phy0.in" || origin == "https://k.phy0.in" || origin == "http://localhost:5173" || origin == "http://127.0.0.1:5173" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		}
		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Start(address string, port int, s *shortener.Shortener) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slug := r.URL.Path[1:]
		longURL, err := s.Retrieve(slug)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		http.Redirect(w, r, longURL, http.StatusSeeOther)
	})

	mux.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
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

	log.Info("API server listening on " + address + ":" + strconv.Itoa(port))
	err := http.ListenAndServe(address+":"+strconv.Itoa(port), corsMiddleware(mux))
	if err != nil {
		log.Info("Server failed to start: " + err.Error())
	}
}
