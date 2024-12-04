package storage

import (
	"errors"
)

type Storage interface {
    Save(slug, longURL string) (string, error)
    Get(slug string) (string, error)
	Close() error
	// Exists(longURL string) (string, error)
}

var ErrNotFound = errors.New("Slug not found")
var ErrExists = errors.New("Slug exists")
