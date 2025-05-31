package main

import (
	"github.com/gin-gonic/gin"
	"minq-backend/internal/shortener"
)

func GetRoutes(r *gin.Engine) {
	shortener.SetupRoutes(r)
}

func main() {
	r := gin.Default()

	GetRoutes(r)

	r.Run()
}
