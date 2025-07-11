package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var DB *sql.DB

func ConnectDB() error {
	db_user := os.Getenv("DB_USERNAME")
	db_name := os.Getenv("DB_NAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", db_user, db_password, db_host, db_name))
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Настройка пула соединений
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(0)

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping db: %w", err)
	}

	DB = db
	return nil
}
