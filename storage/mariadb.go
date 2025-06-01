package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	db_user := os.Getenv("DB_USERNAME")
	db_name := os.Getenv("DB_NAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", db_user, db_password, db_host, db_name))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}
