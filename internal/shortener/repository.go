package shortener

import (
	"fmt"
	"minq-backend/storage"
	"time"
)

func SaveShortenUrl(data ShortenUrlAttributes) (ShortenUrlAttributes, error) {
	db, err := storage.ConnectDB()
	if err != nil {
		return data, err
	}
	defer db.Close()
	query := "INSERT INTO shorten_links (shorten_url, original_url, redirect_count, created_at) VALUES (?, ?, ?, ?)"

	_, err = db.Exec(query, data.ShortenURL, data.OriginalURL, 0, time.Now())
	if err != nil {
		return data, fmt.Errorf("failed to insert shorten URL: %w", err)
	}

	return data, nil
}

func GetUrlByOriginUrl(data *ShortenUrlAttributes) (*ShortenUrlAttributes, error) {
	db, err := storage.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT original_url, shorten_url, redirect_count FROM shorten_links WHERE original_url = ?"
	rows, err := db.Query(query, data.OriginalURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get shorten URL: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var shortenUrl ShortenUrlAttributes
	err = rows.Scan(&shortenUrl.OriginalURL, &shortenUrl.ShortenURL, &shortenUrl.RedirectCount)
	if err != nil {
		return nil, fmt.Errorf("failed to scan shorten URL: %w", err)
	}

	return &shortenUrl, nil
}

func GetUrlByShortenUrl(shortenUrl string) (*string, error) {
	db, err := storage.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT original_url FROM shorten_links WHERE shorten_url = ?"
	rows, err := db.Query(query, shortenUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to get shorten URL: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("url not found")
	}

	var originalUrl string
	err = rows.Scan(&originalUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to scan shorten URL: %w", err)
	}

	return &originalUrl, nil
}
