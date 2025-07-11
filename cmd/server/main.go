package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"minq-backend/internal/shortener"
	"minq-backend/storage"
)

func GetRoutes(r *gin.Engine) {
	shortener.SetupRoutes(r)
}

func main() {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("error while loading .env")
	}

	if err := storage.ConnectDB(); err != nil {
		log.Fatalf("DB init failed: %v", err)
	}
	defer storage.DB.Close()

	GetRoutes(r)

	r.Run()
}
