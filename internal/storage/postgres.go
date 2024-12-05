package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/mradigen/short/internal/logger"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgres(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the database: %v", err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to ping the database: %v", err)
	}

	// Create the "urls" table if it does not exist:
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS urls (
		slug VARCHAR(255) NOT NULL PRIMARY KEY,
		url TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure `urls` table exists: %v", err)
	}

	logger.Debug("DB connection established")

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Save(slug string, u string) (string, error) {
	_, err := s.db.Exec("INSERT INTO urls (slug, url) VALUES ($1, $2)", slug, u)
	if err != nil {
		return "", ErrExists
	}
	return slug, nil
}

func (s *PostgresStorage) Get(slug string) (string, error) {
	var u string
	err := s.db.QueryRow("SELECT url FROM urls WHERE slug = $1", slug).Scan(&u)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNotFound
		}
		return "", ErrNotFound // TODO: Better error handling
	}
	return u, nil
}

func (s *PostgresStorage) Close() error {
	return s.db.Close()
}
