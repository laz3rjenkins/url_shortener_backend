package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"minq-backend/internal/shortener"
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

	GetRoutes(r)

	r.Run()
}
