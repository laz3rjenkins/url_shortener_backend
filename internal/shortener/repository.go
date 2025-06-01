package shortener

import (
	"fmt"
	"minq-backend/storage"
	"time"
)

func SaveShortenUrl(data ShortenUrlAttributes) error {
	db, err := storage.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	// todo: если такой оргинал юрл уже есть, то не генерить, а просто отдать существующий
	query := "INSERT INTO shorten_links (shorten_url, original_url, redirect_count, created_at) VALUES (?, ?, ?, ?)"

	_, err = db.Exec(query, data.ShortenURL, data.OriginalURL, 0, time.Now())
	if err != nil {
		return fmt.Errorf("failed to insert shorten URL: %w", err)
	}

	return nil
}
