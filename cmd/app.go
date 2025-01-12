package main

import (
	"github.com/mradigen/chota/internal/api"
	"github.com/mradigen/chota/internal/config"
	"github.com/mradigen/chota/internal/shortener"
	"github.com/mradigen/chota/internal/storage"
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
