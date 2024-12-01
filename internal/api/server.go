package api

import (
	"net/http"
	"github.com/mradigen/short/internal/shortener"
	"github.com/mradigen/short/internal/storage"
)

func Start() {
	// TODO: Refractor
	m := storage.NewMemory()
	s := shortener.New(m)

	http.HandleFunc("GET /", Resolve(s))
	http.HandleFunc("GET /shorten", Shorten(s))

    http.ListenAndServe(":8080", nil)
}
