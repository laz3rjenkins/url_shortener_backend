package shortener

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type ShortenRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *ShortenRepository {
	return &ShortenRepository{db: db}
}

func (r *ShortenRepository) SaveShortenURL(ctx context.Context, data ShortenUrlAttributes) error {
	query := `INSERT INTO shorten_links (shorten_url, original_url, redirect_count, created_at)
	          VALUES (?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, data.ShortenURL, data.OriginalURL, 0, time.Now())
	if err != nil {
		return fmt.Errorf("failed to insert shorten URL: %w", err)
	}

	return nil
}

func (r *ShortenRepository) GetByOriginalURL(ctx context.Context, originalURL string) (*ShortenUrlAttributes, error) {
	query := `SELECT shorten_url, redirect_count FROM shorten_links WHERE original_url = ?`

	row := r.db.QueryRowContext(ctx, query, originalURL)
	var data ShortenUrlAttributes
	data.OriginalURL = originalURL

	err := row.Scan(&data.ShortenURL, &data.RedirectCount)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get shorten URL: %w", err)
	}

	return &data, nil
}

func (r *ShortenRepository) GetByShortenURL(ctx context.Context, shortenURL string) (*string, error) {
	query := `SELECT original_url FROM shorten_links WHERE shorten_url = ?`

	var originalURL string
	err := r.db.QueryRowContext(ctx, query, shortenURL).Scan(&originalURL)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get original URL: %w", err)
	}

	return &originalURL, nil
}
