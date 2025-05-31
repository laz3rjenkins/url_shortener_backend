package shortener

import "github.com/gin-gonic/gin"

const prefix = "/shortener"

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	r.POST(prefix+"/generate", GenerateUrl)
}
