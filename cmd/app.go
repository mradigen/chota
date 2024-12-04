package main

import (
	"github.com/mradigen/short/internal/api"
	"github.com/mradigen/short/internal/shortener"
	"github.com/mradigen/short/internal/storage"
	"github.com/mradigen/short/internal/config"
)

func main() {
	cfg := config.Load()

	var store storage.Storage
	if cfg.STORAGE_MODE == "memory" {
		store = storage.NewMemory()
	} else if cfg.STORAGE_MODE == "postgres" {
		store, _ = storage.NewPostgres(cfg.DATABASE_URL) // TODO: Error handling
	}
	s := shortener.New(store)

	api.Start(cfg.BIND_ADDRESS, cfg.PORT, s)

	// TODO: Handle error on close
	store.Close()
}
