package shortener

import (
	"github.com/mradigen/short/internal/storage"
	"net/url"
	"errors"
	"crypto/rand"
	"encoding/base64"
)

var ErrInvalidURL = errors.New("Invalid URL")

type Shortener struct {
	store storage.Storage
}

func New(store storage.Storage) *Shortener {
	return &Shortener{store: store}
}

// Should probably move this somewhere
func generateRandomString(length int) string {
	bytes := make([]byte, (length*3)/4+1)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length]
}

func (s *Shortener)Shorten(u string) (string, error) {

	// Parsing
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return "", ErrInvalidURL
	}

	slug, err := s.store.Save(generateRandomString(4), u)
	if err != nil { return "", nil }

	return slug, nil
} 


func (s *Shortener)Retrieve(slug string) (string, error) {

	u, err := s.store.Get(slug)
	if err != nil {
		return "", err
	}

	return u, nil
}
